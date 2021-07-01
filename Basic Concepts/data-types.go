//https://tour.golang.org/list
Data Types
      //We used int to declare integer values
var i int = 8 

      /*Other common types Go supports.
        float32 - a single-precision floating point value.
        float64 - a double-precision floating point value.
        string - a text value.
        bool - Boolean true or false.*/

      /*The difference between float32 and float64 is the precision, 
        meaning that float64 will represent the number with higher accuracy, 
        but will take more space in memory.*/

      /*In Go language, 
        the type is divided into four categories which are as follows: 
        1.Basic type     ~It includes number,string and boolean 
        2.Aggregate type ~It includes Array and structs 
        3.Reference type ~It includes Pointers, slices, maps, functions, and channels 
        4.Interface type

        Float32 is also called sometime FP32. 
        These are the capacities of memory(data) stored in the hard disk drive of the computer, also USB, or SD card. 
        We often say in bytes. For example: megabytes, gigabytes etc
 
        float32 e float64 equals float and double in other languages
 
        bool string byte (uint8 alias repr. binary data) 
        int int8 int16 int32 int64 
        uint uint8 uint16 uint32 uint64 uintptr 
        float32 float64 
        complex64 complex128 (irreal numbers) 
        error rune (int32 alias to repr. char codepoints) 
        interface (any/generic type) 
        struct (user aggregated types) Prepend brackets and you get a sliced type 
        E.g.: []int Put a number between and gets a fixed array. 
        E.g.: [5]int equiv. [...]int{0, 0, 0, 0, 0} value types: ============= nil 
        true, false (as bool) 
        'a' (as rune) 
        "text" (as string) 
        `line1 line2` (as multiline string) */
 
        /*
        float32 if you want to save memory. 
        float64 if you want to increase precision. 
                If compromising on precision is not an option, then increase your computer's RAM.*/
        //https://golang.org/pkg/time/#Time


        //https://golang.org/pkg/time/#Time
        //Characters in Go are typed by "rune", enclosed in single quotes. 
        //Although any int32 could be represented as it because both types are compatible.

        //Here are some example of the data types in code:
package main

import "fmt"

func main() {
  var x int = 42
  var y float32 = 1.37
  var name string = "James"
  var online bool = true

  fmt.Println(name)
  fmt.Println(x)
  fmt.Println(y)
  fmt.Println(online)
}

         //Since we used initializers when declaring the variables, 
         //we could have omitted the types, as they would be taken from the values assigned. 
         //We included them to demonstrate the different data types. 

         /*
         Another interesting feature of Go are zero values: 
         variables that are declared without a value take the zero value of their type:
         0 for numeric types, false for the boolean type, "" for strings.
         */
var x bool
fmt.Println(x) //output: False

          //This is a bit of a dirty trick. A zero value is the default value for any type. 
          //It's so named because an integer that has been declared (var x int) 
          //but not initialized (x = 10) will have a value of 0. 
          //For booleans or type bool, the zero value is false

          //Bool == 0 and 1 0 == false

          //0 for numeric types, false for the boolean type, "" for strings. 
          //Hence 0==false and since we our bool is not equated to 1, we have a default of 0 hence var bool ==false


