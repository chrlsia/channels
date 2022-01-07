package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

// Declare that KeyPressChan will be chan rune
var keyPressChan chan rune //1


func main() {
	// Use keyPressChan
	keyPressChan = make(chan rune) //2
	go listenForKeyPress()//3

	fmt.Println("Press any key or q to quit")//4

	//5
	_=keyboard.Open()
	defer func(){
		keyboard.Close()
	}()
	
	//6
	for{
		//get a key and if it is q or Q quit
		char, _, _:=keyboard.GetSingleKey()
		if char == 'q' || char=='Q' {
			break
		}
		//key is not q or Q so push it into channel
		keyPressChan <- char
	}

}

func listenForKeyPress(){
	//7
	for{
		// push the conetens of channel to key
		// and display it
		key:= <-keyPressChan
		fmt.Println("You pressed ",string(key))
	}

}

