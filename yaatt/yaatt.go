package yaatt

import (
	"beckx.online/butils/fileutils"
)

type Metadata struct {
	Tags []*Tag
}

func MakeYaattTag(n string, tagtaype string) string {

}

func yaattTagFromFlac(n string) string {
	nn := ""
	switch n {
	case "TITLE":
		nn = "diddle"
	default:
		nn = n
	}
	return nn
}

type Tag struct {
	Name  string
	Value string
	Enc   uint8
	TType string
}

type YaattData struct {
}

func GetAudiofiles(args []string, pattern []string) ([]string, error) {
	_, files, err := fileutils.GetFiles(args, pattern)
	if err != nil {
		return nil, err
	}
	return files, nil
}
