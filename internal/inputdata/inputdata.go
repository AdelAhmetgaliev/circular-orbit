package inputdata

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"

	"github.com/AdelAhmetgaliev/circular-orbit/internal/angle"
)

// Структура которая хранит координаты
type Coordinates struct {
	X float64
	Y float64
	Z float64
}

// Структура которая хранит входные данные
type InputData struct {
	// Время наблюдения
	Time float64

	// Экваториальные координаты
	RightAscension angle.Angle // Прямое восхождение
	Declination    angle.Angle // Склонение

	// Геоцентрические экваториальные координаты
	GeocentricCoords Coordinates
}

// Функция считывает входные данные из csv-файла, лежащего по пути filePath
// Фукнция возвращает два экземпляра структуры InputData или вызывает панику
func ReadInputData(filePath string) (InputData, InputData) {
	inputFile, err := os.Open(filePath)
	if err != nil {
		panic("Can't open file")
	}
	defer inputFile.Close()

	csvReader := csv.NewReader(inputFile)
	records, err := csvReader.ReadAll()
	if err != nil {
		panic("Can't read data from file")
	}

	timeArr := records[1]
	rightAscensionArr := records[2]
	declinationArr := records[3]
	xArr := records[4]
	yArr := records[5]
	zArr := records[6]

	time1, _ := strconv.ParseFloat(strings.TrimSpace(timeArr[1]), 64)
	time2, _ := strconv.ParseFloat(strings.TrimSpace(timeArr[2]), 64)

	rightAscensionArr1 := strings.Split(strings.TrimSpace(rightAscensionArr[1]), " ")
	rightAscensionArr2 := strings.Split(strings.TrimSpace(rightAscensionArr[2]), " ")
	rightAscension1 := parseAngleFromHours(rightAscensionArr1)
	rightAscension2 := parseAngleFromHours(rightAscensionArr2)

	declinationArr1 := strings.Split(strings.TrimSpace(declinationArr[1]), " ")
	declinationArr2 := strings.Split(strings.TrimSpace(declinationArr[2]), " ")
	declination1 := parseAngleFromDegrees(declinationArr1)
	declination2 := parseAngleFromDegrees(declinationArr2)

	x1, _ := strconv.ParseFloat(strings.TrimSpace(xArr[1]), 64)
	y1, _ := strconv.ParseFloat(strings.TrimSpace(yArr[1]), 64)
	z1, _ := strconv.ParseFloat(strings.TrimSpace(zArr[1]), 64)
	coords1 := Coordinates{x1, y1, z1}

	x2, _ := strconv.ParseFloat(strings.TrimSpace(xArr[2]), 64)
	y2, _ := strconv.ParseFloat(strings.TrimSpace(yArr[2]), 64)
	z2, _ := strconv.ParseFloat(strings.TrimSpace(zArr[2]), 64)
	coords2 := Coordinates{x2, y2, z2}

	inputData1 := InputData{time1, rightAscension1, declination1, coords1}
	inputData2 := InputData{time2, rightAscension2, declination2, coords2}

	return inputData1, inputData2
}

func parseAngleFromHours(hoursArr []string) angle.Angle {
	hour, _ := strconv.ParseUint(hoursArr[0], 10, 64)
	minute, _ := strconv.ParseUint(hoursArr[1], 10, 64)
	second, _ := strconv.ParseFloat(hoursArr[2], 64)

	hours := float64(hour) + float64(minute)/60.0 + second/3600.0

	return angle.FromDegrees(hours * 15.0)
}

func parseAngleFromDegrees(degreesArr []string) angle.Angle {
	degree, _ := strconv.ParseInt(degreesArr[0], 10, 64)
	minute, _ := strconv.ParseInt(degreesArr[1], 10, 64)
	second, _ := strconv.ParseFloat(degreesArr[2], 64)

	degrees := float64(degree) + float64(minute)/60.0 + second/3600.0

	return angle.FromDegrees(degrees)
}
