Feature: Display Hello World
  In order to test godog
  As a developer
  I need to display "Hi, Stein , Welcome"
  
  Scenario: Display Hello World
    Given application is developed
    When run application
    Then display "Hi, Stein. Welcome!"
     