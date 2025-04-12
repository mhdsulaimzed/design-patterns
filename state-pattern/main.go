package main

import "fmt"

//state design pattern comes under behavioural design pattern
//this pattern is used wen we have to change the behavious of the object based on its internal state
//Eg mobile alert system i.e Ring Silent or vibrate

//state pattern has 3 main participants
//1.Context: Interface Implementation for the client to interact it contain references to the concerete state objects and define their current state
//2.State: Interface for the object state .
//3.Concreate State : Concrete implementation of the state

//state
type tvState interface{
	state()
}

//concrete implementation of state
type on struct{}
func (o *on)state()  {
	fmt.Println("tv is on")
	
}
type off struct{}
func (o *off)state()  {
	fmt.Println("tv is off")
	
}

//context
type StateContext struct{
	currentTVState tvState
}


func NewContext()*StateContext  {
	return&StateContext{
		currentTVState: &off{},
	}
	
}

func (sc *StateContext)setState(state tvState)  {
	sc.currentTVState=state
	
}
func (sc *StateContext)getState()  {
	sc.currentTVState.state()
	
}


func main()  {
tvContext := NewContext()
tvContext.getState()
tvContext.setState(&on{})
tvContext.getState()
}
