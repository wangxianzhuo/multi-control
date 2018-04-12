package control

// Unit is the baisc control unit
type Unit interface {
	Start()
	Stop()
	Info() interface{}
	Restart()
	Reload()

	// ask machine to remove this unit
	EndUnit(id chan string)

	Close()

	ID() string
}
