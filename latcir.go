package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	input_raw, _ := reader.ReadString('\n')

	lat := []rune("ABCČĆDĐEFGHIJKLMNOPRSŠTUVZŽabcčćdđefghijklmnoprsštuvzž")
	cir := []rune("АБЦЧЋДЂЕФГХИЈКЛМНОПРСШТУВЗЖабцчћдђефгхијклмнопрсштувзж")
	input := []rune(input_raw)

	var digraf bool = false

	for i, c := range input {

		if c == rune('L') && (input[i+1] == rune('J') || input[i+1] == 'j') {
			fmt.Printf("Љ")
			digraf = true
		} else if c == rune('l') && (input[i+1] == rune('J') || input[i+1] == 'j') {
			fmt.Printf("љ")
			digraf = true
		} else if c == rune('N') && (input[i+1] == rune('J') || input[i+1] == 'j') {
			fmt.Printf("Њ")
			digraf = true
		} else if c == rune('n') && (input[i+1] == rune('J') || input[i+1] == 'j') {
			fmt.Printf("њ")
			digraf = true
		} else if c == rune('D') && (input[i+1] == rune('Ž') || input[i+1] == 'ž') {
			fmt.Printf("Џ")
			digraf = true
		} else if c == rune('d') && (input[i+1] == rune('Ž') || input[i+1] == 'ž') {
			fmt.Printf("џ")
			digraf = true
		} else {
			if digraf {
				digraf = false
				continue
			}
			for k, slovo := range lat {
				if input[i] == slovo {
					fmt.Printf("%c", cir[k])
				}
			}
		}
	}

}
