package utils

import (
	"os"
	"strconv"
	"strings"
)

type BidirectionalMap struct {
	ForwardMap  map[uint32]string
	BackwardMap map[string]uint32
}

func LoadDict() BidirectionalMap {
	mapInstance := BidirectionalMap{
		ForwardMap:  make(map[uint32]string),
		BackwardMap: make(map[string]uint32),
	}
	fileContents, err := os.ReadFile("dict.txt")
	Check(err)
	contentsSplit := strings.Split(string(fileContents), "\n")
	for _, entry := range contentsSplit {
		numberWordSplit := strings.Split(entry, ":")
		numConverted, _ := strconv.Atoi(numberWordSplit[0])
		converted32 := uint32(numConverted)
		mapInstance.ForwardMap[converted32] = numberWordSplit[1]
		mapInstance.BackwardMap[numberWordSplit[1]] = converted32
	}
	return mapInstance
}
