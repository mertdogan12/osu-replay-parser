package main

import (
	"fmt"
	"log"
	"os"

	"github.com/blamebutton/orpa/parser"
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
	replay, err := parser.GetFileReplay(filePath)
	if err != nil {
		fmt.Println("Filepath:", filePath)
		log.Fatal(err)
		return
	}

	fmt.Println(replay.ReplayData)
}
