package control

import (
	"fmt"
	"sync"

	"github.com/labstack/echo"
)

// Machine ...
type Machine struct {
	echo.Context
	Units       sync.Map
	unitEndChan chan string
	stopSign    chan struct{}
}

// RegisterAsEchoMiddleware ...
func (m *Machine) RegisterAsEchoMiddleware(e *echo.Echo) {
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			m.Context = c
			cc := m
			return h(cc)
		}
	})
}

// NewMachine ...
func NewMachine(endSign chan string) (*Machine, error) {
	return &Machine{
		unitEndChan: make(chan string, 1),
	}, nil
}

// Listen control sign
func (m *Machine) Listen() {
	defer m.close()

	for {
		select {
		case <-m.stopSign:
			m.stop()
			return
		case id := <-m.unitEndChan:
			m.RemoveUnit(id)
		}
	}
}

func (m *Machine) stop() {
	m.Units.Range(func(key, value interface{}) bool {
		u := value.(Unit)
		u.Stop()
		u.Close()
		return true
	})
}

// Stop can stop the machine and stop all of it's units
func (m *Machine) Stop() {
	m.stopSign <- struct{}{}
}

func (m *Machine) close() {
	close(m.stopSign)
	close(m.unitEndChan)
}

func (m *Machine) getUnit(id string) (*Unit, error) {
	uu, ok := m.Units.Load(id)
	if !ok {
		return nil, fmt.Errorf("the unit [" + id + "] is not exists")
	}
	u := uu.(*Unit)
	return u, nil
}

// AddUnit ...
func (m *Machine) AddUnit(u Unit) {
	id := u.ID()
	m.Units.Store(id, &u)
}

// RemoveUnit ...
func (m *Machine) RemoveUnit(id string) {
	m.Units.Delete(id)
}

// UnitInfo ...
func (m *Machine) UnitInfo(id string) (interface{}, error) {
	uu, ok := m.Units.Load(id)
	if !ok {
		return nil, fmt.Errorf("the unit [" + id + "] is not exists")
	}
	u := uu.(*Unit)
	return (*u).Info(), nil
}

// UnitStart ...
func (m *Machine) UnitStart(id string) error {
	u, err := m.getUnit(id)
	if err != nil {
		return err
	}

	(*u).Start()
	return nil
}

// UnitStop ...
func (m *Machine) UnitStop(id string) error {
	u, err := m.getUnit(id)
	if err != nil {
		return err
	}

	(*u).Stop()
	return nil
}

// UnitRestart ...
func (m *Machine) UnitRestart(id string) error {
	u, err := m.getUnit(id)
	if err != nil {
		return err
	}

	(*u).Restart()
	return nil
}

// UnitReload ...
func (m *Machine) UnitReload(id string) error {
	u, err := m.getUnit(id)
	if err != nil {
		return err
	}

	(*u).Reload()
	return nil
}

// AllUnits ...
func (m *Machine) AllUnits() []*Unit {
	var units []*Unit
	m.Units.Range(func(key, value interface{}) bool {
		u := value.(*Unit)
		units = append(units, u)
		return true
	})

	return units
}

// StartAllUnits ...
func (m *Machine) StartAllUnits() {
	m.Units.Range(func(key, value interface{}) bool {
		u := value.(*Unit)
		(*u).Start()
		return true
	})
}

// StopAllUnits ...
func (m *Machine) StopAllUnits() {
	m.Units.Range(func(key, value interface{}) bool {
		u := value.(*Unit)
		(*u).Stop()
		return true
	})
}

// RestartAllUnits ...
func (m *Machine) RestartAllUnits() {
	m.Units.Range(func(key, value interface{}) bool {
		u := value.(*Unit)
		(*u).Restart()
		return true
	})
}

// ReloadAllUnits ...
func (m *Machine) ReloadAllUnits() {
	m.Units.Range(func(key, value interface{}) bool {
		u := value.(*Unit)
		(*u).Reload()
		return true
	})
}

// RemoveAllUnits ...
func (m *Machine) RemoveAllUnits() {
	units := m.AllUnits()
	for _, u := range units {
		m.Units.Delete((*u).ID())
	}
}
