package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getTargetPath() string {
	return "/"
}

func main() {
	target := getTargetPath()
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("OneRM by FengPwner\n")
	fmt.Printf("==================\n")

	fmt.Printf("确认递归删除 %s 及其所有内容 ? [Y/N]: ", target)
	answer, err := reader.ReadString('\n')
	if err != nil && answer == "" {
		fmt.Println("读取输入失败:", err)
		return
	}

	if strings.TrimSpace(strings.ToUpper(answer)) != "Y" {
		fmt.Println("已取消删除。")
		return
	}

	if err := os.RemoveAll(target); err != nil {
		fmt.Println("删除失败:", err)
		return
	}

	fmt.Printf("已成功递归删除: %s\n", target)
}

