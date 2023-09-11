package game

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

// _____
// _____
// _##$_
// _____
// _#___

func (game *Game) RunGame() {
  go keyboardListen(game)
  for {
    clearScreen()
    point := Point{}
    point.X = game.body[len(game.body)-1].X
    point.Y = game.body[len(game.body)-1].Y
    if game.currDir == Up {
      point.X--
    } else if game.currDir == Down {
      point.X++
    } else if game.currDir == Right {
      point.Y++
    } else {
      point.Y--
    }
    game.advance(point)
    game.render()
    time.Sleep(time.Second)
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
  game.points = canvas
  game.body = []Point{
    { X: 4, Y: 0, PointType: SnakeBody },
    { X: 4, Y: 1, PointType: SnakeBody },
    { X: 4, Y: 2, PointType: SnakeBody },
    { X: 4, Y: 3, PointType: SnakeBody },
    { X: 4, Y: 4, PointType: SnakeBody },
  }
  game.currDir = Right
  game.canGoUp = true
  return game
}

func (game *Game) render() {
  for _, val := range game.body {
    game.points[val.X][val.Y].PointType = SnakeBody
  }
  for i := range game.points {
    for j := range game.points {
      fmt.Print(game.points[i][j].PointType)
      game.points[i][j].PointType = Ground
    }
    fmt.Println()
  }
  fmt.Println("Your score:", game.score)
}

func (game *Game) advance(point Point) {
  if point.X == 10 || point.Y == 10 || point.X < 0 || point.Y < 0 {
    game.gameOver()
  }
  temp := point
  var temp2 Point
  for i := len(game.body) - 1; i >= 0; i-- {
    temp2 = game.body[i]
    game.body[i] = temp
    temp = temp2
  }
  for i := 0; i < len(game.body) - 1; i++ {
    if point == game.body[i] {
      game.gameOver()
    }
  }
}

func clearScreen() {
  cmd := exec.Command("clear")
  cmd.Stdout = os.Stdout
  cmd.Run()
}

func keyboardListen(game *Game) {
  exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
  // do not display entered characters on the screen
  exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

  var b []byte = make([]byte, 1)
  for {
    os.Stdin.Read(b)
    input := string(b)
    if input == Down || input == Up {
      if game.canGoUp {
        game.currDir = input
        game.canGoUp = false
      }
    } else {
      if !game.canGoUp {
        game.currDir = input
        game.canGoUp = true
      }
    }
  }
}

func (game *Game) gameOver() {
    fmt.Println("Game over.\nYour score was ", game.score)
    os.Exit(0)
}
