package internal

type Screen interface {
	Fetch()
	Render()
	ChangeState(string) State
}

type Display struct {
	screen Screen
}

func NewDisplay(s Screen) *Display {
	return &Display{
		screen: s,
	}
}

func (d *Display) Fetch() {
	d.screen.Fetch()
}

func (d *Display) Render() {
	d.screen.Render()
}

func (d *Display) ChangeState(s string) State {
	return d.screen.ChangeState(s)
}
