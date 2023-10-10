package main

import (
	"context"
	"fmt"
	"time"
)

type Response struct {
	value int
	err   error
}

func main() {
	userId := 10
	now := time.Now()
	ctx := context.Background()
	value, err := fetchData(ctx, userId)
	fmt.Println(value, err)
	fmt.Println(time.Since(now))
}

func fetchData(ctx context.Context, userId int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()
	resp := make(chan Response)
	go func() {
		value, err := fetchDataSlowly(userId)
		resp <- Response{
			value: value,
			err:   err,
		}
	}()
	for {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("took too long")
		case resp := <-resp:
			if resp.err != nil {
				return 0, fmt.Errorf("error happened while fetching")
			} else {
				return resp.value, resp.err
			}
		}
	}
}

func fetchDataSlowly(userId int) (int, error) {
	time.Sleep(time.Millisecond * 100)
	return userId, nil
}
