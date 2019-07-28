# Variable 

## Variable Declaration
- Variable Init ":=" 
- Variable Assigned "="

## Functions call in the same package
- Files in the same package can freely call functions defined in other files. Check "state.go"

```
go run cards\main.go cards\state.go
```

## Slices and For Loops
- Slice: is an array that can grow and shrink.
- Slice: every element in a slice must be of same type.
- Slice range: [startIndexIncluding:upToNotIncluding] --> ```card[:2]``` --> ```card[0] & card[1]```
- For Loop: the "index" variable must be used somewhere within the loop.

## OO Approach vs GO Approach
- Custom Type Declaration: check out ```deckCards``` in "main.go" and ```deck``` data type in "deck.go"
```
go run cards\main.go cards\state.go cards\deck.go
```
## Random Number Generation
- Create NewSource with different Seed to produce different random number --> Read the pkg documentations