package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	inChan := make(chan int)  // Channel du maître vers ses ouvriers
	outChan := make(chan int) // Channel des ouvriers vers le maître

	nbOuvrier := 10

	for i := 1; i <= nbOuvrier; i++ {
		wg.Add(1)
		go maitreGiveRandom(inChan)
		go ouvrierWork(outChan, <-inChan, i)
		go fmt.Println(time.Now(), "| Ouvrier N°", i, "a envoyé le nombre", <-outChan)
		go wg.Done()
	}

	wg.Wait()

}

func maitreGiveRandom(inC chan int) {
	inC <- rand.Intn(10)
}

func ouvrierWork(outC chan int, nb int, ouvrier int) {
	fmt.Println(time.Now(), "| Ouvrier N°", ouvrier, "a reçu le nombre", nb)
	outC <- nb * nb
}
