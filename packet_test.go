package goz

import (
	"reflect"
	"strconv"
	"testing"
)

func TestNewPacket(t *testing.T) {
	expectedPacketIndex := 0
	expectedRoute := NewRoute("test")
	expectedSegments := []string{"test"}

	packet := NewPacket(expectedRoute, expectedSegments)

	if packet.index != expectedPacketIndex {
		t.Errorf("packet.index expected to be %d, got %d", expectedPacketIndex, packet.index)
	}

	if packet.route != expectedRoute {
		t.Error("packet.route expected equal supplied route")
	}

	if !reflect.DeepEqual(packet.segments, expectedSegments) {
		t.Error("packet.segments expected to supplied segments")
	}
}

func TestAddVariable(t *testing.T) {
	expectedRoute := NewRoute("test")
	expectedSegments := []string{"test1", "test2"}

	packet := NewPacket(expectedRoute, expectedSegments)
	packet.AddVariable()

	if packet.variableMap[strconv.Itoa(packet.index)] != expectedSegments[packet.index] {
		t.Errorf("packet.index expected to be %s but got %s",
			packet.variableMap[strconv.Itoa(packet.index)], expectedSegments[packet.index])
	}
}
