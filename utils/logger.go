package utils

import (
	"fmt"
	"time"
)

func Log(message string) {
	fmt.Println("[" + time.Now().String() + "] LOG  ==> " + message)
}

func Warn(message string) {
	fmt.Println("[" + time.Now().String() + "] WARNING  ==> " + message)
}

func Error(message string) {
	fmt.Println("[" + time.Now().String() + "] ERROR  ==> " + message)
}
