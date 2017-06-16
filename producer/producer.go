package main

import (
	"github.com/xtracdev/pgconn"
	"log"
	"github.com/xtracdev/synthpub/synthevent"
	"github.com/xtracdev/goes"
	"database/sql"
)

func makeEvent()(*goes.Event, error) {
	synthEvent, err := synthevent.NewSyntheticEvent()
	if err != nil {
		return nil, err
	}

	event, err := synthEvent.ToGoESEvent()
	if err != nil {
		return nil, err
	}

	return event,nil
}

func storeEventInPublishTable(db *sql.DB, event * goes.Event) error {
	_,err := db.Exec("insert into t_aepb_publish (aggregate_id, version, typecode, payload) values ($1, $2, $3, $4)",
				event.Source, event.Version, event.TypeCode, event.Payload)
	return err
}

func main() {
	eventConfig, err := pgconn.NewEnvConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	pgdb, err := pgconn.OpenAndConnect(eventConfig.ConnectString(), 3)
	if err != nil {
		log.Fatal(err.Error())
	}


	event, err := makeEvent()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = storeEventInPublishTable(pgdb.DB, event)
	if err != nil {
		log.Fatal(err.Error())
	}


	pgdb.DB.Close()
}
