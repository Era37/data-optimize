package public

import (
	"data-optimize/utils"
	"fmt"
	"strings"

	"encoding/binary"
)

func Encode(input string, dataMap *utils.BidirectionalMap) []byte {
	inputSplit := strings.Split(input, " ")
	for _, word := range inputSplit {
		val, err := (*dataMap).BackwardMap[word]
		if !err {
			continue
		}
		byteArray := utils.BinaryToByteArray(utils.ToBinary(val))
		input = strings.ReplaceAll(input, word, string(byteArray))
	}
	return []byte(input)
}

func Decode(input string, dataMap *utils.BidirectionalMap) []byte {
	//var letterTracker int
	inputSplit := strings.Split(input, " ")
	for _, word := range inputSplit {
		if len(word) != 3 {
			continue
		}
		wordByte := append([]byte{0}, []byte(word)...)
		integer := binary.BigEndian.Uint32(wordByte)
		wordConverted, err := (*dataMap).ForwardMap[integer]
		if !err {
			continue
		}
		fmt.Println(wordConverted, err)
		input = strings.ReplaceAll(input, word, wordConverted)
		//letterTracker += len(word)
	}
	return []byte(input)
}
