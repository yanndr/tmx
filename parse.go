package tmx

import (
	"encoding/xml"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
)

// TMXURL is the URL to your TMX file. If it uses external files, the sources
// given are relative to the location of the TMX file. This should be set if
// you use external tilesets.
var TMXURL string

var FileSystem fs.FS

// Parse returns the Map encoded in the reader
func Parse(r io.Reader) (Map, error) {
	if FileSystem == nil {
		FileSystem = os.DirFS("")
	}
	var m Map
	d, err := ioutil.ReadAll(r)
	if err != nil {
		return m, err
	}
	err = xml.Unmarshal(d, &m)
	return m, err
}
