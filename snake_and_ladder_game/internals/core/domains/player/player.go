package player

type Player struct {
	id       int
	name     string
	email    string
	isActive bool
}

var generator = 0

func GetNewPlayer(name, email string) *Player {
	return &Player{id: generator + 1, name: name, email: email, isActive: true}
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) GetId() int {
	return p.id
}
