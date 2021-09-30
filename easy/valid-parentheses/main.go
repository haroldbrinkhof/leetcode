package main
import "strings"



func pushPrevious(s string, v rune) string {
    var b strings.Builder
    b.WriteString(s)
    b.WriteRune(v)
    return b.String()
}

func popPrevious(s string) (string, rune) {
    if len(s) == 0 {
        return "", rune(0)
    }
    return  s[0:len(s) - 1], rune(s[len(s) - 1])
}

func isValid(s string) bool {
    var curly_braces int = 0
    var brackets int = 0
    var parens int = 0

    var previous string
    var previousRune rune
    for _,v := range s{
        switch v {
            case '(':
                previous = pushPrevious(previous, v)
                parens++
            case ')':
                previous, previousRune = popPrevious(previous) 
                if previousRune != '('{
                    return false
                }
                parens--
                if parens < 0 {
                    return false
                }
            case '[':
                previous = pushPrevious(previous, v)
                brackets++
            case ']':
                previous, previousRune = popPrevious(previous) 
                brackets--
                if previousRune != '[' || brackets < 0{
                    return false
                }
            case '{':
                previous = pushPrevious(previous, v)
                curly_braces++
            case '}':
                previous, previousRune = popPrevious(previous) 
                curly_braces--
                if previousRune != '{' || curly_braces < 0{
                    return false
                }
            default:
        }
       
    }
    return parens + brackets + curly_braces == 0
}
