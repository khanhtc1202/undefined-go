package main

import "fmt"

func main() {
	perm([]string{"A", "B", "C"}, []string{})
}

func perm(s1 []string, s2 []string) {
	if len(s1) == 0 {
		for _, c := range s2 {
			fmt.Print(c)
		}
		fmt.Println()
		return
	}
	for i := 0; i < len(s1); i++ {
		s3 := append(append([]string{}, s1[:i]...), s1[i+1:]...)
		s4 := append(s2, s1[i])
		perm(s3, s4)
	}
}
