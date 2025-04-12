package main

import (
	"fmt"
	"log"
)

//Srategy pattern /Policy Pattern

//suppose i have a client class that want to call based on the user behaviour in the ui
//this client class doent want to call diffrent client class based on which data base that it want to connect
//niether it want to a switch case nor it have a switch condition or if else condition to decide
//client has to decide dynamically on runtime

type IDBconnection interface { // interface can be assined any type in runtime
	Connect()
}

type DBConnection struct {
	Db IDBconnection
}

func (con *DBConnection) Connect() {
	con.Db.Connect()

}

// implement First behavious to connect to mysql
type MySQLConnection struct {
	ConnectionString string
}

func (con *MySQLConnection) Connect() {

	fmt.Printf("Mysql conneted %s", con.ConnectionString)

}

// second behavour
type PostgresConnection struct {
	ConnectionString string
}

func (con *PostgresConnection) Connect() {

	fmt.Printf("Potgres conneted %s", con.ConnectionString)

}

// third behavour
type MongoDBConnection struct {
	ConnectionString string
}

func (con *MongoDBConnection) Connect() {

	fmt.Printf("Mongo db conneted %s", con.ConnectionString)

}

type Order struct {
	Status *Status
}
type Status struct {
	Name string
}

func (s *Status) Change() {
	s.Name = "Sulaim"
}

func (s Status) ChangeName() {
	s.Name = "Inshad"
	d := &s
	log.Println("^^^^^^^^^^^^", &d)
}

func main() {


	

	// order := &Order{Status: &Status{}}
	// order.Status.Name = "sulaim's orger status"

	// status := order.Status.Name
	// fmt.Println("hhhhhhhhhhhhhhhhhhhhhhhh",status)
	// log.Println("((((((((((((((()))))))))))))))", order.Status.Name)
	// Completed := &Status{Name: "Guru"}
	// log.Println("^^^^^^^^^^^^^", &Completed)
	// log.Println("----------->", Completed)
	// Completed.Change()
	// log.Println("^^^^^^^^^^^^^", &Completed)
	// log.Println("----------->", Completed)
	// Completed.ChangeName()
	// log.Println("----------->", Completed)

	// Mysqlconn := MySQLConnection{ConnectionString: "mysql://root:password@localhost:3306 \n"}
	// con := DBConnection{Db: &Mysqlconn}
	// con.Connect()

	// postgressqlconn := PostgresConnection{ConnectionString: "postgres://root:password@localhost:5432 \n"}
	// con = DBConnection{Db: &postgressqlconn}
	// con.Db = &Mysqlconn
	// con.Connect()

	// MongoDB := MongoDBConnection{ConnectionString: "mongodb://root:password@localhost:idk \n"}
	// con = DBConnection{Db: &MongoDB}
	// con.Connect()

}
