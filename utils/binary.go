package utils

import (
	"strconv"
)

func ToBinary(num uint32) []string {

	var binarySlice []string
	var localByteGroup string
	for i := 31; i >= 0; i-- {
		k := num >> i
		bit := "0"
		if k&1 == 1 {
			bit = "1"
		}
		localByteGroup += bit
		if len(localByteGroup) == 8 || i == 0 {
			binarySlice = append(binarySlice, localByteGroup)
			localByteGroup = ""
		}
	}
	for _, binSlice := range binarySlice {
		if binSlice == "00000000" {
			binarySlice = binarySlice[1:]
		}
	}
	return binarySlice
}

func BinaryToByteArray(bin []string) []byte {
	var chars []byte
	for _, element := range bin {
		convertValue, _ := strconv.ParseInt(element, 2, 32)
		chars = append(chars, byte(convertValue))
	}
	return chars
}
