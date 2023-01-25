package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	h0 = "0x6a09e667"
	h1 = "0xbb67ae85"
	h2 = "0x3c6ef372"
	h3 = "0xa54ff53a"
	h4 = "0x510e527f"
	h5 = "0x9b05688c"
	h6 = "0x1f83d9ab"
	h7 = "0x5be0cd19"
)

var _K = []string{
	"0x428a2f98",
	"0x71374491",
	"0xb5c0fbcf",
	"0xe9b5dba5",
	"0x3956c25b",
	"0x59f111f1",
	"0x923f82a4",
	"0xab1c5ed5",
	"0xd807aa98",
	"0x12835b01",
	"0x243185be",
	"0x550c7dc3",
	"0x72be5d74",
	"0x80deb1fe",
	"0x9bdc06a7",
	"0xc19bf174",
	"0xe49b69c1",
	"0xefbe4786",
	"0x0fc19dc6",
	"0x240ca1cc",
	"0x2de92c6f",
	"0x4a7484aa",
	"0x5cb0a9dc",
	"0x76f988da",
	"0x983e5152",
	"0xa831c66d",
	"0xb00327c8",
	"0xbf597fc7",
	"0xc6e00bf3",
	"0xd5a79147",
	"0x06ca6351",
	"0x14292967",
	"0x27b70a85",
	"0x2e1b2138",
	"0x4d2c6dfc",
	"0x53380d13",
	"0x650a7354",
	"0x766a0abb",
	"0x81c2c92e",
	"0x92722c85",
	"0xa2bfe8a1",
	"0xa81a664b",
	"0xc24b8b70",
	"0xc76c51a3",
	"0xd192e819",
	"0xd6990624",
	"0xf40e3585",
	"0x106aa070",
	"0x19a4c116",
	"0x1e376c08",
	"0x2748774c",
	"0x34b0bcb5",
	"0x391c0cb3",
	"0x4ed8aa4a",
	"0x5b9cca4f",
	"0x682e6ff3",
	"0x748f82ee",
	"0x78a5636f",
	"0x84c87814",
	"0x8cc70208",
	"0x90befffa",
	"0xa4506ceb",
	"0xbef9a3f7",
	"0xc67178f2",
}

func BinaryToHex(bin string) string {
	hex, err := strconv.ParseUint(bin, 2, 64)
	if err != nil {
		panic("Unable to convert binary to hexa")
	}
	return fmt.Sprintf("%x", hex)
}

func HexToBinary(hex string) string {
	// hexStr := strconv.Itoa(hex)
	binStr := ""
	BCD := map[uint8]string{
		'0': "0000",
		'1': "0001",
		'2': "0010",
		'3': "0011",
		'4': "0100",
		'5': "0101",
		'6': "0110",
		'7': "0111",
		'8': "1000",
		'9': "1001",
		'a': "1010",
		'b': "1011",
		'c': "1100",
		'd': "1101",
		'e': "1110",
		'f': "1111",
	}
	// fmt.Println("Hexa -> ", hex)
	for i := 2; i < len(hex); i++ {
		binStr += BCD[hex[i]]
	}
	return binStr
}

func ConvertMsgToBinary(msg string) string {
	binaryMsg := ""
	for _, c := range msg {
		binaryMsg = fmt.Sprintf("%s%.8b", binaryMsg, c)
	}
	return binaryMsg
}

func LeftRotate(s string, n int) string {
	temp := s + s
	l := len(s)
	return temp[n : l+n]
}

func RightRotate(s string, n int) string {
	return LeftRotate(s, len(s)-n)
}

func RightShift(s string, n int) string {
	str := ""
	str = s[0 : len(s)-n]
	z := ""
	for i := 0; i < n; i++ {
		z += "0"
	}
	str = z + str
	return str
}

func XorOperation(s1, s2 string) string {
	ans := ""
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			ans += "0"
		} else {
			ans += "1"
		}
	}
	return ans
}

func AndOperation(s1, s2 string) string {
	// fmt.Println(len(s1), len(s2))
	ans := ""
	for i := 0; i < len(s1); i++ {
		if s1[i] == '1' && s2[i] == '1' {
			ans += "1"
		} else {
			ans += "0"
		}
	}
	return ans
}

func NotOperation(s string) string {
	ans := ""
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			ans += "0"
		} else {
			ans += "1"
		}
	}
	return ans
}

func PreprocessingMsg(msg string) string {

	binaryMsg := ConvertMsgToBinary(msg)
	binaryMsg += "1"
	lengthOfBinaryMsg := len(binaryMsg)
	nextMultipleOf512 := 512 * (lengthOfBinaryMsg/512 + 1)
	requiredLength := nextMultipleOf512 - 64

	for i := lengthOfBinaryMsg; i < requiredLength; i++ {
		binaryMsg += "0"
	}

	lengthInBinary := strconv.FormatInt(int64(lengthOfBinaryMsg-1), 2)

	for i := len(lengthInBinary); i < 64; i++ {
		lengthInBinary = "0" + lengthInBinary
	}

	return binaryMsg + lengthInBinary
}

func MakingChunks(msg string) string {
	preProcessedMsg := PreprocessingMsg(msg)
	chunks := ""
	for i := 0; i < len(preProcessedMsg); i += 512 {
		chunks += preProcessedMsg[i:512]
	}
	return chunks
}

func BinaryToInteger(bin string) int64 {
	num, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		panic("Unable to convert binary to integer")
	}
	return num
}

func AddBinary(a, b, c, d, e string) string {
	total := BinaryToInteger(a) + BinaryToInteger(b) + BinaryToInteger(c) + BinaryToInteger(d) + BinaryToInteger(e)
	total = total % int64(math.Pow(2.0, 32.0))
	sum := fmt.Sprintf("%.32b", total)
	return sum
}

func MessageSchedule(msg string) []string {
	preProcessedMsg := PreprocessingMsg(msg)
	wordArray := []string{}
	var s string
	for i := 0; i < len(preProcessedMsg); i += 32 {
		s = ""
		s += string(preProcessedMsg[i : i+32])
		wordArray = append(wordArray, s)
	}
	l := len(wordArray)
	for i := len(wordArray); i < 64; i++ {
		s = ""
		for j := 0; j < 32; j++ {
			s += "0"
		}
		wordArray = append(wordArray, s)
	}

	for i := l; i < 64; i++ {
		res1 := XorOperation(RightRotate(wordArray[i-15], 7), RightRotate(wordArray[i-15], 18))
		s0 := XorOperation(res1, RightShift(wordArray[i-15], 3))
		res2 := XorOperation(RightRotate(wordArray[i-2], 17), RightRotate(wordArray[i-2], 19))
		s1 := XorOperation(res2, RightShift(wordArray[i-2], 10))
		wordArray[i] = AddBinary(wordArray[i-16], s0, wordArray[i-7], s1, "0")
	}

	return wordArray
}

func Compression(wordArray []string) string {
	a := HexToBinary(h0)
	b := HexToBinary(h1)
	c := HexToBinary(h2)
	d := HexToBinary(h3)
	e := HexToBinary(h4)
	f := HexToBinary(h5)
	g := HexToBinary(h6)
	h := HexToBinary(h7)
	// fmt.Println("a->", len(a))
	// a, b, c, d, e, f, g, h := h0, h1, h2, h3, h4, h5, h6, h7
	for i := 0; i < 64; i++ {
		s1 := XorOperation(XorOperation(RightRotate(e, 6), RightRotate(e, 11)), RightRotate(e, 25))
		ch := XorOperation(AndOperation(e, f), AndOperation(NotOperation(e), g))
		// ch = (e and f) xor ((not e) and g)
		temp1 := AddBinary(h, s1, ch, HexToBinary(_K[i]), wordArray[i])
		s0 := XorOperation(XorOperation(RightRotate(a, 2), RightRotate(a, 13)), RightRotate(a, 22))
		// S0 = (a rightrotate 2) xor (a rightrotate 13) xor (a rightrotate 22)
		maj := XorOperation(XorOperation(AndOperation(a, b), AndOperation(a, c)), AndOperation(b, c))
		// maj = (a and b) xor (a and c) xor (b and c)
		temp2 := AddBinary(s0, maj, "0", "0", "0")
		// temp2 := S0 + maj
		h = g
		g = f
		f = e
		e = AddBinary(d, temp1, "0", "0", "0")
		d = c
		c = b
		b = a
		a = AddBinary(temp1, temp2, "0", "0", "0")
	}

	H0 := AddBinary(HexToBinary(h0), a, "0", "0", "0")
	H1 := AddBinary(HexToBinary(h1), b, "0", "0", "0")
	H2 := AddBinary(HexToBinary(h2), c, "0", "0", "0")
	H3 := AddBinary(HexToBinary(h3), d, "0", "0", "0")
	H4 := AddBinary(HexToBinary(h4), e, "0", "0", "0")
	H5 := AddBinary(HexToBinary(h5), f, "0", "0", "0")
	H6 := AddBinary(HexToBinary(h6), g, "0", "0", "0")
	H7 := AddBinary(HexToBinary(h7), h, "0", "0", "0")

	return BinaryToHex(H0) + BinaryToHex(H1) + BinaryToHex(H2) + BinaryToHex(H3) +
		BinaryToHex(H4) + BinaryToHex(H5) + BinaryToHex(H6) + BinaryToHex(H7)
	// fmt.Println(H0 == "10111001010011010010011110111001")
	// fmt.Println(H1 == "10010011010011010011111000001000")
	// fmt.Println(H2 == "10100101001011100101001011010111")
	// fmt.Println(H3 == "11011010011111011010101111111010")
	// fmt.Println(H4 == "11000100100001001110111111100011")
	// fmt.Println(H5 == "01111010010100111000000011101110")
	// fmt.Println(H6 == "10010000100010001111011110101100")
	// fmt.Println(H7 == "11100010111011111100110111101001")
}

func main() {
	msg := "hello world"
	messageSchedule := MessageSchedule(msg)
	sha256HashVal := Compression(messageSchedule)
	fmt.Println(strings.ToUpper(sha256HashVal))
}
