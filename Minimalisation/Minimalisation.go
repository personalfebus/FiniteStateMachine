package main

import "fmt"

type state struct{
	parent, depth, q int
}

func main() {
	var n, m, q int
	fmt.Scan(&n, &m, &q)
	Tr := make([][]int, n)
	Vert := make([]state, n)
	Pi := make([]int, n)
	for i := 0; i < n; i++ {
		Tr[i] = make([]int, m)
		Vert[i].q = i
		Vert[i].depth = 0
		Vert[i].parent = i
		Pi[i] = 0
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
	//fmt.Printf("\n")
	l := aufenkamphonn1(Tr, Out, Vert, Pi)

	//fmt.Printf("\n%d\n", l)

	Tr2 := make([][]int, l)
	Tr3 := make([][]int, l)
	Out2 := make([][]string, l)
	Out3 := make([][]string, l)
	Vert2 := make ([]int, l + 1)
	for i := 0; i < l; i++{
		Tr2[i] = make([]int, m)
		Tr3[i] = make([]int, m)
		Out2[i] = make([]string, m)
		Out3[i] = make([]string, m)
		Vert2[i] = -1
	}
	Vert2[l] = 0
	qq := aufenkamphonn(Tr, Out, Vert, Pi, Tr2, Out2, q)

	//fmt.Printf("\n\n%d\n%d\n0\n", l, m)
	//
	//for i := 0; i < l; i++{
	//	for j := 0; j < m; j++{
	//		fmt.Printf("%d ", Tr2[i][j])
	//	}
	//	fmt.Printf("\n")
	//}
	//fmt.Printf("\n")
	//for i := 0; i < l; i++{
	//	for j := 0; j < m; j++{
	//		fmt.Printf("%s ", Out2[i][j])
	//	}
	//	fmt.Printf("\n")
	//}
	doublecheck := make([]int, l)
	for i := range doublecheck{
		doublecheck[i] = 0
	}
	dfs1(Tr2, Tr3, Out2, Out3, Vert2, qq, doublecheck)
	//for i := 0; i < len(Vert2) - 1; i++{
	//	if Vert2[i] == -1{
	//		//fmt.Printf("%d %d\n", i, Vert2[i])
	//		dfs1(Tr2, Tr3, Out2, Out3, Vert2, i)
	//	}
	//}
	s := '"'
	fmt.Printf("digraph {\n\trankdir = LR\n\tdummy [label = %c%c, shape = none]\n\t", s, s)
	for i := 0; i < l; i++{
		if doublecheck[i] > 0 {
			fmt.Printf("%d [shape = circle]\n\t", i)
		}
	}
	fmt.Printf("dummy -> %d\n", 0)
	for i := 0; i < l; i++{
		if doublecheck[i] > 0 {
			for j := 0; j < m; j++ {
				tok := 'a' + j
				fmt.Printf("\t%d -> %d [label = %c%c(%s)%c]\n", i, Tr3[i][j], s, tok, Out3[i][j], s)
			}
		}
	}
	fmt.Printf("}")
}

func union(x, y int, vert[] state)  {
	rootx := find(x, vert)
	rooty := find(y, vert)
	if vert[rootx].depth < vert[rooty].depth{
		vert[rootx].parent = rooty
	} else{
		vert[rooty].parent = rootx
		if (vert[rootx].depth == vert[rooty].depth) && (vert[rootx].q != vert[rooty].q){
			vert[rootx].depth = vert[rootx].depth + 1
		}
	}
}

func find(x int, vert[] state) int {
	if vert[x].parent == x{
		return x
	} else{
		vert[x].parent = find(vert[x].parent, vert)
		return vert[x].parent
	}
}

func split1(in[][] int, out[][] string, vert[] state, pi[] int) int {
	m := len(vert)
	for i := range vert{
		vert[i].parent = i
		vert[i].depth = 0
		vert[i].q = i
	}
	for i := 0; i < len(vert); i++{
		for j := i + 1; j < len(vert); j++{
			//fmt.Printf("%d(par = %d) and %d(par = %d) ", i, find(i, vert), j, find(j, vert))
			if find(i, vert) != find(j, vert){
				eq := true
				for k := 0; k < len(out[i]); k++{
					if out[i][k] != out[j][k]{
						eq = false
						//fmt.Printf(" FALSE")
						break
					}
				}
				if eq{
					//fmt.Printf(" TRUE")
					//pi[i]=pi[j]
					union(i, j, vert)
					m--
				}
			}
			//fmt.Printf("\n")
		}
	}
	for i, x := range vert{
		pi[x.q] = find(i, vert)
	}
	return m
}

func split(in[][] int, out[][] string, vert[] state, pi[] int) int {
	m := len(vert)
	for i := range vert{
		vert[i].parent = i
		vert[i].depth = 0
		vert[i].q = i
		//fmt.Printf("%d == %d\n", x.parent, i)
	}
	//fmt.Printf("%d == %d\n", vert[3].parent, find(3, vert))
	for i := 0; i < len(vert); i++ {
		for j := i + 1; j < len(vert); j++ {
			//fmt.Printf("%d vs %d :: %d ? %d && %d ? %d\n", i, j, pi[vert[i].q], pi[vert[j].q], find(i, vert), find(j, vert))
			if (pi[vert[i].q] == pi[vert[j].q]) && (find(i, vert) != find(j, vert)){
				//fmt.Printf("huh?\n")
				eq := true
				for k := 0; k < len(in[i]); k++{
					w1 := in[i][k]
					w2 := in[j][k]
					if pi[w1] != pi[w2]{
						//fmt.Printf("Yikes: %d !+ %d", pi[w1], pi[w2])
						eq = false
						break
					}
				}
				if eq{
					//fmt.Printf("GOTTEM %d\n", m)
					union(i, j, vert)
					m--
				}
			}
		}
	}
	for i := range vert{
		pi[i] = find(i, vert)
	}
	return m
}

func aufenkamphonn1(in[][] int, out[][] string, vert[] state, pi[] int) int {
	m := split1(in, out, vert, pi)
	for ;;{
		m1 := split(in, out, vert, pi)
		//fmt.Printf("%d %d\n", m, m1)
		if m == m1{
			break
		}
		m = m1
	}
	return m
}

func aufenkamphonn(in[][] int, out[][] string, vert[] state, pi[] int, in2[][] int, out2[][] string, q int) int {
	for i := 0; i < len(in2); i++{
		for j := 0; j < len(in2[i]); j++{
			in2[i][j] = -1
			out2[i][j] = ""
		}
	}

	mover := make([]int, len(pi))
	negro := 0

	for i := range pi{
		//fmt.Printf("%d belongs to %d, ", i, pi[i])
		if pi[i] == i{
			mover[i] = negro
			negro++
		} else {
			mover[i] = -1
		}
	}

	//for i := range vert{
	//	q2 := pi[i]
	//	fmt.Printf("%d class = %d\n", i, q2)
	//	if (q2 > -1) && (q2 < len(in2)) && (in2[q2][0] == -1){
	//		for j := 0; j < len(in2[0]); j++{
	//			fmt.Printf("%d -> ", j)
	//			fmt.Printf("%d %s\n", pi[in[i][j]], out[i][j])
	//			in2[q2][j] = pi[in[i][j]]
	//			out2[q2][j] = out[i][j]
	//		}
	//	}
	//}

	for i := range mover{
		//fmt.Printf("%d", i, )
		if mover[i] > -1{
			for j := 0; j < len(in2[0]); j++{
				in2[mover[i]][j] = mover[pi[in[i][j]]]
				out2[mover[i]][j] = out[i][j]
			}
		}
	}
	return mover[pi[q]]
}

func dfs1(in1 [][]int, in2 [][]int, out1 [][]string, out2 [][]string, vert []int, curr int, dch []int) {
	vert[curr] = vert[len(vert) - 1]
	vert[len(vert) - 1]++
	dch[vert[curr]] = 1
	for i := 0; i < len(in1[0]); i++{
		//fmt.Printf("Curr = %d, Ind = %d ", curr, i)
		u := in1[curr][i]
		//fmt.Printf("%d\n", u)
		if vert[u] == -1{
			dfs1(in1, in2, out1, out2, vert, u, dch)
		}
		in2[vert[curr]][i] = vert[u]
		out2[vert[curr]][i] = out1[curr][i]
	}
}