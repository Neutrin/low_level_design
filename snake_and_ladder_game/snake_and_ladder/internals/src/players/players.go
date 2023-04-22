package players

import "sync/atomic"

var id uint64 = 0

type Player struct {
	id     int
	name   string
	widget *Widget
}

func NewPlayer(name string, widget *Widget) *Player {
	atomic.AddUint64(&id, 1)
	return &Player{
		id:     int(id),
		name:   name,
		widget: widget,
	}
}

func (p *Player) GetId() int {
	return p.id
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) GetWidget() *Widget {
	return p.widget
}
