package main

import (
	"fmt"
	"math"
	"path/filepath"

	"github.com/AdelAhmetgaliev/circular-orbit/internal/angle"
	"github.com/AdelAhmetgaliev/circular-orbit/internal/constants"
	"github.com/AdelAhmetgaliev/circular-orbit/internal/inputdata"
)

func main() {
	inputDataFilePath := filepath.Join("data", "input.csv")
	inputData1, inputData2 := inputdata.ReadInputData(inputDataFilePath)

	tempValueA1 := inputData1.RightAscension.Cos() * inputData1.Declination.Cos()
	tempValueB1 := inputData1.RightAscension.Sin() * inputData1.Declination.Cos()
	tempValueC1 := inputData1.Declination.Sin()

	tempValueA2 := inputData2.RightAscension.Cos() * inputData2.Declination.Cos()
	tempValueB2 := inputData2.RightAscension.Sin() * inputData2.Declination.Cos()
	tempValueC2 := inputData2.Declination.Sin()

	if (tempValueA1*tempValueA1+tempValueB1*tempValueB1+tempValueC1*tempValueC1)-1.0 > constants.Epsilon {
		panic("The first temporary values are calculated incorrectly")
	}

	if (tempValueA2*tempValueA2+tempValueB2*tempValueB2+tempValueC2*tempValueC2)-1.0 > constants.Epsilon {
		panic("The second temporary values are calculated incorrectly")
	}

	rCos_1 := -(tempValueA1*inputData1.GeocentricCoords.X +
		tempValueB1*inputData1.GeocentricCoords.Y +
		tempValueC1*inputData1.GeocentricCoords.Z)
	r2_1 := inputData1.GeocentricCoords.X*inputData1.GeocentricCoords.X +
		inputData1.GeocentricCoords.Y*inputData1.GeocentricCoords.Y +
		inputData1.GeocentricCoords.Z*inputData1.GeocentricCoords.Z
	rSin2_1 := r2_1 - rCos_1*rCos_1

	rCos_2 := -(tempValueA2*inputData2.GeocentricCoords.X +
		tempValueB2*inputData2.GeocentricCoords.Y +
		tempValueC2*inputData2.GeocentricCoords.Z)
	r2_2 := inputData2.GeocentricCoords.X*inputData2.GeocentricCoords.X +
		inputData2.GeocentricCoords.Y*inputData2.GeocentricCoords.Y +
		inputData2.GeocentricCoords.Z*inputData2.GeocentricCoords.Z
	rSin2_2 := r2_2 - rCos_2*rCos_2

	a, a1 := inputData1.SemiMajorAxis, inputData1.SemiMajorAxis+0.1
	f := angle.FromRadians(0.0)
	x1, y1, z1 := 0.0, 0.0, 0.0
	x2, y2, z2 := 0.0, 0.0, 0.0
	for {
		ro11 := math.Sqrt(a*a-rSin2_1) - rCos_1
		x11 := tempValueA1*ro11 - inputData1.GeocentricCoords.X
		y11 := tempValueB1*ro11 - inputData1.GeocentricCoords.Y
		z11 := tempValueC1*ro11 - inputData1.GeocentricCoords.Z

		ro21 := math.Sqrt(a*a-rSin2_2) - rCos_2
		x21 := tempValueA2*ro21 - inputData2.GeocentricCoords.X
		y21 := tempValueB2*ro21 - inputData2.GeocentricCoords.Y
		z21 := tempValueC2*ro21 - inputData2.GeocentricCoords.Z

		sin_fg11 := math.Sqrt((1 / (4.0 * a * a)) * ((x21-x11)*(x21-x11) + (y21-y11)*(y21-y11) + (z21-z11)*(z21-z11)))
		fg11 := angle.Asin(sin_fg11)
		fd11 := constants.GravitationalConstant * (inputData2.Time - inputData1.Time) / (2 * a * math.Sqrt(a))

		ro12 := math.Sqrt(a1*a1-rSin2_1) - rCos_1
		x12 := tempValueA1*ro12 - inputData1.GeocentricCoords.X
		y12 := tempValueB1*ro12 - inputData1.GeocentricCoords.Y
		z12 := tempValueC1*ro12 - inputData1.GeocentricCoords.Z

		ro22 := math.Sqrt(a1*a1-rSin2_2) - rCos_2
		x22 := tempValueA2*ro22 - inputData2.GeocentricCoords.X
		y22 := tempValueB2*ro22 - inputData2.GeocentricCoords.Y
		z22 := tempValueC2*ro22 - inputData2.GeocentricCoords.Z

		sin_fg12 := math.Sqrt((1 / (4.0 * a1 * a1)) * ((x22-x12)*(x22-x12) + (y22-y12)*(y22-y12) + (z22-z12)*(z22-z12)))
		fg12 := angle.Asin(sin_fg12)
		fd12 := constants.GravitationalConstant * (inputData2.Time - inputData1.Time) / (2 * a1 * math.Sqrt(a1))

		delta1 := float64(fg11) - fd11
		delta2 := float64(fg12) - fd12

		a2 := a1 - delta2*(a1-a)/(delta2-delta1)

		a1 = a
		a = a2

		ro1 := math.Sqrt(a*a-rSin2_1) - rCos_1
		x1 = tempValueA1*ro1 - inputData1.GeocentricCoords.X
		y1 = tempValueB1*ro1 - inputData1.GeocentricCoords.Y
		z1 = tempValueC1*ro1 - inputData1.GeocentricCoords.Z

		ro2 := math.Sqrt(a*a-rSin2_2) - rCos_2
		x2 = tempValueA2*ro2 - inputData2.GeocentricCoords.X
		y2 = tempValueB2*ro2 - inputData2.GeocentricCoords.Y
		z2 = tempValueC2*ro2 - inputData2.GeocentricCoords.Z

		sin_fg := math.Sqrt((1 / (4.0 * a * a)) * ((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1) + (z2-z1)*(z2-z1)))
		fg := angle.Asin(sin_fg)
		fd := constants.GravitationalConstant * (inputData2.Time - inputData1.Time) / (2 * a * math.Sqrt(a))

		f = (fg + angle.FromRadians(fd)) / angle.Angle(2.0)

		if math.Abs(a2-a1) < constants.Epsilon {
			break
		}
	}

	pX := (x1 + x2) / (2.0 * f.Cos() * a)
	pY := (y1 + y2) / (2.0 * f.Cos() * a)
	pZ := (z1 + z2) / (2.0 * f.Cos() * a)

	if math.Abs((pX*pX+pY*pY+pZ*pZ)-1.0) > constants.Epsilon {
		panic("The first vector element is calculated incorrectly")
	}

	qX := (x2 - x1) / (2.0 * f.Sin() * a)
	qY := (y2 - y1) / (2.0 * f.Sin() * a)
	qZ := (z2 - z1) / (2.0 * f.Sin() * a)

	if math.Abs((qX*qX+qY*qY+qZ*qZ)-1.0) > constants.Epsilon {
		panic("The second vector element is calculated incorrectly")
	}

	if math.Abs(pX*qX+pY*qY+pZ*qZ) > constants.Epsilon {
		panic("Vector elements are not orthogonal")
	}

	eclipticTilt := angle.FromDegrees(constants.EclipticTiltDegrees)
	time0 := (inputData1.Time + inputData2.Time) / 2.0

	sinIsinU := pZ*eclipticTilt.Cos() - pY*eclipticTilt.Sin()
	sinIcosU := qZ*eclipticTilt.Cos() - qY*eclipticTilt.Sin()

	sinI := math.Sqrt(sinIsinU*sinIsinU + sinIcosU*sinIcosU)
	sinU := sinIsinU / sinI
	cosU := sinIcosU / sinI

	u0 := angle.Atan2(sinU, cosU)

	sinW := (pY*cosU - qY*sinU) / eclipticTilt.Cos()
	cosW := pX*cosU - qX*sinU
	ascendingNode := angle.Atan2(sinW, cosW)

	cosI := -(pX*sinU + qX*cosU) / sinW

	inclination := angle.Atan2(sinI, cosI)

	fmt.Printf("e = 0\n")
	fmt.Printf("a = %.8f a.e.\n", a)
	fmt.Printf("i = %.8f°\n", inclination.Degrees())
	fmt.Printf("Ω = %.8f°\n", ascendingNode.Degrees())
	fmt.Printf("u = %.8f°\n", u0.Degrees())
	fmt.Printf("t = %.5f\n", time0)
}
