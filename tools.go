package abcktools

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

func Jrs(t int) ([]int, error) {
	randomFactor := int(math.Floor(100000*rand.Float64() + 10000))
	multiplied := strconv.Itoa(t * randomFactor)
	isLong := len(multiplied) >= 18
	index := 0
	floatArr := make([]float64, 0, 6)
	for len(floatArr) < 6 {
		num, err := strconv.Atoi(multiplied[index:(index + 2)])
		if err != nil {
			return nil, err
		}
		floatArr = append(floatArr, float64(num))
		if isLong {
			index = index + 3
		} else {
			index = index + 2
		}
	}
	dis, err := CalcDis(floatArr)
	if err != nil {
		return nil, err
	}
	return []int{randomFactor, dis}, nil
}

func CalcDis(floatArr []float64) (int, error) {
	if len(floatArr) < 6 {
		return 0, fmt.Errorf("input list should contain at least 6 elements")
	}
	a := floatArr[0] - floatArr[1]
	e := floatArr[2] - floatArr[3]
	n := floatArr[4] - floatArr[5]
	return int(math.Sqrt(a*a + e*e + n*n)), nil
}

func Ab(input string) int {
	sum := 0
	for _, n := range []rune(input) {
		if n < 128 {
			sum += int(n)
		}
	}
	return sum
}

func GenerateSeparator(inputStrings []string, number int) string {
	alphabet := "abcdefghijklmnopaqrstuvxyzABCDEFGHIJKLMNOPAQRSTUVXYZ!@#%&-_=;:<>,~"
	stringSet := make(map[string]struct{}, len(inputStrings))
	for _, str := range inputStrings {
		stringSet[str] = struct{}{}
	}

	var builder strings.Builder
	if number >= 0 && number <= 9 {
		for i := 0; i <= 9; i++ {
			if i != number {
				builder.WriteString(strings.Repeat(strconv.Itoa(i), 20))
			}
		}
		alphabet += builder.String()
	}

	for {
		builder.Reset()
		builder.WriteRune(',')
		for i := 0; i < rand.Intn(3)+3; i++ {
			for j := 0; j < rand.Intn(3)+3; j++ {
				builder.WriteRune(rune(alphabet[rand.Intn(len(alphabet))]))
			}
			builder.WriteRune(',')
		}
		randomStr := builder.String()
		if _, found := stringSet[randomStr]; !found {
			return randomStr
		}
	}
}

func Od(t string, a string) string {
	runes := []rune(t)
	length := len(a)
	if length == 0 {
		return t
	}

	result := make([]string, len(runes))
	for index, char := range runes {
		subChar := a[index%length]
		char = rune(Rir(int(char), 47, 57, int(subChar)))
		result[index] = string(char)
	}

	return strings.Join(result, "")
}

func Rir(t int, a int, e int, n int) int {
	if t > a && t <= e {
		t += n % (e - a)
		if t > e {
			t = t - e + a
		}
	}
	return t
}
