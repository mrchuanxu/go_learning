package main

import "fmt"

type g2048 [4][4]int

func (t *g2048) MirrorV(){
	tn := new(g2048)
	for i,line := range t{
		for j,num := range line{
			tn[len(t)-i-1][j] = num
		}
	}
	*t = *tn
}
func (t *g2048) Left90() {
    tn := new(g2048)
    for i, line := range t {
        for j, num := range line {
            tn[len(line)-j-1][i] = num
        }
    }
    *t = *tn
}

func (t *g2048) Right90(){
	tn := new(g2048)
	for i,line := range t{
		for j,num := range line{
			tn[j][len(t)-i-1] = num
		}
	}
	*t = *tn
}

func (g *g2048) R90(){
	tn := new(g2048)
	for x,line := range g{
		for y,_ := range line{
			tn[x][y]= g[len(line)-1-y][x]
		}
	}
	*g = *tn
}

func (t *g2048) Right180(){
	tn := new(g2048)
	for i,line := range t{
		for j,num := range line{
			tn[len(line)-i-1][len(line)-j-1] = num
		}
	}

	*t = *tn
}

func (t *g2048)Print(){
	for _,line := range t {
		for _,number := range line{
			fmt.Printf("%2d",number)
		}
		fmt.Println()
	}
	fmt.Println()
	tn:=g2048{{1,2,3,4},{5,8},{9,10,11},{13,14,16}}
	*t = tn
}




func main(){
	fmt.Println("origin")
    t := g2048{{1, 2, 3, 4}, {5, 8}, {9, 10, 11}, {13, 14, 16}}
    t.Print()
    fmt.Println("mirror")
    t.MirrorV()
    t.Print()
    fmt.Println("Left90")
    t.Left90()
    t.Print()
    fmt.Println("Right90")
    t.R90()
    t.Print()
    fmt.Println("Right180")
    t.Right180()
    t.Print()
}