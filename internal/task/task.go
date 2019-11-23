package task

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

const (
	inputExample1Selector = "#task-statement > span > span.lang-ja > div:nth-child(7) > section > pre"
	inputExample2Selector = "#task-statement > span > span.lang-ja > div:nth-child(10) > section > pre"
	inputExample3Selector = "#task-statement > span > span.lang-ja > div:nth-child(13) > section > pre"
	inputExample4Selector = "#task-statement > span > span.lang-ja > div:nth-child(16) > section > pre"

	outputExample1Selector = "#task-statement > span > span.lang-ja > div:nth-child(8) > section > pre"
	outputExample2Selector = "#task-statement > span > span.lang-ja > div:nth-child(11) > section > pre"
	outputExample3Selector = "#task-statement > span > span.lang-ja > div:nth-child(14) > section > pre"
	outputExample4Selector = "#task-statement > span > span.lang-ja > div:nth-child(17) > section > pre"
)

type Example struct {
	Input1  string
	Output1 string

	Input2  string
	Output2 string

	Input3  string
	Output3 string

	Input4  string
	Output4 string
}

// GetExample return Example
// TODO 41 以下のコンテストでもできるようにする
// TODO URLとコンテスト番号がマッチしていないものにも対応したい
// なぜ goquery はpre#pre-pre-sample0 を探せないのか
func GetExample(contest, question string) (*Example, error) {
	url := fmt.Sprintf(`https://atcoder.jp/contests/%[1]s/tasks/%[1]s_%s`, contest, question)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}

	return &Example{
		Input1:  doc.Find(inputExample1Selector).Text(),
		Output1: doc.Find(outputExample1Selector).Text(),
		Input2:  doc.Find(inputExample2Selector).Text(),
		Output2: doc.Find(outputExample2Selector).Text(),
		Input3:  doc.Find(inputExample3Selector).Text(),
		Output3: doc.Find(outputExample3Selector).Text(),
		Input4:  doc.Find(inputExample4Selector).Text(),
		Output4: doc.Find(outputExample4Selector).Text(),
	}, nil
}
