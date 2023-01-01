package main

type Subject interface {
	registry(o Observer)
	unRegister(o Observer)
	notify()
}

type Observer interface {
	update(string)
	getID() string
}
