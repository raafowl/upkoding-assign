package main

import (
  "fmt"
  "os"
)

/**
 * \run     $ go run count-unique-char.go [INPUT]
 *          $ go run count-unique-char.go "UUpPkodingGG"
 * \hasil
 *          huruf U : 2
 *          huruf p : 1
 *          huruf P : 1
 *          huruf k : 1
 *          huruf o : 1
 *          huruf d : 1
 *          huruf i : 1
 *          huruf n : 1
 *          huruf g : 1
 *          huruf G : 2
 */

func main() {
  var input = os.Args[1];
  var magic[0x100] uint32
  for i:=0; i<len(input); i++ {
    if( (input[i]>='A' && input[i]<='Z') ||
        (input[i]>='a' && input[i]<='z') ) {
      magic[input[i]]++
    }
  }
  for i:=0; i<len(input); i++ {
    if( magic[input[i]]!=0 && (magic[input[i]]>>31)==0 ) {
      fmt.Printf("huruf %c : %d\n", input[i], magic[input[i]])
      magic[input[i]] |= 1<<31
    }
  }
}
