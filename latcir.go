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

		if c == 'L' && (input[i+1] == 'J' || input[i+1] == 'j') {
			fmt.Printf("Љ")
			digraf = true
		} else if c == 'l' && (input[i+1] == 'J' || input[i+1] == 'j') {
			fmt.Printf("љ")
			digraf = true
		} else if c == 'N' && (input[i+1] == 'J' || input[i+1] == 'j') {
			fmt.Printf("Њ")
			digraf = true
		} else if c == 'n' && (input[i+1] == 'J' || input[i+1] == 'j') {
			fmt.Printf("њ")
			digraf = true
		} else if c == 'D' && (input[i+1] == 'Ž' || input[i+1] == 'ž') {
			fmt.Printf("Џ")
			digraf = true
		} else if c == 'd' && (input[i+1] == 'Ž' || input[i+1] == 'ž') {
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
