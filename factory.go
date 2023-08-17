package main

import "fmt"

type vehiculo interface {
	numeroDePasajeros() int
	ponerEnMarcha() avanzar
}

type avanzar interface {
	encenderMotor(string)
	acelerar(string)
}

type carro struct{}
type avion struct{}
type moto struct{}
type ponerEnMarcha struct{}

func (c ponerEnMarcha) encenderMotor(vehiculo string) {
	fmt.Printf("Encendiendo motor %v \n", vehiculo)
}

func (c ponerEnMarcha) acelerar(vehiculo string) {
	fmt.Printf("acelarando %v \n", vehiculo)
}

func (c carro) numeroDePasajeros() int {
	return 5
}

func (c carro) ponerEnMarcha() avanzar {
	return ponerEnMarcha{}
}

func (c avion) numeroDePasajeros() int {
	return 100
}

func (c avion) ponerEnMarcha() avanzar {
	return ponerEnMarcha{}
}

func (c moto) numeroDePasajeros() int {
	return 2
}

func (c moto) ponerEnMarcha() avanzar {
	return ponerEnMarcha{}
}

func elegirVehiculo(tipo string) (vehiculo, error) {
	switch tipo {
	case "carro":
		return carro{}, nil
	case "avion":
		return avion{}, nil
	case "moto":
		return moto{}, nil
	default:
		return nil, fmt.Errorf("vehiculo no encontrado")
	}
}

func main() {
	var seleccionVehiculo string
	fmt.Print("Ingresa un vehiculo: ")
	fmt.Scanln(&seleccionVehiculo)
	fmt.Println("Ingresaste:", seleccionVehiculo)

	vehiculo, error := elegirVehiculo(seleccionVehiculo)
	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Println("El vehiculo tiene", vehiculo.numeroDePasajeros(), "pasajeros")
	vehiculo.ponerEnMarcha().encenderMotor(seleccionVehiculo)
	vehiculo.ponerEnMarcha().acelerar(seleccionVehiculo)
}
