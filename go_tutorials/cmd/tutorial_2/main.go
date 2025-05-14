package main

import "fmt"

func main(){
	var intNum int = 32767
	fmt.Println(intNum)

	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)
	
	var floatNum32 float32 = 10.1
	var intNum8 int8 = 2
	var result float32 = floatNum32 + float32(intNum8)
	fmt.Println(result)

	var intNum1 int = 5
	var intNum2 int = 4
	fmt.Println(intNum1/intNum2)
	fmt.Println(intNum1%intNum2)

	var aString string = "Hello World"
	var newlineString string = `Hello 
	World` 
	var nextlineString string = "Hello \nWorkld"
	fmt.Println(aString)
	fmt.Println(newlineString)
	fmt.Println(nextlineString)

	fmt.Println(len("Length outputs in bytesize."))
	fmt.Println(utf.RuneCountInString("String count in characters."))	

	var myBoolean bool = false
	fmt.Println(myBoolean)

	var intNum3 rune
	fmt.Println("Default value for rune is" + intNum3)

	var myVar = "text"
	StringVariable := "text"
	fmt.Println(myVar + StringVariable)
	
	var1, var2 := 1,
	fmt.Println(var1, var2)
	
	const myConst string = "I can't change it's value once created"
	const pi float32 = 3.14
	fmt.Println(myConst + pi)
}

