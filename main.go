package main

import (
	"data-optimize/public"
	"data-optimize/utils"
	"fmt"
)

func main() {
	biMap := utils.LoadDict()
	data := "I love you mommy"
	//fmt.Println([]byte(data))
	encoded := public.Encode(data, &biMap)
	fmt.Println(len(encoded), len([]byte(data)))
	/*myInt := uint32(1434534342) // You can replace this with the desired 32-bit integer
	decimalValue := utils.BinaryToByteArray(utils.ToBinary(myInt))

	str := []byte("I love you mommyjsndrkjn ")
	str = append(str, decimalValue...)

	fmt.Println(str)
	fmt.Println(string(str))*/
}
