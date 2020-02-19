Definition:
- Channels are a typed conduit through which you can send and receive values with the channel operator, <-
- First in, first out
- By default, sends and receives block until the other side is ready.

Buffered Channels:
- Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel:
    ch := make(chan int, 100)  
- Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.

Select:
- The select statement lets a goroutine wait on multiple communication operations.
- A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.

Example:
- ch <- v    // Send v to channel ch.
- v := <-ch  // Receive from ch, and
             // assign value to v.