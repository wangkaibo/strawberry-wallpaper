package bootstrap

import (
	"github.com/gin-gonic/gin"
)

type Bootstrap struct {
	Engine *gin.Engine
	AppName string
}

type Configurator func(bootstrap *Bootstrap)

func NewApp(appName string) *Bootstrap {
	b := &Bootstrap{
		AppName: appName,
		Engine: gin.Default(),
	}

	return b
}

// Configure accepts configurations and runs them inside the Bootstraper's context.
func (b *Bootstrap) Configure(cs ...Configurator) {
	for _, c := range cs {
		c(b)
	}
}