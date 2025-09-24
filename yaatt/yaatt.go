package yaatt

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"beckx.online/butils/fileutils"
	"github.com/rs/zerolog/log"
)

type TagMap struct {
	YaatToId323  map[string]string
	Id323ToYatt  map[string]string
	YaatToVorbis map[string]string
	VorbisToYaat map[string]string
}

type YaattData struct {
	Tagmap *TagMap

	Files     []string
	MetaDatas map[string]*MetaData
}

func NewYaattData(args []string, confpath string) (*YaattData, error) {
	confpath = confpath + "/tagdef.csv"
	tm, err := newTagMap(confpath)
	if err != nil {
		return nil, err
	}
	yd := &YaattData{
		Tagmap:    tm,
		MetaDatas: make(map[string]*MetaData),
	}

	yd.Files, err = GetAudiofiles(args, []string{".mp3", ".flac"})
	if err != nil {
		return yd, err
	}

	var errs error
	for i, fp := range yd.Files {
		md, err := ReadMetadata(fp, *yd.Tagmap)
		if err != nil {
			errs = errors.Join(errs, err)
		} else {
			log.Debug().Msgf("Read file", i+1, fp)
			yd.MetaDatas[fp] = md
		}
	}
	return yd, errs
}

func newTagMap(fp string) (*TagMap, error) {
	// TODO check csv-cols; check csv rows, should be nine
	f, err := os.Open(fp)
	if err != nil {
		return nil, err
	}
	defer f.Close()
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

func GetAudiofiles(args []string, pattern []string) ([]string, error) {
	log.Info().Msg("looking for some files....")
	_, files, err := fileutils.GetFiles(args, pattern)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func (yd YaattData) PrintMetadata() string {
	txt := ""
	for _, fp := range yd.Files {
		txt = txt + filepath.Base(fp) + " Metadata:\n"
		md, ok := yd.MetaDatas[fp]
		if !ok {
			continue
		}
		for _, k := range md.TextTagIndex {
			texttags, ok := md.TextTags[k]
			if !ok {
				txt = txt + "\t" + k + ":\n"
			} else {
				for _, texttag := range texttags {
					txt = txt + "\t" + k + ": " + texttag.Value + "\n"
				}
			}
		}
	}
	return txt
}
