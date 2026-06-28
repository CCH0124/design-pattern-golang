package builder

type IServerBuilder interface {
	SetPort(port int) IServerBuilder
	SetTLS(enabled bool) IServerBuilder
	SetLogLevel(level string) IServerBuilder
	SetMaxConnections(max int) IServerBuilder
	GetResult() ServerConfig
}
