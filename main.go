package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	parser "github.com/mertdogan12/osu-replay-converter/pkg/osu-replay-parser"
)

func main() {
	// Gets command line arguments
	if len(os.Args) != 2 {
		log.Fatal("Filepath is missing \n",
			"[command] <filepath>")
		return
	}

	// Parse the replay
	filePath := os.Args[1]
	replay, err := parser.ConvertToObject(filePath)

	if err != nil {
		log.Fatal(err)
		return
	}

	json, err := json.MarshalIndent(replay, "\t", "\t")

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(string(json))
}
