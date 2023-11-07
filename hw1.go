package main

import "fmt"

// The main function calls each of the three functions and prints to the screen.
func main() {

	// function 1 call.
	pressure := CalcPressure(1.0, 1.0, 298.15)
	fmt.Printf(" %.4f pascals\n", pressure)

	//funtion 2 call
	c := map[byte]byte{
		'e': 'u', 'h': 'f', 'l': 'n', 'o': 'y',
	}
	fmt.Println(string(Decode(c, []byte("hello"))))

	// Fuction 3 call
	a, b := OddParity([]int{1})
	fmt.Println(a, b)

}

// CalcPressure calculates and returns the pressure using the ideal gas law. The function takes in the 3 parameters which are the v(volume),
// n(molarity) and t(temperature).
func CalcPressure(v, n, t float64) float64 {

	var pressure float64
	const R = 8.3144598
	pressure = (n * R * t) / v
	return pressure
}

// The Decode functions purpose is to decrypt a message. It takes in a map and a slice as its parameters
// and returns the decrypted message which is a slice of type byte.
func Decode(c map[byte]byte, e []byte) []byte {
	for i, v := range e {
		b, ok := c[v]
		if ok {
			e[i] = b
		}
	}
	return e

}

// The OddParity function returns two bool values. The first bool determines whether the integer value passed in
// has an odd or even amount of one's. Odd amount of one's will return true. The second bool result will return
// true if the integer contains only zeros and or one's.
func OddParity(nums []int) (bool, bool) {
	result1 := true
	result2 := true
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++

		} else if nums[i] != 0 {
			result2 = false

		}
	}
	if count%2 == 0 {
		result1 = false
	}

	return result1, result2

}
