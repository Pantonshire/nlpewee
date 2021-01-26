package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/pantonshire/nlpewee/core"
	pb "github.com/pantonshire/nlpewee/proto"
	"github.com/pantonshire/nlpewee/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type config struct {
	Port           uint          `yaml:"port"`
	ConnectTimeout time.Duration `yaml:"connect_timeout"`
	TLS            tlsConfig     `yaml:"tls"`
}

type tlsConfig struct {
	Enabled bool   `yaml:"enabled"`
	Crt     string `yaml:"crt"`
	Key     string `yaml:"key"`
}

type options struct {
	ConfigPath string `short:"c" long:"config" required:"yes" description:"Path to the config file to use"`
}

func main() {
	var opts options
	if _, err := flags.Parse(&opts); err != nil {
		if flagErr, ok := err.(*flags.Error); ok {
			if flagErr.Type == flags.ErrHelp {
				os.Exit(0)
			} else {
				os.Exit(1)
			}
		} else {
			panic(err)
		}
	}

	bytes, err := ioutil.ReadFile(opts.ConfigPath)
	if err != nil {
		panic(err)
	}

	var conf config
	if err := yaml.Unmarshal(bytes, &conf); err != nil {
		panic(err)
	}

	core.Init()

	address := fmt.Sprintf(":%d", conf.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	fmt.Println("Created TCP listener at " + address)

	var serverOptions []grpc.ServerOption

	if conf.ConnectTimeout > 0 {
		serverOptions = append(serverOptions, grpc.ConnectionTimeout(conf.ConnectTimeout))
	}

	if conf.TLS.Enabled {
		creds, err := credentials.NewServerTLSFromFile(conf.TLS.Crt, conf.TLS.Key)
		if err != nil {
			panic(err)
		}
		serverOptions = append(serverOptions, grpc.Creds(creds))
	}

	grpcServer := grpc.NewServer(serverOptions...)
	nlpeweeServer := server.NLPeweeServer{}
	pb.RegisterNLPeweeServer(grpcServer, nlpeweeServer)

	fatal := make(chan error)

	go func() {
		fmt.Println("Starting server")
		if err := grpcServer.Serve(listener); err != nil {
			fatal <- err
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(interrupt)

	select {
	case <-interrupt:
		fmt.Println("Interrupted, stopping server")
		grpcServer.GracefulStop()
		fmt.Println("Stopped server")
	case err := <-fatal:
		panic(err)
	}
}
