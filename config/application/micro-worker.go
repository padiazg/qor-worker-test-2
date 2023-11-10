package application

// MicroWorkerInterface micro worker interface
type MicroWorkerInterface interface {
	ConfigureWorker(*Application)
}
