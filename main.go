package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	f, err := os.Create("cpu.pprof")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if err := pprof.StartCPUProfile(f); err != nil {
		fmt.Println(err.Error())
		return
	}
	defer pprof.StopCPUProfile()

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("Hello from goroutine")
			time.Sleep(1 * time.Second)
		}()
	}

	fmt.Println("Main function exited.")
	time.Sleep(2 * time.Second) // ゴルーチンが完了する前にメイン関数が終了しないようにする
}
