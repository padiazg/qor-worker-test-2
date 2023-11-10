package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/padiazg/qor-worker-test/app"
	a "github.com/padiazg/qor-worker-test/config/application"
	settings "github.com/padiazg/qor-worker-test/config/settings"
)

var (
	ctx, cancel = context.WithCancel(context.Background())
	osSignal    = make(chan os.Signal, 1)
	done        = make(chan bool)
	stng        = &settings.Settings{}
)

func main() {
	signal.Notify(osSignal,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	if err := stng.Read(); err != nil {
		log.Fatalf("reading settings: %+v", err)
	}

	application := a.New(&a.Config{
		HostPort: 9090,
		Domain:   "http://localhost:9090",
		Settings: stng,
	})

	app.Mount(application, ctx)

	go application.Run()

	for {
		select {
		case <-done:
			log.Println("Exiting...")
			if application.Server != nil {
				if err := application.Shutdown(ctx); err != nil {
					log.Fatalf("shutting down server: %+v", err)
				}
			}
			return
		case <-osSignal:
			log.Println("Ctrl+C pressed")
			close(done)
		}
	}
}
