package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("estBissextile(2000) =>", estBissextile(2000))
	fmt.Println("estPremier(11) =>", estPremier(11))
	fmt.Println("premiersNombresPremiers(5) =>", premiersNombresPremiers(5))

	tableau := initTableau(5)
	tableau2 := initTableau(30)
	toSearch := rand.Intn(100)

	fmt.Println("initTableau(5) =>", tableau)
	fmt.Println("bubbleSort(tableau) =>", bubbleSort(tableau))
	fmt.Println("selectionSort(tableau2) =>", selectionSort(tableau2))
	fmt.Printf("dicoSearch(tableau2, %d) => %v\n", toSearch, dicoSearch(tableau2, toSearch))

}

func estBissextile(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func estPremier(n int) bool {
	if n <= 1 {
		return false
	}
	for i := n - 1; i > 1; i-- {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func premiersNombresPremiers(n int) []int {

	var premiers []int

	for i := 3; len(premiers) <= n; i++ {
		if estPremier(i) {
			premiers = append(premiers, i)
		}
	}

	return premiers
}

func initTableau(n int) []int {
	var tableau []int
	for i := 0; i < n; i++ {
		tableau = append(tableau, rand.Intn(100))
	}
	return tableau
}

func bubbleSort(tab []int) []int {
	for i := 0; i < len(tab); i++ {
		for j := 0; j < len(tab)-1; j++ {
			if tab[j] > tab[j+1] {
				tab[j], tab[j+1] = tab[j+1], tab[j]
			}
		}
	}

	return tab
}

func selectionSort(tab []int) []int {
	n := len(tab)
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if tab[j] < tab[min] {
				min = j
			}
		}
		tab[i], tab[min] = tab[min], tab[i]
	}

	return tab
}

func dicoSearch(tab []int, val int) bool {
	n := len(tab)
	if n == 0 {
		return false
	}
	mid := n / 2
	if tab[mid] == val {
		return true
	}
	if tab[mid] < val {
		return dicoSearch(tab[mid+1:], val)
	}
	return dicoSearch(tab[:mid], val)
}
