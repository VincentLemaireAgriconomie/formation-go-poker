## struct and function

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

for example a card :
``` go
type card struct{
	name  string
	unicode  rune
}
```