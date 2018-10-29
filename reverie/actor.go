package reverie

type ActorBehavior interface {
	Name() string
	Close()
} 

type Actor struct {
}

func (a *Actor) Name() string {
	return ""
}

func (a *Actor) Close() {
}