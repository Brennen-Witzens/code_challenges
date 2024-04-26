package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	fmt.Print("Please input a string. ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	obfuscated_str := obfuscate(input.Text())
	deobfuscated_str := reverse(deobfuscate(obfuscated_str))
	fmt.Println("Obfuscated: ", obfuscated_str)
	fmt.Println("Deobfuscated: ", deobfuscated_str)
}

func obfuscate(str string) string {

	reversed_str := reverse(str)
	rs := make([]rune, 0, utf8.RuneCount([]byte(reversed_str)))
	for _, s := range reversed_str {
		//fmt.Println("String value: ", string(s))
		changed_val := (s + 15)
		rs = append(rs, changed_val)
	}

	return string(rs)
}

func deobfuscate(str string) string {

	n := make([]rune, 0, utf8.RuneCount([]byte(str)))
	for _, s := range str {
		//fmt.Println("String value: ", string(s))
		changed_value := (s - 15)
		n = append(n, changed_value)
	}

	return string(n)
}

func reverse(str string) (result string) {

	for _, v := range str {
		result = string(v) + result
	}
	return
}
