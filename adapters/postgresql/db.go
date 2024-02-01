package postgresql

import "github.com/lib/pq"

type DB struct {
	client pq.Driver
}

func (DB) Implementation() string { return "Postgresql" }
