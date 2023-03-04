package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCity(t *testing.T) {
	body, err := ioutil.ReadFile("city_test_data.html")
	if err != nil {
		panic(err)
	}

	const resultSize = 20
	parseResult := ParseCity(body)
	if len(parseResult.Requests) != resultSize {
		t.Errorf("Incorrect Requests size. Expected %d but got %d", resultSize, len(parseResult.Items))
	}

	if len(parseResult.Items) != resultSize {
		t.Errorf("Incorrect Items size. Expected %d but got %d", resultSize, len(parseResult.Requests))
	}

}
