package util

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// ArrayToString 配列を文字列に変換する。
// 配列の要素中の値は絶対値をとった状態で置き換えられる
func ArrayToString(a []int) string {
	for i, v := range a {
		a[i] = int(math.Abs(float64(v)))
	}
	s := fmt.Sprint(a)
	s = strings.Replace(s, " ", "", -1)
	return strings.Trim(s, "[]")
}

// ArrayToInt 配列を数値に変換する。
// 配列の要素中の値は絶対値をとった状態で置き換えられる
func ArrayToInt(a []int) int {
	n := 0.0
	for i, v := range a {
		digit := len(a) - i - 1
		n += math.Abs(float64(v)) * math.Pow(10.0, float64(digit))
	}
	return int(n)
}

// IntToArray 与えられた数値を長さ length の配列に変換する。
// length<0 の場合、数値の桁の大きさと等しい大きさの配列を作成する。
// n<0の場合、絶対値をとった状態で配列を作成する。
func IntToArray(n int, length int) []int {
	n = int(math.Abs(float64(n)))

	if length < 0 {
		nStr := strconv.Itoa(n)
		length = len(nStr)
	}

	a := make([]int, length)
	for i := 0; i < length; i++ {
		digit := length - i - 1
		digit = int(math.Pow(10.0, float64(digit)))
		v := n / digit
		a[i] = v

		if n > digit {
			n -= v * digit
		}
	}
	return a
}
