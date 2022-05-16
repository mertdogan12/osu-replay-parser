# osu-replay-parser

Parses [.osr](https://osu.ppy.sh/wiki/en/Client/File_formats/Osr_%28file_format%29) files

## Use it

```bash
go get github.com/mertdogan12/osu-replay-parser@master
```

## Code example

```bash
package main

import (
	"os"

	parser "github.com/mertdogan12/osu-replay-parser"
)

func main() {
	// Read data from replay file
	data, err := os.ReadFile("path to file")
	if err != nil {
		panic(err)
	}

	// Parse the data
	parser.Parse(data)

	// Parse the data direct from a file
	parser.ParseFile("path to file")
}
```
