package controllers

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/abarrios1/HomonitorBackend/models"
	"github.com/labstack/echo/v4"
)

// GetDevice ...
func GetDevice(ctx echo.Context) error {


	return ctx.JSON(http.StatusOK, "")
}

// GetAllDevices ...
func GetAllDevices(ctx echo.Context) error {


	return ctx.JSON(http.StatusOK, "")
}

// CreateDevice ...
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
func DeviceInterfaceEntryStats (ctx echo.Context) error {
	var deviceInterfaceEntryStats models.DeviceInterfaceEntryStats


	return ctx.JSON(http.StatusOK, deviceInterfaceEntryStats)
}
