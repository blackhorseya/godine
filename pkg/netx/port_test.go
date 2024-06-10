package netx

import (
	"net"
	"strconv"
	"testing"
)

func TestGetAvailablePortReturnsValidPort(t *testing.T) {
	port := GetAvailablePort()
	if port < _startPort || port > _endPort {
		t.Errorf("Invalid port: got %v, want a port between %v and %v", port, _startPort, _endPort)
	}

	// Check if the port is really available
	address := ":" + strconv.Itoa(port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		t.Errorf("Port is not available: got error %v", err)
	}
	listener.Close()
}
