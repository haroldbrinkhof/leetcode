package main
import "strings"

type sequence struct {
	value string
	openingNr int
	closingNr int
	limit int	
}

func calcNextSequence(sequences []sequence) (result []sequence){
	for _, seq := range sequences {
		if seq.openingNr + 1 <= seq.limit {
			var b strings.Builder
			b.WriteString(seq.value)
			b.WriteString("(")
			openingSeq := sequence{ b.String() , seq.openingNr + 1, seq.closingNr, seq.limit }
			result = append(result, openingSeq)
		}

		if seq.closingNr + 1 <= seq.limit && seq.closingNr + 1 <= seq.openingNr {
			var b strings.Builder
			b.WriteString(seq.value)
			b.WriteString(")")
			closingSeq := sequence{ b.String() , seq.openingNr, seq.closingNr + 1, seq.limit }
			result = append(result, closingSeq)
		}
	}

	if result[0].closingNr == result[0].limit {
		return
	}
	
	return calcNextSequence(result)
}

func generateParenthesis(n int) []string {
	var outcome []string
	var sequences []sequence = []sequence{ sequence{"(",1,0,n} }
	for _,seq := range calcNextSequence(sequences){
		outcome = append(outcome,seq.value)
	}
	return outcome
}

