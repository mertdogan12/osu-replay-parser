package tests

import (
	"encoding/json"
	"testing"

	parser "github.com/mertdogan12/osu-replay-converter/pkg/osu-replay-parser"
)

func osuReplayParser(t *testing.T) {
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

	t.Log(string(json))
}
