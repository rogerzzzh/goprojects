package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	body, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	parseResult := ParseCityList(body)
	const resultSize = 470
	if len(parseResult.Items) != 470 {
		t.Errorf("Incorrect Items size. Expected %d but got %d", resultSize, len(parseResult.Items))
	}
	if len(parseResult.Requests) != 470 {
		t.Errorf("Incorrect Requests size. Expected %d but got %d", resultSize, len(parseResult.Requests))
	}
}
