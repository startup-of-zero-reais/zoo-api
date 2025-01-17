package providers

import (
	"github.com/goravel/framework/contracts/foundation"
	"github.com/goravel/framework/facades"

	"github.com/startup-of-zero-reais/zoo-api/app/grpc"
	"github.com/startup-of-zero-reais/zoo-api/routes"
)

type GrpcServiceProvider struct {
}

func (receiver *GrpcServiceProvider) Register(app foundation.Application) {
	// Add Grpc interceptors
	kernel := grpc.Kernel{}
	facades.Grpc().UnaryServerInterceptors(kernel.UnaryServerInterceptors())
	facades.Grpc().UnaryClientInterceptorGroups(kernel.UnaryClientInterceptorGroups())
}

func (receiver *GrpcServiceProvider) Boot(app foundation.Application) {
	// Add routes
	routes.Grpc()
}
