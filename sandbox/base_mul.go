package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	ele1 := "29123841234812351239412"
	ele2 := "3496123842341123491234123"
	fmt.Printf("ele1 * ele2 = %s\n", mul(ele1, ele2, 9))
	//
	ele3 := "2912384123481235123941252362745232"
	ele4 := "34961238423411234912341237262725728683"
	fmt.Printf("ele3 * ele4 = %s\n", mul(ele3, ele4, 9))
}

func mul(ele1, ele2 string, b int) string {
	ele1BaseB := baseConv(ele1, b)
	ele2BaseB := baseConv(ele2, b)
	fmt.Printf("converted ele = %v\n", ele1BaseB)
	fmt.Printf("converted ele = %v\n", ele2BaseB)

	bot, up := 0, 0
	flag := int64(0)
	base := int64(math.Pow(10, float64(b)))
	result := make([]int64, len(ele1BaseB)+len(ele2BaseB))

	for i := len(ele1BaseB) - 1; i >= 0; i-- {
		for j := len(ele2BaseB) - 1; j >= 0; j-- {
			mulUnit := ele1BaseB[i]*ele2BaseB[j] + flag + result[up]
			flag = 0
			if mulUnit >= base {
				flag = mulUnit / base
				mulUnit = mulUnit % base
			}
			result[up] = mulUnit
			if j == 0 && flag != 0 {
				result[up+1] = flag
				flag = 0
			}
			up += 1
		}
		bot += 1
		up = bot
	}

	return convertBack(result, b)
}

func convertBack(num []int64, base int) string {
	var result string
	fmt.Println(num)
	for i := len(num) - 1; i >= 0; i-- {
		curUnit := fmt.Sprintf("%d", num[i])
		if i == len(num)-1 {
			result = result + curUnit
		} else {
			result = result + addZero(curUnit, base)
		}
	}
	return result
}

func addZero(num string, base int) string {
	if len(num) > base {
		panic("error unit bigger than base")
	}
	if len(num) == base {
		return num
	}
	for len(num) < base {
		num = "0" + num
	}
	return num
}

func baseConv(num string, base int) []int64 {
	result := make([]int64, 0)
	remainder := len(num) % base
	if remainder > 0 {
		result = append(result, parseInt64(num[:remainder]))
		num = num[remainder:]
	}
	for len(num) > 0 {
		result = append(result, parseInt64(num[:base]))
		num = num[base:]
	}
	return result
}

func parseInt64(num string) int64 {
	out, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		panic(err.Error())
	}
	return out
}
