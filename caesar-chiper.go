package main

import (
  "fmt"
  "os"
  "strconv"
)

/**
 * \run    $ go run caesar-chiper.go [OFFSET] [PESAN]
           $ go run caesar-chiper.go 2 "hello"
 * \hasil  "jgnnq"
 */

func shifted_char(char byte, offset int) byte {
  const base = "abcdefghijklmnopqrstuvwxyz1234567890"
  if( char>='a' && char<='z' ) {
    char -= 'a'
  } else if( char>='1' && char<='9' ) {
    char = (char-'1')+26
  } else if( char=='0' ) {
    char = 35
  }
  var appro = (int(char)+offset)%36
  if( appro<0 ) {
    char = byte(36+appro)
  } else {
    char = byte(appro)
  }
  return base[char]
}

func main() {
  offset, ret := strconv.Atoi(os.Args[1])
  if( ret!=nil ) {
    os.Exit(1)
  }
  var input = os.Args[2];
  for i:=0; i<len(input); i++ {
    if( (input[i]>='a' && input[i]<='z') ||
        (input[i]>='0' && input[i]<='9') ) {
      fmt.Printf("%c", shifted_char(input[i], offset))
    } else {
      fmt.Printf("%c", input[i])
    }
  }
  fmt.Printf("\n")
}
