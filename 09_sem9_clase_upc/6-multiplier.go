package main

import "fmt"

var end chan bool

func zero(row int, west chan float64) {
	for i := 0; i < row; i++ {
		west <- 0.0
	}
	close(west)
}

func result(ans [][]float64, i int, east chan float64){
	j := 0
	for value := range east {
		ans[i][j] = value
		j++
	}
	end <- true
}

func source(datarow []float64, south chan float64) {
	for _, value := range datarow {
		south <- value
	}
	close(south)
}

func sink(north chan float64) {
	for range north {
	}
}

func multiplier(firstElement float64, north, east, south, west chan float64) {
	cnt := 0
	for {
		secondElement := <- north
		sum := <- east
		
		sum = sum + firstElement*secondElement
		
		south <- secondElement
		west <- sum

		cnt += 1
		if cnt == 3 {
			break
		}
	}
	close(south)
	close(west)
}

func printMatrix(title string, matrix [][]float64) {
	fmt.Println(title)

	for _, v := range matrix {
		fmt.Println(v)
	}
}

func main() {
	a := [][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	b := [][]float64{{1, 0, 2}, {0, 1, 2}, {1, 0, 0}}
	ans := [][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}

	row := len(a)
	col := len(a[0])

	// channel end
	end = make(chan bool)

	// south channels (4x3)
	sch := make([][]chan float64, row+1)
	for i := range sch {
		sch[i] = make([]chan float64, col)
		for j := range sch[i] {
			sch[i][j] = make(chan float64)
		}
	}

	// west channels (3x4)
	wch := make([][]chan float64, row)
	for i := range wch {
		wch[i] = make([]chan float64, col+1)
		for j := range wch[i] {
			wch[i][j] = make(chan float64)
		}
	}

	// zero & result processes
	for i := 0; i < row; i++ {
		go zero(row, wch[i][col])
		go result(ans, i, wch[i][0])
	}

	// source & sink processes
	for j := 0; j < col; j++ {
		go source(b[j], sch[0][j])
		go sink(sch[row][j])
	}

	// multiplier processes
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			go multiplier(a[i][j], // value
				sch[i][j],   // North
				wch[i][j+1], // East
				sch[i+1][j], // South
				wch[i][j])   // West
		}
	}

	for i := 0; i < row; i++ {
		<-end
	}
	printMatrix("Matrix A", a)
	printMatrix("Matrix B", b)
	printMatrix("Multiplication result", ans)
}