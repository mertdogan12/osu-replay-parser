package parser

type OsrObject struct {
	// Gamemode, Game -& Replayinformations
	Gametype    string
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
	// Online Score Id
	OnlineScoreId uint64
}
