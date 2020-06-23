package controllers

import (
	"bytes"
	"github.com/abarrios1/HomonitorBackend/lib"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/abarrios1/HomonitorBackend/models"
	"github.com/gocql/gocql"
	"github.com/labstack/echo/v4"
)

// GetDevice ... Collect individual device information
func GetDevice(ctx echo.Context) error {
	cqlHost := os.Getenv("CASSANDRA_HOST")
	cluster := gocql.NewCluster(cqlHost)
	cluster.Keyspace = "Homonitor"
	cluster.Consistency = gocql.Quorum
	session, _ := cluster.CreateSession()
	defer session.Close()

	session.Query(`SELECT * FROM device WHERE mac_address = ?`, "ValueInsertFromCtx").Iter()

	return ctx.JSON(http.StatusOK, "")
}

// GetAllDevices ... Collect all devices information
func GetAllDevices(ctx echo.Context) error {
	cqlHost := os.Getenv("CASSANDRA_HOST")
	cluster := gocql.NewCluster(cqlHost)
	cluster.Keyspace = "Homonitor"
	cluster.Consistency = gocql.Quorum
	session, _ := cluster.CreateSession()
	defer session.Close()

	session.Query(`SELECT * FROM device`).Iter()

	return ctx.JSON(http.StatusOK, "")
}

// CreateDevice ... Create a new device and insert into C*
func CreateDevice(ctx echo.Context) error {
	device := new(models.Device)

	// Read body content
	var bodyBytes []byte
	if ctx.Request().Body != nil {
		bodyBytes, _ := ioutil.ReadAll(ctx.Request().Body)
	} else {
		return ctx.JSON(http.StatusNoContent, device)
	}

	ctx.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	er := ctx.Bind(device)
	if er != nil {
		panic(er)
	}

	session, err := lib.ConnectCassandraCluster()
	defer session.Close()

	// insert a tweet
	if err := session.Query(`INSERT INTO device (hostname, device_type, ip_address, serial_number, mac_address) VALUES (?, ?, ?, ?, ?)`,
		"", "", "", "", "").Exec(); err != nil {
		log.Fatal(err)
	}

	return ctx.JSON(http.StatusOK, device)
}

// UpdateDevice ...
func UpdateDevice(ctx echo.Context) error {

	return ctx.JSON(http.StatusOK, "")
}

// DeleteDevice ...
func DeleteDevice(ctx echo.Context) error {

	return ctx.JSON(http.StatusOK, "")
}

// DeviceInterfaceEntryStats ...
func DeviceInterfaceEntryStats(ctx echo.Context) error {
	var deviceInterfaceEntryStats models.DeviceInterfaceEntryStats

	return ctx.JSON(http.StatusOK, deviceInterfaceEntryStats)
}
