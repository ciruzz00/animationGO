package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"time"
)

// ASCII chars from darkest to lightest
const asciiChars = " .:-=+*#%@"

// Convert pixel brightness to ASCII character
func brightnessToASCII(brightness float64) byte {
	scale := float64(len(asciiChars)-1) * brightness
	return asciiChars[int(scale)]
}

// Convert image to ASCII frame
func imageToASCII(img image.Image, width, height int) []string {
	bounds := img.Bounds()
	ascii := make([]string, height)

	for y := 0; y < height; y++ {
		line := ""
		for x := 0; x < width; x++ {
			// Calculate corresponding pixel position
			px := bounds.Min.X + x*bounds.Dx()/width
			py := bounds.Min.Y + y*bounds.Dy()/height

			// Get pixel color
			r, g, b, _ := img.At(px, py).RGBA()
			// Convert to grayscale (0-1)
			brightness := (float64(r)*0.299 + float64(g)*0.587 + float64(b)*0.114) / 65535.0
			line += string(brightnessToASCII(brightness))
		}
		ascii[y] = line
	}
	return ascii
}

// Clear terminal screen
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run ascii_animation.go <image_path>")
		os.Exit(1)
	}

	// Open image file
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Error opening image: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Decode image
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Printf("Error decoding image: %v\n", err)
		os.Exit(1)
	}

	// Terminal dimensions
	width := 80
	height := 40

	// Animation parameters
	frames := 30
	duration := 100 * time.Millisecond

	// Animation loop
	for i := 0; i < frames; i++ {
		clearScreen()

		// Create simple animation effect by shifting brightness
		shiftedImg := image.NewGray(img.Bounds())
		for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
			for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
				r, g, b, _ := img.At(x, y).RGBA()
				brightness := (float64(r)*0.299 + float64(g)*0.587 + float64(b)*0.114) / 65535.0
				// Create pulsing effect
				pulse := 0.5 + 0.5*(float64(i)/float64(frames))
				adjusted := brightness * pulse
				if adjusted > 1 {
					adjusted = 1
				}
				grayVal := uint8(adjusted * 255)
				shiftedImg.SetGray(x, y, color.Gray{Y: grayVal})
			}
		}

		// Convert to ASCII and display
		asciiFrame := imageToASCII(shiftedImg, width, height)
		for _, line := range asciiFrame {
			fmt.Println(line)
		}

		time.Sleep(duration)
	}
}

//Ciruzz00