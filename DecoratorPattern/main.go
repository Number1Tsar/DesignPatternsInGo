package main

import "DesignPatternsInGo/DecoratorPattern/internal"

func main(){
	p := internal.NewClient(5)
	p = internal.NewPerformerLogger(p)
	p.PerformAction()
}