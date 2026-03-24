package game

type BaseDestination string

const (
	DestinationOut    BaseDestination = "out"
	DestinationFirst  BaseDestination = "first"
	DestinationSecond BaseDestination = "second"
	DestinationThird  BaseDestination = "third"
	DestinationHome   BaseDestination = "home"
)

type Advancement struct {
	Batter     *BaseDestination
	FromFirst  *BaseDestination
	FromSecond *BaseDestination
	FromThird  *BaseDestination
}

type Play struct {
	Type     PlayType
	Override *Advancement
}
