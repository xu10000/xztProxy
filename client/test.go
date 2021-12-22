import (
	"testing"
	"encoding/json"
	jsoniter "github.com/json-iterator/go"
)

type Car struct {
	Name  string
	Name1 string
	Name2 string
	Name3 string
}

func BenchmarkStructJsoniter(b *testing.B) {
	c := Car{Name1: "xxxxxxxxxxxxx", Name2: "xxxxxxxxxxxxxxxxxxxx", Name3: "xxxxxxxxxxxxxxxxxx", Name: "buickxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"}
	for i := 0; i < b.N; i++ { //use b.N for looping
		jsoniter.Marshal(&c)
	}
}

func BenchmarkStructStd(b *testing.B) {
	c := Car{Name1: "xxxxxxxxxxxxx", Name2: "xxxxxxxxxxxxxxxxxxxx", Name3: "xxxxxxxxxxxxxxxxxx", Name: "buickxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"}
	for i := 0; i < b.N; i++ { //use b.N for looping
		json.Marshal(&c)
	}
}

func BenchmarkMapJsoniter(b *testing.B) {
	mymap := make(map[string]string, 10000)
	mymap["Name1"] = "xxxxxxxxxxxxx"
	mymap["Name2"] = "xxxxxxxxxxxxxxxxxxxx"
	mymap["Name3"] = "xxxxxxxxxxxxxxxxxx"
	mymap["Name"] = "buickxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	for i := 0; i < b.N; i++ { //use b.N for looping
		jsoniter.Marshal(&mymap)
	}
}

func BenchmarkMapStd(b *testing.B) {
	mymap := make(map[string]string, 10000)
	mymap["Name1"] = "xxxxxxxxxxxxx"
	mymap["Name2"] = "xxxxxxxxxxxxxxxxxxxx"
	mymap["Name3"] = "xxxxxxxxxxxxxxxxxx"
	mymap["Name"] = "buickxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	for i := 0; i < b.N; i++ { //use b.N for looping
		json.Marshal(&mymap)
	}
}