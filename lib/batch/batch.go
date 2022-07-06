package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	res = make([]user, n)

	const paralel = 10

	var wg sync.WaitGroup

	sem := make(chan struct{}, paralel)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			sem <- struct{}{}
			getOne(pool)
			<-sem
			wg.Done()
		}()
	}

	wg.Wait()
	return res
}

//func work() {
//	time.Sleep(1 * time.Second)
//	fmt.Println("Work is done")
//}
