package api

import (
	"github.com/oze4/godaddygo/pkg/http"
)

// Gateway allows you to select different versions of the GoDaddy API
type Gateway interface {
	V1Getter
}

// versions implements APIInterface
type gateway struct {
	*http.Request
}

// V1 returns the V1 section of the GoDaddy API
func (g *gateway) V1() V1 {
	g.attach("1")
	return &v1{g.Request}
}

// V2 returns the V2 section of the GoDaddy API
func (g *gateway) V2() {
	g.attach("2")
	panic("V2 not implemented yet!")
}

func (g *gateway) attach(version string) {
	g.URL = g.URL + "/v" + version
}