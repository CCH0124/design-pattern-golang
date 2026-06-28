package builder

type GRPCServerBuilder struct {
	config ServerConfig
}

func NewGRPCServerBuilder() IServerBuilder {
	return &GRPCServerBuilder{
		config: ServerConfig{
			Protocol: "gRPC",
		},
	}
}

func (g *GRPCServerBuilder) SetPort(port int) IServerBuilder {
	g.config.Port = port
	return g
}

func (g *GRPCServerBuilder) SetTLS(enabled bool) IServerBuilder {
	g.config.TLSEnabled = enabled
	return g
}

func (g *GRPCServerBuilder) SetLogLevel(level string) IServerBuilder {
	g.config.LogLevel = level
	return g
}

func (g *GRPCServerBuilder) SetMaxConnections(max int) IServerBuilder {
	g.config.MaxConnections = max
	return g
}

func (g *GRPCServerBuilder) GetResult() ServerConfig {
	return g.config
}
