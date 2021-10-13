package parser

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"ekyu.moe/leb128"
)

func ConvertToJson(filePath string) (*OsrObject, error) {
	data, err := os.ReadFile(filePath)
	var osrJson OsrObject

	if err != nil {
		return nil, err
	}

	// Gametype
	switch data[0] {
	case 0:
		osrJson.Gametype = "osu! Standard"
	case 1:
		osrJson.Gametype = "Taiko"
	case 2:
		osrJson.Gametype = "Catch the Beat"
	case 3:
		osrJson.Gametype = "osu!mania"
	default:
		err = errors.New("Osu gametype is not identifiable")
	}
	data = data[1:]

	// Version, BeatmapHash, PlayerName, ReplayHash
	osrJson.Version = binary.LittleEndian.Uint32(data[:4])
	osrJson.BeatmapHash, data = convertFirstString(data[4:])
	osrJson.PlayerName, data = convertFirstString(data)
	osrJson.ReplayHash, data = convertFirstString(data)

	// Scoreinformation, Mods
	osrJson.ThreeHunreds = binary.LittleEndian.Uint16(data[:2])
	osrJson.Hunreds = binary.LittleEndian.Uint16(data[2 : 2+2])
	osrJson.Fifths = binary.LittleEndian.Uint16(data[4 : 4+2])
	osrJson.Gekis = binary.LittleEndian.Uint16(data[6 : 6+2])
	osrJson.Katus = binary.LittleEndian.Uint16(data[8 : 8+2])
	osrJson.Misses = binary.LittleEndian.Uint16(data[10 : 10+2])
	osrJson.Score = binary.LittleEndian.Uint32(data[12 : 12+4])
	osrJson.Combo = binary.LittleEndian.Uint16(data[16 : 16+2])
	osrJson.FullCombo = (data[18] == 1)
	osrJson.Mods = binary.LittleEndian.Uint32(data[19 : 19+4])
	data = data[23:]

	// Lifebar
	lifeBarString, data := convertFirstString(data)
	lifeBarArray := strings.Split(lifeBarString, ",")
	osrJson.Lifebar = make([][]float32, len(lifeBarArray))
	for i, element := range lifeBarArray {
		elements := strings.Split(element, "|")
		u, err := strconv.ParseFloat(elements[0], 32)
		v, err := strconv.ParseFloat(elements[1], 32)

		if err != nil {
			return nil, err
		}

		osrJson.Lifebar[i] = []float32{float32(u), float32(v)}
	}

	// TimeStamp
	osrJson.TimeStamp = binary.LittleEndian.Uint64(data[:8])
	data = data[8:]

	// Replay Data
	dataLenght := binary.LittleEndian.Uint32(data[:4])
	compressedData := data[4 : dataLenght+4]
	data = data[dataLenght+4:]
	os.WriteFile("out.lzma", compressedData, 0644)
	// TODO decompress replay data

	// Online Score Id
	osrJson.OnlineScoreId = binary.LittleEndian.Uint64(data[:8])

	return &osrJson, err
}

func convertFirstString(data []byte) (string, []byte) {
	if data[0] == 0x0b {
		dataLenght, n := leb128.DecodeUleb128(data[1:])
		return string(data[1+n : dataLenght+2]), data[1+uint64(n)+dataLenght:]
	} else {
		fmt.Println(data[0])
		return "", data
	}
}
