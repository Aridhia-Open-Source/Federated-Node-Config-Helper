package helpers

type state struct {
	Namespace     string
	CloudPlatform string
}

var State = &state{
	Namespace:     "default",
	CloudPlatform: "Local",
}
