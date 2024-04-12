package main

import (
	"game/pkg/models"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	playerSprite, daggerSprite, swordSprite, axeSprite, bowSprite, daggerUp, swordUp, axeUp, bowUp := models.Create()
	err := ebiten.RunGame(&models.Game{
		PS:  playerSprite,
		DS:  daggerSprite,
		SS:  swordSprite,
		AS:  axeSprite,
		BS:  bowSprite,
		DUS: daggerUp,
		SUS: swordUp,
		AUS: axeUp,
		BUS: bowUp,
	})
	if err != nil {
		panic(err)
	}
}
