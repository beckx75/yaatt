package yaatt

import (
	"encoding/csv"
	"fmt"
	"os"

	"beckx.online/butils/fileutils"
	"github.com/rs/zerolog/log"
)

type TagMap struct {
	YaatToId323  map[string]string
	Id323ToYatt  map[string]string
	YaatToVorbis map[string]string
	VorbisToYaat map[string]string
}

func newTagMap(fp string) (*TagMap, error) {
	// TODO check csv-cols; check csv rows, should be nine
	f, err := os.Open(fp)
	if err != nil {
		return nil, err
	}
	c := csv.NewReader(f)
	c.Comma = ';'
	rec, err := c.ReadAll()
	if err != nil {
		return nil, err
	}
	if len(rec[0]) != 3 {
		return nil, fmt.Errorf("malformed tagmap: search for 3 columns, got %d", len(rec[0]))
	}
	tm := &TagMap{
		YaatToId323:  make(map[string]string),
		Id323ToYatt:  make(map[string]string),
		YaatToVorbis: make(map[string]string),
		VorbisToYaat: make(map[string]string),
	}
	for i := 1; i < len(rec); i++ {
		tm.YaatToId323[rec[i][0]] = rec[i][1]
		tm.Id323ToYatt[rec[i][1]] = rec[i][0]
		tm.YaatToVorbis[rec[i][0]] = rec[i][2]
		tm.VorbisToYaat[rec[i][2]] = rec[i][0]
	}
	return tm, nil
}

type YaattData struct {
	Tagmap *TagMap
}

func NewYaattData(args []string, confpath string) (*YaattData, error) {

	yd := &YaattData{}

	return yd, nil
}

func GetAudiofiles(args []string, pattern []string) ([]string, error) {
	log.Info().Msg("looking for some files....")
	_, files, err := fileutils.GetFiles(args, pattern)
	if err != nil {
		return nil, err
	}
	return files, nil
}
