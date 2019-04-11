package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for i := 0; i < 101; i++ {
			fmt.Println("[程序加载进度>>>>>>>>>>>>" + strconv.Itoa(i) + "%]")
			time.Sleep(time.Millisecond * 100)
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			_ = cmd.Run()

		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("程序加载完毕，发动机正在启动...")
}
