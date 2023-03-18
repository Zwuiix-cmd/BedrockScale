package main

import (
	"LSD-Scale/scale"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/TheTitanrain/w32"
	"github.com/kbinani/win"
	"golang.org/x/sys/windows"
	"net/url"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

var isToggledGuiScale bool = false

var (
	c = make(chan os.Signal)
)

func main() {
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	handler()
}

func handler() *scale.Handler {
	h := scale.New()

	win.SetConsoleTitle("")
	win.SetConsoleIcon(win.HICON(h.Handle()))

	open := func() bool {
		return w32.FindWindowW(nil, windows.StringToUTF16Ptr("Minecraft")) != 0
	}

	for !open() {
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

	start(h)
	return h
}

func start(h *scale.Handler) error {
	a := app.New()
	w := a.NewWindow("LSD-Scale - Zwuiix#0001")

	var URL, _ = url.Parse("https://discord.gg/mpuWfXXwbD")
	message := widget.NewHyperlink("Discord", URL)

	separatorGuiScale := widget.NewSeparator()
	toggleGuiScale := widget.NewCheck("Toggle Small GuiScale", func(bool2 bool) {
		isToggledGuiScale = bool2
		if isToggledGuiScale {
			scale.GuiScale{}.SetGuiScale(h, 7.0)
		} else {
			scale.GuiScale{}.SetGuiScale(h, 3.0)
		}
	})

	w.SetContent(container.NewVBox(
		message,
		separatorGuiScale,
		toggleGuiScale,
	))

	w.Resize(fyne.NewSize(400, 100))
	w.SetFixedSize(true)
	w.SetMaster()
	w.CenterOnScreen()
	w.ShowAndRun()

	return nil
}
