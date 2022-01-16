package satisfiabilityofequailityequation

func EquationsPossible(equations []string) bool {
	a := make([]int, 26)
	for i := 0; i < 26; i++ {
		a[i] = i
	}
	for _, e := range equations {
		if e[1] == '=' {
			f := e[0] - 'a'
			l := e[3] - 'a'
			a[findRoot(a, int(l))] = findRoot(a, int(f))
		}
	}
	for _, e := range equations {
		if e[1] == '!' {
			f := e[0] - 'a'
			l := e[3] - 'a'
			if findRoot(a, int(f)) == findRoot(a, int(l)) {
				return false
			}
		}
	}
	return true
}

func findRoot(a []int, i int) int {
	if a[i] == i {
		return i
	} else {
		return findRoot(a, a[i])
	}
}
