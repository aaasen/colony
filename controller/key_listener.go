package controller

type KeyListener struct {
	KeyBindings map[KeyEvent]int
	Subscribers []chan<- int

	rawEvents <-chan KeyEvent
}

func NewKeyListener() *KeyListener {
	return &KeyListener{
		KeyBindings: defaultKeyBindings,
		Subscribers: make([]chan<- int, 0),
		rawEvents:   keyEventChannel,
	}
}

func (self *KeyListener) Subscribe(eventChan chan<- int) {
	self.Subscribers = append(self.Subscribers, eventChan)
}

func (self *KeyListener) Listen() {
	for {
		select {
		case rawEvent := <-self.rawEvents:
			if event, ok := self.KeyBindings[rawEvent]; ok {
				go func(event int, subs []chan<- int) {
					for _, sub := range subs {
						sub <- event
					}
				}(event, self.Subscribers)
			}
		}
	}
}
