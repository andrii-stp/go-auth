package service_manager

import "context"

// Resource interface is designed to be as an abstraction layer for services to have resources running
// in separate goroutines (e.g. telemetry, cache, message queue)
type ServiceManager interface {
	Start()
	Stop(ctx context.Context)
}

func Run(services []ServiceManager) {
	for _, service := range services {
		go service.Start()
	}
}
