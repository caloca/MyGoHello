package stringutil

func Reverse(str string) string {
	revstr := []rune(str)
	strlen := len(revstr)

	for i,j := 0,strlen-1 ; i < strlen/2 ; i,j = i+1,j-1 {
		revstr[i] , revstr[j] = revstr[j] , revstr[i]
	}

	return string(revstr)
}
