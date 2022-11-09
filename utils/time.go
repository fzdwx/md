package utils

import (
	"fmt"
	"log"
	"time"
)

func TimeIt(f func() string) string {
	now := time.Now()
	res := f()
	duration := time.Now().Sub(now)
	val := fmt.Sprintf("cost: %s", duration)
	log.Println(val)
	return res
}
