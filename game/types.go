package game

type Point struct {
  X           int
  Y           int
  PointType   string
}

type Game struct {
  Points    [10][10]Point
  Body      []Point
  canGoUp   bool
  food      Point
  score     int
}


const (
  SnakeBody         = "#"
  SnakeHead         = "$"
  Food              = "&"
  Ground            = "_"
)

// scenarios:
// 1 The point is an empty space
// 2 The point is where the body of the snake is at
// 3 The point where the food is
