package main

import (
	"fmt"
	"time"

	"github.com/jucassoli/localdate"
)

func main() {
	now := time.Now()

	res := localdate.Format(now)

	fmt.Printf("-- %v  -- > %T", res, now)
}
