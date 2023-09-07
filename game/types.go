package game

type Point struct {
  Y           int
  X           int
  PointType   string
}

type Game struct {
  Points    [10][10]Point
  Body      []Point
  food      Point
  score     int
  currDir   string
}

const (
  SnakeBody = "#"
  Food      = "&"
  Ground    = "_"
  
  // Directions
  Up        = "k"
  Down      = "j"
  Left      = "h"
  Right     = "l"
)

// scenarios:
// 1 The point is an empty space
// 2 The point is where the body of the snake is at
// 3 The point where the food is
