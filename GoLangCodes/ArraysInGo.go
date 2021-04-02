//Without Array:
/*
package main

import (
	"fmt"
)

func main() {
	grade1 := 97
	grade2 := 85
	grade3 := 93
	fmt.Printf("Grades: %v, %v, %v ", grade1, grade2, grade3)
}
*/

//With array:
/*
package main

import (
	"fmt"
)

func main() {
	grades := [3]int{97, 85, 93} //Declaring size of array as[3] then datatype then data
	fmt.Printf("Grades: %v", grades)
}
*/

//Example2 Array
package main

import (
	"fmt"
)

func main() {
	var students [3]string
	fmt.Printf("Students: %v\n", students)
	//Inserting element now in array
	students[1] = "Lisa"
	students[2] = "Ahmad"
	students[0] = "Arnold"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Number Of Students: %v\n", len(students))
}
