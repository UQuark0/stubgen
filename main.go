package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Pattern string `json:"pattern"`
	Values [][]interface{} `json:"values"`
}

func main() {
	if len(os.Args) < 2 {
		panic("must specify pattern file")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	config := Config{}

	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}

	for _, values := range config.Values {
		fmt.Printf(config.Pattern + "\n\n", values...)
	}
}
