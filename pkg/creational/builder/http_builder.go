package builder

type HTTPServerBuilder struct {
	config ServerConfig
}

func NewHTTPServerBuilder() IServerBuilder {
	return &HTTPServerBuilder{
		config: ServerConfig{
			Protocol: "HTTP",
		},
	}
}

func (h *HTTPServerBuilder) SetPort(port int) IServerBuilder {
	h.config.Port = port
	return h
}

func (h *HTTPServerBuilder) SetTLS(enabled bool) IServerBuilder {
	h.config.TLSEnabled = enabled
	return h
}

func (h *HTTPServerBuilder) SetLogLevel(level string) IServerBuilder {
	h.config.LogLevel = level
	return h
}

func (h *HTTPServerBuilder) SetMaxConnections(max int) IServerBuilder {
	h.config.MaxConnections = max
	return h
}

func (h *HTTPServerBuilder) GetResult() ServerConfig {
	return h.config
}
