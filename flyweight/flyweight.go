package flyweight

import "fmt"

const (
	TerroristDressType = "tDress"
	CounterTerroristDressType = "ctDress"
)

var (
	dressFactorySingleInstance = &dressFactory{
		dressMap: make(map[string]dress),
	}
)
type dressFactory struct {
	dressMap map[string]dress
}

func (d *dressFactory)getDressByType(dressType string)(dress, error)  {
	if d.dressMap[dressType] != nil {
		return  d.dressMap[dressType], nil
	}
	if dressType == TerroristDressType {
		d.dressMap[dressType] = newTerroristDress()
		return d.dressMap[dressType], nil
	}
	if dressType == CounterTerroristDressType {
		d.dressMap[dressType] = newCounterTerroristDress()
		return d.dressMap[dressType], nil
	}
	return nil, fmt.Errorf("Wrong dress type passed")
}
func getDressFactorySingleInstance() *dressFactory {
	return dressFactorySingleInstance
}
type dress interface {
	getColor() string
}
type terroristDress struct {
	color string
}
func(t *terroristDress)getColor()string {
	return t.color
}
func newTerroristDress()*terroristDress {
	return &terroristDress{color: "red"}
}
type counterTerroristDress struct {
	color string
}

func (c *counterTerroristDress)getColor() string {
	return c.color
}
func newCounterTerroristDress() *counterTerroristDress{
	return &counterTerroristDress{color: "green"}
}

type player struct {
	dress dress
	playerType string
	lat int
	long int
}
func newPlayer(playerType, dressType string) *player {
	dress,_ := getDressFactorySingleInstance().getDressByType(dressType)
	return &player{
		dress:      dress,
		playerType: playerType,
	}
}
func (p *player)newLocation(lat, long int)  {
	p.lat = lat
	p.long = long
}

type game struct {
	terrorist []*player
	counterTerrorist []*player
}
func newGame()*game{
	return &game{
		terrorist:        make([]*player,1),
		counterTerrorist: make([]*player,1),
	}
}
func (g *game)AddTerrorist(dressType string)  {
	player := newPlayer("T", dressType)
	g.terrorist = append(g.terrorist,player)
	return
}
func (g *game)AddCounterTerrorist(dressType string)  {
	player := newPlayer("CT", dressType)
	g.counterTerrorist= append(g.counterTerrorist,player)
	return
}

func TestFlyweight(){
	game := newGame()
	game.AddTerrorist(TerroristDressType)
	game.AddTerrorist(TerroristDressType)
	game.AddTerrorist(TerroristDressType)
	game.AddTerrorist(TerroristDressType)
	
	game.AddCounterTerrorist(CounterTerroristDressType)
	game.AddCounterTerrorist(CounterTerroristDressType)
	game.AddCounterTerrorist(CounterTerroristDressType)
	dressFactorySingleInstance := getDressFactorySingleInstance()
	for dressType, dress := range dressFactorySingleInstance.dressMap{
		fmt.Printf("DressColorType: %s\nDressColor: %s\n",dressType, dress.getColor())
	}
	
	
}