package recursion

import "fmt"

func PrintInteger(n int) {
	if n < 10 {
		fmt.Printf("%d", n)
	} else {
		PrintInteger(n / 10)
		fmt.Printf("%d", n%10)
	}
}

// The Tower of Hanoi
func MoveTower(n int, start, finish, tmp string) {
	if n == 1 {
		fmt.Printf("%s --> %s\n", start, finish)
	} else {
		MoveTower(n-1, start, tmp, finish)
		fmt.Printf("%s --> %s\n", start, finish)
		MoveTower(n-1, tmp, finish, start)
	}
}

// func ListPermutations(s string) {
// 	listPermutationsWithPrefix("", s)
// }

// func listPermutationsWithPrefix(prefix, s string) {
// 	if len(s) == 0 {
// 		// fmt.Println(prefix)
// 	} else {
// 		for i := 0; i < len(s); i++ {
// 			ch := s[i]
// 			rest := s[i+1:]
// 			listPermutationsWithPrefix(fmt.Sprintf("%s%s", prefix, string(ch)), rest)
// 		}
// 	}
// }

func ReverseString(s string) string {
	if len(s) == 1 {
		return s
	}
	return ReverseString(s[1:]) + string(s[0])
}

func CountX(s string) int {
	if len(s) == 1 {
		if string(s[0]) == "x" {
			return 1
		}
		return 0
	}

	if string(s[0]) == "x" {
		return 1 + CountX(s[1:])
	} else {
		return CountX(s[1:])
	}
}
