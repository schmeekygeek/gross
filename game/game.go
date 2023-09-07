package game

import (
	"fmt"
	"os"
	"os/exec"
)

// _____
// _____
// _##$_
// _____
// _#___

func (game *Game) RunGame() {
  input := make(chan string)
  for {
    clearScreen()
    point := Point{}
    point.X = game.Body[len(game.Body)-1].X
    point.Y = game.Body[len(game.Body)-1].Y
    fmt.Println("advance point ", point)
    if game.currDir == Up {
      point.X--
    }
    if game.currDir == Right {
      point.Y++
    }
    if game.currDir == Down {
      point.X++
    }
    if game.currDir == Left {
      point.Y--
    }
    fmt.Println(point)
    game.render()
    game.advance(point)
    go keyboardListen(&input)
    game.currDir = <-input
  }
}

func InitGame() Game {
  game := Game{}
  canvas := [10][10]Point{}
  for i := 0; i < 10; i++ {
    for j := 0; j < 10; j++ {
      canvas[j][i] = Point{j, i, Ground}
    }
  }
  game.Points = canvas
  game.Body = []Point{
    {
      X: 4, Y: 3, PointType: SnakeBody,
    },
    {
      X: 4, Y: 4, PointType: SnakeBody,
    },
    {
      X: 4, Y: 5, PointType: SnakeBody,
    },
  }
  game.currDir = Up
  return game
}

func (game *Game) render() {
  for _, val := range game.Body {
    game.Points[val.X][val.Y].PointType = SnakeBody
  }
  for i := range game.Points {
    for j := range game.Points {
      fmt.Print(game.Points[i][j].PointType)
      game.Points[i][j].PointType = Ground
    }
    fmt.Println()
  }
  fmt.Printf("Your score: %v", game.score)
  fmt.Println()
}

func (game *Game) advance(point Point) {
  if point.X == 10 || point.Y == 10 || point.X < 0 || point.Y < 0 {
    fmt.Println("Game over.\nYour score was ", game.score)
    os.Exit(0)
  }
  temp := point
  var temp2 Point
  for i := len(game.Body) - 1; i >= 0; i-- {
    temp2 = game.Body[i]
    game.Body[i] = temp
    temp = temp2
  }
  fmt.Println(game.Body)
}

func clearScreen() {
  cmd := exec.Command("clear")
  cmd.Stdout = os.Stdout
  cmd.Run()
}

func keyboardListen(input *chan string) {

  exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
  // do not display entered characters on the screen
  exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

  var b []byte = make([]byte, 1)
  fmt.Println("hi")
  for {
    os.Stdin.Read(b)
    *input<-string(b)
  }
}
