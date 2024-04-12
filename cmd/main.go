package main

import (
	"game/pkg/models"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	playerSprite, daggerSprite, swordSprite, axeSprite, bowSprite, daggerUp, swordUp, axeUp, bowUp := models.Create()
	err := ebiten.RunGame(&models.Game{
		PlayerSprite:   &playerSprite,
		DaggerSprite:   &daggerSprite,
		SwordSprite:    &swordSprite,
		AxeSprite:      &axeSprite,
		BowSprite:      &bowSprite,
		DaggerUpSprite: &daggerUp,
		SwordUpSprite:  &swordUp,
		AxeUpSprite:    &axeUp,
		BowUpSprite:    &bowUp,
	})
	if err != nil {
		panic(err)
	}
}
