package logicforsignal

import (
	"fmt"
	ma "programm/Ma/MaFirst"
)

func DetermineMovingAverageSignal(){
	shortMA := ma.MovingAverage()
	longMA := ma.MovingAverageSecond()
    if shortMA > longMA {
        fmt.Println("buy")
    } else if shortMA < longMA {
        fmt.Println("sell")
    } else {
        fmt.Println("waight")
    }
}