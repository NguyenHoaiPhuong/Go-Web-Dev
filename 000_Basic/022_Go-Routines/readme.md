Definition:
- Go-routines are functions or methods that run concurrently with other functions or methods.

Race Condition:
- Some goroutines access one variable and modify the value of this varible
- To check race condition, add keyword '-race'. For example, 'go run -race main.go'
- One way to solve the race condition is to pass the value of the variable into the args of the function

Wait Group:
- Synchronize multiple go routines together

Mutex:
- Lock and Unlock: protect part of codes that one entity can be manipulating that codes at a time
- RWMutex: 
    Many entities can read data, but only one can write it at a time
    Infinite number of readers, but only one writer
