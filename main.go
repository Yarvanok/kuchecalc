package main

import (
	"fmt"
	"math"
)

var (
	// Для кухонного оборудования
	// Длина и ширина в метрах
	A float64 = 0.55
	B float64 = 0.9

	Q        float64 = 14   // Мощность в кВт
	Z        float64 = 0.65 // Высота от поверхности до отсоса в метрах
	position int     = 0    // 0 - свободностоящая, 1 - у стены, 2 - в углу

	k1 float64 = 90  // Таблица А. Доля явных выделений от установочной мощности
	k2 float64 = 0.5 // Доля конвективных выделений от явных выделений
	k3 float64 = 0.7 // Таблица Б. Коэффициент одновременности работы

	// Для местных отсосов
	k4 float64 = 0.85 // Коэффециент эффективности местного отсоса
	k5 float64 = 1.25 // Поправочный коэф. по Таблице 2.
	// При перемешиваемой: через решётки на стене: 1.25, через плафоны на потолке 1.2
	// При подаче в рабочей зоне 1.05, на потолке 1.10
)

func CalcAirUnderKitchenware() float64 {
	const k float64 = 0.005

	q := CalcKitchenwareHeat()

	d := (2 * A * B) / (A + B)

	var r float64 = 1.0
	switch position {
	case 1:
		a := (0.63 * B) / A
		if a < 0.63 {
			r = 0.63
		} else {
			r = a
		}
	case 2:
		r = 0.4
	}

	return k * math.Sqrt(math.Sqrt(q)) * math.Sqrt(math.Sqrt((Z+1.7*d)*(Z+1.7*d)*(Z+1.7*d)*(Z+1.7*d)*(Z+1.7*d))) * r
}

func CalcKitchenwareHeat() float64 {
	return Q * k1 * k2 * k3
}

func CalcAirOutcomingByLocal() float64 {
	L1 := CalcAirUnderKitchenware()
	L2 := 3.75 * 0.0000001 * Q * k3

	return (L1 + L2) * (k5 / k4)
}

func CalcAirIncoming() float64 {
 return 0
}

func main() {
	a := CalcAirOutcomingByLocal()
	b := CalcAirIncoming()
	fmt.Printf("Расход воздуха равен: %.3f м3/с\n", a)
	fmt.Printf("В пересчёте: %.3f м3/ч\n", a*60*60)
	fmt.Printf("Приток воздуха: %.3f м3/с\n", b)
	fmt.Printf("Приток воздуха: %.3f м3/ч\n", b*60*60)
}
