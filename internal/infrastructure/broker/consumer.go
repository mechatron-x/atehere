package broker

type Consumer[TEvent Event] interface {
	ProcessEvent(event TEvent) error
}
