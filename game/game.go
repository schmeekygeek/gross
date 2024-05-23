package game

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"sync"
	"time"
)

// _____
// _____
// _##$_
// _____
// _#___

func (game *Game) RunGame() {
  game.hasLost = false
  wg := sync.WaitGroup{}
  wg.Wait()
  go keyboardListen(game, &wg)
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
    if game.hasLost {
      return
    }
    game.render()
    time.Sleep(150 * time.Millisecond)
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
  game.canvas = canvas
  game.body = []Point{
    { X: 4, Y: 1, PointType: SnakeBody },
  }
  game.currDir = Right
  game.canGoUp = true
  game.food = Point{7, 5, Food}
  game.canvas[5][7] = game.food
  return game
}

func (game *Game) render() {
  for _, val := range game.body {
    game.canvas[val.X][val.Y].PointType = SnakeBody
  }
  for i := range game.canvas {
    for j := range game.canvas {
      fmt.Print(game.canvas[i][j].PointType)
      if game.canvas[i][j].PointType != Food {
        game.canvas[i][j].PointType = Ground
      }
    }
    fmt.Println()
  }
  game.canvas[game.food.X][game.food.Y].PointType = Food
  fmt.Println("Your score:", game.score)
}

func (game *Game) advance(point Point) {
  if point.X == 10 || point.Y == 10 || point.X < 0 || point.Y < 0 {
    game.gameOver()
    return
  } else if (point.X == game.food.X && point.Y == game.food.Y) {
    game.ateFood(point)
  }
  temp := point
  var temp2 Point
  for i := len(game.body) - 1; i >= 0; i-- {
    if point == game.body[i] && i != len(game.body) - 1 {
      game.gameOver()
    }
    temp2 = game.body[i]
    game.body[i] = temp
    temp = temp2
  }
}

func clearScreen() {
  cmd := exec.Command("clear")
  cmd.Stdout = os.Stdout
  cmd.Run()
}

func keyboardListen(game *Game, wg *sync.WaitGroup) {
  exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
  paused := false
  var b []byte = make([]byte, 1)
  for {
    os.Stdin.Read(b)
    input := string(b)
    if input == Down || input == Up {
      if game.canGoUp {
        game.currDir = input
        game.canGoUp = false
      }
    } else if input == " " {
      if !paused {
        wg.Add(1)
        paused = true
      } else {
        wg.Done()
        paused = false
      }
      fmt.Println("Game paused")
    } else {
      if !game.canGoUp {
        game.currDir = input
        game.canGoUp = true
      }
    }
  }
}

func (game *Game) ateFood(point Point) {
  y, x := rand.Intn(10), rand.Intn(10)
  game.food = Point{y, x, Food}
  game.canvas[point.X][point.Y].PointType = Food
  game.body = append(game.body, point)
  game.score++
}

func (game *Game) gameOver() {
  fmt.Println("Game over.\nYour score was", game.score)
  game.hasLost = true
}
