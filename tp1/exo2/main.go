package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	conoway := initEmptyConoway(16, 16)

	// Glinder
	generateGlider(conoway, 0, 0)

	printer := map[int]string{
		0: " ",
		1: "#",
	}

	for i := 0; i < 50; i++ {
		printConoway(conoway, printer)
		conoway = update(conoway)
		time.Sleep(250 * time.Millisecond)
	}

}

func initConoway(n int, m int) [][]int {
	var grille [][]int
	for i := 0; i < n; i++ {
		var ligne []int
		for j := 0; j < m; j++ {
			ligne = append(ligne, rand.Intn(2))
		}
		grille = append(grille, ligne)
	}
	return grille
}

func initEmptyConoway(n int, m int) [][]int {
	var grille [][]int
	for i := 0; i < n; i++ {
		var ligne []int
		for j := 0; j < m; j++ {
			ligne = append(ligne, 0)
		}
		grille = append(grille, ligne)
	}
	return grille
}

func countNeighbour(conoway [][]int, i int, j int) int {
	count := 0
	for x := i - 1; x <= i+1; x++ {
		for y := j - 1; y <= j+1; y++ {
			if x >= 0 && y >= 0 && x < len(conoway) && y < len(conoway[0]) && (x != i || y != j) {
				if conoway[x][y] == 1 {
					count += conoway[x][y]
				}
			}
		}
	}
	return count
}

func update(conoway [][]int) [][]int {
	var newConoway [][]int
	for i := 0; i < len(conoway); i++ {
		var ligne []int
		for j := 0; j < len(conoway[0]); j++ {
			count := countNeighbour(conoway, i, j)
			if conoway[i][j] == 1 {
				if count < 2 || count > 3 {
					ligne = append(ligne, 0)
				} else {
					ligne = append(ligne, 1)
				}
			} else {
				if count == 3 {
					ligne = append(ligne, 1)
				} else {
					ligne = append(ligne, 0)
				}
			}
		}
		newConoway = append(newConoway, ligne)
	}

	return newConoway
}

func printConoway(conoway [][]int, printer map[int]string) {

	for i := 0; i < len(conoway[0]); i++ {
		fmt.Print("---")
	}

	fmt.Println()

	for i := 0; i < len(conoway); i++ {
		for j := 0; j < len(conoway[0]); j++ {
			fmt.Print(" ", printer[conoway[i][j]], " ")
		}
		fmt.Println()
	}

	for i := 0; i < len(conoway[0]); i++ {
		fmt.Print("---")
	}
	fmt.Println()
}

func generateGlider(conoway [][]int, x int, y int) {
	if x+2 < len(conoway) && y+2 < len(conoway[0]) {
		conoway[x][y+1] = 1
		conoway[x+1][y+2] = 1
		conoway[x+2][y] = 1
		conoway[x+2][y+1] = 1
		conoway[x+2][y+2] = 1
	}
}

func generateBlinker(conoway [][]int, x int, y int) {
	if x+2 < len(conoway) && y < len(conoway[0]) {
		conoway[x][y] = 1
		conoway[x+1][y] = 1
		conoway[x+2][y] = 1
	}
}
