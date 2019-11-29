package main

import "fmt"

func main()  {
	var n, m, q int

	fmt.Scan(&n, &m)

	Links := make([][]int, m)
	Signals := make([]string, m)
	Alphabet := make([]string, m)
	alphsize := 0
	lambdacounter := 0

	for i := range Links{
		Links[i] = make([]int, 2)
		for j := range Links[i]{
			fmt.Scan(&Links[i][j])
		}
		fmt.Scan(&Signals[i])
		if Signals[i] == "lambda"{
			lambdacounter++
		} else{
			unic := true
			for k := 0; k < alphsize; k++{
				if Alphabet[k] == Signals[i]{
					unic = false
					break
				}
			}
			if unic{
				Alphabet[alphsize] = Signals[i]
				alphsize++
			}
		}
	}

	Final := make([]int, n)

	for i := range Final{
		fmt.Scan(&Final[i])
	}

	fmt.Scan(&q)
	//fmt.Printf("\n")
	//for i := 0; i < alphsize; i++{
	//	fmt.Printf("%s\n", Alphabet[i])
	//}

	Tr := make([][][]int, n)
	for i := range Tr{
		Tr[i] = make([][]int, alphsize + 1)
		for j := range Tr[i]{
			Tr[i][j] = make([]int, n)
			for k := range Tr[i][j]{
				Tr[i][j][k] = -1
			}
		}
	}

	for i := range Links{
		pos := 0
		if Signals[i] == "lambda"{
			pos = alphsize
		} else {
			for j := range Alphabet {
				if Alphabet[j] == Signals[i] {
					pos = j
					break
				}
			}
		}
		for j := range Tr[Links[i][0]][pos]{
			if Tr[Links[i][0]][pos][j] == -1{
				Tr[Links[i][0]][pos][j] = Links[i][1]
				break
			}
		}
	}

	//for i := range Tr{
	//	for j := range Tr[i]{
	//		for k := range Tr[i][j]{
	//			if Tr[i][j][k] == -1{
	//				break
	//			}
	//			fmt.Printf("%d ", Tr[i][j][k])
	//		}
	//		fmt.Printf("| ")
	//	}
	//	fmt.Printf("\n")
	//}

	Result := Det(Alphabet, Tr, Final, q)
	//s := '"'
	fmt.Printf("digraph {\n\trankdir = LR\n")
	fmt.Printf("\tdummy [label = \"\", shape = none]\n")
	for i := 0; i < Result[3][0][0]; i++{
		//0 [label = "[0]", shape = circle]
		fmt.Printf("\t%d [label = \"[", i)
		for j := range Result[0][i]{
			if Result[0][i][j] == -1{
				break
			}
			if j == 0{
				fmt.Printf("%d", Result[0][i][j])
			} else {
				fmt.Printf(" %d", Result[0][i][j])
			}
		}
		if Result[2][0][i] == 1{
			fmt.Printf("]\", shape = doublecircle]\n")
		} else {
			fmt.Printf("]\", shape = circle]\n")
		}
	}
	fmt.Printf("\tdummy -> %d\n", 0)
	for i := 0; i < Result[3][0][1]; i++{
		if Result[1][i][0] == -2{
			continue
		}
		fmt.Printf("\t%d -> %d [label = \"%s", Result[1][i][0], Result[1][i][1], Alphabet[Result[1][i][2]])
		for j := i + 1; j < Result[3][0][1]; j++{
			if (Result[1][i][0] == Result[1][j][0]) && (Result[1][i][1] == Result[1][j][1]){
				fmt.Printf(", %s", Alphabet[Result[1][j][2]])
				Result[1][j][0] = -2
			}
		}
		fmt.Printf("\"]\n")
	}
	fmt.Printf("}")
}

func Closure(Tr [][][]int, z []int) []int {
	c := make([]int, len(z))
	for i := range c {
		c[i] = -1
	}

	for i := range z {
		if z[i] == -1 {
			break
		}
		Dfs(Tr, z[i], c)
	}
	//fmt.Printf("Closure Ans = [")
	for i := range c{
		if c[i] == -1{
			break
		}
		//fmt.Printf(" %d", c[i])
	}
	//fmt.Printf("]\n")
	return c
}

func Dfs(Tr [][][]int, q int, c []int) []int {
	unic := true
	pos := -1
	for i := range c{
		if c[i] == -1{
			pos = i
			break
		}
		if c[i] == q{
			unic = false
			break
		}
	}
	//fmt.Printf("%d %d\n", q, pos)
	if unic{
		c[pos] = q
		for i := range Tr[q][len(Tr[q]) - 1]{
			if Tr[q][len(Tr[q]) - 1][i] == -1{
				break
			}
			Dfs(Tr, Tr[q][len(Tr[q]) - 1][i], c)
		}
	}
	return c
}

func Det(Alphabet []string, Tr [][][]int, Final []int, q int) [][][]int {
	tmp := make([]int, len(Tr))
	tmp[0] = q
	for i := 1; i < len(tmp); i++{
		tmp[i] = -1
	}

	q0 := Closure(Tr, tmp)
	Sort(q0)
	//for i := range q0{
	//	fmt.Printf("%d ", q0[i])
	//}

	Answer := make([][][]int, 4)

	//Answer[0] = Q
	Answer[0] = make([][]int , len(Tr)*len(Alphabet))
	//Answer[1] = DELTA
	Answer[1] = make([][]int , len(Tr)*len(Tr)*len(Alphabet))
	//Answer[2] = New Finals
	Answer[2] = make([][]int, 1)
	Answer[2][0] = make([]int, len(Tr)*len(Alphabet))
	//Answer[3] = True Sizes
	Answer[3] = make([][]int, 1)
	Answer[3][0] = make([]int, 2)
	Answer[3][0][0] = 0
	Answer[3][0][1] = 0

	for i := range Answer[2][0]{
		Answer[2][0][i] = 0
	}

	Answer[0][Answer[3][0][0]] = q0
	Answer[3][0][0]++

	Stack := make([]int, len(Tr)*len(Tr)*len(Alphabet))
	stacksize := 0
	Units := make([][]int, len(Tr)*len(Tr)*len(Alphabet))
	unitsize := 0
	//for i := range Units{
	//	Units[i] = make([]int, len(Tr))
	//}
	Units[unitsize] = q0
	Stack[stacksize] = unitsize
	unitsize++
	stacksize++

	for stacksize > 0{
		//fmt.Printf("NEW ITERATION WITH VERTEX = [")
		stacksize--
		z := Units[Stack[stacksize]]
		//for kk := range z{
		//	if z[kk] == -1{
		//		break
		//	}
		//	fmt.Printf(" %d", z[kk])
		//}
		//fmt.Printf("]\n")

		pos_z_in_Q := 0

		for j := 0; j < Answer[3][0][0]; j++{
			tst := true
			for k := range Answer[0][j]{
				if Answer[0][j][k] == -1{
					if z[k] != -1{
						tst = false
					}
					break
				}
				if z[k] != Answer[0][j][k]{
					tst = false
					break
				}
			}
			if tst{
				pos_z_in_Q = j
				break
			}
		}

		for i := range z{
			if z[i] == -1{
				break
			}
			if Final[z[i]] == 1{
				//for j := range z{
				//	if z[j] == -1{
				//		break
				//	}
				//	Answer[2][0][z[j]] = 1
				//}
				//fmt.Printf("POSITION = %d\n", pos_z_in_Q)
				Answer[2][0][pos_z_in_Q] = 1
				break
			}
		}

		for i := 0; i < len(Tr[0]) - 1; i++{
			verts := make([]int, len(Tr)*len(Tr))
			curr := 0
			for j := range verts{
				verts[j] = -1
			}

			for j := range z{
				if z[j] == -1{
					break
				}
				//fmt.Printf(" -- %d --\n", z[j])
				for k := range Tr[z[j]][i]{
					if Tr[z[j]][i][k] == -1{
						break
					}
					verts[curr] = Tr[z[j]][i][k]
					curr++
				}
			}

			z1 := Closure(Tr, verts)
			Sort(z1)
			unic := true
			for j := 0; j < Answer[3][0][0]; j++{
				tst := true
				for k := range Answer[0][j]{
					if Answer[0][j][k] == -1{
						if z1[k] != -1{
							tst = false
						}
						break
					}
					if z1[k] != Answer[0][j][k]{
						tst = false
						break
					}
				}
				if tst{
					unic = false
					break
				}
			}
			if unic{
				//fmt.Printf("NEW VERTEX -> [")
				//for kk := range z1{
				//	if z1[kk] == -1{
				//		break
				//	}
				//	fmt.Printf("%d ", z1[kk])
				//}
				//fmt.Printf("]\n")

				Answer[0][Answer[3][0][0]] = z1
				Answer[3][0][0]++

				Units[unitsize] = z1
				Stack[stacksize] = unitsize
				unitsize++
				stacksize++
			}
			Answer[1][Answer[3][0][1]] = make([]int, 3)
			Answer[1][Answer[3][0][1]][0] = pos_z_in_Q

			pos_z1_in_Q := 0

			for j := 0; j < Answer[3][0][0]; j++{
				tst := true
				for k := range Answer[0][j]{
					if Answer[0][j][k] == -1{
						if z1[k] != -1{
							tst = false
						}
						break
					}
					if z1[k] != Answer[0][j][k]{
						tst = false
						break
					}
				}
				if tst{
					pos_z1_in_Q = j
					break
				}
			}

			Answer[1][Answer[3][0][1]][1] = pos_z1_in_Q
			Answer[1][Answer[3][0][1]][2] = i
			Answer[3][0][1]++
		}
	}
	return Answer
}

func Sort(arr []int)  {
	pos := 0
	len_true := 0

	//for i := range arr{
	//	if arr[i] == -1{
	//		break
	//	}
	//	fmt.Printf("%d ", arr[i])
	//}
	//fmt.Printf("\n")

	for i := range arr{
		if arr[i] == -1{
			len_true = i
			break
		}
	}
	for pos < len_true - 1{
		j := pos + 1
		for j > 0{
			if arr[j] < arr[j - 1]{
				tmp := arr[j]
				arr[j] = arr[j - 1]
				arr[j - 1] = tmp
			}
			j--
		}
		pos++
	}
}
