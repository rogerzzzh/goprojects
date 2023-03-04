package persist

import "log"

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	counter := 0
	go func() {
		for item := range out {
			counter++
			log.Printf("Saver Info: Got #%d item to save, %s", counter, item)
		}
	}()
	return out
}
