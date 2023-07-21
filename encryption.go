package abcktools

import "strings"

var Alphabet = " !#$%&()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[]^_`abcdefghijklmnopqrstuvwxyz{|}~"
var AlphabetIndex = []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 0, 1, -1, 2, 3, 4, 5, -1, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, -1, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91}
var SeedConstants = []int{1061, 300000, 100000, 10000, 8888888, 7777777, 999999, 3600000, 65535, 65793, 4294967295, 4282663, 8388607, 4064256, 3000}

func Encrypt(sensor string, keys []int) string {
	shuffledSensor := ShuffleString(sensor, keys[0])
	return EncryptString(shuffledSensor, keys[1])
}

func ShuffleString(input string, seed int) string {
	inputArray := strings.Split(input, ",")
	length := len(inputArray)

	for i := 0; i < length; i++ {
		seed = updateSeed(seed)
		randomIndex := seed % length

		seed = updateSeed(seed)
		temporaryValue := seed % length

		inputArray[randomIndex], inputArray[temporaryValue] = inputArray[temporaryValue], inputArray[randomIndex]
	}

	return strings.Join(inputArray, ",")
}

func updateSeed(seed int) int {
	seed *= SeedConstants[9]
	seed &= SeedConstants[10]
	seed += SeedConstants[11]
	seed &= SeedConstants[12]
	return (seed >> 8) & SeedConstants[8]
}

func EncryptString(input string, key int) string {
	var currentChar, encodedCharIndex int
	outputText := make([]rune, len(input))

	for currentChar = 0; currentChar < len(input); currentChar++ {
		seed := (key >> 8) & SeedConstants[8]
		key *= SeedConstants[9]
		key &= SeedConstants[10]
		key += SeedConstants[11]
		key &= SeedConstants[12]

		encodedCharIndex = AlphabetIndex[int(input[currentChar])]
		if encodedCharIndex >= 0 {
			encodedCharIndex += seed % len(Alphabet)
			encodedCharIndex %= len(Alphabet)
			outputText[currentChar] = rune(Alphabet[encodedCharIndex])
		} else {
			outputText[currentChar] = rune(input[currentChar])
		}
	}

	return string(outputText)
}
