package parser

import (
	"goprojects/crawler/engine"
	"goprojects/crawler/zhenai"
	"regexp"
	"strconv"
)

// example: <td><span class="label">年龄：</span>97岁</td>
const ageRe = `<td><span class=\"label\">年龄：</span>([0-9]+)岁</td>`

// example: <td><span class="label">性别：</span><span field="">男</span></td>
const genderRe = `<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`

// example: <h1 class="ceiling-name ib fl fs24 lh32 blue">学霸的芯迁就</h1>
const nameRe = `<h1 class="ceiling-name ib fl fs24 lh32 blue">([^<]+)</h1>`

// example: <td><span class="label">身高：</span>31CM</td>
const heightRe = `<td><span class="label">身高：</span>([0-9]+)CM</td>`

// example: <td><span class="label">月收入：</span>1-2000元</td>
const incomeRe = `<td><span class="label">月收入：</span>([^<]+)</td>`

// example: <td><span class="label">婚况：</span>离异</td>
const marriageRe = `<td><span class="label">婚况：</span>([^<]+)</td>`

// example: <td><span class="label">学历：</span>大学</td>
const educationRe = `<td><span class="label">学历：</span>([^<]+)</td>`

// example: <td><span class="label">职业： </span>金融</td>
const occupationRe = `<td><span class="label">职业： </span>([^<]+)</td>`

// exmaple: <td><span class="label">体重：</span><span field="">33KG</span></td>
const weightRe = `<td><span class="label">体重：</span><span field="">([0-9]+)KG</span></td>`

// example: http://localhost:8080/mock/album.zhenai.com/u/2858527851503377365
const idUrelRe = `http://localhost:8080/mock/album.zhenai.com/u/([0-9]+)`

func ParseProfile(body []byte, url string) engine.ParseResult {
	age, _ := strconv.Atoi(extractMatch(body, ageRe))
	gender := extractMatch(body, genderRe)
	name := extractMatch(body, nameRe)
	height, _ := strconv.Atoi(extractMatch(body, heightRe))
	income := extractMatch(body, incomeRe)
	marriage := extractMatch(body, marriageRe)
	education := extractMatch(body, educationRe)
	occupation := extractMatch(body, occupationRe)
	weight, _ := strconv.Atoi(extractMatch(body, weightRe))

	profile := zhenai.UserProfile{
		Age:        age,
		Gender:     gender,
		Name:       name,
		Height:     height,
		Income:     income,
		Marriage:   marriage,
		Education:  education,
		Occupation: occupation,
		Weight:     weight,
	}
	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Id:      extractMatch([]byte(url), idUrelRe),
				Payload: profile,
			},
		},
	}
	return result
}

func extractMatch(body []byte, re string) string {
	r := regexp.MustCompile(re)
	match := r.FindSubmatch(body)
	if match == nil {
		return ""
	}

	return string(match[1])
}
