package sqip
import (
	"image"
)
// Run takes a file and primitve related config properties and creates a SVG-based LQIP image.
func Run(file string, workSize, count, mode, alpha, repeat, workers int, background string) (out string, width, height int, err error) {
	// Load image
	image, err := LoadImage(file)
	if err != nil {
		return "", 0, 0, err
	}

	return RunLoaded(image, workSize, count, mode, alpha, repeat, workers, background)
}

//RunLoaded takes an already loaded image and config properties to generate an SVG LQIP
func RunLoaded(image image.Image, workSize, count, mode, alpha, repeat, workers int, background string) (out string, width, height int, err error) {
	// Use image-size to retrieve the width and height dimensions of the input image
	// We need these sizes to pass to Primitive and to write the SVG viewbox
	w, h := ImageWidthAndHeight(image)
	// Since Primitive is only interested in the larger dimension of the input image, let's find it
	outputSize := largerOne(w, h)

	// create primitive
	svg, err := Primitive(image, workSize, outputSize, count, mode, alpha, repeat, workers, background)
	if err != nil {
		return "", 0, 0, err
	}

	// resize BG to match original
	// Ensures agreement between viewBox and actual content
	svg = Refit(svg, w, h)

	// minify svg
	svg, err = Minify(svg)
	if err != nil {
		return "", 0, 0, err
	}

	// blur the svg
	svg, err = Blur(svg, w, h)
	if err != nil {
		return "", 0, 0, err
	}
	return svg, w, h, nil
}
