# Concurrency

**This [page] links to resources for learning about concurrency in Go. The items are presented in order, from beginner material to advanced topics.**

[page]:https://github.com/golang/go/wiki/LearnConcurrency

[](https://yourbasic.org/golang/goroutines-explained/)
*Making progress on more than one task simultaneously is known as concurrency. Go has rich support for concurrency using goroutines and channels.*  
```
import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
/*
    Will print:
    world
    hello
    hello
    world
    world
    hello
    hello
    world
    world
    hello
*/
```

With a goroutine we return immediately to the next line and don't wait for the function to complete.

### Channels
Channels are a typed conduit through which you can send and receive values with the channel operator, <-.
```
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v.
```
           
#### Links

- [Concurrency & goroutines, Golang Book](https://www.golang-book.com/books/intro/10)  
- [Concurrency & channels, GolangBootCamp](http://www.golangbootcamp.com/book/concurrency)
- [GoRoutines, A Tour of go](https://tour.golang.org/concurrency/1)

#### Vídeos  

- Visualizing Concurrency in Go, Ivan Danyliuk [(Video)](https://www.youtube.com/watch?v=KyuFeiG3Y60) [(Post)](http://divan.github.io/posts/go_concurrency_visualize/)
- [Rob Pike: Concurrency is not parallelism](https://www.youtube.com/watch?v=cN_DpYBzKso)