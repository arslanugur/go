Imports
      //Each package has exported names, which you can access after importing them.

      //In Go, a name is exported if it begins with a capital letter.
      //You can access the exported names using the package name, a dot, and the exported name.

      //In our case, we access the Println() function of the "fmt" package, 
      //which is used to generate the output:
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
      //The Println() is an exported name, which is why it starts with a capital letter.
      //We provide the output we want to generate inside the parentheses enclosed in quotes.

      //In Go, a name is exported if it begins with a capital letter, 
      //so function, structure, etc can be accessed in other files. 
      //We have neither classes nor access modifiers in Go, 
      //but when we name a function, etc with a exported name, 
      //that function become public but when we don't, that function can't be accessed in others files (private). 
      //There are no classes in Go, so there's no 'protected' access modifier.

      //Example:
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
    fmt.Println("5+7")
    fmt.Println (5+7)
    fmt.Println ("5"+"7")
    fmt.Println ('a' + 'b')
    fmt.Println ('A' + 'B')
    fmt.Println ('*')
    fmt.Println ('&')
}

      //Similar to other programming languages 
      //such as Java or C++, func main() is the entry point of our program. 
      //It is the function that gets executed when we run our program. 

      /*You can see that the code:
        1. defines the main package
        2. imports the "fmt" package used for output
        3. defines the main() function.
        4. Uses the Println() function of the fmt package to generate the desired output.

        We will learn more about functions in the coming modules.*/


        
