package internal

import (
	"github.com/jackc/pgproto3"
	"github.com/pg-sharding/spqr/internal/qrouter"
	"github.com/wal-g/tracelog"
)

func frontend(rt qrouter.Qrouter, cl Client, cmngr ConnManager) error {

	tracelog.InfoLogger.Printf("process frontend for user %s %s", cl.Usr(), cl.DB())

	rst := &RelayState{
		ActiveShards: nil,
		TxActive:     false,
	}

	for {
		msg, err := cl.Receive()
		if err != nil {
			tracelog.ErrorLogger.Printf("failed to recieve msg %w", err)
			return err
		}

		tracelog.InfoLogger.Printf("recieved msg %v", msg)

		switch v := msg.(type) {
		case *pgproto3.Query:
			if cmngr.ValidateReRoute(rst) {
				tracelog.InfoLogger.Println("Try reroute")
				if err := rst.reroute(rt, cl, cmngr, v); err != nil {
					return err
				}
			}

			var txst byte
			if txst, err = cl.ProcQuery(v); err != nil {
				return err
			}

			if err := rst.completeRelay(cl, cmngr, txst); err != nil {
				return err
			}

		default:
		}
	}
}
