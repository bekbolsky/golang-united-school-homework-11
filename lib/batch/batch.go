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

// getBatch takes two arguments - the number of users and the number of goroutines in which users will concurrently load.
// It returns an array of received users.
func getBatch(n int64, pool int64) (res []user) {

	ch := make(chan user)

	for i := int64(0); i < pool; i++ {
		go func(id int64) {
			ch <- getOne(id)
		}(i)
	}

	for i := int64(0); i < n; i++ {
		res = append(res, <-ch)
	}

	return res
}
