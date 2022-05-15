package parser

type OsrObject struct {
	// Gamemode, Game -& Replayinformations
	Gametype    Gametype
	Version     uint32
	BeatmapHash string
	PlayerName  string
	ReplayHash  string
	// Scoreinformations
	ThreeHunreds uint16
	Hunreds      uint16
	Fifths       uint16
	Gekis        uint16
	Katus        uint16
	Misses       uint16
	Score        uint32
	Combo        uint16
	FullCombo    bool
	Mods         uint32
	// Lifebar | 0 --> time in mil, 1 --> amount of life (0 - 1)
	Lifebar [][]float32
	// Time stamp
	TimeStamp uint64
	// Replay Data
	ReplayData []ReplayData
	// Online Score Id
	OnlineScoreId uint64
}

type ReplayData struct {
	W uint64
	X float32
	Y float32
	Z uint32
}

type Gametype struct {
	id   int8
	name string
}
