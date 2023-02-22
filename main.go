package main

import (
	"fmt"
	"github.com/TheTitanrain/w32"
	"github.com/kbinani/win"
	"github.com/logrusorgru/aurora"
	"github.com/mattn/go-colorable"
	"github.com/zwuiix-cmd/guiscale-client/guiscale"
	"golang.org/x/sys/windows"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	log.SetOutput(colorable.NewColorableStdout())

}

var (
	c = make(chan os.Signal)
)

func main() {
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)

	handler()
}

func handler() *guiscale.Handler {
	h := guiscale.New()

	win.SetConsoleTitle("LSD GuiScale - Diango#0001")
	win.SetConsoleIcon(win.HICON(h.Handle()))

	fmt.Println()
	fmt.Println(aurora.Green("██████╗ ██╗   ██╗    ██╗     ███████╗██████╗ \n██╔══██╗╚██╗ ██╔╝    ██║     ██╔════╝██╔══██╗\n██████╔╝ ╚████╔╝     ██║     ███████╗██║  ██║\n██╔══██╗  ╚██╔╝      ██║     ╚════██║██║  ██║\n██████╔╝   ██║       ███████╗███████║██████╔╝\n╚═════╝    ╚═╝       ╚══════╝╚══════╝╚═════╝ \n                                             "))
	fmt.Println()

	open := func() bool {
		return w32.FindWindowW(nil, windows.StringToUTF16Ptr("Minecraft")) != 0
	}

	for !open() {
		fmt.Println(aurora.Red("Sorry, we could not find your game, we will open your game, then reopen the software..."))
		cmd := exec.Command("explorer.exe", "shell:appsFolder\\Microsoft.MinecraftUWP_8wekyb3d8bbwe!App")
		if err := cmd.Run(); err != nil {
			os.Exit(-1)
		}

		for !open() {
			time.Sleep(time.Second)
		}
		time.Sleep(time.Second * 1)
		os.Exit(-1)
	}
	lunch(h)
	return h
}

func lunch(h *guiscale.Handler) {
	start(h)
}

func start(h *guiscale.Handler) {
	guiscale.GuiScale{}.SetGuiScale(h, 7.0)
	fmt.Println(aurora.Green("Injection done, the size of the interface has been reduced!"))
	time.Sleep(time.Second * 10)
}
