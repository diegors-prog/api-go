package main

import (
	"database/sql"
	"fmt"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

//A GO não é orientado a objetos, mas tem a sacada das structs
//para conseguir programar parecido a orientação a objetos
//A GO junta funcional procedural e orientação a objetos.

type Car struct {
	Name  string
	Model string
	Price float64
}

//slice - é como se fosse um arraycom tamanho infinito, ou seja,
//ele pode crescer a todo momento.

var cars []Car

func generateCars() {
	cars = append(cars, Car{Name: "Parati", Model: "VW", Price: 20000})
	cars = append(cars, Car{Name: "Tempra", Model: "Fiat", Price: 10000})
	cars = append(cars, Car{Name: "Escort", Model: "Ford", Price: 15000})
}

//interfaces, apenas implementando o metodo que existe
//na interface, vc ja está implementando ela

//interface Vehiculo {
//	Andar()
//}

// comportamento em uma struct, basicamente add métodos
func (c Car) Andar() {
	fmt.Println("O", c.Name, "está andando")
}

func soma(a, b int) (int, error) {
	if a+b > 10 {
		return 0, fmt.Errorf("Soma maior que 10")
	} else {
		return a + b, nil
	}
}

func main() {
	result, err := soma(1, 9)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
		carro := Car{"Fusca", "Audi", 5000.90}
		carro.Model = "VW"
		fmt.Println(carro.Name, carro.Model)
		carro.Andar()
	}

	generateCars()
	e := echo.New()
	e.GET("/cars", getCars)
	e.POST("/cars", createCar)
	e.Logger.Fatal(e.Start(":8080"))
}

//em memoria

func getCars(c echo.Context) error {
	return c.JSON(200, cars)
}

func createCar(c echo.Context) error {
	car := new(Car)
	if err := c.Bind(car); err != nil {
		return err
	}
	cars = append(cars, *car)
	saveCar(*car)
	return c.JSON(200, cars)
}

//conexão com o banco

func saveCar(car Car) error {
	db, err := sql.Open("sqlite3", "cars.db")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO cars (name, model, price) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(car.Name, car.Model, car.Price)
	if err != nil {
		return err
	}
	return nil
}
