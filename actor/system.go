package actor

type actorSystem struct {
	actors map[string]ActorBehavior
}
