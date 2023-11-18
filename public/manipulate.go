package public

import (
	"data-optimize/utils"
	"fmt"
	"strings"
)

func Encode(input string, dataMap *utils.BidirectionalMap) []byte {
	inputSplit := strings.Split(input, " ")
	for _, word := range inputSplit {
		val, err := (*dataMap).BackwardMap[strings.ToLower(word)]
		if !err {
			continue
		}
		fmt.Println(word, val, err)
		byteArray := utils.BinaryToByteArray(utils.ToBinary(val))
		input = strings.ReplaceAll(input, word, string(byteArray))
	}
	return []byte(input)
}

func Decode(input []byte, dataMap *utils.BidirectionalMap) []byte {
	//var filtered []string
	return []byte{}
}
