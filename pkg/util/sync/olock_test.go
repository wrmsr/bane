package sync

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestOMutex(t *testing.T) {
	l := OMutex{}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(l.Lock(420))
			time.Sleep(1 * time.Millisecond)
		}
		time.Sleep(2 * time.Second)
		for i := 0; i < 3; i++ {
			fmt.Println(l.Unlock(420))
		}
		wg.Done()
	}()
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println(l.Lock(421))
		fmt.Println(l.Unlock(421))
		wg.Done()
	}()
	wg.Wait()
}
