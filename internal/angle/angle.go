package angle

import "math"

// Тип Angle синоним типа float64
type Angle float64

// Конструктор Angle, который принимает градусы в формате float64
// Возвращает экземпляр типа Angle
func FromDegrees(d float64) Angle {
	return Angle(d * (math.Pi / 180.0))
}

// Конструктор Angle, который принимает радианы в формате float64
// Возвращает экземпляр типа Angle
func FromRadians(r float64) Angle {
	return Angle(r)
}

func Acos(c float64) Angle {
	return Angle(math.Acos(c))
}

func Asin(s float64) Angle {
	return Angle(math.Asin(s))
}

func Atan(t float64) Angle {
	return Angle(math.Atan(t))
}

// Функция возвращающая угол в правильной четверти
// Принимает значение синуса s и косинуса c
// Возвращает Angle, лежащий в пределах [0; 2*pi]
func Atan2(s float64, c float64) Angle {
	if c >= 0.0 && s >= 0.0 {
		return Asin(s)
	}

	if c <= 0.0 && s >= 0.0 {
		return Acos(c)
	}

	if c <= 0.0 && s <= 0.0 {
		return Angle(math.Pi) - Asin(s)
	}

	return Asin(s) + Angle(2.0*math.Pi)
}

func (a Angle) Degrees() float64 {
	return float64(a * (180.0 / math.Pi))
}

func (a Angle) Radians() float64 {
	return float64(a)
}

func (a Angle) Cos() float64 {
	return math.Cos(float64(a))
}

func (a Angle) Sin() float64 {
	return math.Sin(float64(a))
}

func (a Angle) Tan() float64 {
	return math.Tan(float64(a))
}
