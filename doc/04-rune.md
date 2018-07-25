# Rune literals
https://golang.org/ref/spec#Rune_literals
What's a rune: 
- Is Go terminology for a single Unicode code point
- An integer value identifying a Unicode code point
- It's a character
- Alias for int32 ([Numeric types](https://golang.org/ref/spec#Numeric_types))
    - 4 bytes -> 4 * 8 = 32
    - UTF-8 is a 4 byte coding scheme
    - with int 32 we have numbers for each of the code points
    
** from Todd McLeod
    
A rune literal is expressed as one or more characters enclosed in single quotes, as in 'x' or '\n'.

```
var aRuneVar rune = 'h'
```

```
fmt.Println([]byte("Hello"))

//[72 101 108 108 111] 1 byte per character
```
https://en.wikipedia.org/wiki/ASCII#Printable_characters

```
fmt.Println([]byte("世界"))

//[228 184 150 231 149 140] 3 bytes per character
```

[Rune & UTF-8 example](../todd-mcleod/01-fundamentals/rune.go)