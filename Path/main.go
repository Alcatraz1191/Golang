package main

import (
	"fmt"
	"strings"
)

/*
Given a string path, where path[i] = 'N', 'S', 'E' or 'W', each representing moving one unit north, south, east, or west, respectively. You start at the origin (0, 0) on a 2D plane and walk on the path specified by path.

Return true if the path crosses itself at any point, that is, if at any time you are on a location you have previously visited. Return false otherwise.
*/

func isPathCrossing(path string) bool {
    temp := make([][]int,len(path) + 1)
    temp [0] = make([]int,2)
    for i := 1; i < len(temp); i++  {
      temp[i] = make([]int, 2)
    }
    n := strings.Split(path, "")
    for i, _ := range path{
        if n[i] == "N"{
            if i > 0{
                copy(temp[i], temp[i-1])
            }
            temp[i][0]++
        }
        if n[i] == "S"{
            if i > 0{
                copy(temp[i], temp[i-1])
            }
            temp[i][0]--
        }
        if n[i] == "E"{
            if i > 0{
                copy(temp[i], temp[i-1])
            }
            temp[i][1]++
        }
        if n[i] == "W"{
            if i > 0 {
                copy(temp[i], temp[i-1])
            }
            temp[i][1]--
        }
    }
    cross := false
    for i := 0; i < len(temp); i++{
        for j := i+1; j < len(temp); j++{
            if Equal(temp[i], temp[j]){
                cross = true
            }
        }
    }
    return cross
}

func Equal(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}

func main() {
	n := "NESWW"

	is := isPathCrossing(n)

	fmt.Println(is)
}