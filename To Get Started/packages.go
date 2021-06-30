Packages
      //Every Go program is made up of packages.
      //A Go program starts running in the main package.
      //This is why we need to declare our code as the main package 
      //-- to make it run and create the output:
package main //the package is used when you run a Go program

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
      //The first line of the code defines that this is the main package.

      //https://medium.com/golangspec

Imports
      //Apart from main, Go has many packages 
      //that can be imported and used in the code to accomplish different tasks.

      //One of the most popular packages is "fmt", 
      //which stands for format, and provides input and output functionality.

      //To import a package, we use the import statement
      
      //You can import multiple packages at once, using parentheses. 
      //For example:
import (
"fmt"
"math"  //to import the math pack
)

      //Example:
package main 
import(
     "fmt"
     "math"
)
func main() { 
    
    c := math.Exp2(7) 

fmt.Println(c)
} //output: 128

      //You can import multiple packages using parenthesis notation. 
      //import "fmt" import "strings" is the same as import ("fmt" "strings") 
      //New lines are important. Quotes too! Is the way of distinguish from a variable/alias. 
      //By default, to access to content you uses the same name as variable 
      //but you can also define an alias for a package to avoid collisions: 
import ( "strings" mystr "myprogram/strings" ) 

      //Even, able to import all content as globals (but take care with collisions). 
      //Use dot as alias: 
import ( . "fmt" . "math" ) 

      //One more tip!! External packages are referenced by its repo url 
      //and underscore notation can be used to avoid declare variable (usefull bootstraping drivers) 
import ( . "fmt" _ "github.com/awesome-drivers/go-driver" "time" "sort" strs "strings" ) 

      //Explore core and more community packages at: - https://golang.org/pkg/ - https://pkg.go.dev/




