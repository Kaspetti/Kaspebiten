package game

import "github.com/hajimehoshi/ebiten"


// Window properties
const (
    WINDOW_WIDTH = 1280
    WINDOW_HEIGHT = 720
    VIEWPORT_WIDTH = 128
    VIEWPORT_HEIGHT = 72
    WINDOW_TITLE = "Kaspebiten Game"
)


type Game struct {

}


func (g *Game) Update(screen *ebiten.Image) error {
    return nil
}


func (g *Game) Draw(screen *ebiten.Image) {

}


func (g *Game) Layout(windowWidth, windowHeight int) (viewportWidth, viewportHeight int) {
    return VIEWPORT_WIDTH, VIEWPORT_HEIGHT
}


func StartGame() {
    ebiten.SetWindowSize(WINDOW_WIDTH, WINDOW_HEIGHT)
    ebiten.SetWindowTitle(WINDOW_TITLE)
    ebiten.SetWindowResizable(false)

    game := &Game{}

    if err := ebiten.RunGame(game); err != nil {
        panic(err)
    }
} 
