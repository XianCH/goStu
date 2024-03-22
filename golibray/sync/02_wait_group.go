package syncTest

import (
	"log"
	"sync"
)

func TestWait() {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(&wg)
	go worker(&wg)
	wg.Wait()

}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	var sum int
	for i := 0; i < 50; i++ {
		sum += 1
	}
	log.Println(sum)
}

func palyg(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		log.Println("palyg say hey to :", i)
	}
}
