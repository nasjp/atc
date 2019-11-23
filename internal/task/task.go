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
	Input  string
	Output string
}

type Examples []*Example

// GetExample return Example
// TODO 41 以下のコンテストでもできるようにする
// TODO URLとコンテスト番号がマッチしていないものにも対応したい
// なぜ goquery はpre#pre-pre-sample0 を探せないのか
func GetExamples(contest, question string) (Examples, error) {
	url := fmt.Sprintf(`https://atcoder.jp/contests/%[1]s/tasks/%[1]s_%s`, contest, question)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}

	es := make(Examples, 0, 4)

	i1 := doc.Find(inputExample1Selector).Text()
	o1 := doc.Find(outputExample1Selector).Text()
	es = append(es, &Example{i1, o1})

	i2 := doc.Find(inputExample2Selector).Text()
	o2 := doc.Find(outputExample2Selector).Text()
	es = append(es, &Example{i2, o2})

	i3 := doc.Find(inputExample3Selector).Text()
	o3 := doc.Find(outputExample3Selector).Text()
	if i3 == "" {
		return es, nil
	}
	es = append(es, &Example{i3, o3})

	i4 := doc.Find(inputExample4Selector).Text()
	o4 := doc.Find(outputExample4Selector).Text()
	if i4 == "" {
		return es, nil
	}
	es = append(es, &Example{i4, o4})
	return es, nil
}
}
