package concurrency

import "sync"

type concurrency struct {

}
func task(){}
func semaphore (){
	var wg sync.WaitGroup
	sem := make(chan struct{}, 3)
	for i :=0; i< 10; i++ {
		wg.Add(1)
		sem <- struct{}{}
		go func() {
			defer wg.Done()
			defer func(){
				<-sem
				task()
			}()
		}()
	}
	wg.Wait()
	close(sem)
}