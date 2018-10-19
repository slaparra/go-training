# Concurrency

**This [page] links to resources for learning about concurrency in Go. The items are presented in order, from beginner material to advanced topics.**

[page]:https://github.com/golang/go/wiki/LearnConcurrency

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


##### Concurrency vs. parallelism

![Picture](https://birdchan.files.wordpress.com/2017/05/concurrency_vs_parallelism.png)

### Channels
Channels are a typed conduit through which you can send and receive values with the channel operator, <-.
```
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v.
```
*Channels provide a way for two goroutines to communicate with one another and synchronize their execution:*
[Channels example](../src/10-concurrency/channels.go)           

Go has a special statement called select which works like a switch but for channels:
[select example](../src/10-concurrency/select.go)           

##### Close a channel
A sender can close a channel to indicate that no more values will be sent. Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: after

```
//from: https://tour.golang.org/concurrency/4
v, ok := <-ch
```
[Close example](../src/10-concurrency/range-close.go)  

**ok** is false if there are no more values to receive and the channel is closed.  
**Note**: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.


#### Links

- [**Concurrency & goroutines, Golang Book**](https://www.golang-book.com/books/intro/10)  
- [Concurrency & channels, GolangBootCamp](http://www.golangbootcamp.com/book/concurrency)
- [Step-by-step guide to concurrency](https://yourbasic.org/golang/concurrent-programming/)
- [Go routines explained](https://yourbasic.org/golang/goroutines-explained/) 
- [GoRoutines, A Tour of go](https://tour.golang.org/concurrency/1)
- [Graceful upgrades in go](https://blog.cloudflare.com/graceful-upgrades-in-go/)

#### Vídeos  

- Visualizing Concurrency in Go, Ivan Danyliuk [(Video)](https://www.youtube.com/watch?v=KyuFeiG3Y60) [(Post)](http://divan.github.io/posts/go_concurrency_visualize/)
- [Rob Pike: Concurrency is not parallelism](https://www.youtube.com/watch?v=cN_DpYBzKso)
- [Rob Pike: Go concurrency patterns](https://www.youtube.com/watch?v=f6kdp27TYZs)
- [Sameer Ajmani: Advanced Go Concurrency Patterns](https://www.youtube.com/watch?v=QDDwwePbDtw)