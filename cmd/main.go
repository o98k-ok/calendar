package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/o98k-ok/calendar/internal/date"
	"github.com/o98k-ok/lazy/v2/alfred"
)

func main() {
	app := alfred.NewApp("日历助手")

	// 设置主题 && 字体颜色
	if os.Getenv("THEME") == "white" {
		date.MODE = "black"
	}

	app.Bind("all", func(s []string) {
		dates := date.GetDates()
		items := &date.Items{
			Preselect: time.Now().Format("2006-01-02"),
		}
		for _, date := range dates {
			items.Items = append(items.Items, date.ToAlfredElem())
		}
		data, _ := json.Marshal(items)
		fmt.Println(string(data))
	})

	app.Bind("detail", func(s []string) {
		detail := date.Detail(s[0])
		fmt.Println(detail.DetailFilter().Encode())
	})

	app.Run(os.Args)
}
