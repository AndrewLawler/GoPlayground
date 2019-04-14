package main

import (
	"fmt"
)

func main(){
	a := 4.0;
	b := 6.0;
	fmt.Println("\nMath Examples\n- - - - - - -");
	addition(a,b);
	subtraction(a,b);
	multiplication(a,b);
	division(a,b);	
}

func addition(a float64, b float64){
	calc := a+b;
	fmt.Println(a,"+",b," =",calc);
}

func subtraction(a float64, b float64){
	calc := a-b;
	fmt.Println(a,"-",b," =",calc);
}

func multiplication(a float64, b float64){
	calc := a*b;
	fmt.Println(a,"*",b," =",calc);
}

func division(a float64, b float64){
	calc := a/b;
	calcRounded := fmt.Sprintf("%.2f", calc);
	fmt.Println(a,"/",b," =",calcRounded);
}
