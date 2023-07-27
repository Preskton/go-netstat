//go:build darwin
// +build darwin

package netstat

import (
	"errors"
)

const (
	pathTCPTab  = "/proc/net/tcp"
	pathTCP6Tab = "/proc/net/tcp6"
	pathUDPTab  = "/proc/net/udp"
	pathUDP6Tab = "/proc/net/udp6"

	ipv4StrLen = 8
	ipv6StrLen = 32
)

// Socket states
const (
	Established SkState = 0x01
	SynSent             = 0x02
	SynRecv             = 0x03
	FinWait1            = 0x04
	FinWait2            = 0x05
	TimeWait            = 0x06
	Close               = 0x07
	CloseWait           = 0x08
	LastAck             = 0x09
	Listen              = 0x0a
	Closing             = 0x0b
)

var skStates = [...]string{
	"UNKNOWN",
	"ESTABLISHED",
	"SYN_SENT",
	"SYN_RECV",
	"FIN_WAIT1",
	"FIN_WAIT2",
	"TIME_WAIT",
	"", // CLOSE
	"CLOSE_WAIT",
	"LAST_ACK",
	"LISTEN",
	"CLOSING",
}

// Errors returned by gonetstat
var (
	ErrNotEnoughFields = errors.New("gonetstat: not enough fields in the line")
)


// doNetstat - collect information about network port status
func doNetstat(path string, fn AcceptFn) ([]SockTabEntry, error) {
	return nil, errors.New("not implemented")
}

// TCPSocks returns a slice of active TCP sockets containing only those
// elements that satisfy the accept function
func osTCPSocks(accept AcceptFn) ([]SockTabEntry, error) {
	return doNetstat(pathTCPTab, accept)
}

// TCP6Socks returns a slice of active TCP IPv4 sockets containing only those
// elements that satisfy the accept function
func osTCP6Socks(accept AcceptFn) ([]SockTabEntry, error) {
	return doNetstat(pathTCP6Tab, accept)
}

// UDPSocks returns a slice of active UDP sockets containing only those
// elements that satisfy the accept function
func osUDPSocks(accept AcceptFn) ([]SockTabEntry, error) {
	return doNetstat(pathUDPTab, accept)
}

// UDP6Socks returns a slice of active UDP IPv6 sockets containing only those
// elements that satisfy the accept function
func osUDP6Socks(accept AcceptFn) ([]SockTabEntry, error) {
	return doNetstat(pathUDP6Tab, accept)
}
