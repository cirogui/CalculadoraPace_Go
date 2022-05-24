package main

import (
	"fmt"
)

func main() {

	fmt.Println("Calculadora de Pace")
	fmt.Println("-------------------")

	calcVm()

}

func calcVm() {
	distancia := recebeDistancia()
	horas := recebeHora()
	minutos := recebeMin()
	segundos := recebeSeg()
	vm := (distancia / (horas + minutos + segundos)) * 3.6
	pace := ((horas + minutos + segundos) / 60) / distancia / 1000
	fmt.Println("Distância total = ", distancia/1000, " km")
	fmt.Println("Tempo total = ", horas/3600, " horas", minutos/60, " minutos e ", segundos, " segundos")
	fmt.Println("Velocidade média = ", vm, "km/h")
	fmt.Println("Pace = ", pace, " min")

}

func recebeDistancia() float32 {

	var dist float32
	fmt.Println("Digite a distância em Km")
	fmt.Scan(&dist)
	return dist * 1000
}

func recebeHora() float32 {

	var hora float32
	fmt.Println("Digite a quantidade de horas")
	fmt.Scan(&hora)
	return hora * 3600
}

func recebeMin() float32 {

	var min float32
	fmt.Println("Digite a quantidade de minutos")
	fmt.Scan(&min)
	return min * 60
}

func recebeSeg() float32 {

	var seg float32
	fmt.Println("Digite a quantidade de segundos")
	fmt.Scan(&seg)
	return seg

}
