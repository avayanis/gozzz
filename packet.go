package goz

import (
	"strconv"
)

// Packet describes a GoApp request as it routes
type Packet struct {
	index       int
	segments    []string
	route       *Route
	variableMap map[string]string
}

// NewPacket constructs and returns an initialized Packet
func NewPacket(route *Route, segments []string) *Packet {
	packet := new(Packet)

	packet.index = 0
	packet.route = route
	packet.segments = segments

	packet.variableMap = make(map[string]string)

	return packet
}

// AddVariable inserts a new varname into the variableMap at the current index.
func (packet *Packet) AddVariable() {
	index := strconv.Itoa(packet.index)

	if _, ok := packet.variableMap[index]; !ok {
		packet.variableMap[index] = packet.segments[0]
	}
}
