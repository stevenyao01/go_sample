package main

import (
	"context"
	//"fmt"
	"os"
	"os/exec"
	"time"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	cmd := exec.CommandContext(ctx, "ls", "-l")
	//cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Stdout = os.Stdout
	cmd.Start()

	time.Sleep(100 * time.Second)
	// fmt.Println("退出程序中...", cmd.Process.Pid)
	fmt.Println("sdfkjsdkfl")
	cancel()

	cmd.Wait()
}
