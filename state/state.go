package state

type state struct {
	Namespace     string
	CloudPlatform string
	Error         string
}

var State = &state{
	Namespace:     "default",
	CloudPlatform: "Local",
	Error:         "",
}
