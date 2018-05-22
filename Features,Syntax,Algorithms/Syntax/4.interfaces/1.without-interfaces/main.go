/*
 * We know  |-> Every value HAS A TYPE
 *          |-> Every function HAS TO SPECIFY the type of its arguments
 *
 * hmm.. then,
 * Every function we ever write HAS TO BE REWRITTEN to accommodate different type?
 * Even if the logic in it is IDENTICAL ?
 */

/*
 * example :
 *
 *   |  	English Bot			       			  |		Spanish Bot                           |
 *   |--------------------------------------------|-------------------------------------------|
 * 1 | type englishBot struct                     |  type spanishBot struct					  |
 * 2 | func (englishBot) getGreeting() string {}  |  func (spanishBot) getGreeting() string {}|
 * 3 | func printGreeting(eb englishBot) {}       |  func printGreeting(sb spanishBot) {}     |
 *
 *
 * note :
 * 2 -> probably have DIFFERENT LOGIC in these func. its ok.
 * 3 -> THAT'S IT. These probably have IDENTICAL logic!
 */

package main

import (
	"fmt"
)

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreetingEB(eb)
	printGreetingES(sb)
}

func (englishBot) getGreeting() string {
	return "this is english bot => Hello!"
}

func (spanishBot) getGreeting() string {
	return "esto es bot español => Hola!"
}

func printGreetingEB(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

// printGreetingES has identical logic like printGreetingEB. Redundant!
func printGreetingES(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

// now  imagine we have 100 different bots. Duh! 🙄
