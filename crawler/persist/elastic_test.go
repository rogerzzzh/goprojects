package persist

import (
	"bytes"
	"encoding/json"
	"goprojects/crawler/engine"
	"goprojects/crawler/zhenai"
	"testing"
)

func TestElasticConnection(t *testing.T) {
	e8, err := GetClient()
	if err != nil {
		t.Errorf("Not able to create ES client.")
		panic(err)
	}
	res, err := e8.Info()
	defer res.Body.Close()
	if err != nil {
		t.Errorf("Not able to get ES info.")
		panic(err)
	}

	if res.IsError() {
		t.Errorf("Error saving the data, got response %s", res)
	}

	profile := engine.Item{
		Id:  "abc",
		Url: "http://www.baidu.com",
		Payload: zhenai.UserProfile{
			Age:        15,
			Gender:     "Male",
			Name:       "Mundo",
			Height:     220,
			Income:     "3000",
			Marriage:   "",
			Education:  "",
			Occupation: "",
			Weight:     200,
		},
	}

	const index = "test_index"
	err = Save(e8, index, profile)
	if err != nil {
		panic(err)
	}

	res, err = e8.GetSource("test_index", profile.Id)
	if res.IsError() {
		t.Errorf("Error accessing the data, got response %s", res)
	}
	defer res.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	var actual engine.Item
	var actualProfile zhenai.UserProfile
	if err = json.Unmarshal(buf.Bytes(), &actual); err != nil {
		t.Errorf("Error parsing the response body: %s", err)
	} else {
		actualProfile, err = zhenai.FromJsonObj(actual.Payload)
		if actualProfile != profile.Payload {
			t.Errorf("Got unexpected payload, expect %s, got %s", profile, actual)
		}
	}
}
