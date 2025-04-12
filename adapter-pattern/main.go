package main

import "fmt"

//structural design pattern
// Its also know as Wrapper
//used so that 2 unrelated objects can work togather using an adapter
//the thing that joins these unrelated objects is called an Adapter

//There are Four Participants
//Target interface: This is the interface which will be used by the client
//Adapter: this is a wrapper which implements the target interface and modifies the specific request available from Adaptee class
//Adaptee: This is the object which is used by the Adapter to reuse the existing fuctionality and modify thrn for desired use
//Client : This will interact with the Adapter

//WHY
//when we dont need to change the existing object or interface rather wants to add new functionality on top of what is existing

//target interface
type mobile interface{
	chargeAppleMobile()
}
//prototype
type apple struct {}

func (a *apple)chargeAppleMobile()  {
	fmt.Println("Charging apple mobile!!!")
	
}
//clirnt
type client struct {}
func (c *client)ChargeMobile(mob mobile)  {
	mob.chargeAppleMobile()
	
}


//Adaptee
type android struct{}

func (ad *android)ChargeAndroidMobile()  {
	fmt.Println("Chanrginh ANDROID mobile")
	
}


type android_adapter struct{
	android *android
}

//Adapter
func(ad *android_adapter) chargeAppleMobile()  {
 ad.android.ChargeAndroidMobile()
	
}

func main()  {
//intial requirement
	apple:= &apple{}
	client:= &client{}
	client.ChargeMobile(apple)

//after that client need to charge android mobile too

android:= &android{}
android_adapter:= &android_adapter{
	android: android,
}
client.ChargeMobile(android_adapter)
	
	
}