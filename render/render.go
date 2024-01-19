package render

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	"github.com/make-42/cpu3d/utils"

	"github.com/StephaneBunel/bresenham"
	"github.com/fogleman/gg"
)

func RenderScreenLines(resolution utils.IntPair, lines *[]utils.CameraLine, outFile string) {
	var imgRect = image.Rect(0, 0, resolution.X, resolution.Y)
	var img = image.NewRGBA(imgRect)
	var bg = color.RGBA{41, 41, 41, 255}
	var pink = color.RGBA{255, 64, 105, 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{bg}, image.ZP, draw.Src)

	for _, line := range *lines {
		bresenham.DrawLine(img, int(line.A.X*float64(resolution.X)), int(line.A.Y*float64(resolution.Y)), int(line.B.X*float64(resolution.X)), int(line.B.Y*float64(resolution.Y)), pink)
	}
	toimg, _ := os.Create(outFile)
	defer toimg.Close()
	png.Encode(toimg, img)
}

/*
func RenderScreenLines(strokeWidth float64, resolution utils.IntPair, lines *[]utils.CameraLine, outFile string) {
	dc := gg.NewContext(resolution.X, resolution.Y)
	dc.SetRGB(0.161, 0.161, 0.161)
	dc.Clear()
	dc.SetRGB(1, 0.251, 0.412)
	dc.SetLineWidth(strokeWidth)
	for _, line := range *lines {
		dc.DrawLine(line.A.X*float64(resolution.X), line.A.Y*float64(resolution.Y), line.B.X*float64(resolution.X), line.B.Y*float64(resolution.Y))
	}
	dc.Stroke()
	dc.SavePNG(outFile)
}
*/

func RenderScreenPoints(strokeWidth float64, resolution utils.IntPair, pointCloudScreenCoords *[]utils.CameraCoords, outFile string) {
	dc := gg.NewContext(resolution.X, resolution.Y)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(strokeWidth)
	for _, coords := range *pointCloudScreenCoords {
		dc.DrawPoint(coords.X*float64(resolution.X), coords.Y*float64(resolution.Y), strokeWidth)
	}
	dc.Stroke()
	dc.SavePNG(outFile)
}
