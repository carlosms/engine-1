package main

import (
	"net"
	"strings"

	flags "github.com/jessevdk/go-flags"
	"github.com/sirupsen/logrus"
	api "github.com/src-d/engine-cli/api"
	"github.com/src-d/engine-cli/cmd/srcd-server/engine"
	grpc "google.golang.org/grpc"
)

var version = "undefined"

func main() {
	var options struct {
		Addr    string `long:"address" short:"a" default:"0.0.0.0:4242"`
		Workdir string `long:"workdir" short:"w" default:""`
	}

	_, err := flags.Parse(&options)
	if err != nil {
		logrus.Fatal(err)
	}

	workdir := strings.TrimSpace(options.Workdir)
	if workdir == "" {
		logrus.Fatal("No work directory provided!")
	}

	l, err := net.Listen("tcp", options.Addr)
	if err != nil {
		logrus.Fatal(err)
	}

	srv := grpc.NewServer()
	api.RegisterEngineServer(srv, engine.NewServer(version, workdir))

	logrus.Infof("listening on %s", options.Addr)
	if err := srv.Serve(l); err != nil {
		logrus.Fatal(err)
	}
}