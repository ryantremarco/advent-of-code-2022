package main

import (
	_ "embed"
	"errors"
)

//go:embed input
var input string

func main() {
	signal := input
	packetMarker, err := getMarker(signal, 4)
	if err != nil {
		panic(err)
	}

	messageMarker, err := getMarker(signal, 14)
	if err != nil {
		panic(err)
	}

	println("Packet Marker", packetMarker)
	println("Message Marker", messageMarker)
}

func getMarker(signal string, uniqueInMarker int) (int, error) {
	signalRunes := []rune(signal)
	for i := uniqueInMarker - 1; i < len(signalRunes); i++ {
		chars := make(map[rune]bool)
		for j := 0; j < uniqueInMarker; j++ {
			chars[signalRunes[i-j]] = true
		}
		if len(chars) == uniqueInMarker {
			return i + 1, nil
		}
	}

	return 0, errors.New("failed to find marker in signal")
}
