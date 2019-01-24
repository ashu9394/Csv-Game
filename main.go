package main

import (
	"bufio"
	"os"
	"fmt"
)


func main(){
	if(os.Args[1] == "withTimer"){
		evalWithTimer()
	} 
	evalNormal()
}

func evalWithTimer(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter Something")
	text,_ := reader.ReadString('\n')
	fmt.Printf(text)
}

func evalNormal(){
	
}