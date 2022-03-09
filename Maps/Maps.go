package main

import (
	"fmt"
)

func main(){
	
	// Define a map with "var" keyword.
	var myMap map[string]int
	
	// Print the map that is created above.
	fmt.Println(myMap)
	
	// Define another map with short-hand decleration.
	myMap2 := map[string][]string{
		"Nick": []string{"cat", "car", "dog"},
		"Joe": []string{"house", "pen"},
	}
	
	// Print the map.
	fmt.Println(myMap2)
	
	// Check the type of value.
	a := myMap2["Nick"]
	fmt.Printf("The type of value is: %T\n", a)
	
	// READING and WRITING MAPS
	
	// Create students exam notes.
	myMap3 := map[string]int{}
	
	// Assign student notes with names.
	myMap3["Nick"] = 90
	myMap3["Joe"] = 70
	myMap3["Chloe"] = 50
	myMap3["Sarah"] = 100
	
	// Print the notes.
	fmt.Println(myMap3)
	
	// Reassing some notes
	// Change Nick's and Chloe's notes.
	myMap3["Nick"] = 80
	myMap3["Chloe"] = 65
	
	// Print the notes again.
	fmt.Println(myMap3)
	
	// Check a student's note who is not in the class.
	fmt.Println(myMap3["Emily"])
	
	// Delete some students.
	delete(myMap3, "Nick")
	
	// Print myMap3 again.
	fmt.Println(myMap3)
	
	// CHECK an ELEMENT EXISTANCE.
	
	// Define top 5 gold medals in Tokio 2020 Olympic Games.
	olympicMedals := map[string]int{
		"USA": 39,
		"China": 38,
		"Japan": 27,
		"UK": 22,
		"ROC": 20,
	}
	
	// Check if USA is top 5 county in Olympics
	// And if yes, give how many medals USA had.
	vUSA, ok := olympicMedals["USA"]
	fmt.Println(vUSA, ok)
	
	// Check France for medals.
	vFrance, ok := olympicMedals["France"]
	fmt.Println(vFrance, ok)
}
