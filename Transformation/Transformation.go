package main

import "fmt"

func main()  {
	var n, m, q int
	fmt.Scan(&n)
	AlphIn := make([]string, n)
	for i := range AlphIn{
		fmt.Scan(&AlphIn[i])
	}
	fmt.Scan(&m)
	AlphOut := make([]string, m)
	for i := range AlphOut{
		fmt.Scan(&AlphOut[i])
	}
	fmt.Scan(&q)
	Tr := make([][]int, q)
	for i := range Tr{
		Tr[i] = make([]int, n)
		for j := range Tr[i]{
			fmt.Scan(&Tr[i][j])
		}
	}
	Out := make([][]string, q)
	for i := range Out{
		Out[i] = make([]string, n)
		for j := range Out[i]{
			fmt.Scan(&Out[i][j])
		}
	}
	Pairs := make([][]int, q)
	PairsNum := make([][]int, q)
	numbers := 0
	for k := range Pairs{
		count := 0
		Pairs[k] = make([]int, m)
		PairsNum[k] = make([]int, m)

		for i := range Pairs[k]{
			Pairs[k][i] = -1
			PairsNum[k][i] = numbers
			numbers++
		}

		for i := range Tr{
			for j := range Tr[i]{
				if Tr[i][j] == k{
					for l := range AlphOut{
						if AlphOut[l] == Out[i][j]{
							unic := true
							for b := 0; b < count; b++{
								if Pairs[k][b] == l{
									unic = false
									break
								}
							}
							if unic {
								Pairs[k][count] = l
								count++
							}
							break
						}
					}
				}
			}
		}
	}

	s := '"'
	fmt.Printf("digraph {\n\trankdir = LR\n")

	for i := range Pairs{
		for j := range Pairs[i]{
			if Pairs[i][j] != -1 {
				//0 [label = "(0,0)"]
				fmt.Printf("\t%d [label = \"(%d,%s)\"]\n", PairsNum[i][j], i, AlphOut[Pairs[i][j]])
			}
		}
	}

	for i := range Tr{
		for j := range Tr[i]{
			for k := range Pairs[i]{
				if Pairs[i][k] == -1{
					break
				}
				pos := 0
				for b := range Pairs[Tr[i][j]]{
					if AlphOut[Pairs[Tr[i][j]][b]] == Out[i][j]{
						pos = b
						break
					}
				}
				fmt.Printf("\t%d -> %d [label = %c%s%c]\n", PairsNum[i][k], PairsNum[Tr[i][j]][pos], s, AlphIn[j], s)
			}
		}
	}
	fmt.Printf("}")
}
