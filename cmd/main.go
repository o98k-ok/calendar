package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
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
		items := alfred.NewItems().WithPreselect(time.Now().Format("2006-01-02"))
		for _, date := range dates {
			items.Append(date.ToAlfredElem())
		}
		items.Show()
	})

	app.Bind("detail", func(s []string) {
		detail := date.Detail(s[0])
		data, _ := json.Marshal(detail.DetailFilter())
		fmt.Println(string(data))
	})

	app.Bind("note", func(s []string) {
		filename := os.Getenv(date.NOTE_DATE_KEY)
		if len(filename) == 0 {
			alfred.Log("filename is empty")
			return
		}

		newContent := strings.TrimSpace(s[0])
		if len(newContent) != 0 && newContent != "note" {
			file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
			if err != nil {
				alfred.Log("open file error: %v", err)
				return
			}
			defer file.Close()
			if _, err := file.WriteString("* " + newContent + "\n"); err != nil {
				alfred.Log("write file error: %v", err)
				return
			}
		}

		content, _ := os.ReadFile(filename)
		result := map[string]any{
			"variables": map[string]string{
				date.NOTE_DATE_KEY: filename,
			},
			"response": string(content),
			"behaviour": map[string]string{
				"scroll": "end",
			},
		}
		data, _ := json.Marshal(result)
		fmt.Println(string(data))
	})

	app.Bind("cat", func(s []string) {
		filename := os.Getenv(date.NOTE_DATE_KEY)
		if len(filename) == 0 {
			alfred.Log("filename is empty")
			return
		}

		content, _ := os.ReadFile(filename)
		result := map[string]any{
			"response": string(content),
			"variables": map[string]string{
				"note_date": filename,
			},
		}
		data, _ := json.Marshal(result)
		fmt.Println(string(data))
	})

	app.Run(os.Args)
}
