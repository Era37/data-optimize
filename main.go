package main

import (
	"data-optimize/public"
	"data-optimize/utils"
	"os"
)

func main() {
	biMap := utils.LoadDict()
	data, _ := os.ReadFile("./IMDB Dataset.txt")
	public.Encode(string(data), &biMap)
	dec := public.Decode(string(data), &biMap)
	f, _ := os.Create("test imdb2.txt")
	f.WriteString(string(dec))
	/*myInt := uint32(1434534342) // You can replace this with the desired 32-bit integer
	decimalValue := utils.BinaryToByteArray(utils.ToBinary(myInt))

	str := []byte("I love you mommyjsndrkjn ")
	str = append(str, decimalValue...)

	fmt.Println(str)
	fmt.Println(string(str))*/
}
