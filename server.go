package control

import (
	"net/http"

	"github.com/labstack/echo"
)

// Routes ...
func Routes(e *echo.Echo) {
	e.GET("/units", allUnits)
	e.DELETE("/units", removeAll)

	e.GET("/unit/:id", info)
	e.POST("/unit/:id", addUnit)
	e.DELETE("/unit/:id", removeUnit)

	e.POST("/unit/:id/start", start)
	e.POST("/unit/:id/stop", stop)
	e.POST("/unit/:id/restart", restart)
	e.POST("/unit/:id/reload", reload)

	e.POST("/units/start", startAll)
	e.POST("/units/stop", stopAll)
	e.POST("/units/restart", restartAll)
	e.POST("/units/reload", reloadAll)
}

// addUnit ...
// not support current!
func addUnit(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, errMessage("add unit current is not supported"))
}

// removeUnit ...
func removeUnit(c echo.Context) error {
	id := c.Param("id")
	m := c.(*Machine)
	m.RemoveUnit(id)
	return c.JSON(http.StatusOK, map[string]string{})
}

// info ...
func info(c echo.Context) error {
	id := c.Param("id")
	m := c.(*Machine)
	info, err := m.UnitInfo(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, errMessage(err.Error()))
	}
	return c.JSON(http.StatusOK, info)
}

// start ...
func start(c echo.Context) error {
	id := c.Param("id")
	m := c.(*Machine)
	m.UnitStart(id)
	return c.JSON(http.StatusOK, map[string]string{})
}

// stop ...
func stop(c echo.Context) error {
	id := c.Param("id")
	m := c.(*Machine)
	m.UnitStop(id)
	return c.JSON(http.StatusOK, map[string]string{})
}

// restart ...
func restart(c echo.Context) error {
	id := c.Param("id")
	m := c.(*Machine)
	m.UnitRestart(id)
	return c.JSON(http.StatusOK, map[string]string{})
}

// reload ...
func reload(c echo.Context) error {
	id := c.Param("id")
	m := c.(*Machine)
	m.UnitReload(id)
	return c.JSON(http.StatusOK, map[string]string{})
}

// allUnits ...
func allUnits(c echo.Context) error {
	m := c.(*Machine)
	us := m.AllUnits()
	return c.JSON(http.StatusOK, us)
}

// startAll ...
func startAll(c echo.Context) error {
	m := c.(*Machine)
	m.StartAllUnits()
	return c.JSON(http.StatusOK, map[string]string{})
}

// stopAll ...
func stopAll(c echo.Context) error {
	m := c.(*Machine)
	m.StopAllUnits()
	return c.JSON(http.StatusOK, map[string]string{})
}

// restartAll ...
func restartAll(c echo.Context) error {
	m := c.(*Machine)
	m.RestartAllUnits()
	return c.JSON(http.StatusOK, map[string]string{})
}

// reloadAll ...
func reloadAll(c echo.Context) error {
	m := c.(*Machine)
	m.ReloadAllUnits()
	return c.JSON(http.StatusOK, map[string]string{})
}

// removeAll ...
func removeAll(c echo.Context) error {
	m := c.(*Machine)
	m.RemoveAllUnits()
	return c.JSON(http.StatusOK, map[string]string{})
}

func errMessage(msg string) interface{} {
	return map[string]string{"msg": msg}
}
