package builder

type ServerDirector struct {
	builder IServerBuilder
}

func NewServerDirector(builder IServerBuilder) *ServerDirector {
	return &ServerDirector{
		builder: builder,
	}
}

// allow change builder
func (d *ServerDirector) SetBuilder(b IServerBuilder) {
	d.builder = b
}

// encapsulation standard server build steps
func (d *ServerDirector) ConstructStandardServer() {
	d.builder.SetPort(8080).
		SetLogLevel("DEBUG").
		SetMaxConnections(1000)
}

// encapsulation secure server build steps
func (d *ServerDirector) ConstructSecureServer() {
	d.builder.SetPort(8443).
		SetTLS(true).
		SetLogLevel("WARN").
		SetMaxConnections(50000)
}
