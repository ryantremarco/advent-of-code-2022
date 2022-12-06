package main

import (
	"fmt"
	"testing"
)

func TestGetMarker(t *testing.T) {
	exampleSignals := []struct {
		signal        string
		packetMarker  int
		messageMarker int
	}{
		{
			"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			7,
			19,
		},
		{
			"bvwbjplbgvbhsrlpgdmjqwftvncz",
			5,
			23,
		},
		{
			"nppdvjthqldpwncqszvftbrmjlhg",
			6,
			23,
		},
		{
			"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			10,
			29,
		},
		{
			"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			11,
			26,
		},
	}

	for _, exampleSignal := range exampleSignals {
		t.Run(fmt.Sprintf("signal %s packet marker", exampleSignal.signal), func(t *testing.T) {
			got, err := getMarker(exampleSignal.signal, 4)
			if err != nil {
				t.Fatalf("unexpected err: %s", err)
			}
			if exampleSignal.packetMarker != got {
				t.Fatalf("expected packetMarker of %d, got %d", exampleSignal.packetMarker, got)
			}

		})

		t.Run(fmt.Sprintf("signal %s message marker", exampleSignal.signal), func(t *testing.T) {
			got, err := getMarker(exampleSignal.signal, 14)
			if err != nil {
				t.Fatalf("unexpected err: %s", err)
			}
			if exampleSignal.messageMarker != got {
				t.Fatalf("expected messageMarker of %d, got %d", exampleSignal.messageMarker, got)
			}

		})
	}
}
