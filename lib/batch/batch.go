package batch

import (
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

	// channel to receive the users
	ch := make(chan user)

	// create a goroutine for each user
	for i := int64(0); i < pool; i++ {
		go func(i int64) {
			for j := i; j < n; j += pool {
				ch <- getOne(j)
			}
		}(i)
	}

	// collect the users
	for i := int64(0); i < n; i++ {
		res = append(res, <-ch)
	}

	return res
}
