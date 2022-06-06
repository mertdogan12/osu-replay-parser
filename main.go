package parser

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ulikunitz/xz/lzma"
)

func ParseFile(filePath string) (*OsrObject, error) {
	data, err := os.ReadFile(filePath)

	if err != nil {
		return nil, err
	}

	return Parse(data)
}

func Parse(data []byte) (*OsrObject, error) {
	var osrObject OsrObject

	osrObject.Gametype.id = int8(data[0])

	// Gametype
	switch data[0] {
	case 0:
		osrObject.Gametype.name = "osu! Standard"
	case 1:
		osrObject.Gametype.name = "Taiko"
	case 2:
		osrObject.Gametype.name = "Catch the Beat"
	case 3:
		osrObject.Gametype.name = "osu!mania"
	default:
		return nil, errors.New("Osu gametype is not identifiable")
	}
	data = data[1:]

	// Version, BeatmapHash, PlayerName, ReplayHash
	var err error
	osrObject.Version = binary.LittleEndian.Uint32(data[:4])
	osrObject.BeatmapHash, data, err = convertFirstString(data[4:])
	osrObject.PlayerName, data, err = convertFirstString(data)
	osrObject.ReplayHash, data, err = convertFirstString(data)
	if err != nil {
		return nil, err
	}

	// Scoreinformation, Mods
	osrObject.ThreeHunreds = binary.LittleEndian.Uint16(data[:2])
	osrObject.Hunreds = binary.LittleEndian.Uint16(data[2 : 2+2])
	osrObject.Fifths = binary.LittleEndian.Uint16(data[4 : 4+2])
	osrObject.Gekis = binary.LittleEndian.Uint16(data[6 : 6+2])
	osrObject.Katus = binary.LittleEndian.Uint16(data[8 : 8+2])
	osrObject.Misses = binary.LittleEndian.Uint16(data[10 : 10+2])
	osrObject.Score = binary.LittleEndian.Uint32(data[12 : 12+4])
	osrObject.Combo = binary.LittleEndian.Uint16(data[16 : 16+2])
	osrObject.FullCombo = (data[18] == 1)
	osrObject.Mods = binary.LittleEndian.Uint32(data[19 : 19+4])
	data = data[23:]

	// Lifebar
	lifeBarString, data, err := convertFirstString(data)
	if err != nil {
		return nil, err
	}
	if lifeBarString != "" {
		lifeBarArray := strings.Split(lifeBarString, ",")
		osrObject.Lifebar = make([][]float32, len(lifeBarArray))
		for i, element := range lifeBarArray {
			elements := strings.Split(element, "|")
			if len(elements) < 2 {
				continue
			}
			u, err := strconv.ParseFloat(elements[0], 32)
			v, err := strconv.ParseFloat(elements[1], 32)

			if err != nil {
				return nil, err
			}

			osrObject.Lifebar[i] = []float32{float32(u), float32(v)}
		}
	} else {
		osrObject.Lifebar = make([][]float32, 0)
	}

	// TimeStamp
	osrObject.TimeStamp = binary.LittleEndian.Uint64(data[:8])
	data = data[8:]

	// Replay Data
	// TODO check dataLenght --> error
	dataLenght := binary.LittleEndian.Uint32(data[:4])
	if dataLenght+4 >= uint32(len(data[4:])) {
		return nil, errors.New(fmt.Sprintf("Parsed replay data lenght is to high, %d, %#x %#x %#x %#x", dataLenght, data[0], data[1], data[2], data[3]))
	}
	compressedData := data[4 : dataLenght+4]
	data = data[dataLenght+4:]

	r, err := lzma.NewReader(bytes.NewReader(compressedData))
	osrObject.ReplayData, err = convertReplayString(streamToString(r))
	if err != nil {
		return nil, err
	}

	// Online Score Id
	osrObject.OnlineScoreId = binary.LittleEndian.Uint64(data[:8])

	return &osrObject, err
}
