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
			time.Sleep(time.Millisecond * 50)
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			_ = cmd.Run()
			fmt.Println("[程序加载进度>>>>>>>>>>>>" + strconv.Itoa(i) + "%]")

		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("程序加载完毕，发动机正在启动...")
	wg.Add(1)
	go func() {
		time.Sleep(time.Second * 1)
		fmt.Println("硬件系统自检完成...")
		wg.Done()
	}()
	wg.Wait()

	wg.Add(1)
	go func() {
		time.Sleep(time.Second * 1)
		fmt.Println("撞针系统自检完成...")
		wg.Done()
	}()
	wg.Wait()

	wg.Add(1)
	go func() {
		time.Sleep(time.Second * 1)
		fmt.Println("发动机正在启动...")
		wg.Done()
	}()
	wg.Wait()

	wg.Add(1)
	go func() {
		for i := 0; i <= 10; i++ {
			time.Sleep(time.Second * 1)
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			_ = cmd.Run()
			fmt.Printf("行星转向发动机正在启动中............%d/100\n", i*10)
		}
		wg.Done()
	}()
	wg.Wait()

}
