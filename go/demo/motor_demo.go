package main

//This is a demo of how our motor interface works

import (
	"../motor"
	"../motor/fake"
	"log"
)

func main() {
	//First, lets make a fake motor
	log.Println("Making a fake motor.")
	m := fake.NewFake("example", motor.DC, 0)

	//Set motor speed
	m.Set(255)

	//Print info
	log.Println("Motor type: " + m.GetMotorType().TypeName())
	log.Printf("Motor speed: %d\n", m.State())
}
