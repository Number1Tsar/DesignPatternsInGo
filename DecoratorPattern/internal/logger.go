package internal

import (
	"log"
	"time"
)

// performerLogger is a decorator of Performer. It is a Performer and it also has a Performer.
type performerLogger struct {
	performer Performer
}

// returns instance
func NewPerformerLogger(performer Performer) Performer {
	return &performerLogger{performer:performer}
}

func(pl *performerLogger)PerformAction(){
	start := time.Now()
	pl.performer.PerformAction()
	log.Printf("Time taken to execute method %v\n",time.Since(start))
}
