package public

import (
	"data-optimize/utils"
	"fmt"
	"strings"

	"encoding/binary"
)

func Encode(input string, dataMap *utils.BidirectionalMap) []byte {
	inputSplit := strings.Split(input, " ")
	fmt.Println([]byte(input))
	for _, word := range inputSplit {
		val, err := (*dataMap).BackwardMap[strings.ToLower(word)]
		if !err {
			continue
		}
		fmt.Println(val)
		byteArray := utils.BinaryToByteArray(utils.ToBinary(val))
		input = strings.ReplaceAll(input, word, string(byteArray))
	}
	return []byte(input)
}

func Decode(input string, dataMap *utils.BidirectionalMap) []byte {
	fmt.Println(input)
	inputSplit := strings.Split(input, " ")
	for _, word := range inputSplit {
		if len(word) != 3 {
			continue
		}
		wordByte := append([]byte{0}, []byte(word)...)
		integer := binary.BigEndian.Uint32(wordByte)
		fmt.Println(integer, wordByte)
	}
	return []byte{}
}
