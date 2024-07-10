package main

import (
	"fmt"
	"strconv"
)

func byteToint(b rune) int{
	res, _ := strconv.Atoi(string(b))
	return res
}
func thaiid(id string) bool {
	sum := 0
	for i, v := range id[:12] {
		value := byteToint(v)
		sum += (13 - i) * value
	}
	return (11-sum%11)%10 == byteToint(rune(id[12]))
}

func main() {
	fmt.Print(thaiid("1234567890121"))
}
