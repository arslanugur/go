Variables
      //Variables are used to store values.

      //In Go, the var keyword is used to declare variables.
      //For example:
var i int

      //The code above declares a variable named i of type int.
      //int stands for integer and represents whole numbers.

      //We can assign the variable a value and output it:
package main

import "fmt"

func main() {
  var i int = 8
  fmt.Println(i) 
} //output: 8

      //We will learn about the other variable types that Go supports in the next lesson.


      //to declare a variable called "x" of type int, assign it the value 42 and output it.
  var x int = 42
  fmt.Println(x)

      /*A variable is just a piece of code that *is* something (as opposed to functions, which *do* something). 
        Variables hold a value. We call it variable because you can change the value after declaring it. 
        For instance, after declaring var x int = 42, later we can change our mind and say x = 20 */

      //You can also declare multiple variables on one line and assign values to them:
package main

import "fmt"

func main() {
  var i, j int = 8, 42
  fmt.Println(i) 
  fmt.Println(j)
} //output: 8, 42

      //If you assign a value to the variable, you can omit the type declaration, 
      //as Go will automatically take the type of the value assigned
package main

import "fmt"

func main() {
  var i, j = 8, 42
  fmt.Println(i) 
  fmt.Println(j)
}  //output: 8, 42
      //The process of assigning values to variables is called "initialization".

      //Initialization - process of assigning values to the variables
      //We can declare multiple variables on one line and we can assign values respectively. 
      //Like var x, y int = 10, 12 So here x = 10 and y = 12 
      //We can ommit int, as it will automatically take the type. var x, y = 10, 12

      //Go supports short variable declaration using :=
package main

import "fmt"

func main() {
  i := 8
  fmt.Println(i)
} 
      //It can also be used to declare and initialize multiple variables on one line:
package main

import "fmt"

func main() {
  x, y := 10, 5 
  fmt.Println(x)
  fmt.Println(y)
} 

      //The := operator automatically defines and initialized variables with the given value.
      //As if it is Python's walrus operator.

      //There are 3 ways of declare/initializing variables. 
      //var x | var y = x = | z := with any, 
      //you can also handle multiple in the same line (except when type doesn't match / mixing ways) 
      //var x, y int = 5, 6 xx, yy := 7, 8 var a, b int, c string = 69, 96, "WRONG"! // Error i, j := 5 
      // Error. 2 vars, 1 value A variable can contain a value of other type only if one are super equivalent of another. 
      //Eg. int / int32 / rune Test it + 4 BONUS 
      //Example:
//Variable declaration/initialization cases
package main

import . "fmt"

// global variables always start with "var"
var glo1 string = "I'm a global variable not exportable"
var Glo1 string = "I can be visible in other packages"
// constants
const C1 = 10

func main() {
  // 1. single line declare + single init
  var a int
  a = 3
  
  // 2. single line declare/init
  var aa int = 4
  
  // 3. One line multi declare, single init
  var b, bb int
  b = 33
  bb = 34
  
  // 4. One line declare + one line init
  var c, cc int
  c, cc = 63, 64
  
  // 5. One line multi declare/init
  var d, dd int = 93, 94
  
  // 6. One line declare/init no-types
  var e = 122
  var ee, eee = 123, 124
  
  // 7. One line no-var/no-type (shortened)
  g := 'G'
  gg, ggo := 'O', '🐹'
  
  Println(`00. ------------------
     `, glo1, C1, Glo1,
      `
01. ------------------
     `, a,
      `
02. ------------------
     `, aa,
      `
03. ------------------
     `, b, bb,
      `
04. ------------------
     `, c, cc,
      `
05. ------------------
     `, d, dd,
      `
06. ------------------
     `, e, ee, eee,
      `

07. SHORTEN ----------
     `, g, gg, ggo)
  
  
  // BONUS 8. Assign by reference: "&"
  var rgo int
  _, err := Scan(&rgo)
  
  Println(`
08. INPUT BY REF -----
     `, err, rgo)
  
  
  // BONUS 9. Omit positional assign: "_"
  var go1, go2, go3, go4 rune
  _, go1 = '🐹', 'G'
  go2, _, go3, _ = 'o', '!', '🐹', '!'
  
  Println(`
09. POSITIONAL -------
     `, go1, go2, go3,
      // uninitialized
      go4)
  
  
  // BONUS 10. swap
  go1, go3 = go3, go1
  
  Println(`
10. SWAP ABOVE -------
     `, go1, go2, go3, go4)
  
  
  // BONUS 11. retyping??
  // variables cannot change runtime type, in any of the ways you declare them. An exception are equivalent types
  var vv int = 71
  var rr rune = 'o'
  Printf(`
11. RETYPE -----------
  Before:
      vv%v
      rr%v
`, describe(vv), describe(rr))
  //vv = "GO"        // compile error
  vv, rr = '🐹', 33  //OK rune eqv int32/int
  Printf(`  After:
      vv%v
      rr%v
`, describe(vv), describe(rr))
} 


// Describe any variable kind, printing their type, value and then quoted value
func describe(v interface{}) string {
    return Sprintf(" <%5[1]T> = %6[1]v  %8[1]q", v)
}



