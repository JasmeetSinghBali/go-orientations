> ## GOPHER's Hub
> **Breakthrough/Guides/Walkthrough/BestPractice in Go**



---

> 1. (composing-interfaces)[https://github.com/JasmeetSinghBali/go-orientations/tree/stable/gophershub/composinginterfaces]


> 2. (when&how-to-use-pointers)[https://github.com/JasmeetSinghBali/go-orientations/tree/stable/gophershub/effectiveuseofpointers]

```
reff: https://go101.org/article/pointer.html 

# pointer
size: 8-bytes


# When

- when we need to update/mutate a state, never use pointer for a function that only act as getter


NOTE- for all other cases DONT USE POINTERS UNLESS YOU ARE GIGACHAD, bcz heavy use of pointer throughout application can cause unwanted issues Another case is Nil pointer reff are annoying and harder to debug as nil pointer can gen at runtime!

💫 
The below is fine though, bcz if *User is replaced with User then the return statement has to be
return User{},fmt.Errorf("bar") which means returning a empty User struct while that shud not be the case if no user is found also this cause waste in memory as the User struct will have default values while returning a nil pointer wont cost any memory

func getUser() (*User error){
    return nil,fmt.Errorf("bar")
}

# Escape Analysis
# for any program to know where does the variable goes in stack or heap?
go build -gcflags="-m -l"
reff: https://www.youtube.com/watch?v=ZMZpH4yT7M0&list=PL_Qhsq78MDGa7S5KoI-GA6lo6UC5-BM1q&index=8

To summarize- escape analysis is a automated allocation of the variable to the heap i.e escape to the heap to avoid dangling pointers issue in case a variable is going to get refferenced after func in which it is defined returns & in rest of the cases 'typically' it remains on the stack frame that moves up and down while replacing the invalid section variables automatically with the valid sections variables as and when needed.
while the final answer to know the location wud be only escape analysis by compiler that you can determine with go build -gcflags="-m -l"

example

func main (){
    n := answer()
    println(*n/2)
}

func answer() *int {
    x := 42
    return &x
}

Now since the x variable is going to get refferenced after the func answer() returns hence &x i.e address of x escapes to the heap.


So When are values constructed on the heap?

1. When a value could possibly be referenced after the function that constructed the value returns.
2. When the compiler determines a value is too large to fit in stack frame of that function
3. when the compler does not know the size of the value at the compile time- MOST COMMON, example a []bytes that is going to determined at runtime , so now way for compiler to know its small enough for stack frame hence automatically goes to the heap


Commonly allocated/escaped values that goes to heap

1. values shared with pointers
2. variables stored in interface variables
3. backing data for maps,channels,slices & strings([]byte)


💫 EXAMPLE

# escapes to heap                               # stay on stack
func main(){                                    func main(){
                                                    b := make([]byte,32)
                                                    read(b)
                                                    // use b
                                                }
    b := read()                                             
    // use b
}

func read() []byte{                               func read(b []byte){
                                                    //write into slice
                                                    //return this slice 
                                                }
    // return a new slice
    b := make([]byte, 32)
    return b
}   

- as every time read() is called the variable b gets refferenced after the read function returns


💡 Also the above example gives proof that why the io.Reader interface is the way it is!!
type Reader interface {
    Read(p []byte) (n int, err error)
}

# instead of , where it wud return the slice of bytes that we read which is not known hence escapes to heap
type Reader interface{
    Read(n int) (b []byte,err error)
}

So the idea is that you make a slice of bytes and pass it to the Read method of io.Reader interface which then returns the how much of the supplied slice got filled as n int so that compiler can understand the size of the slice and keep that in stack frame if it is relatively small


💫 Optimize code facts for go:

1. "Go only puts func variable on stack  if it can prove a variable is not used after the function returns" so help compiler accordingly.

2. sharing down typically stay on stack i.e pointers,refferences passed to some function remains on stack but with some exceptions

3. DONT GUESS , USE THE TOOLING !!!


```

> 3. (Generics >= v1.18 released 2022)[https://github.com/JasmeetSinghBali/go-orientations/tree/stable/gophershub/generics_1.18_2022]

- the aim of generics is to write generic function for repeatable logic based code by specifying generic type during function defination powered by custom types

```
package main

import "fmt"

// 🎈 duplicate version of the same logic of adding
func AddInt(a int, b int) int {
	return a + b
}
func AddFloat(a float64, b float64) float64 {
	return a + b
}

// 🎈

// 💫
// ~ helps to use any type associated to the underlyi9ng type CustomNum when calling AddGeneric
type NumberID int
type CustomNum interface {
	~int | ~float64
}

// 💫 generic function comes to the rescue
// with custom generic type defined  in square brackets [] & then used
func AddGeneric[T CustomNum](a T, b T) T {
	return a + b
}

func main() {
	result := AddInt(1, 2)
	result2 := AddFloat(1.1, 2.2)
	a := NumberID(7)
	b := NumberID(6)
	resultGeneric := AddGeneric(3.4, 2)
	resultAliasType := AddGeneric(a, b)
	fmt.Println(result)
	fmt.Println(result2)
	fmt.Println(resultGeneric)
	fmt.Println(resultAliasType)
}
```