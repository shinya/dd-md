package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/ini.v1"
)

const (
	targetMark             = ":"
	markerSection          = "marker"
	variablesSection       = "variables"
	settingFileName        = "setting.ini"
	inputeMarkdownFileName = "template.md"
	// outputDirectory        = "output"
	// outputMarkdownFileName = "output.md"
)

func main() {

	// カレントディレクトリの取得
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	//	fmt.Println(dir)

	// // 実行モジュールのパスを取得
	// exe, err := os.Executable()
	// if err != nil {
	// 	log.Fatal(err)
	// 	os.Exit(1)
	// }
	// exePath := filepath.Dir(exe)

	var (
		templateMarkdownFilePath = flag.String("i", dir+"/"+inputeMarkdownFileName, "読み込む雛形Markdown")
		outputFilePath           = flag.String("o", "", "出力するMarkdown")
		settingFilePath          = flag.String("c", dir+"/"+settingFileName, "読み込む設定ファイル")
	)
	// コマンドライン引数の取得
	flag.Parse()

	// テンプレート存在確認
	if _, err := os.Stat(*templateMarkdownFilePath); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// 設定ファイル存在確認
	if _, err := os.Stat(*settingFilePath); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	variables := readSetting(*settingFilePath, variablesSection)
	markers := readSetting(*settingFilePath, markerSection)
	markdown := readMd(*templateMarkdownFilePath)

	startMaker := markers.Key("start").String()
	if startMaker == "" {
		startMaker = targetMark
	}

	endMaker := markers.Key("end").String()
	if endMaker == "" {
		endMaker = targetMark
	}

	// 設定数の分置換処理を実行する
	for _, v := range variables.KeyStrings() {
		markdown = strings.ReplaceAll(markdown, startMaker+v+endMaker, variables.Key(v).String())
	}

	// 指定があれば書き込み。なければ標準出力
	if *outputFilePath == "" {
		fmt.Println(markdown)
	} else {
		fmt.Println("Output Path: " + dir + "/" + *outputFilePath)
		// ファイルへの書き込み
		writeFunc(dir+"/"+*outputFilePath, markdown)
	}

}

// 設定ファイルの読み込み
func readSetting(fileName string, sectionName string) *ini.Section {
	cfg, err := ini.Load(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return cfg.Section(sectionName)
}

// テンプレートファイルの読み込み
func readMd(fileName string) string {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return string(b)
}

// 出力ファイルへの書き込み
func writeFunc(fileName string, outputString string) {
	err := ioutil.WriteFile(fileName, []byte(outputString), 0666)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func helpMessage() {
	helpString := `Usage:
	dd-md [OPTIONS]
  
  Application Options:
	-i  Import template read file (If not specified, read the file name template.md)
	-c  Configuration read file  (If not specified, read the file name setting.ini)
	-o  Output file name (If not specified, output to standard output)
`
	fmt.Printf("%s", helpString)
}
