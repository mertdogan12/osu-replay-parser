package main

import "github.com/mertdogan-org/osu-replay-converter/pkg/opengl"

func main() {
	// Gets command line arguments
	/*if len(os.Args) != 2 {
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

	os.WriteFile("out.json", json, 0644)*/
	opengl.Init(1280, 720)
}
