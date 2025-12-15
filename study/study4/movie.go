package main

import (
	"fmt"
	"log"

	"github.com/goccy/go-json"
)

// 这里是js的键值对，默认是字段名，改掉后例如released对应的就是Year，也会显示released
// omitempty的0值不显示，如果布尔类型是true才会显示
type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true, Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

func main() {
	//将go切片转换为js格式切片（[]byte）
	//如果不懂前缀和缩进是什么，不妨运行并修改一下，就明白了
	data, err := json.MarshalIndent(movies, "前缀", "缩进")
	if err != nil {
		//打印带时间戳的错误信息并且停止运行（类似os.exit(1)）
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles)
}
