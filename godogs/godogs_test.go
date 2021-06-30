package main

import (
         "github.com/cucumber/godog"
         "github.com/stretchr/testify/assert"
         "example.com/greetings"
       )

type scenario struct{}
func (_ *scenario) assert(a assertion, expected, actual interface{}, msgAndArgs ...interface{}) error {
    var t asserter
    a(&t, expected, actual, msgAndArgs...)
    return t.err
}

func applicationIsDeveloped() error {
	return nil
}

func (sc *scenario)display(message string) error {
	return sc.assert(assert.Equal,message, greetings.Hello("Hello World"))
}

func runApplication() error {
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^application is developed$`, applicationIsDeveloped)
	ctx.Step(`^display "([^"]*)"$`, display)
	ctx.Step(`^run application$`, runApplication)
}
