Definition:
- Go-routines are functions or methods that run concurrently with other functions or methods.

Race Condition:
- Some goroutines access one variable and modify the value of this varible
- To check race condition, add keyword '-race'. For example, 'go run -race main.go'
- One way to solve the race condition is to pass the value of the variable into the args of the function

Wait Group:
- Synchronize multiple go routines together

Mutex:
- Lock and Unlock: protect block of codes that only one entity can be manipulating that codes at a time
- RWMutex: 
    Many entities can read data, but only one can write it at a time
    Infinite number of readers, but only one writer

runtime.GOMAXPROCS:
- runtime.GOMAXPROCS(-1): By default, returns the number of OS threads which is equal to the number of cores
- runtime.GOMAXPROCS(n): change the number of threads to n

Best Practices:
- Don't create go routines in libraries. Let consumer control concurrency
- When creating a go routine, know how it will end.
    Avoid subtle memory leaks
- Check for race conditions at compile time