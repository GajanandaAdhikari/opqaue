package main

import (
	"fmt"
	"errors"
)

func main(){
	printMe()

	var numerator int = 12
	var denominator int = 4
	var result, remainder, err = intDivision(numerator, denominator)
	if err!=nil{
		fmt.Printf(err.Error())
	}else if remainder = 0{
		fmt.Printf("The result of the divisoin is %v.",result)
	}else{
	fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	}

	// if	conditions given above can be re-written as switch 
	// {
	/*
	switch{
	case err!=nil:
		fmt.Println()
	case remainder == 0:
		Println
	default:
		the default print as normal operation
	}
*/
	
// conditional swetich is also available.
	
	sweitch remainder{
	case 0:
		fmt.Println("The division is exact.")
	case 1,2:
		fmt.Println("The division is close.")
	default:
		fmt,Println("The division is not close.")
	}
}


func printMe(){
	fmt.Println("Hellow World")
}

func intDivision(numerator int, denominator int) (int, int){

	var err error
	if denominator == 0{
		err = errors.New("Cannot Divide by Zero")
		return 0,0, err
	}

	var result int = numerator/denominator
	var remainder int = numerator%denominator
	return result, remainderr 
}
