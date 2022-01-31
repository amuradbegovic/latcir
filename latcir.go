package latcir

var lat = []rune("ABCČĆDĐEFGHIJKLMNOPRSŠTUVZŽabcčćdđefghijklmnoprsštuvzž")
var cir = []rune("АБЦЧЋДЂЕФГХИЈКЛМНОПРСШТУВЗЖабцчћдђефгхијклмнопрсштувзж")

func Latcir(input []rune) string {

	var output string
	var digraf, grafem bool = false, false

	for i, c := range input {
		if i < len(input)-1 {
			if c == 'L' && (input[i+1] == 'J' || input[i+1] == 'j') {
				output = output + "Љ"
				digraf = true
			} else if c == 'l' && (input[i+1] == 'J' || input[i+1] == 'j') {
				output = output + "љ"
				digraf = true
			} else if c == 'N' && (input[i+1] == 'J' || input[i+1] == 'j') {
				output = output + "Њ"
				digraf = true
			} else if c == 'n' && (input[i+1] == 'J' || input[i+1] == 'j') {
				output = output + "њ"
				digraf = true
			} else if c == 'D' && (input[i+1] == 'Ž' || input[i+1] == 'ž') {
				output = output + "Џ"
				digraf = true
			} else if c == 'd' && (input[i+1] == 'Ž' || input[i+1] == 'ž') {
				output = output + "џ"
				digraf = true
			}
		}
		if digraf {
			digraf = false
			continue
		}
		for k, slovo := range lat {
			if input[i] == slovo {
				output = output + string(cir[k])
				grafem = true
				break
			}
		}
		if !grafem {
			output = output + string(c)
		} else {
			grafem = false
		}

	}
	return output
}

// Originalna funkcija se skoro potpuno ponavlja. Moram pronaći pametnije rješenje.
func Cirlat(input []rune) string {

	var output string
	var digraf, grafem bool = false, false

	for i, c := range input {
		if c == 'Љ' {
			output = output + "Lj"
			digraf = true
		} else if c == 'љ' {
			output = output + "lj"
			digraf = true
		} else if c == 'Њ' {
			output = output + "Nj"
			digraf = true
		} else if c == 'њ' {
			output = output + "nj"
			digraf = true
		} else if c == 'Џ' {
			output = output + "Dž"
			digraf = true
		} else if c == 'џ' {
			output = output + "dž"
			digraf = true
		}
		if digraf {
			digraf = false
			continue
		}
		for k, slovo := range cir {
			if input[i] == slovo {
				output = output + string(lat[k])
				grafem = true
				break
			}
		}
		if !grafem {
			output = output + string(c)
		} else {
			grafem = false
		}
	}

	return output
}
