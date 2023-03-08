package zhenai

import "encoding/json"

type UserProfile struct {
	Age        int
	Gender     string
	Name       string
	Height     int
	Income     string
	Marriage   string
	Education  string
	Occupation string
	Weight     int
}

func FromJsonObj(o interface{}) (UserProfile, error) {
	var res UserProfile
	s, err := json.Marshal(o)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(s, &res)
	return res, err
}
