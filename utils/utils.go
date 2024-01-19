package utils

import "log"

type IntPair struct {
	X int
	Y int
}

type CameraCoords struct {
	X float64 // 0 to 1
	Y float64 // 0 to 1
}

type PlaneCoords struct {
	X float64 // m
	Y float64 // m
}

type SpaceCoords struct {
	X float64 // m
	Y float64 // m
	Z float64 // m
}

type CameraLine struct {
	A CameraCoords
	B CameraCoords
}

type SpaceEdge struct {
	A SpaceCoords
	B SpaceCoords
}

type RotationMatrix struct {
	A float64
	B float64
	C float64
	D float64
	E float64
	F float64
	G float64
	H float64
	I float64
}

type RotationCoords struct {
	X float64 // rad
	Y float64 // rad
	Z float64 // rad
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
