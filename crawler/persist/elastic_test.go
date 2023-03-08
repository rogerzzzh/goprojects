package persist

import (
	"encoding/json"
	"goprojects/crawler/zhenai"
	"testing"
)

func TestElasticConnection(t *testing.T) {
	e8, err := getClient()
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

	profile := zhenai.UserProfile{
		Name:   "Mundo",
		Age:    15,
		Gender: "Male",
	}

	id, err := save(profile)
	if err != nil {
		panic(err)
	}

	res, err = e8.GetSource("test_index", id)
	if res.IsError() {
		t.Errorf("Error accessing the data, got response %s", res)
	}
	defer res.Body.Close()

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		t.Errorf("Error parsing the response body: %s", err)
	} else {
		if r["Name"] != profile.Name {
			t.Errorf("Got unexpected name, expect %s, got %s", profile.Name, r["Name"])
		}
	}
}
