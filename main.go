package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/ini.v1"
)

const (
	targetMark               = ":"
	settingFileName          = "setting.ini"
	templateMarkdownFileName = "template.md"
	sectionName              = "variables"
	outputDirectory          = "output"
	outputMarkdownFileName   = "output.md"
)

func main() {

	settings := readSetting(settingFileName)
	markdown := readMd(templateMarkdownFileName)

	// 設定の数分置換処理を実行する
	for k, v := range settings.KeyStrings() {
		fmt.Println(k, v, settings.Key(v))

		markdown = strings.ReplaceAll(markdown, targetMark+v+targetMark, settings.Key(v).String())
	}

	//	fmt.Println(markdown)

	// ファイルへの書き込み
	writeFunc(outputMarkdownFileName, markdown)

}

// 設定ファイルの読み込み
func readSetting(fileName string) *ini.Section {
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
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}

	return string(b)
}

// 出力ファイルへの書き込み
func writeFunc(fileName string, outputString string) {
	err := ioutil.WriteFile(outputDirectory+"/"+fileName, []byte(outputString), 0666)
	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}
