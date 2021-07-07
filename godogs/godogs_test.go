package main

import (
	     "os"
         "github.com/cucumber/godog"
         "github.com/stretchr/testify/assert"
         "testing"
         "fmt"
         "github.com/cucumber/godog/colors"
	      flag "github.com/spf13/pflag"
       )


func applicationIsDeveloped() error {
	return nil
}


func display(expectedMessage string) error {
	
	var actualMessage =  hello("Stein")
	return  assertExpectedAndActual(
		assert.Equal, expectedMessage, actualMessage,
		"Expected messag to be %d  but actualMessage  is %d", 
		expectedMessage, actualMessage,
	)
	
}


func runApplication() error {
	return nil
}


func InitializeScenario(ctx *godog.ScenarioContext) {
	
	ctx.Step(`^application is developed$`, applicationIsDeveloped)
	ctx.Step(`^display "([^"]*)"$`, display)
	ctx.Step(`^run application$`, runApplication)
	
}

//Linking Godog to testify

var opts = godog.Options{
	Output: colors.Colored(os.Stdout),
	Format: "progress", // can define default values
}

func init() {
	godog.BindCommandLineFlags("godog.", &opts)        // godog v0.11.0 (latest)
}

func TestMain(m *testing.M) {
    flag.Parse()
	opts.Paths = flag.Args()

    status := godog.TestSuite{
        Name: "godogs",
		ScenarioInitializer:  InitializeScenario,
		Options: &opts,     
	   }.Run()
	   
	if st := m.Run(); st > status {
		status = st
	}
	    os.Exit(status)
}


func assertExpectedAndActual(a expectedAndActualAssertion, expected, actual interface{}, msgAndArgs ...interface{}) error{
	
	var t asserter
	a(&t, expected, actual, msgAndArgs...)
	return t.err
	
}


type expectedAndActualAssertion func(t assert.TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool


type asserter struct {
	err error
}


func (a *asserter) Errorf(format string, args ...interface{}) {
	
	a.err = fmt.Errorf(format, args...)
	
}

