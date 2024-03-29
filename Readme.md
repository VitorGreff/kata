# Kata 

Kata's a straightfoward library that aims to implement some functional constructs from other languages and bring them to golang. Many approaches are possible when adapting functions and methods from one language to another, but, as we are playing with functional patterns, I chose not to use any pointers at all and decided that all functions should return a completely new value/slice. 

## How to use it
Make sure you have go 1.22 or some newer version, with that said, just import the code using go module system:
<br></br>

```
go get github.com/VitorGreff/kata 
```

<hr></hr>

## Generics

Generics is a concept widely used on this library, as a way to use type inference to grant every function many ways to be used based on its parameters. This library just used 2 type constraints, aside of primitive types. which are:

|            |                                                                                                                    |
| ---------- | ------------------------------------------------------------------------------------------------------------------ |
| any        | contains every single primitive type                                                                               |
| comparable | contains every type that can be comparable with any logical/relational operator, was used as a map-key generic<br> |

<hr></hr>

## Functions Description

| METHOD    | PARAMETERS                | RETURN             | Description                                                                                                               |
| --------- | ------------------------- | ------------------ | ------------------------------------------------------------------------------------------------------------------------- |
| Map       | slice[T], func(T) T       | slice[T]           | Apply any function to every slice's element. Returns a new slice                                                          |
| Filter    | slice[T], func(T) bool    | slice[T]           | Filter a slice based on a boolean function                                                                                |
| Foldr     | slice[T], func(T, T) T, T | T                  | Reduce a slice to a single variable based on a function and an accumulator                                                |
| Compose   | ...func(T) T              | func(T)            | Returns a function made of an undefined number of functions                                                               |
| ForEach   | slice[T], func(T) T       | ----------         | Apply any function to every slice's element. Does not return a new slice                                                  |
| FindIndex | slice[T], func(T) bool    | (int, error)       | Find the index of the first element of a given slice that resolves the parameter function                                 |
| Take      | slice[T], n uint          | slice[T]           | Return a slice with **n** first elements                                                                                  |
| TakeWhile | slice[T], func(T) bool    | slice[T]           | Return a slice with **n** first elements that resolve the given function                                                  |
| Drop      | slice[T], n uint          | slice[T]           | Return a slice without **n** first elements                                                                               |
| DropWhile | slice[T], func(T) bool    | slice[T]           | Return a slice without **n** first elements that resolve the given function                                               |
| FlatMap   | matrix[T][T], func(T) T   | slice[T]           | Take a 2D slice, apply a function to every element and returns a 1D slice                                                 |
| GroupBy   | slice[T], func(T) T       | map[T] -> slice[T] | Creates a map based on a given function and correlate each element to a key                                               |
| Zip       | slice[T], slice2[T]       | matrix[T][T]       | Return a tuple array, which each stores every zip pairs. The size of the array is the size of the smaller parameter slice |

