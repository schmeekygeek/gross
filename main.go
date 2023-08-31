package main

import (
  "snake/game"
  "fmt"
  "os"
  "os/exec"
)

func main() {
  instance := game.InitGame()
  pt1 := game.Point{
    X: 0,
    Y: 1,
    PointType: game.Ground,
  }
  pt2 := game.Point{
    X: 1,
    Y: 2,
    PointType: game.Ground,
  }
  instance.Advance(pt1)
  fmt.Println(instance.Body)
  instance.Advance(pt2)
  fmt.Println(instance.Body)
}


func ListenEvents() {

  exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
  // do not display entered characters on the screen
  exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

  var b []byte = make([]byte, 1)
  fmt.Println("hi")
  for {
    os.Stdin.Read(b)
    fmt.Println("got ", string(b))
  }
}
