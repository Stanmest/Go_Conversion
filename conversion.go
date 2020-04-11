package main
import (
	"fmt"
	"math"
	)


//Converter for conversion
type Converter struct{
	feet Feet
	centimeter Centimeters 
	minutes Minutes
	seconds Seconds
	milliseconds Milliseconds
	celsius Celsius
	fahrenheit Fahrenheit
	radian Radian
	degree Degree
	kilogram Kilogram
	pounds Pounds
	meters Meters
	millimeters Millimeters

}

type (
	//Feet for conversion
Feet float64
//Centimeters for conversion
Centimeters float64
//Minutes for conversion
Minutes float64
//Seconds for conversion
Seconds float64
//Milliseconds for conversion
Milliseconds float64
//Celsius for conversion
Celsius float64
//Fahrenheit for conversion
Fahrenheit float64
//Radian for conversion
Radian float64
//Degree for conversion
Degree float64
//Kilogram for conversion
Kilogram float64
//Pounds for conversion
Pounds float64
//Meters for conversion
Meters float64
//Millimeters for conversion
Millimeters float64 
)

func (cvr Converter)minutesToSeconds(min Minutes)Seconds{
	    return Seconds(min * 60)
}

func (cvr Converter)secondsToMinutes(sec Seconds)Minutes{
	return Minutes(sec / 60)
}

func (cvr Converter)secondsToMilliseconds(sec Seconds)Milliseconds{
	return Milliseconds(sec * 1000)
}

func (cvr Converter)millisecondsToSeconds(milsec Milliseconds)Seconds{
	return Seconds(milsec / 1000)
}

func (cvr Converter)centimetersToFeet(centi Centimeters)Feet{
	return Feet(centi / 30.48)
}

func (cvr Converter)feetToCentimeters(ft Feet)Centimeters{
	return Centimeters(ft * 30.48)
}

func (cvr Converter)celsiusToFahrenheit(cels Celsius)Fahrenheit{
	return Fahrenheit(cels * 9/5 + 32)
}

func (cvr Converter)fahrenheitToCelsius(fah Fahrenheit)Celsius{
	return Celsius((fah - 32) * 5/9)
}

func (cvr Converter)radianToDegree(rad Radian)Degree{
	return Degree(rad * 180/math.Pi )
}

func (cvr Converter)degreeToRadian(deg Degree)Radian{
	return Radian(deg * math.Pi/180 )
}

func (cvr Converter)kilogramToPounds(kilo Kilogram)Pounds{
	return Pounds(kilo * 2.2046226218)
}
func (cvr Converter)poundsToKilogram(lbs Pounds)Kilogram{
	return Kilogram(lbs / 2.2046226218)
}

func (cvr Converter)metersToMillimeters(m Meters)Millimeters{
	return Millimeters(m * 1000)
}

func (cvr Converter)millimetersToMeters(mm Millimeters)Meters{
	return Meters(mm / 1000)
}

func main(){
	cvr := Converter{}
	fmt.Println(cvr.minutesToSeconds(2),"Seconds" )
	fmt.Println(cvr.secondsToMinutes(120),"Minutes")
	fmt.Println(cvr.secondsToMilliseconds(2),"Milliseconds")
	fmt.Println(cvr.millisecondsToSeconds(2000),"Seconds")
	fmt.Println(cvr.centimetersToFeet(90),"Feet")
	fmt.Println(cvr.feetToCentimeters(10),"Centimeters")
	fmt.Println(cvr.celsiusToFahrenheit(25),"Fahrenheit")
	fmt.Println(cvr.fahrenheitToCelsius(98.6),"Celsius")
	fmt.Println(cvr.radianToDegree(3.147),"Degree")
	fmt.Println(cvr.degreeToRadian(180),"Radians")
	fmt.Println(cvr.kilogramToPounds(2),"pounds")
	fmt.Println(cvr.poundsToKilogram(4),"Kilogram")
	fmt.Println(cvr.metersToMillimeters(4),"Millimeter")
	fmt.Println(cvr.millimetersToMeters(4000),"Meters")
	




}