package main

import (
  "fmt"
  "os"
  // Won't compile with unused packages
  //"io"
  //"strings"
)

func main() {

  /* Basics */
  const name, id = "Nanako", 17
  err := fmt.Errorf("user %q (id %d) not found", name, id)
  // Generous error printing behavior
  fmt.Println(err)
  // user "Nanako" (id 17) not found
  fmt.Println(err.Error())
  // user "Nanako" (id 17) not found
  
  const age = 22
  // Formatting
  n, err := fmt.Fprint(os.Stdout, name, " is actually ", age, " years old.\n")
  fmt.Fprintln(os.Stdout, name, "is actually", age, "years old.")
  // Nanako is actually 22 years old.
  fmt.Fprintf(os.Stdout, name, " is actually ", age, " years old.\n")
  // Nanako%!(EXTRA string= is actually , int=22, string= years old.
  fmt.Fprintf(os.Stdout, "%s is actually %d years old.\n", name, age)
  // Nanako is actually 22 years old.

  if err != nil {
    fmt.Fprintf(os.Stderr, "print: %v\n", err)
  }
  fmt.Print(n, " bytes written.\n")
  // 33 bytes written.
  fmt.Printf("%d bytes written.\n", n)
  // 33 bytes written.

}