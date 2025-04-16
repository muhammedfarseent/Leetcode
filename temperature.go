package main 
import "fmt"


func convertTemperature(celsius float64) []float64 {
    Kelvin := celsius + 273.15
    Fahrenheit := celsius * 1.80 + 32.00
    return []float64{Kelvin, Fahrenheit}

}
func main(){
	new:= 25.5
	fmt.Println(convertTemperature(new))
	fmt.Println("this shows both the Fahrenheit and Kelvin")
}