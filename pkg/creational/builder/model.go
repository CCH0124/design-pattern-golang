package builder

import "fmt"

type ServerConfig struct {
	Protocol       string
	Port           int
	TLSEnabled     bool
	LogLevel       string
	MaxConnections int
}

func (s *ServerConfig) String() string {
	return fmt.Sprintf("Protocol: %s, Port: %d, TLSEnabled: %t, LogLevel: %s, MaxConnections: %d", s.Protocol, s.Port, s.TLSEnabled, s.LogLevel, s.MaxConnections)
}
