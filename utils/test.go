package utils

import (
	"fmt"
	"os"
	"strings"
)

func arrHas32(arr []byte) bool {
	flag := false
	for _, el := range arr {
		if el == 32 {
			flag = true
		}
	}
	return flag
}

func Sort() {
	var final []string
	var goodNumbers []uint32
	for start := 210_000; start < 1_000_000; start++ {
		if len(ToBinary(uint32(start))) != 3 || arrHas32(BinaryToByteArray(ToBinary(uint32(start)))) {
			//fmt.Println(start, BinaryToByteArray(ToBinary(uint32(start))))
			continue
		}
		goodNumbers = append(goodNumbers, uint32(start))
	}
	file, _ := os.ReadFile("./words.txt")
	fileSplit := strings.Split(string(file), "\n")
	fmt.Println(fileSplit)
	for _, word := range fileSplit {
		final = append(final, fmt.Sprintf("%d:%s", goodNumbers[0], word))
		goodNumbers = goodNumbers[1:]
	}
	f, _ := os.Create("dict.txt")
	f.WriteString(strings.Join(final, "\n"))
}
