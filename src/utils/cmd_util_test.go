package utils

import (
	"fmt"
	"testing"
	"time"
)

// 测试清除终端数据
func TestCallClear(t *testing.T) {
	fmt.Println("2秒后清除终端数据内容")
	time.Sleep(2 * time.Second)
	CallClear()
	fmt.Println("清除完成！")
}
