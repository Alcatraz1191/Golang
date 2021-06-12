package main

import (
	"fmt"
)

func main() {
	fmt.Println(DuplicateCharOneLoop("aaidd"))
}


//Order of characters doesn't matter
//Eg = "aabbcdd", "abcabcd" give valid results
func DuplicateCharTwoLoop(s string) string {
	sMap := make(map[string]int)
	for i := 0; i < len(s); i++ {
		sMap[string(s[i])] = sMap[string(s[i])] + 1
	}
	for i := 0; i < len(s); i++{
		if sMap[string(s[i])] == 1{
			return string(s[i])
		}
	}
	return "Unique character not found"
}


//Order of characters matter
//Eg = "aabbc" valid, "abcacbe" invalid result
func DuplicateCharOneLoop(s string) string {
	m := make(map[byte]int)
	sli := []byte{}
	for i := 0; i < len(s); i++ {
		//val := int( - 'a')
		if _, ok := m[s[i]]; !ok {
			sli = append(sli, s[i])
			m[s[i]] = len(sli)
		} else {
		sli = append(sli[:m[s[i]]-1], sli[m[s[i]]:]...)
		}
	}
	
	return string(sli[0])
}
