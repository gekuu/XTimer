package main

import (
	"XTimer/TimeUtility"
	"log"
	"time"
)

// this is our main method!
func main() {
	timer := TimeUtility.NewTimerPlus(2);
	go func(tp *TimeUtility.TimerPlus) {
		// it will stop in 10 seconds.( it's just a simulation)
		time.Sleep(time.Duration(10) * time.Second)
		tp.Stop();
	}(timer)
	timer.Interval(func(argument chan TimeUtility.TimerData) {
		log.Printf("Time: %v \n", (<-argument).Time);
	});
	/*
	timer.Interval(MyInterval);
	*/
	timer.Start();
}
/*
func MyInterval(argument chan TimeUtility.TimerData) {
	log.Printf("Time: %v \n", (<-argument).Time);
}
*/


