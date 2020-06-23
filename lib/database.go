package lib

import (
	"github.com/gocql/gocql"
	"os"
)

// ConnectCassandraCluster ... Establish connection to C* Cluster
func ConnectCassandraCluster() (*gocql.Session, error) {
	cqlHost := os.Getenv("CASSANDRA_HOST")
	cluster := gocql.NewCluster(cqlHost)
	cluster.Keyspace = "Homonitor"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()

	return session, err
}
