# Go Programming Gems and Classic Problems

The goal is to compile hidden gems and off-the-beaten path Go features and patterns. This is not meant to be a library or reference implementation.


 * [add_two_numbers](./add_two_numbers)
   * [main.go](./add_two_numbers/main.go)
 * [binary_uint8_reversal](./binary_uint8_reversal)
   * [main.go](./binary_uint8_reversal/main.go)
 * [caesar_cipher](./caesar_cipher)
   * [main.go](./caesar_cipher/main.go)
 * [composite_literals](./composite_literals)
    
    It is important to understand what **composite literals** mean in Go through some interesting use-cases
 
   * [main.go](./composite_literals/main.go)
 * [eratosthenes](./eratosthenes)
    
    A Classic problem but I wanted to dig deeper into the number theory behind it. I suggest reading 
    [Why in Sieve of Erastothenes of 𝑁 number you need to check and cross out numbers up to 𝑁‾‾√? How it's proved?](https://math.stackexchange.com/questions/58799/why-in-sieve-of-erastothenes-of-n-number-you-need-to-check-and-cross-out-numbe)
    
   * [main.go](./eratosthenes/main.go)
 * [for_with_label](./for_with_label)
    
    For loops can have break and continue statements with **labels**. This is a useful but less well-known feature.
 
   * [main.go](./for_with_label/main.go)
 * [go_routines_gotchas](./go_routines_gotchas)
   * [main.go](./go_routines_gotchas/main.go)
 * [is_rotation](./is_rotation)
   * [main.go](./is_rotation/main.go)
 * [longest_substring](./longest_substring)
   * [main.go](./longest_substring/main.go)
 * [median_sorted_arrays](./median_sorted_arrays)
   * [main.go](./median_sorted_arrays/main.go)
 * [one_away](./one_away)
   * [main.go](./one_away/main.go)
 * [palindromes](./palindromes)
   * [main.go](./palindromes/main.go)
 * [permutations](./permutations)
   * [main.go](./permutations/main.go)
 * [runes_bytes_loops](./runes_bytes_loops)
 
     A must-understand for serious programmers. Simple program demonstrating (sometimes tricky) concepts found in https://blog.golang.org/strings. Do you know that 'a' is a int32, str[0] is a uint8, the value of position 0 in a range loop is a rune? 

 
   * [main.go](./runes_bytes_loops/main.go)
 * [slice_tricks](./slice_tricks)
 
   Implementation of https://github.com/golang/go/wiki/SliceTricks. Simple functions that demonstrate interesting slice operations such as **extending, expanding and inserting** elements.
 
   * [main.go](./slice_tricks/main.go)
 * [two_sum](./two_sum)
   * [main.go](./two_sum/main.go)
 * [unique_chars](./unique_chars)
   * [main.go](./unique_chars/main.go)
 * [urlify](./urlify)
   * [main.go](./urlify/main.go)
 * [variadic_functions](./variadic_functions)
 
    Variadic functions in Go are certainly off-the-beaten path. One hidden gem is **interface{} variadic functions**
 
   * [main.go](./variadic_functions/main.go)
