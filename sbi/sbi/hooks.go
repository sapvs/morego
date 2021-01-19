package sbi

//State type for state of SBI
type State int

const (
	//DOWN SBI is down
	DOWN State = iota
	//WAITING SBI is waiting, may be for deps
	WAITING
	//STARTING SBI is STARTING
	STARTING
	//UP SBI is up
	UP
)

//SBI represents a SBI
type SBI struct {
	//Name name of the SBI as the system knows it
	Name  string
	State State
}

//Handler interface for entities to susbscribe, check, and hanle updates fro SBIs
type Handler interface {
	//WaitForSBI blocking call, returns when the SBI is in UP state
	WaitForSBI(sbi *SBI)
	//WaitForSBIState blocking call, returns when the SBI is in given state
	WaitForSBIState(sbi *SBI, state State)
	//SubscribeSBIUpdates subscribe to receive the updates of SBI.
	SubscribeSBIUpdates(sbi *SBI)
	HandleSBIUpdate()
}
