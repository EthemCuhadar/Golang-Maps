# Golang

![Golang Image](golang.png)

---------------------------------------------------------------------

## Slices

* Maps are another composite type in Go programming language. Working principle of maps depends on **key** and **value** like Python.
* **Key** and **value** types should be decleared while defining a map.

```go
var myMap map[string]int
fmt.Println(myMap)
```

```console
map[]
```

* Value type would be any type in Go programming language like int, string, slice etc.

```go
myMap2 := map[string][]string{
    "Nick": []string{"cat", "car", "dog"},
    "Joe": []string{"house", "pen"},
}
fmt.Println(myMap2)
```

```console
map[Joe:[house pen] Nick:[cat car dog]]
```

* In this example, the types of values are **string slices**. Let's check the types to see more clearly.

```go
a := myMap["Nick"]
fmt.Printf("The type of value is: %T\n", a)
```

```console
The type of value is: []string
```

-----------------------------------------------------------

## Reading and writing a map

* Now let's look at reading and writing a map. What we will do is, simple read and print the keys and values. After that we will change the values if we can do. Let's create student notes in a class.

```go
myMap3 := map[string]int{}

myMap3["Nick"] = 90
myMap3["Joe"] = 70
myMap3["Chloe"] = 50
myMap3["Sarah"] = 100

fmt.Println(myMap3)
```

```console
map[Chloe:50 Joe:70 Nick:90 Sarah:100]
```

* Now let's think that teacher had some mistake with notes and he or she wants to change some notes.

```go
myMap3["Nick"] = 80
myMap3["Chloe"] = 65

    // Print the notes again.
    fmt.Println(myMap3)
```

```consoleag-0-1ftn1b9fnag-1-1ftn1b9fn
map[Chloe:65 Joe:70 Nick:80 Sarah:100]
```

* Now Nick's and Chloe's notes have changed.

* Let's check a student's note who is not in the class.

```go
fmt.Println(myMap3["Emily"])
```

```
0
```

* As you can notice, the program does not give an error. On the contrary, it gives 0 which is not in the map. 

* Let's delete some students in the class. To do that, **delete()** built-in function can be used.

```go
delete(myMap3, "Nick")
fmt.Println(myMap3)
```

```console
map[Chloe:65 Joe:70 Sarah:100]
```

----------------------------------------------

## Checking map elements

* Sometimes a map may have hundreds elements and we would check whether if the element is exist or not.

* Go programming language provides very practical usage for that. Let's look at example below.

```go
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
```

```console
39 true
```

* USA is the top 5 golod medal country in Tokio 202 Olympics. Hence,  the statement is true and 39 is the medal number of USA.

* Let's look at another country that is not top 5 gold medal Country.

```go
// Check France for medals.
vFrance, ok := olympicMedals["France"]
fmt.Println(vFrance, ok)
```

```console
0 false
```

* As you have noticed, France is not in top 5 gold medal list. Therefore, statement is **false** and the map has no idea about how many gold medals France have. Hence, number of medals is 0. 
