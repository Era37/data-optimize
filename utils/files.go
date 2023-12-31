package utils

import (
	"os"
	"strconv"
	"strings"
)

func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}

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
		entry = entry[:len(entry)-1]
		numberWordSplit := strings.Split(entry, ":")
		numConverted, _ := strconv.Atoi(numberWordSplit[0])
		converted32 := uint32(numConverted)
		mapInstance.ForwardMap[converted32] = numberWordSplit[1]
		mapInstance.BackwardMap[numberWordSplit[1]] = converted32
	}
	return mapInstance
}
