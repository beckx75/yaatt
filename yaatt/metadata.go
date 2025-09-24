package yaatt

import (
	"encoding/binary"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/bogem/id3v2/v2"
	"github.com/go-flac/flacvorbis/v2"
	"github.com/go-flac/go-flac/v2"
	"github.com/rs/zerolog/log"
)

type TagType uint8

const (
	TT_UNDEF TagType = iota
	TT_NOTAG
	TT_ID3V23
	TT_VORBIS
)

type TextTag struct {
	Name    string
	Value   string
	OrgName string
	Enc     string
}

type MetaData struct {
	TagType      TagType
	TextTagIndex []string
	TextTags     map[string]*TextTag
}

func ReadMetadata(fp string, tm TagMap) (*MetaData, error) {
	md := &MetaData{
		TextTags:     make(map[string]*TextTag),
		TextTagIndex: []string{},
	}
	for k := range tm.YaatToVorbis {
		md.TextTagIndex = append(md.TextTagIndex, k)
	}
	sort.Strings(md.TextTagIndex)

	var err error
	md.TagType, err = getTagType(fp)
	if err != nil {
		return nil, err
	}
	switch md.TagType {
	case TT_VORBIS:
		err = md.readVorbisMetadata(fp, tm)
		if err != nil {
			return nil, err
		}
	case TT_ID3V23:
		err = md.readId3v2Metadata(fp, tm)
		if err != nil {
			return nil, err
		}
	case TT_NOTAG:
		return nil, fmt.Errorf("file is mp3-file and has no ID3v2x Tag: %s", fp)
	case TT_UNDEF:
		return nil, fmt.Errorf("could not read metadata because of unknown TagType in '%s': %v",
			fp, md.TagType)
	}
	if err != nil {
		panic(err)
	}
	return md, err
}

func (md *MetaData) readVorbisMetadata(fp string, tm TagMap) error {
	log.Debug().Msgf("reading Vorbis-Metadata from file %f", fp)
	f, err := flac.ParseFile(fp)
	if err != nil {
		return err
	}
	var vorbisblock *flacvorbis.MetaDataBlockVorbisComment
	for idx, meta := range f.Meta {
		log.Debug().Msgf("VorbisBlock-Index: %d", idx)
		if meta.Type == flac.VorbisComment {
			vorbisblock, _ = flacvorbis.ParseFromMetaDataBlock(*meta)
			for _, cmt := range vorbisblock.Comments {
				keyVal := strings.Split(cmt, "=")
				tt := &TextTag{
					OrgName: keyVal[0],
					Value:   keyVal[1],
				}
				name, ok := tm.VorbisToYaat[tt.OrgName]
				if ok {
					tt.Name = name
				} else {
					tt.Name = tt.OrgName
					md.TextTagIndex = append(md.TextTagIndex, tt.Name)
				}
				_, ok = md.TextTags[tt.Name]
				if ok {
					return fmt.Errorf("found non unuique TagName %s", tt.Name)
				} else {
					md.TextTags[tt.Name] = tt
				}
			}
		}
	}

	return nil
}

func (md *MetaData) readId3v2Metadata(fp string, tm TagMap) error {
	tag, err := id3v2.Open(fp, id3v2.Options{Parse: true})
	if err != nil {
		return err
	}
	defer tag.Close()

	framer := tag.AllFrames()
	for fn, frs := range framer {
		if fn == "TXXX" {
			log.Warn().Msg("TXXX Frames not supported yet :(")
		} else if fn[0] == 'T' {
			for _, fr := range frs {
				tf, ok := fr.(id3v2.TextFrame)
				if !ok {
					return fmt.Errorf("could not typecast frame to TextFrame: %s", fn)
				}
				tt := &TextTag{
					OrgName: fn,
					Value:   tf.Text,
					Enc:     tf.Encoding.Name,
				}
				name, ok := tm.Id323ToYatt[fn]
				if ok {
					tt.Name = name
				} else {
					tt.Name = fn
					md.TextTagIndex = append(md.TextTagIndex, tt.Name)
				}
				_, ok = md.TextTags[tt.Name]
				if ok {
					return fmt.Errorf("found non unique TextTag with name '%s'", tt.Name)
				} else {
					md.TextTags[tt.Name] = tt
				}
			}
		} else {
			log.Warn().Msgf("unsupported ID3v2-Frame '%s' in '%s'", fn, fp)
		}
	}
	return nil
}

func getTagType(fp string) (TagType, error) {
	file, err := os.OpenFile(fp, os.O_RDONLY, 0644)
	if err != nil {
		return TT_UNDEF, err
	}
	defer file.Close()
	var b byte
	// first 3 bytes of audiofile (ID3, fLa)
	var fileIdentifier string
	var bs [3]uint8
	for i := 0; i < 3; i++ {
		err = binary.Read(file, binary.BigEndian, &b)
		if err != nil {
			return TT_UNDEF, fmt.Errorf("could not get TagType %v", err)
		}
		bs[i] = b
		fileIdentifier = fmt.Sprintf("%s%s", fileIdentifier, string(b))
	}
	switch fileIdentifier {
	case "ID3":
		return TT_ID3V23, nil
	case "fLa":
		return TT_VORBIS, nil
	default:
		if bs[0] == 0xff && bs[1] == 0xfb {
			return TT_NOTAG, nil
		} else {
			return TT_UNDEF, nil
		}
	}
}
