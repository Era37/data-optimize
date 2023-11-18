package public

import (
	"data-optimize/utils"
	"strings"

	"encoding/binary"
)

func Encode(input string, dataMap *utils.BidirectionalMap) []byte {
	inputSplit := strings.Split(input, " ")
	for _, word := range inputSplit {
		val, err := (*dataMap).BackwardMap[strings.ToLower(word)]
		if !err {
			continue
		}
		byteArray := utils.BinaryToByteArray(utils.ToBinary(val))
		input = strings.ReplaceAll(input, word, string(byteArray))
	}
	return []byte(input)
}

func Decode(input string, dataMap *utils.BidirectionalMap) []byte {
	inputSplit := strings.Split(input, " ")
	for _, word := range inputSplit {
		if len(word) != 3 {
			continue
		}
		binary.BigEndian.Uint32([]byte(word))
	}
	return []byte{}
}
