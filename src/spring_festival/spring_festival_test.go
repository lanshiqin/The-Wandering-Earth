package main

import (
	"sync"
	"testing"
)

// 测试加载程序
func TestLoader(t *testing.T) {
	var wg sync.WaitGroup
	Loader(&wg)
}

// 测试硬件系统
func TestHardware(t *testing.T) {
	var wg sync.WaitGroup
	Hardware(&wg)
}

// 测试撞针系统
func TestFiringPin(t *testing.T) {
	var wg sync.WaitGroup
	FiringPin(&wg)
}

// 测试火焰系统发射
func TestFlameLaunch(t *testing.T) {
	Flame(true)
}

// 测试火焰系统暂停
func TestFlamePause(t *testing.T) {
	Flame(false)
}

// 测试发动机底座
func TestEngine(t *testing.T) {
	Engine()
}

// 测试启动发动机
func TestStartEngine(t *testing.T) {
	var wg sync.WaitGroup
	StartEngine(&wg)
}
