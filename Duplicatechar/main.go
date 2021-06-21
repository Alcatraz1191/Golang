package main

import (
	"fmt"
)

type Values struct {
	Index int
	occ   int
}

func main() {
	fmt.Println(DuplicateCharOneLoop("c"))
}

//Order of characters doesn't matter
//Eg = "aabbcdd", "abcabcd" give valid results
func DuplicateCharTwoLoop(s string) string {
	sMap := make(map[string]int)
	for i := 0; i < len(s); i++ {
		sMap[string(s[i])] = sMap[string(s[i])] + 1
	}
	for i := 0; i < len(s); i++ {
		if sMap[string(s[i])] == 1 {
			return string(s[i])
		}
	}
	return "Unique character not found"
}

//Order of characters matter
//Eg = "aabbc" valid, "abcacbe" invalid result
func DuplicateCharOneLoop(s string) string {
	m := make(map[byte]Values)
	sli := []byte{}
	for i := 0; i < len(s); i++ {
		//val := int( - 'a')
		if _, ok := m[s[i]]; !ok {
			sli = append(sli, s[i])
			m[s[i]] = Values{
				Index: len(sli),
				occ:   1,
			}
		} else if len(sli) > 0 && m[s[i]].occ < 2 {
			va := m[s[i]]
			va.occ++
			m[s[i]] = va
			sli = append(sli[:m[s[i]].Index-1], sli[m[s[i]].Index:]...)
		}
	}

	if len(sli) != 0 {
		return string(sli[0])
	} else {
		return "Unique character not found"
	}
}
