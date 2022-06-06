package parser_test

import (
	"io/ioutil"
	"testing"

	parser "github.com/mertdogan12/osu-replay-parser"
)

func TestParse(t *testing.T) {
	replayDir := "./replays/"
	files, err := ioutil.ReadDir(replayDir)
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		_, err := parser.ParseFile(replayDir + f.Name())
		if err != nil {
			t.Error(err)
		}
	}
}
