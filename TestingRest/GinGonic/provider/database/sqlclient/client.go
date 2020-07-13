package sqlclient

import "database/sql"

var dbClient *sql.DB

type SqlClientInterface interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

type sqlClient struct{}

func init() {
	dbClient, _ = sql.Open("mysql", "conecctionString")
}

func (c *sqlClient) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return dbClient.Query(query, args)
}
