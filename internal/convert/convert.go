package convert


func CelsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9.0 / 5.0) + 32.0
}

func KphToMph(kph float64) float64 {
	return kph * 0.62137119
}
