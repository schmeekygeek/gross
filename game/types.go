package game

type Point struct {
  Y           int
  X           int
  PointType   string
}

type Game struct {

  // game canvas
  canvas    [10][10]Point

  // snake body
  body      []Point

  // if the snake can go up (snake can't go left or right when it's already going either left or right)
  // similarly, the snake can't go up and down if it's either going up or down already
  canGoUp   bool

  // food
  food      Point
  score     int

  // current direction where the snake is headed
  currDir   string

  // has lost
  hasLost   bool
}

const (
  SnakeBody = "#"
  Food      = "&"
  Ground    = "_"
  
  // Directions
  Up        = "w"
  Down      = "s"
  Left      = "a"
  Right     = "d"
)

// scenarios:
// 1 The point is an empty space
// 2 The point is where the body of the snake is at
// 3 The point where the food is
