package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	var chunkX float32 = 10
	var chunkY float32 = 10
	var numberOfCunksInARow float32 = 4
	var numberOfCunksInACol float32 = 3
	var numberOfCunks float32 = numberOfCunksInARow * numberOfCunksInACol
	var sizeOfCunk float32 = chunkX * chunkY
	var worldX float32 = chunkX * numberOfCunksInARow
	var worldY float32 = chunkY * numberOfCunksInACol
	var worldCellCount float32 = worldX * worldY
	var worldXY rl.Vector2 = rl.Vector2{X: worldX, Y: worldY}
	fmt.Println("the world map is made of", numberOfCunks, "Chunks", "with a size of ", sizeOfCunk, "for a total size of", worldCellCount, "it has a X of ", worldXY.X, "and a Y of", worldXY.Y)
}
