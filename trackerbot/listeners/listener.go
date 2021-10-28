package listeners

type Listener struct {
	prefix string
}

func New(prefix string) *Listener {
	return &Listener{prefix: prefix}
}
