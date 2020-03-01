# Go Programming Gems and Classic Problems

The goal is to compile hidden gems and off-the-beaten path Go features and patterns. This is not meant to be a library or reference implementation.

## Gems and Off-the-beaten Path

 * [closures](./closures)
    
    A dramatic example of Go Closures. A two-level closure pattern.

 * [composite_literals](./composite_literals)
    
    It is important to understand what **composite literals** mean in Go through some interesting use-cases

 * [for_with_label](./for_with_label)
    
    For loops can have break and continue statements with **labels**. This is a useful but less well-known feature.
    
  * [runes_bytes_loops](./runes_bytes_loops)
  
      A must-understand for serious programmers. Simple program demonstrating tricky [concepts](https://blog.golang.org/strings). Do you know that 'a' is a int32, str[0] is a uint8, the value of position 0 in a _range loop_ is a rune? 
      
   * [slice_gotchas](./slice_gotchas)
   
   Do you know that this does not cause an out of bounds error?
   
 ```go
 	c := make([]int, 3)
 	fmt.Println(c[len(c):len(c)])
 ```

 * [slice_tricks](./slice_tricks)
 
   Implementation of simple functions that demonstrate interesting slice operations such as [**extending, expanding and inserting**](https://github.com/golang/go/wiki/SliceTricks) elements.
   
  * [variadic_functions](./variadic_functions)
  
     Variadic functions in Go are certainly off-the-beaten path. One hidden gem is **interface{} variadic functions**
     
## Classic Problems 

 * [binary_uint8_reversal](./binary_uint8_reversal)
 * [caesar_cipher](./caesar_cipher)
 * [eratosthenes](./eratosthenes)
    
    A Classic problem but I wanted to dig deeper into the number theory behind it. I suggest reading 
    [Why in Sieve of Erastothenes of 𝑁 number you need to check and cross out numbers up to 𝑁‾‾√? How it's proved?](https://math.stackexchange.com/questions/58799/why-in-sieve-of-erastothenes-of-n-number-you-need-to-check-and-cross-out-numbe)
    
 * [go_routines_gotchas](./go_routines_gotchas)
 * [is_rotation](./is_rotation)
 * [longest_substring](./longest_substring)
 * [median_sorted_arrays](./median_sorted_arrays)
 * [one_away](./one_away)
 * [palindromes](./palindromes)
 * [permutations](./permutations)
 * [two_sum](./two_sum)
 * [unique_chars](./unique_chars)
 * [urlify](./urlify)

 
