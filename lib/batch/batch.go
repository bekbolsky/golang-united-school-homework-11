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

// getBatch concurrently loads n users from the database.
// The number of goroutines used to load the users is determined by the pool parameter.
// The function returns a slice of users.
// The function should return a slice of users with the same length as the n parameter.
func getBatch(n int64, pool int64) (res []user) {
	// create a slice of users with n capacity
	res = make([]user, 0, n)

	// create variable to hold the wait group
	// and add the number of goroutines to it
	// so that we know when all goroutines have finished
	var wg sync.WaitGroup

	// add a number of goroutines to the wait group
	// the number of goroutines is determined by the pool parameter
	wg.Add(int(pool))

	// iterate over the number of goroutines
	for i := 0; i < int(pool); i++ {
		// create a gorroutine that loads batch of users
		go func(i int) {
			// iterate over the batch of users
			// the number of users is determined by the n parameter
			// the number of users loaded by each goroutine is determined by the pool parameter
			for j := int64(i); j < n; j += int64(pool) {
				// append the result of getOne() to the slice
				// this will block until the goroutine is finished
				res = append(res, getOne(j))
			}
			// when the goroutine is done, decrement the wait group
			wg.Done()
		}(i)
	}
	// wait for all goroutines to finish
	// this blocks the main goroutine
	// until all goroutines have finished
	wg.Wait()

	return res
}
