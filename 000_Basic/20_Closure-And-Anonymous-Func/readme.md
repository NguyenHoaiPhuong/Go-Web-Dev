Definition:
- A closure is a function value that references variables from outside its body.
- The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables

5 Useful Ways to Use Closures in Go:
- Isolating data
    Create a function that has access to data that persists even after the function exists.
    Don't want anyone else to have access to that data (so they can't accidentially change it)
    For example, count how many times the function has been called, or create a fibonacci number generator
- Wrapping functions and creating middleware
    Functions in Go are first-class citizens. What this means is that you can not only create anonymous functions dynamically, but you can also pass functions as parameters to a function. For example, when creating a web server it is common to provide a function that processes a web request to a specific route.
    Middleware is basically a fancy term for reusable function that can run code both before and after your code designed to handle a web request.
- Accessing data that typically isn’t available
    A closure can also be used to wrap data inside of a function that otherwise wouldn’t typically have access to that data.
- Binary searching with the sort package
    Closure are also often needed to use packages in the standard library, such as the sort package.
- Deferring work

Explanation:
- 01_Get-Started: The add() function returns a closure. Each closure is bound to its own sum variable
- 02_Isolating-Data: the func Fibonacci() return a function (a closure) that returns successive fibonacci numbers {0, 1, 1, 2, 3, 5...}
- 03_Wrapping-Functions: func timed returns a closure function which is used in http.HandleFunc
- 04_Accessing-Unavailable-Data: Handler function hello receives the arg Database
- 05_Sort-Package
- 06_Deffering-Work