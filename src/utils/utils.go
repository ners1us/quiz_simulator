package utils

import (
	"fmt"
	"os"
	"time"
)

func Exit(msg string) {
	fmt.Println(msg)

	os.Exit(1)
}

func CheckTime(timeVal *int) error {
	if *timeVal < 0 {
		return fmt.Errorf("incorrect value of time: %d", *timeVal)
	}
	return nil
}

func StartTimer(seconds int) *time.Timer {
	return time.NewTimer(time.Duration(seconds) * time.Second)
}
