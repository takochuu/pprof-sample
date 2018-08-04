package main

import (
	"encoding/json"
	_ "net/http/pprof"
	"strings"

	"github.com/golang/glog"
)

var data = `
[
   {
      "Name": "takochuu",
      "Sex": "Male",
      "State": {
         "Height": 170,
         "Weight": 60,
         "Money": 0,
         "Country": "Japan"
      },
   },
   {
      "Name": "takochuu1",
      "Sex": "Male",
      "State": {
         "Height": 170,
         "Weight": 60,
         "Money": 0,
         "Country": "Japan"
      },
   },
   {
      "Name": "takochuu2",
      "Sex": "Male",
      "State": {
         "Height": 170,
         "Weight": 60,
         "Money": 0,
         "Country": "Japan"
      },
   },
]
`

func main() {
	for {
		parse(map[string]bool{
			"type": true,
		})
	}
}

func parse(option map[string]bool) {
	if option["type"] {
		decoder()
	} else {
		marshal()
	}
}

func marshal() {
	_, err := json.Marshal(data)
	if err != nil {
		glog.Errorln(err)
	}
}

type State struct {
	Height  int
	Weight  int
	Money   int
	Country string
}

type testDataMapper struct {
	Name  string
	Sex   string
	State State
}

func decoder() {
	decoder := json.NewDecoder(strings.NewReader(data))
	mapper := testDataMapper{}
	decoder.Decode(&mapper)
}
