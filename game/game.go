package game

import "fmt"

// _____
// _____
// _##$_
// _____
// _#___

func (game *Game) RunGame() {
  for {

  }
}

func InitGame() *Game {

  game := Game{}
  game.Body = make([]Point, 3)
  for i := 0; i < 5; i++ {
    for j := 0; j < 5; j++ {
      game.Points[i][j].PointType = Ground
    }
  }
  game.Body[0].PointType = SnakeBody
  game.Body[1].PointType = SnakeBody
  game.Body[2].PointType = SnakeHead
  game.canGoUp = false
  return &game
}

func (game *Game) Render() {
  for i := range game.Points {
    for j := range game.Points {
      fmt.Print(game.Points[i][j].PointType)
    }
    fmt.Println()
  }
  fmt.Printf("Your score: %s", fmt.Sprint(game.score))
  fmt.Println()
}

func (game *Game) Advance(point Point) {
  temp := point
  var temp2 Point
  for i := 0; i < len(game.Body); i++ {
    temp2 = game.Body[i]
    game.Body[i] = temp
    temp = temp2
  }
}
