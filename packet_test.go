package goz

import (
  "reflect"
  "strconv"
  "testing"
)

func TestNewPacket(t *testing.T) {
  route := NewRoute("test")
  segments := []string{"test"}

  packet := NewPacket(route, segments)

  if packet.index != 0 {
    t.Error("packet.index expected to be 0, got ", packet.index)
  }

  if packet.route != route {
    t.Error("packet.route expected equal supplied route")
  }

  if !reflect.DeepEqual(packet.segments, segments) {
    t.Error("packet.segments expected to supplied segments")
  }
}

func TestAddVariable(t *testing.T) {
  route := NewRoute("test")
  segments := []string{"test1", "test2"}

  packet := NewPacket(route, segments)
  packet.AddVariable()

  if packet.variableMap[strconv.Itoa(packet.index)] != segments[packet.index] {
    t.Error("packet.index expected to be ", packet.variableMap[strconv.Itoa(packet.index)],
      ", but got ", segments[packet.index])
  }
}
