package main

import (
	"lab5/lab5_car"
	_ "lab5/lab5_car"
	"lab5/lab5_car/headlights"
	_ "lab5/lab5_car/headlights"
	"lab5/lab5_car/stereo"
	_ "lab5/lab5_car/stereo"
	"lab5/lab5_car/wheels"
	_ "lab5/lab5_car/wheels"
)

func main() {

	lab5_car.OpenDoor()
	headlights.TurnOn()
	stereo.TurnOn()
	stereo.BoostBass()
	wheels.Steer()
	wheels.Accelerate()


}
