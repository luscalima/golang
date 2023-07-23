# Types

## Vigenère Cipher

> The challenge was write two programs to cipher and decipher a text using the Vigenère Cipher. To keep it simple, all characters are uppercase English letters for both the text and keyword.
>
> Examples of challenge output results can be found at [cipher output](./output_cipher.txt) and [decipher output](./output_decipher.txt).

Below is the list of topics covered by the challenge with a brief explanation of how each fits into the language.

Go is a statically typed language, which means the types are known at compile-time rather than being inferred at runtime. Each variable and expression in Go has a specific type that must be declared explicitly or implicitly during the variable declaration.

Some of the basic built-in types in Go include:

1. Numeric types:

   - int: Integer type (platform-dependent size).
   - float32, float64: Floating-point types for single-precision and double-precision floating-point numbers.

2. Boolean type:

   - bool: Represents true or false.

3. String type:
   - string: Represents a sequence of characters.

Type conversion, also known as type casting, is the process of converting a value from one type to another. In Go, explicit type conversion is required when assigning a value of one type to a variable of another type that is not directly compatible.

To convert a value from one type to another, the syntax is as follows:

```go
var a int = 42
var b float64 = float64(a) // Converting 'a' from int to float64
```
