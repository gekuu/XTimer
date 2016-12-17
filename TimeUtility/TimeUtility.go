package TimeUtility

import (
	"time"
	"log"
)

var _delegate func(chan TimerData)
var _data chan TimerData;

var _isDone bool;

type TimerPlus struct {
	Second int
	Data   chan string
}

type TimerData struct {
	Time time.Time
}

func NewTimerPlus(second int) *TimerPlus {
	return &TimerPlus{Second:second};
}

func (t TimerPlus) Interval(delegate func(chan TimerData)) {
	_delegate = delegate;
}

func (t TimerPlus) Start() {
	_data = make(chan TimerData,1);
	for {
		if (!_isDone) {
			_data <- TimerData{Time:time.Now()};
			_delegate(_data);
			time.Sleep(time.Duration(t.Second) * time.Second)
		} else {
			break;
		}
	}
}

func (t TimerPlus) Stop() {
	if (_data != nil) {
		log.Print("DONE!\n");
		close(_data);
		_isDone = true;
	}
}


