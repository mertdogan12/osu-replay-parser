package parser

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"ekyu.moe/leb128"
	"github.com/ulikunitz/xz/lzma"
)

func ConvertToObject(filePath string) (*OsrObject, error) {
	data, err := os.ReadFile(filePath)
	var osrObject OsrObject

	if err != nil {
		return nil, err
	}

	// Gametype
	switch data[0] {
	case 0:
		osrObject.Gametype = "osu! Standard"
	case 1:
		osrObject.Gametype = "Taiko"
	case 2:
		osrObject.Gametype = "Catch the Beat"
	case 3:
		osrObject.Gametype = "osu!mania"
	default:
		err = errors.New("Osu gametype is not identifiable")
	}
	data = data[1:]

	// Version, BeatmapHash, PlayerName, ReplayHash
	osrObject.Version = binary.LittleEndian.Uint32(data[:4])
	osrObject.BeatmapHash, data = convertFirstString(data[4:])
	osrObject.PlayerName, data = convertFirstString(data)
	osrObject.ReplayHash, data = convertFirstString(data)

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
	lifeBarString, data := convertFirstString(data)
	if lifeBarString != "" {
		lifeBarArray := strings.Split(lifeBarString, ",")
		osrObject.Lifebar = make([][]float32, len(lifeBarArray))
		for i, element := range lifeBarArray {
			elements := strings.Split(element, "|")
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
	dataLenght := binary.LittleEndian.Uint32(data[:4])
	compressedData := data[4 : dataLenght+4]
	data = data[dataLenght+4:]
	os.WriteFile("out.lzma", compressedData, 0644)

	r, err := lzma.NewReader(bytes.NewReader(compressedData))
	fmt.Println(streamToString(r))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Online Score Id
	osrObject.OnlineScoreId = binary.LittleEndian.Uint64(data[:8])

	return &osrObject, err
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

func streamToString(stream io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.String()
}
