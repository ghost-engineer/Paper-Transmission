package main

import (
	"flag"
	"fmt"
	"math"
	"os"
)

func main() {
	message := flag.String("m", "", "Message to encrypt")
	output := flag.String("o", "crease_pattern.svg", "SVG output file for crease pattern")
	flag.Parse()

	if *message == "" {
		fmt.Println("Please provide a message to encrypt using -m.")
		os.Exit(1)
	}

	creasePattern := EncryptToCreasePattern(*message)
	err := os.WriteFile(*output, []byte(creasePattern), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Crease pattern saved to %s\n", *output)
}

// EncryptToCreasePattern encrypts a message into an SVG origami crease pattern
func EncryptToCreasePattern(msg string) string {
	// Simple example: each character is encoded as a line at a different angle
	creaseLines := ""
	w, h := 500, 500
	cx, cy := w/2, h/2
	angleStep := 360.0 / float64(len(msg)+1)
	for i := range msg {
		angle := angleStep * float64(i)
		rad := angle * math.Pi / 180.0
		x := float64(cx) + 200.0*math.Cos(rad)
		y := float64(cy) + 200.0*math.Sin(rad)
		creaseLines += fmt.Sprintf("<line x1=\"%d\" y1=\"%d\" x2=\"%.1f\" y2=\"%.1f\" stroke=\"black\" stroke-width=\"2\" />\n", cx, cy, x, y)
	}
	return fmt.Sprintf(`<svg width=\"%d\" height=\"%d\" xmlns=\"http://www.w3.org/2000/svg\">\n%s</svg>`, w, h, creaseLines)
}
