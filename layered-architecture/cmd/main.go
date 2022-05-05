package main

import (
	backend "github.com/mmmommm/go-sample"
)

func main() {
	entry, cleanup, err := backend.NewEntryPoint()
	log := entry.Logger
	if err != nil {
		log.Errorw("init error",
			"msg", err.Error())
	}
	defer cleanup()
	log.Infow("Server running...",
		"addr", entry.Config.ServerAddr(),
	)
	if err := entry.Srv.Start(entry.Config.ServerAddr()); err != nil {
		log.Errorw("server error",
			"msg", err.Error())
	}
}