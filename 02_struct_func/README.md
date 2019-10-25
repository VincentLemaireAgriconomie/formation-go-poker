## struct and function

#### struct

It's time to start our poker game.

what do we need for a good poker game ?

- a board with
    - a deck
        - cards 
    - some players
        - a stack of pieces
        - a bet
        - some cards
        
first create [the structs](https://gobyexample.com/structs) 

for the cards, 

for example a card :
``` go
type player struct {
	name  string
    ...
}
```

to make an array, the syntax is `[]string`

for the cards, the simplest way is to make an alias with the rune as a rune is an unicode char and cards game are available in unicode.

``` go
 type card = rune 
```

when the types are good, we need to instanciate : golang doesn't have constructor, 

we have two way : 

```go
var p player
p.name = "A"
```

or fastest, only if p doenst exist : 

```go
p := player{
    name:a
}
```

to check the content of your object, the simplext way is : ( [doc](https://golang.org/pkg/fmt/) )

```go
fmt.Printf("%+v\n",p)
```

as you can see, the values that you dont defined have a default one : nil for the array, 0 for the int, "" for strings. 
you cannot get a go objecct that is not valid.


#### functions

the first function that we want is to draw a cart from a deck, as go is strongly typed, we need to tell what do we expect as input and what do we guarantee as output.

```go
func draw(deck []card) card{
	return '🃕'
}
```

or we can add a function to a type :

```go

func(d deck) draw() card{
	return '🃕'
}

```
