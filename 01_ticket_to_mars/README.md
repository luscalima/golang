# Imperative Programming

## Ticket to Mars

> When planning a trip to Mars, it would be handy to have ticket pricing from multiple spacelines in one place. Websites exist that aggregate ticket prices for airlines, but so far nothing exists for spacelines.
>
> An example of the result of the find challenge in [output](./output.txt).

Below is the list of topics covered by the challenge with a brief explanation of how each fits into the language.

1. **Variables**:

   - Variables are used to store data in memory during the execution of a program.
   - In Go, you declare a variable using the `var` keyword, followed by the variable name and its type (explicitly or implicitly).
   - Example:

     ```go
     var age int // Declaration of an integer variable named "age"
     age = 30    // Assignment of a value to the "age" variable
     ```

2. **Constants**:

   - Constants are like variables but their values cannot be changed once defined.
   - In Go, you declare a constant using the `const` keyword, followed by the constant name and its value.
   - Example:

     ```go
     const pi = 3.14159 // Declaration of a constant named "pi" with a value of 3.14159
     ```

3. **Switch**:

   - `switch` is a control flow statement used to make decisions based on the value of an expression.
   - It provides an alternative to long sequences of `if-else` statements.
   - Example:

     ```go
     var dayOfWeek = 2
     switch dayOfWeek {
     case 1:
         fmt.Println("Sunday")
     case 2:
         fmt.Println("Monday")
     case 3:
         fmt.Println("Tuesday")
     default:
         fmt.Println("Unknown day")
     }
     ```

4. **Conditional**:

   - `if` is a conditional statement used to execute a block of code if a certain condition is true.
   - It can be used alone or in combination with `else` and `else if` clauses.
   - Example:

     ```go
     var age = 18
     if age >= 18 {
         fmt.Println("You are an adult.")
     } else {
         fmt.Println("You are a minor.")
     }
     ```

5. **Loop**:

   - `for` is a loop statement used to repeatedly execute a block of code while a condition is true.
   - It can also be used as a simple while loop by omitting the loop condition.
   - Example:

     ```go
     for i := 1; i <= 5; i++ {
         fmt.Println(i) // Prints numbers 1 to 5
     }
     ```
