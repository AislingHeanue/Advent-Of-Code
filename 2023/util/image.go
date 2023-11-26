package util

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

var Image image.Image
var WindowClosureChan = make(chan struct{})
var WindowBeingUsed = false

type DrawInstance struct {
	Image image.Image
}

func (di *DrawInstance) Update() error {
	di.Image = Image
	if ebiten.IsWindowBeingClosed() {
		close(WindowClosureChan)
		return fmt.Errorf("window closed")
	}
	return nil
}

func (di *DrawInstance) Draw(screen *ebiten.Image) {
	screen.DrawImage(
		ebiten.NewImageFromImage(di.Image),
		&ebiten.DrawImageOptions{},
	)
}

func (di *DrawInstance) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return di.Image.Bounds().Dx(), di.Image.Bounds().Dy()
}

func EbitenSetup() {
	ebiten.SetWindowSize(400, 400)
	ebiten.SetWindowTitle("Output")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowClosingHandled(true)
	ebiten.SetTPS(60)
	di := DrawInstance{}
	di.Image = image.NewRGBA(image.Rect(0, 0, 50, 50))
	go ebiten.RunGame(&di)
	// go autoCloseWindow(1 * time.Second)
}

// func autoCloseWindow(delay time.Duration) {
// 	time.Sleep(delay)
// 	if !WindowBeingUsed {
// 		close(WindowClosureChan)
// 	}
// }
