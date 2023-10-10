package ctx

import (
	"context"
	"fmt"
	"time"
)

func main() {
	userId := 10
	now := time.Now()
	ctx := context.Background()
	value, err := fetchData(ctx, userId)
	fmt.Println(value, err)
	fmt.Println(time.Since(now))
}

func fetchData(ctx context.Context, userId int) (int, error) {

	value, err := fetchDataSlowly(userId)
	if err != nil {
		return 0, fmt.Errorf("error happened while fetching")
	}

	return value, err
}

func fetchDataSlowly(userId int) (int, error) {
	time.Sleep(time.Millisecond * 500)
	return userId, nil
}
