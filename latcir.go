package latcir

func Latcir(input []rune) string  {


	lat := []rune("ABCČĆDĐEFGHIJKLMNOPRSŠTUVZŽabcčćdđefghijklmnoprsštuvzž")
	cir := []rune("АБЦЧЋДЂЕФГХИЈКЛМНОПРСШТУВЗЖабцчћдђефгхијклмнопрсштувзж")

	var output string
	var digraf bool = false

	for i, c := range input {

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
		} else {
			if digraf {
				digraf = false
				continue
			}
			for k, slovo := range lat {
				if input[i] == slovo {
					output = output + string(cir[k])
				}
			}
		}
	}
	return output
}
