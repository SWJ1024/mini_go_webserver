package taskrunner

const (
	ReadyToDispatch = "d"
	ReadyToExecute  = "e"
	CLOSE           = "c"

	VideoPath = "./videos/"
)

type controlChan chan string


type dataChan chan interface{}


type fn func(dc dataChan) error
