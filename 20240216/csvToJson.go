package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	// CSV open
	file, err := os.Open("input.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// create csv reader
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	// json data
	var data []map[string]string
	header := records[0]
	// read csv file
	for i, record := range records {
		if i == 0 {
			continue
		}

		// レコードをマップに変換
		var item = make(map[string]string)

		for j, value := range record {
			item[header[j]] = value
		}
		// json data
		data = append(data, item)
	}
	// json形式に変換
	//jsonData, err := json.Marshal(data)
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))
}
