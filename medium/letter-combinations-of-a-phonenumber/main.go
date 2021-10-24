package main
import "fmt"

func letterCombinations(digits string) []string {
	var combinations map[byte][]string = make(map[byte][]string)
	combinations['2'] = []string{"a","b","c"}
	combinations['3'] = []string{"d","e","f"}
	combinations['4'] = []string{"g","h","i"}
	combinations['5'] = []string{"j","k","l"}
	combinations['6'] = []string{"m","n","o"}
	combinations['7'] = []string{"p","q","r","s"}
	combinations['8'] = []string{"t","u","v"}
	combinations['9'] = []string{"w","x","y","z"}

	


	return  combine(combinations, []string{}, digits)
}

func combine(combinations map[byte][]string, start []string, rest string)[]string {
	if len(rest) == 0 {
		return start
	}
	
	var prefix string = ""
	var result []string
	var i int = 0

	for {
		if i < len(start) {
			prefix = start[i]
		}
		for _, letter := range combinations[rest[0]]{
			result = append(result, fmt.Sprintf("%s%s", prefix, letter))
		}
		i++
		if i >= len(start) {
			break
		}
	}

	return combine(combinations, result, rest[1:])
	
}

