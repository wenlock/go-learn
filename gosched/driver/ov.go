package ov
import (
  "fmt"
  "runtime"
  "github.com/docker/machine/log"
)

func Ov1() (error) {
  runtime.Gosched()
  fmt.Println("in Ov1 another test foo")
  log.Fatalf("the force is with you")
  return nil
}
