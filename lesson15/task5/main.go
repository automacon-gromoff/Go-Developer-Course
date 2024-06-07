package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
)

type temperature struct {
	mu   sync.RWMutex
	temp string
}

type Meteo interface {
	ReadTemp() string
	ChangeTemp(v string)
}

func (t *temperature) ReadTemp() string {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.temp
}

func (t *temperature) ChangeTemp(v string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.temp = v
}

func main() {
	temp := temperature{
		mu:   sync.RWMutex{},
		temp: "0",
	}
	var meteo Meteo = &temp
	wg := sync.WaitGroup{}
	for i := 0; i < 50; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			newTemp := fmt.Sprint(rand.IntN(60) - 30)
			meteo.ChangeTemp(newTemp)
		}()
		go func() {
			defer wg.Done()
			fmt.Printf("Температура окружающей среды %s°C\n", meteo.ReadTemp())
		}()
	}
	wg.Wait()
}
