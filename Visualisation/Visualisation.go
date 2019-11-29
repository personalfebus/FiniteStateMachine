package main

import "fmt"

func main() {
	var n, m, q int;
	fmt.Scan(&n, &m, &q)
	Tr := make([][]int, n)
	for i := 0; i < n; i++ {
		Tr[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&Tr[i][j])
		}
	}
	Out := make([][]string, n)
	for i := 0; i < n; i++ {
		Out[i] = make([]string, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&Out[i][j])
		}
	}
	s := '"'
	fmt.Printf("digraph {\n\trankdir = LR\n\tdummy [label = %c%c, shape = none]\n\t", s, s)
	for i := 0; i < n; i++{
		fmt.Printf("%d [shape = circle]\n\t", i)
	}
	fmt.Printf("dummy -> %d\n\t", q)

	for i := 0; i < n; i++{
		for j := 0; j < m; j++{
			if (i == n - 1) && (j == m -1){
				tok := 'a' + j;
				fmt.Printf("%d -> %d [label = %c%c(%s)%c]\n", i, Tr[i][j], s, tok, Out[i][j], s)
				break
			}
			tok := 'a' + j;
			fmt.Printf("%d -> %d [label = %c%c(%s)%c]\n\t", i, Tr[i][j], s, tok, Out[i][j], s)
		}
	}
	fmt.Printf("}")
}