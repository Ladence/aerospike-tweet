package storage

import aero "github.com/aerospike/aerospike-client-go"

func NewAero() *aero.Client {
	client, err := aero.NewClient("127.0.0.1", 3000)
	if err != nil {
		return nil
	}
	return client
}
