package main

import "fmt"

type state struct{
	parent, depth, q int
}

func main()  {
	var An, Am, Aq int
	fmt.Scan(&An, &Am, &Aq)
	ATr := make([][]int, An)
	AVert := make([]state, An)
	APi := make([]int, An)
	for i := 0; i < An; i++ {
		ATr[i] = make([]int, Am)
		AVert[i].q = i
		AVert[i].depth = 0
		AVert[i].parent = i
		APi[i] = 0
		for j := 0; j < Am; j++ {
			fmt.Scan(&ATr[i][j])
		}
	}
	AOut := make([][]string, An)
	for i := 0; i < An; i++ {
		AOut[i] = make([]string, Am)
		for j := 0; j < Am; j++ {
			fmt.Scan(&AOut[i][j])
		}
	}
	Al := aufenkamphonn1(ATr, AOut, AVert, APi)
	ATr2 := make([][]int, Al)
	ATr3 := make([][]int, Al)
	AOut2 := make([][]string, Al)
	AOut3 := make([][]string, Al)
	AVert2 := make ([]int, Al + 1)
	for i := 0; i < Al; i++{
		ATr2[i] = make([]int, Am)
		ATr3[i] = make([]int, Am)
		AOut2[i] = make([]string, Am)
		AOut3[i] = make([]string, Am)
		AVert2[i] = -1
	}
	AVert2[Al] = 0
	Aqq := aufenkamphonn(ATr, AOut, AVert, APi, ATr2, AOut2, Aq)
	Adoublecheck := make([]int, Al)
	for i := range Adoublecheck{
		Adoublecheck[i] = 0
	}
	dfs1(ATr2, ATr3, AOut2, AOut3, AVert2, Aqq, Adoublecheck)
	All := 0
	for i := range Adoublecheck{
		All += Adoublecheck[i]
	}
	Acount := 0
	ATr4 := make([][]int, All)
	AOut4 := make([][]string, All)
	for i := range Adoublecheck{
		if Adoublecheck[i] == 1{
			ATr4[Acount] = make([]int, Am)
			AOut4[Acount] = make([]string, Am)
			for j := 0; j < Am; j++{
				ATr4[Acount][j] = ATr3[i][j]
				AOut4[Acount][j] = AOut3[i][j]
			}
			Acount++
		}
	}

	var Bn, Bm, Bq int
	fmt.Scan(&Bn, &Bm, &Bq)
	BTr := make([][]int, Bn)
	BVert := make([]state, Bn)
	BPi := make([]int, Bn)
	for i := 0; i < Bn; i++ {
		BTr[i] = make([]int, Bm)
		BVert[i].q = i
		BVert[i].depth = 0
		BVert[i].parent = i
		BPi[i] = 0
		for j := 0; j < Bm; j++ {
			fmt.Scan(&BTr[i][j])
		}
	}
	BOut := make([][]string, Bn)
	for i := 0; i < Bn; i++ {
		BOut[i] = make([]string, Bm)
		for j := 0; j < Bm; j++ {
			fmt.Scan(&BOut[i][j])
		}
	}
	Bl := aufenkamphonn1(BTr, BOut, BVert, BPi)
	BTr2 := make([][]int, Bl)
	BTr3 := make([][]int, Bl)
	BOut2 := make([][]string, Bl)
	BOut3 := make([][]string, Bl)
	BVert2 := make ([]int, Bl + 1)
	for i := 0; i < Bl; i++{
		BTr2[i] = make([]int, Bm)
		BTr3[i] = make([]int, Bm)
		BOut2[i] = make([]string, Bm)
		BOut3[i] = make([]string, Bm)
		BVert2[i] = -1
	}
	BVert2[Bl] = 0
	Bqq := aufenkamphonn(BTr, BOut, BVert, BPi, BTr2, BOut2, Bq)
	Bdoublecheck := make([]int, Bl)
	for i := range Bdoublecheck{
		Bdoublecheck[i] = 0
	}
	dfs1(BTr2, BTr3, BOut2, BOut3, BVert2, Bqq, Bdoublecheck)
	Bll := 0
	for i := range Bdoublecheck{
		Bll += Bdoublecheck[i]
	}
	Bcount := 0
	BTr4 := make([][]int, Bll)
	BOut4 := make([][]string, Bll)
	for i := range Bdoublecheck{
		if Bdoublecheck[i] == 1{
			BTr4[Bcount] = make([]int, Bm)
			BOut4[Bcount] = make([]string, Bm)
			for j := 0; j < Bm; j++{
				BTr4[Bcount][j] = BTr3[i][j]
				BOut4[Bcount][j] = BOut3[i][j]
			}
			Bcount++
		}
	}

	eq := true
	if (All != Bll) || (Am != Bm){
		eq = false
	} else{
		for i := range ATr4{
			for j := range ATr4[0]{
				if (ATr4[i][j] != BTr4[i][j]) || (AOut4[i][j] != BOut4[i][j]){
					eq = false
					break
				}
			}
			if eq == false{
				break
			}
		}
	}
	if eq{
		fmt.Printf("EQUAL")
	} else{
		fmt.Printf("NOT EQUAL")
	}
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