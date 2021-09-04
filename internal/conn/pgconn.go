package conn

import (
	"crypto/tls"
	"encoding/binary"
	"net"

	"github.com/jackc/pgproto3"
	"golang.org/x/xerrors"
)

const SSLPROTO = 80877103

type PgConn interface {
	Send(query pgproto3.FrontendMessage) error
	Receive() (pgproto3.BackendMessage, error)

	ReqBackendSsl(tlscfg *tls.Config) error
}

type PgConnImpl struct {
	conn     net.Conn
	frontend *pgproto3.Frontend
}

func (pgconn *PgConnImpl) Send(query pgproto3.FrontendMessage) error {
	return pgconn.frontend.Send(query)
}

func (pgconn *PgConnImpl) Receive() (pgproto3.BackendMessage, error) {
	return pgconn.frontend.Receive()
}

func NewPgConn(netconn net.Conn) (PgConn, error) {

	pgconn := &PgConnImpl{
		conn: netconn,
	}

	var err error
	pgconn.frontend, err = pgproto3.NewFrontend(pgproto3.NewChunkReader(pgconn.conn), pgconn.conn)
	if err != nil {
		return nil, err
	}

	return pgconn, nil
}

func (pgconn *PgConnImpl) ReqBackendSsl(tlscfg *tls.Config) error {

	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, 8)
	// Gen salt
	b = append(b, 0, 0, 0, 0)
	binary.BigEndian.PutUint32(b[4:], SSLPROTO)

	_, err := pgconn.conn.Write(b)

	if err != nil {
		return xerrors.Errorf("ReqBackendSsl: %w", err)
	}

	resp := make([]byte, 1)

	if _, err := pgconn.conn.Read(resp); err != nil {
		return err
	}

	sym := resp[0]

	if sym != 'S' {
		return xerrors.New("SSL should be enabled")
	}

	pgconn.conn = tls.Client(pgconn.conn, tlscfg)
	return nil
}