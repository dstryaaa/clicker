package models

func Create() (*PSprite, *Weapon, *Weapon, *Weapon, *Weapon, *LevelUp, *LevelUp, *LevelUp, *LevelUp) {
	playerSprite := &PSprite{
		Image:     PlayerSprite,
		X:         50,
		Y:         50,
		Width:     float64(PlayerSprite.Bounds().Dx()),
		Height:    float64(PlayerSprite.Bounds().Dy()),
		IsClicked: false,
	}
	daggerSprite := &Weapon{
		Name:       "Dagger",
		Cost:       10,
		Image:      DaggerSprite,
		X:          450,
		Y:          20,
		Width:      float64(DaggerSprite.Bounds().Dx()) * 0.04,
		Height:     float64(DaggerSprite.Bounds().Dy()) * 0.04,
		IsClicked:  false,
		Damage:     1,
		FullDamage: 0,
		Quantity:   0,
	}
	swordSprite := &Weapon{
		Name:       "Sword",
		Cost:       100,
		Image:      SwordSprite,
		X:          450,
		Y:          140,
		Width:      float64(SwordSprite.Bounds().Dx()) * 0.05,
		Height:     float64(SwordSprite.Bounds().Dy()) * 0.05,
		IsClicked:  false,
		Damage:     30,
		FullDamage: 0,
		Quantity:   0,
	}
	axeSprite := &Weapon{
		Name:       "Axe",
		Cost:       1000,
		Image:      AxeSprite,
		X:          450,
		Y:          260,
		Width:      float64(AxeSprite.Bounds().Dx()) * 0.05,
		Height:     float64(AxeSprite.Bounds().Dy()) * 0.05,
		IsClicked:  false,
		Damage:     150,
		FullDamage: 0,
		Quantity:   0,
	}
	bowSprite := &Weapon{
		Name:       "Bow",
		Cost:       10000,
		Image:      BowSprite,
		X:          470,
		Y:          380,
		Width:      float64(BowSprite.Bounds().Dx()) * 0.07,
		Height:     float64(BowSprite.Bounds().Dy()) * 0.07,
		IsClicked:  false,
		Damage:     400,
		FullDamage: 0,
		Quantity:   0,
	}
	daggerUp := &LevelUp{
		Name:         "DaggerUp",
		Cost:         150,
		Image:        LevelUpSprite,
		LinkedWeapon: daggerSprite,
		X:            530,
		Y:            35,
		Width:        float64(LevelUpSprite.Bounds().Dx()),
		Height:       float64(LevelUpSprite.Bounds().Dy()),
		IsClicked:    false,
		Quantity:     0,
	}
	swordUp := &LevelUp{
		Name:         "SwordUp",
		Cost:         1500,
		Image:        LevelUpSprite,
		LinkedWeapon: swordSprite,
		X:            530,
		Y:            145,
		Width:        float64(LevelUpSprite.Bounds().Dx()),
		Height:       float64(LevelUpSprite.Bounds().Dy()),
		IsClicked:    false,
		Quantity:     0,
	}
	axeUp := &LevelUp{
		Name:         "AxeUp",
		Cost:         15000,
		Image:        LevelUpSprite,
		LinkedWeapon: axeSprite,
		X:            530,
		Y:            265,
		Width:        float64(LevelUpSprite.Bounds().Dx()),
		Height:       float64(LevelUpSprite.Bounds().Dy()),
		IsClicked:    false,
		Quantity:     0,
	}
	bowUp := &LevelUp{
		Name:         "BowUp",
		Cost:         150000,
		Image:        LevelUpSprite,
		LinkedWeapon: bowSprite,
		X:            530,
		Y:            385,
		Width:        float64(LevelUpSprite.Bounds().Dx()),
		Height:       float64(LevelUpSprite.Bounds().Dy()),
		IsClicked:    false,
		Quantity:     0,
	}
	return playerSprite, daggerSprite, swordSprite, axeSprite, bowSprite, daggerUp, swordUp, axeUp, bowUp
}
