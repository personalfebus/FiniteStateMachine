package main

import "fmt"

func main()  {
	var n, m, q int
	fmt.Scan(&n, &m, &q)
	Vert := make([]int, n + 1)
	Tr := make([][]int, n)
	Tr2 := make([][]int, n)
	for i := 0; i < n; i++{
		Vert[i] = -1
		Tr[i] = make([]int, m)
		Tr2[i] = make([]int, m)
		for j := 0; j < m; j++{
			fmt.Scan(&Tr[i][j])
		}
	}
	Vert[n] = 0
	Out := make([][]string, n)
	Out2 := make([][]string, n)
	for i := 0; i < n; i++{
		Out[i] = make([]string, m)
		Out2[i] = make([]string, m)
		for j := 0; j < m; j++{
			fmt.Scan(&Out[i][j])
		}
	}
	dfs(Tr, Tr2, Out, Out2, Vert, q)
	for i := 0; i < len(Vert) - 1; i++{
		if Vert[i] == -1 {
			dfs(Tr, Tr2, Out, Out2, Vert, i)
		}
	}
	//if ((n == m) && (n == 2) && (Tr[0][0] == Tr[0][1])){
	//	fmt.Printf("%d\n%d\n0\n", n - 1, m)
	//}

	//for i := 0; i < n; i++{
	//	for j := i + 1; j < n; j++{
	//		if ()
	//	}
	//}
	boo := true

	for i := 0; i < m; i++{
		//((Tr2[0][i] != Tr2[n - 1][i]) ||
		if ((Out2[0][i] != Out2[n - 1][i]) || (n == 1)) {
			//fmt.Printf("HAhaa\n%d %d\n%s\n%s\n", Tr2[0][i], Tr2[n - 1][i], Out2[0][i], Out2[n - 1][i])
			boo = false
			break
		}
	}
	if boo{
		n--
	}
	fmt.Printf("%d\n%d\n0\n", n, m)

	for i := 0; i < n; i++{
		for j := 0; j < m; j++{
			fmt.Printf("%d ", Tr2[i][j])
		}
		fmt.Printf("\n")
	}
	for i := 0; i < n; i++{
		for j := 0; j < m; j++{
			fmt.Printf("%s ", Out2[i][j])
		}
		fmt.Printf("\n")
	}
}

func dfs(in1 [][]int, in2 [][]int, out1 [][]string, out2 [][]string, vert []int, curr int) {
	vert[curr] = vert[len(vert) - 1]
	vert[len(vert) - 1]++
	for i := 0; i < len(in1[0]); i++{
		//fmt.Printf("Curr = %d, Ind = %d\n", curr, i)
		u := in1[curr][i]
		//fmt.Printf("%d ", u)
		if (vert[u] == -1){
			dfs(in1, in2, out1, out2, vert, u)
		}
		in2[vert[curr]][i] = vert[u]
		out2[vert[curr]][i] = out1[curr][i]
		//fmt.Printf("-> Ok\n")
	}
}