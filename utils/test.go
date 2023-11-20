package utils

import (
	"fmt"
	"os"
	"strings"
)

func arrHasInvalid(arr []byte) bool {
	flag := false
	for _, el := range arr {
		if el < 33 {
			flag = true
		}
	}
	return flag
}

func GenerateWords() {
	var final []string
	s, _ := os.ReadFile("./old_words.txt")
	sSplit := strings.Split(string(s), "\n")
	for _, word := range sSplit {

		if len(word) > 4 {
			final = append(final, word)
			continue
		}
		fmt.Println(word)
	}
	f, _ := os.Create("./words.txt")
	f.WriteString(strings.Join(final, "\n"))
}

func Sort() {
	var final []string
	var goodNumbers []uint32
	for start := 210_000; start < 6_000_000; start++ {
		if len(ToBinary(uint32(start))) != 3 || arrHasInvalid(BinaryToByteArray(ToBinary(uint32(start)))) {
			//fmt.Println(start, BinaryToByteArray(ToBinary(uint32(start))))
			continue
		}
		goodNumbers = append(goodNumbers, uint32(start))
	}
	file, _ := os.ReadFile("./words.txt")
	fileSplit := strings.Split(string(file), "\n")
	fmt.Println(fileSplit)
	for _, word := range fileSplit {
		final = append(final, fmt.Sprintf("%d:%s", goodNumbers[0], strings.ToLower(word)))
		goodNumbers = goodNumbers[1:]
	}
	f, _ := os.Create("dict.txt")
	f.WriteString(strings.Join(final, "\n"))
}
