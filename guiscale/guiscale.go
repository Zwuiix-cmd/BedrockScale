package guiscale

import (
	"github.com/kbinani/win"
	"unsafe"
)

type GuiScale struct{}

func (r GuiScale) SetGuiScale(h *Handler, number float64) {
	value := float32(number)
	var num win.DWORD
	var bytesWritten win.SIZE_T

	address := win.LPVOID(h.GameID() + 0x3E45030)

	win.VirtualProtectEx(h.Handle(), address, 4, 0x40, &num)
	win.WriteProcessMemory(h.Handle(), address, uintptr(unsafe.Pointer(&value)), 4, &bytesWritten)
}
