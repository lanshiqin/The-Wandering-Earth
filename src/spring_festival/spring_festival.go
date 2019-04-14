package main

import (
	"The-Wandering-Earth/src/utils"
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	// 加载程序
	Loader(&wg)
	// 硬件系统自检
	Hardware(&wg)
	// 撞针系统自检
	FiringPin(&wg)
	// 启动发动机
	StartEngine(&wg)
}

// 加载程序
func Loader(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		for i := 0; i < 101; i++ {
			time.Sleep(time.Millisecond * 50)
			utils.CallClear()
			fmt.Println("[程序加载进度>>>>>>>>>>>>" + strconv.Itoa(i) + "%]")

		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("程序加载完毕...")
}

// 硬件系统
func Hardware(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		time.Sleep(time.Second * 1)
		fmt.Println("硬件系统自检完成...")
		wg.Done()
	}()
	wg.Wait()
}

// 撞针系统
func FiringPin(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		time.Sleep(time.Second * 1)
		fmt.Println("撞针系统自检完成...")
		wg.Done()
	}()
	wg.Wait()
}

// 启动发动机
func StartEngine(wg *sync.WaitGroup) {
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
			utils.CallClear()
			fmt.Printf("行星转向发动机正在启动中............%d/100\n", i*10)
		}
		wg.Done()
	}()
	wg.Wait()
	wg.Add(1)
	go func() {
		for i := 0; i < 24; i++ {
			if i%2 == 0 {
				time.Sleep(time.Second * 2)
				utils.CallClear()
				Flame(true)
			} else {
				time.Sleep(time.Second * 1)
				utils.CallClear()
				Flame(false)
			}
			Engine()
		}
		wg.Done()
	}()
	wg.Wait()
}

// 火焰系统
func Flame(status bool) {
	var flameLaunch = "!!!!!" // 火焰发射状态
	var flamePause = "     "  // 火焰暂停状态
	var flameHeight = 18      // 火焰高度
	var currentFlame = flamePause
	if status {
		currentFlame = flameLaunch
	}
	for i := 0; i < flameHeight; i++ {
		fmt.Printf("                 %s               \n", currentFlame)
	}
}

// 发动机底座
func Engine() {
	fmt.Println("              \\__     __/           ")
	fmt.Println("               |_______|           ")
	fmt.Println("               |       |           ")
	fmt.Println("               |       |           ")
	fmt.Println("               |       |           ")
	fmt.Println("               |       |           ")
	fmt.Println("               |_______|           ")
	fmt.Println("              /_________\\           ")
	fmt.Println("            /_____________\\           ")
	fmt.Println("           |_______________|           ")
}
