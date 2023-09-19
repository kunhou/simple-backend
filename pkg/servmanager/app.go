package servmanager

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// Application is the main struct of the application
type Application struct {
	servers []Server
	signals []os.Signal
}

// Option is the option of the application
type Option func(application *Application)

// NewApp creates a new Application
func NewApp(ops ...Option) *Application {
	app := &Application{}
	for _, op := range ops {
		op(app)
	}

	// default accept signals
	if len(app.signals) == 0 {
		app.signals = []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT}
	}
	return app
}

// WithServer application add server
func WithServer(servers ...Server) func(application *Application) {
	return func(application *Application) {
		application.servers = servers
	}
}

// WithSignals application add listen signals
func WithSignals(signals []os.Signal) func(application *Application) {
	return func(application *Application) {
		application.signals = signals
	}
}

// Run application run
func (app *Application) Run(ctx context.Context) error {
	if len(app.servers) == 0 {
		return nil
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, app.signals...)
	errCh := make(chan error, 1)

	for _, s := range app.servers {
		go func(srv Server) {
			if err := srv.Start(); err != nil {
				log.Printf("failed to start server, err: %s", err)
				errCh <- err
			}
		}(s)
	}

	select {
	case err := <-errCh:
		_ = app.Stop()
		return err
	case <-ctx.Done():
		return app.Stop()
	case <-quit:
		return app.Stop()
	}
}

// Stop application stop
func (app *Application) Stop() error {
	wg := sync.WaitGroup{}
	for _, s := range app.servers {
		wg.Add(1)
		go func(srv Server) {
			defer wg.Done()
			if err := srv.Shutdown(); err != nil {
				log.Printf("failed to stop server, err: %s", err)
			}
		}(s)
	}
	// wait all server graceful shutdown
	wg.Wait()
	return nil
}
