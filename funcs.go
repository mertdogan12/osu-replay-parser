package parser

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/mertdogan12/leb128"
)

func convertFirstString(data []byte) (string, []byte, error) {
	if data[0] == 0x0b {
		dataLenght, n, err := leb128.DecodeUnsigned(bytes.NewReader(data[1:]))
		dataLen := dataLenght.Uint64()
		if err != nil {
			return "", make([]byte, 0), err
		}
		return string(data[1+n : dataLen+2]), data[1+uint64(n)+dataLen:], nil
	} else {
		return "", data, nil
	}
}

func streamToString(stream io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.String()
}

func convertReplayString(input string) ([]ReplayData, error) {
	if input == "" {
		fmt.Println("LULLLUL")
	}
	elements := strings.Split(input, ",")
	replayData := make([]ReplayData, len(elements))

	for i, element := range elements {
		rawData := strings.Split(element, "|")
		if len(rawData) != 4 {
			continue
		}

		w, err := strconv.ParseUint(rawData[0], 10, 64)
		x, err := strconv.ParseFloat(rawData[1], 32)
		y, err := strconv.ParseFloat(rawData[2], 32)
		z, err := strconv.ParseUint(rawData[3], 10, 32)

		if err != nil {
			return nil, err
		}

		replayData[i] = ReplayData{w, float32(x), float32(y), uint32(z)}
	}

	return replayData, nil
}
