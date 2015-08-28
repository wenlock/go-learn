package main

import (
//	"os"
//	"path"
    "runtime"
    "fmt"
    "./driver"
    "github.com/docker/machine/log"
)
func intro() {
  runtime.Gosched()
  fmt.Println("another test foo")
  log.Fatalf("test foo")
}

func main() {
//  _ = "breakpoint"
  go intro()
  e := ov.Ov1()
  if e != nil {
    fmt.Println("Ov1 failed:", e)
  } else {
    fmt.Println("Ov1 worked!")
  }
}
