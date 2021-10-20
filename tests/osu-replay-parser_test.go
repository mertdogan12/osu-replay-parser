package tests

import (
	"encoding/json"
	"os"
	"testing"

	parser "github.com/mertdogan-org/osu-replay-converter/pkg/osu-replay-parser"
)

func TestOsuReplayParser(t *testing.T) {
	filePath := "../test-replays/test1.osr"
	replay, err := parser.ConvertToObject(filePath)
	if err != nil {
		t.Error(err)
		return
	}

	json, err := json.MarshalIndent(replay, "\t", "\t")
	if err != nil {
		t.Error(err)
		return
	}

	os.WriteFile("out.json", json, 0644)
}
