package print

import "fmt"

type Print struct {
	PID string
}

func (p *Print) String() string {
	return p.PID
}

func (p *Print) Start() {
	fmt.Println("Start")
}

func (p *Print) Stop() {
	fmt.Println("Stop")
}

func (p *Print) Info() interface{} {
	return p.String()
}

func (p *Print) Restart() {
	fmt.Println("Restart")
}

func (p *Print) Reload() {
	fmt.Println("Reload")
}

func (p *Print) EndUnit(id chan string) {
	fmt.Println("EndUnit")
}

func (p *Print) Close() {
	fmt.Println("Close")
}

func (p *Print) ID() string {
	return p.PID
}
