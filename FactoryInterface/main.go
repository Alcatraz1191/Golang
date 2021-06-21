package main

import "fmt"

type Engine interface{
	setEngineType(eType string)
	setPower(power int)
	getEngineType() string
	getEnginePower() int
}

type engine struct {
	engineType string
	power  int
}

func (e *engine)setEngineType(eType string){
	e.engineType = eType
}

func (e *engine)setPower(power int){
	e.power = power
}

func (e *engine)getEngineType() string{
	return e.engineType
}

func (e* engine)getEnginePower() int {
	return e.power
}

type v8 struct{
	engine
}

type str86 struct{
	engine
}

func newv8() Engine{
	return &v8{
		engine: engine{
			engineType: "V8",
			power: 450,
		},
	}
}

func newstr86() Engine{
	return &str86{
		engine: engine{
			engineType: "Straight 6",
			power: 350,
		},
	}
}

func NewEngine(etype string, power int) engine {
	return engine{
		engineType: etype,
		power:  power,
	}
}

func getEngine(eType string) (Engine, error){
	if eType == "V8"{
		return newv8(), nil
	}
	if eType == "Straight 6"{
		return newstr86(), nil
	}
	return nil, fmt.Errorf("Engine type not found.")
}

func main() {
	mustang, _ := getEngine("V8")
	supra, _ := getEngine("Straight 6")
	printDetails(mustang)
	printDetails(supra)
	supra.setEngineType("V8")
	supra.setPower(500)
	printDetails(supra)
}

func printDetails(e Engine) {
    fmt.Printf("Engine: %s", e.getEngineType())
    fmt.Println()
    fmt.Printf("Power: %d", e.getEnginePower())
    fmt.Println()
}