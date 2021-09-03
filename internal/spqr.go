package internal

import (
	"crypto/tls"
	"net"

	"github.com/pg-sharding/spqr/internal/config"
	"github.com/pg-sharding/spqr/internal/qrouter"
	"github.com/pkg/errors"
	"github.com/wal-g/tracelog"
)

type Spqr struct {
	Router  *Router
	Qrouter *qrouter.QrouterImpl
	SPIexecuter *Executer
}

const defaultProto = "tcp"

func NewSpqr(initialSqlPath string) (*Spqr, error) {
	qrouter := qrouter.NewR()
	router, err := NewRouter(qrouter)
	if err != nil {
		return nil, errors.Wrap(err, "NewRouter")
	}

	shardMapping := config.GetRoutingConfig().ShardMapping
	tracelog.InfoLogger.Printf("%v", shardMapping)

	for name, shard := range shardMapping {
		if shard.TLSCfg.ReqSSL {
			cert, err := tls.LoadX509KeyPair(shard.TLSCfg.CertFile, shard.TLSCfg.KeyFile)
			if err != nil {
				return nil, errors.Wrap(err, "failed to make route failure resp")
			}
			tlscfg := &tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
			tracelog.InfoLogger.Printf("initialising shard tls config for %s", name)

			if err := shard.Init(tlscfg); err != nil {
				return nil, err
			}
		}
		if shard.Proto == "" {
			shard.Proto = defaultProto
		}

		tracelog.InfoLogger.FatalOnError(qrouter.AddShard(name, &shard))
	}

	executer := NewExecuter(initialSqlPath)

	executer.SPIexec(router.ConsoleDB, NewFakeClient())

	return &Spqr{
		Router:      router,
		Qrouter:     qrouter,
		SPIexecuter: executer,
	}, nil
}

func (sg *Spqr) serv(netconn net.Conn) error {

	client, err := sg.Router.PreRoute(netconn)
	if err != nil {
		return err
	}

	cmngr, err := InitClConnection(client)
	if err != nil {
		return err
	}

	return frontend(sg.Qrouter, client, cmngr)
}

func (sg *Spqr) Run(listener net.Listener) error {
	for {
		conn, _ := listener.Accept()

		go func() {
			if err := sg.serv(conn); err != nil {
				tracelog.ErrorLogger.PrintError(err)
			}
		}()
	}
}

func (sg *Spqr) servAdm(netconn net.Conn) error {
	return sg.Router.ServeConsole(netconn)
}

func (sg *Spqr) RunAdm(listener net.Listener) error {
	for {
		conn, err := listener.Accept()
		if err != nil {
			return errors.Wrap(err, "RunAdm failed")
		}
		go func() {
			if err := sg.servAdm(conn); err != nil {
				tracelog.ErrorLogger.PrintError(err)
			}
		}()
	}
}
