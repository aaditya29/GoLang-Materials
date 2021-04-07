/*An interface is a collection of method signatures that a Type can implement(using methods).
Hence interface defines (not declares) the behavior of the object (of the type Type).

For example, a Dog can walk and bark.
If an interface defines method signatures for walk and bark while Dog implements walk and bark methods,
then Dog is said to implement that interface.
*/

package main

import "fmt"

type Shape interface { //Declaring Static Interface
	Area() float64
	Perimeter() float64
}

func main() {
	var s Shape //A variable of an interface type can hold a value of a type that implements the interface
	fmt.Println("value of s is", s)
	fmt.Printf("type of s is %T\n", s)
	//var s will print nil because at this moment we have declared s but did not assign any value
}
