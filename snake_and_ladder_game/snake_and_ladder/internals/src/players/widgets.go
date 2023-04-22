package players

import "sync/atomic"

var widID uint64

type Widget struct {
	id    int
	color string
}

func NewWidget(color string) *Widget {
	atomic.AddUint64(&widID, 1)
	return &Widget{
		id:    int(widID),
		color: color,
	}
}

func (w *Widget) GetID() int {
	return w.id
}

func (w *Widget) Color() string {
	return w.color
}
