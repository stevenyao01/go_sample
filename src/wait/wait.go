package wait

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main1() {
	ctx, cancel := context.WithCancel(context.Background())
	cmd := exec.CommandContext(ctx, "bin/my_exec")
	//cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Stdout = os.Stdout
	cmd.Start()

	time.Sleep(10 * time.Second)
	fmt.Println("退出程序中...", cmd.Process.Pid)
	cancel()

	cmd.Wait()
}
