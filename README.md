# animationGO

A simple Go program that converts an image into ASCII art and animates it in the terminal with a pulsing effect.

## Features
- Supports JPEG and PNG images
- Converts images to ASCII art using a grayscale character scale
- Creates a pulsing animation effect in the terminal
- Configurable terminal output size (default: 80x40 characters)

## Prerequisites
- [Go](https://golang.org/dl/) (version 1.16 or higher)
- An image file (JPEG or PNG format)

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/ciruzz00/animationGO.git
   cd animationGO
   ```
2. Ensure Go is installed and configured properly.

## Usage
1. Place an image file (e.g., `image.jpg` or `image.png`) in the project directory.
2. Run the program with the image path as an argument:
   ```bash
   go run main.go image.jpg
   ```
   Or
   ```bash
   go build && ./animationGO image.jpg
   ```
3. The program will display an ASCII animation of the image in the terminal.

## Configuration
You can modify the following constants in `ascii_animation.go` to customize the output:
- `width` and `height`: Set the terminal output dimensions (default: 80x40).
- `frames`: Number of animation frames (default: 30).
- `duration`: Time between frames (default: 100ms).

## Example
```bash
./animationGO sample.png
```
This will display a 3-second animation of `sample.png` converted to ASCII art with a pulsing brightness effect.

