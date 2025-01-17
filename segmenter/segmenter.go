package segmenter

import (
    "github.com/wangbin/jiebago/posseg"
)

const (
    DEFAULT_DICT_PATH string = "conf/dict.txt"
    USER_DICT_PATH    string = "conf/userdict.txt"
)

var (
    gSeg posseg.Segmenter
)

func init() {
	err := gSeg.LoadDictionary(DEFAULT_DICT_PATH)
	if err != nil {
		panic(err)
	}
	err = gSeg.LoadUserDictionary(USER_DICT_PATH)
	if err != nil {
		panic(err)
	}
}

func GetSegmenter() *posseg.Segmenter {
    return &gSeg
}
