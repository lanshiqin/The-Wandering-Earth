package utils

import (
	"os"
	"os/exec"
	"runtime"
)

var clear map[string]func() // 清除终端方法

func init() {
	clear = make(map[string]func()) // 初始化一个map对象，key为方法名，value为一个方法
	clear["linux"] = func() {
		cmd := exec.Command("clear") // Linux系统执行清理终端的方法
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	}
	clear["darwin"] = func() {
		cmd := exec.Command("clear") // macOS系统执行清理终端的方法
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") // Windows系统执行清理终端的方法
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("暂不支持您的操作系统")
	}
}
