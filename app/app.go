package app

import (
	"net"

	reuse "github.com/libp2p/go-reuseport"
	shhttp "github.com/pg-sharding/spqr/http"
	"github.com/pg-sharding/spqr/internal"
	"github.com/pg-sharding/spqr/internal/config"
	"github.com/wal-g/tracelog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	spqr *internal.Spqr
}

func NewApp(sg *internal.Spqr) *App {
	return &App{
		spqr: sg,
	}
}

// TODO split into separate apps?
func (app *App) ProcPG() error {
	////	listener, err := net.Listen("tcp", "man-a6p8ynmq7hanpybg.db.yandex.net:6432")
	config := config.GetAppConfig()
	listener, err := reuse.Listen(config.PROTO, config.Addr)
	if err != nil {
		return err
	}
	defer listener.Close()
	tracelog.InfoLogger.Printf("ProcPG listening %s by %s", config.Addr, config.PROTO)
	return app.spqr.Run(listener)
}

func (app *App) ProcADM() error {
	//	listener, err := net.Listen("tcp", "man-a6p8ynmq7hanpybg.db.yandex.net:7432")
	config := config.GetAppConfig()
	listener, err := net.Listen(config.PROTO, config.ADMAddr)
	if err != nil {
		return err
	}
	defer listener.Close()
	tracelog.InfoLogger.Printf("ProcADM listening %s by %s", config.ADMAddr, config.PROTO)
	return app.spqr.RunAdm(listener)
}

func (app *App) ServHttp() error {
	config := config.GetAppConfig()
	serv := grpc.NewServer()
	shhttp.Register(serv)
	reflection.Register(serv)
	listener, err := net.Listen("tcp", config.HttpAddr)
	if err != nil {
		return err
	}
	tracelog.InfoLogger.Printf("ServHttp listening %s by tcp", config.HttpAddr)
	return serv.Serve(listener)
}
