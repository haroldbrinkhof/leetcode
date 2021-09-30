package main

const dot rune = rune('.')
const asterisk rune = rune('*')

func isMatch(s string, p string) bool {
    return match(s,p)
}


func match(s,p string) bool {
	if len(s) == 0 && len(p) == 0 {
		return true
	}
	if len(s) > 0 && len(p) == 0 {
		return false
	}


	toMatch, isMulti, sequenceRest := getNextMatchingSequence(p)
	success, nonMatchingRest := matchToInput(toMatch, isMulti, s)
	if !success {
		return false
	}
	
	if success && len(nonMatchingRest) == 0 && len(sequenceRest) == 0 {
		return true
	}

	var matches bool = false
	if success {
		for _,v := range nonMatchingRest {
			matches = matches || match(v,sequenceRest)
		}
	}

	if len(nonMatchingRest) == 0 && len(sequenceRest) > 1 { // check for straggling <character>* at the end
		var onlyMultis bool = true
		for len(sequenceRest) > 0 {
			_, isMulti, sequenceRest = getNextMatchingSequence(sequenceRest)
			onlyMultis = onlyMultis && isMulti
		}
		matches = matches || onlyMultis
	}

	return matches //false 
}

func matchToInput(match rune, isMulti bool, input string) (success bool, nonMatchingRest []string){
	if len(input) == 0 && !isMulti{
		success = false
		return
	}

	if isMulti {
		success = true
		for i,v := range input {
			nonMatchingRest = append(nonMatchingRest, input[i:]) // zero-length match
			if match != dot && v != match {
				break
			}
			if match != dot && match == v && i + 1 == len(input){
				nonMatchingRest = append(nonMatchingRest,"")
			}
		}
		if match == dot {
			nonMatchingRest = append(nonMatchingRest,"")
		}

		return
	}

	if len(input) > 0 {
		if len(input) > 1 {
			nonMatchingRest = []string{ input[1:] }
		}
		success = charMatches(rune(input[0]), match)
	}
	return
}

func charMatches(input, match rune) bool {
	return input == match || match == dot
}

func getNextMatchingSequence(regex string) (match rune, isMulti bool, rest string){
	var size int = 1
	isMulti = false
	match = rune(regex[0])
	if len(regex) > 1 && rune(regex[1]) == asterisk{
		size++
		isMulti = true
	}
	rest = regex[size:]

	return
}

