package internal

import (
	"bufio"
	"crypto/tls"
	"log"
	"os"

	"github.com/jackc/pgproto3"
	"github.com/pg-sharding/spqr/internal/config"
	"github.com/wal-g/tracelog"
)

type FakeClient struct {
}

func NewFakeClient() *FakeClient {
	return &FakeClient{}
}

func (f FakeClient) Server() Server {
	return nil
}

func (f FakeClient) Unroute() {
}

func (f FakeClient) AssignRule(rule *config.FRRule) {
}

func (f FakeClient) AssignRoute(r *Route) {
}

func (f FakeClient) AssignServerConn(srv Server) {
}

func (f FakeClient) ReplyErr(errnsg string) error {
	return nil
}

func (f FakeClient) Init(cfg *tls.Config, reqssl bool) error {
	return nil
}

func (f FakeClient) Auth() error {
	return nil
}

func (f FakeClient) StartupMessage() *pgproto3.StartupMessage {
	return nil
}

func (f FakeClient) Usr() string {
	return defaultUsr
}

func (f FakeClient) DB() string {
	return defaultDB
}

func (f FakeClient) PasswordCT() string {
	return ""
}

func (f FakeClient) PasswordMD5() string {
	return ""
}

func (f FakeClient) DefaultReply() error {
	return nil
}

func (f FakeClient) Route() *Route {
	return nil
}

func (f FakeClient) Rule() *config.FRRule {
	return nil
}

func (f FakeClient) ProcQuery(query *pgproto3.Query) (byte, error) {
	return 0, nil
}

func (f FakeClient) Send(msg pgproto3.BackendMessage) error {
	return nil
}

func (f FakeClient) Receive() (pgproto3.FrontendMessage, error) {
	return &pgproto3.Query{}, nil
}

var _ Client = &FakeClient{}

type Executer struct {
	initSqlPath string
}

func NewExecuter(initSqlPath string) *Executer {
	return &Executer{
		initSqlPath,
	}
}

func (e *Executer) ReadCmds() []string {
	f, err := os.Open(e.initSqlPath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	ret := make([]string, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}

	return ret
}

func (e *Executer) SPIexec(console Console, cl Client) error {
	for _, cmd := range e.ReadCmds() {
		tracelog.InfoLogger.Printf("executing init sql cmd %s", cmd)
		if err := console.processQ(cmd, cl); err != nil {
			tracelog.InfoLogger.PrintError(err)
		}
	}

	return nil
}
