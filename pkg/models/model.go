package models

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"log"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 480
	BarWidth     = 200
	BarHeight    = 20
	MaxHP        = 1000000
)

var (
	TtfFont       font.Face
	Gold          int
	CurrentHp     = MaxHP
	PassiveDamage int
	GoldTime      = time.Now()
)

func init() {
	fontData, err := ioutil.ReadFile("static/assets/PfhighwaysansproBlackitalic.ttf")
	if err != nil {
		log.Fatal(err)
	}

	tt, err := opentype.Parse(fontData)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	TtfFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    12,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

type Sprite interface {
	Draw(screen *ebiten.Image)
}

type PSprite struct {
	Image                  *ebiten.Image
	X, Y                   float64
	Width, Height          float64
	IsClicked              bool
	SpriteClickedStartTime time.Time
}

type Weapon struct {
	Name                   string
	Cost                   float64
	Image                  *ebiten.Image
	X, Y                   float64
	Width, Height          float64
	IsClicked              bool
	SpriteClickedStartTime time.Time
	Damage                 int
	FullDamage             int
	Quantity               int
}

type LevelUp struct {
	Name                   string
	Cost                   float64
	LinkedWeapon           *Weapon
	Image                  *ebiten.Image
	X, Y                   float64
	Width, Height          float64
	IsClicked              bool
	SpriteClickedStartTime time.Time
	Quantity               int
}

var PlayerSprite, _, _ = ebitenutil.NewImageFromFile("static/assets/ufoRed.png")
var DaggerSprite, _, _ = ebitenutil.NewImageFromFile("static/assets/dagger.png")
var SwordSprite, _, _ = ebitenutil.NewImageFromFile("static/assets/sword.png")
var AxeSprite, _, _ = ebitenutil.NewImageFromFile("static/assets/axe.png")
var BowSprite, _, _ = ebitenutil.NewImageFromFile("static/assets/bow.png")
var LevelUpSprite, _, _ = ebitenutil.NewImageFromFile("static/assets/arrow.png")

type Game struct {
	PlayerSprite   *PSprite
	DaggerSprite   *Weapon
	SwordSprite    *Weapon
	AxeSprite      *Weapon
	BowSprite      *Weapon
	DaggerUpSprite *LevelUp
	SwordUpSprite  *LevelUp
	AxeUpSprite    *LevelUp
	BowUpSprite    *LevelUp
}

func (lu *LevelUp) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		// s.Width, s.Height = s.Width*0.05, s.Height*0.05
		mX, mY := ebiten.CursorPosition()
		mouseX, mouseY := float64(mX), float64(mY)

		// Вычисляем координаты центра спрайта
		centerX := lu.X + lu.Width/2
		centerY := lu.Y + lu.Height/2

		// Вычисляем расстояние от центра спрайта до позиции курсора
		distance := math.Sqrt(math.Pow(mouseX-centerX, 2) + math.Pow(mouseY-centerY, 2))

		// Если расстояние меньше или равно радиусу спрайта, считаем клик в пределах спрайта
		if distance <= lu.Width/2 {
			if !lu.IsClicked && Gold >= int(lu.Cost) && lu.Quantity < 3 {

				if lu.Quantity == 0 && lu.LinkedWeapon.Quantity >= 10 {
					Gold -= int(lu.Cost)
					lu.Cost *= 10
					lu.LinkedWeapon.FullDamage *= 2
					lu.LinkedWeapon.Damage *= 2
					lu.Quantity++
					lu.IsClicked = true
				} else if lu.Quantity == 1 && lu.LinkedWeapon.Quantity >= 50 {
					Gold -= int(lu.Cost)
					lu.Cost *= 10
					lu.LinkedWeapon.FullDamage *= 2
					lu.LinkedWeapon.Damage *= 2
					lu.Quantity++
					lu.IsClicked = true
				} else if lu.Quantity == 2 && lu.LinkedWeapon.Quantity >= 100 {
					Gold -= int(lu.Cost)
					lu.Cost *= 10
					lu.LinkedWeapon.FullDamage *= 2
					lu.LinkedWeapon.Damage *= 2
					lu.Quantity++
					lu.IsClicked = true
				}
			}
		} else {
			lu.IsClicked = false
		}
	} else {
		lu.IsClicked = false
	}
	return nil
}

func (s *Weapon) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		// s.Width, s.Height = s.Width*0.05, s.Height*0.05
		mX, mY := ebiten.CursorPosition()
		mouseX, mouseY := float64(mX), float64(mY)

		// Вычисляем координаты центра спрайта
		centerX := s.X + s.Width/2
		centerY := s.Y + s.Height/2

		// Вычисляем расстояние от центра спрайта до позиции курсора
		distance := math.Sqrt(math.Pow(mouseX-centerX, 2) + math.Pow(mouseY-centerY, 2))

		// Если расстояние меньше или равно радиусу спрайта, считаем клик в пределах спрайта
		if distance <= s.Width/2 {
			if !s.IsClicked && Gold >= int(s.Cost) {
				Gold -= int(s.Cost)
				s.Cost = s.Cost * 1.2
				s.FullDamage += s.Damage
				s.Quantity++
				s.SpriteClickedStartTime = time.Now()
				s.IsClicked = true
			}
		} else {
			s.IsClicked = false
		}
	} else {
		s.IsClicked = false
	}
	return nil
}

func (s *PSprite) Update() error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mX, mY := ebiten.CursorPosition()
		mouseX, mouseY := float64(mX), float64(mY)

		// Вычисляем координаты центра спрайта
		centerX := s.X + s.Width/2
		centerY := s.Y + s.Height/2

		// Вычисляем расстояние от центра спрайта до позиции курсора
		distance := math.Sqrt(math.Pow(mouseX-centerX, 2) + math.Pow(mouseY-centerY, 2))

		// Если расстояние меньше или равно радиусу спрайта, считаем клик в пределах спрайта
		if distance <= s.Width/2 {
			if !s.IsClicked {
				Gold++
				CurrentHp--
				s.SpriteClickedStartTime = time.Now()
				s.IsClicked = true
			}
		} else {
			s.IsClicked = false
		}
	} else {
		s.IsClicked = false
	}
	return nil
}

func (g *Game) Update() error {
	err := g.PlayerSprite.Update()
	if err != nil {
		return err
	}
	err = g.DaggerSprite.Update()
	if err != nil {
		return err
	}
	err = g.SwordSprite.Update()
	if err != nil {
		return err
	}
	err = g.AxeSprite.Update()
	if err != nil {
		return err
	}
	err = g.BowSprite.Update()
	if err != nil {
		return err
	}
	err = g.DaggerUpSprite.Update()
	if err != nil {
		return err
	}
	err = g.SwordUpSprite.Update()
	if err != nil {
		return err
	}
	err = g.AxeUpSprite.Update()
	if err != nil {
		return err
	}
	err = g.BowUpSprite.Update()
	if err != nil {
		return err
	}
	if time.Since(GoldTime) >= 1*time.Second {
		Gold += PassiveDamage
		CurrentHp -= PassiveDamage
		// fmt.Println("passive damage: ", passiveDamage)
		// fmt.Printf("cost of the weapon %s : %f \n", g.DaggerSprite.Name, g.DaggerSprite.Cost)
		// fmt.Printf("quantity of the weapon %s : %d\n", g.SwordSprite.Name, g.SwordSprite.Quantity)
		// fmt.Printf("quantity of the weapon %s : %d\n", g.AxeSprite.Name, g.AxeSprite.Quantity)
		// fmt.Printf("quantity of the weapon %s : %d\n", g.BowSprite.Name, g.BowSprite.Quantity)
		// fmt.Println("once per second ", g.DaggerUpSprite.Quantity, g.DaggerUpSprite.Cost)
		GoldTime = time.Now()
	}
	PassiveDamage = g.DaggerSprite.FullDamage + g.SwordSprite.FullDamage + g.AxeSprite.FullDamage + g.BowSprite.FullDamage
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.PlayerSprite.Draw(screen)
	g.DaggerSprite.Draw(screen)
	g.SwordSprite.Draw(screen)
	g.AxeSprite.Draw(screen)
	g.BowSprite.Draw(screen)
	g.DaggerUpSprite.Draw(screen)
	g.SwordUpSprite.Draw(screen)
	g.AxeUpSprite.Draw(screen)
	g.BowUpSprite.Draw(screen)
}
func (lu *LevelUp) Draw(screen *ebiten.Image) {
	msg := fmt.Sprintf("%s's upgrade\nlevel: %d.\nCost : %.2f", lu.LinkedWeapon.Name, lu.Quantity, lu.Cost)
	if lu.Quantity >= 0 && lu.Quantity < 3 {
		opts := &ebiten.DrawImageOptions{}
		// txtopts := &text.DrawOptions{}
		// txtopts.GeoM.Translate(lu.X, lu.Y+45)
		// txtopts.ColorScale.Scale(0, 0, 0, 0)

		text.Draw(screen, msg, TtfFont, int(lu.X), int(lu.Y)+45, color.White)
		opts.GeoM.Translate(lu.X, lu.Y)
		if lu.IsClicked && time.Since(lu.SpriteClickedStartTime) >= 10*time.Millisecond {
			op := &colorm.DrawImageOptions{}
			op.GeoM.Translate(lu.X, lu.Y)
			cm := colorm.ColorM{}
			cm.Scale(1.0, 1.0, 1.0, 0.5)
			colorm.DrawImage(screen, lu.Image, cm, op)
		} else {
			screen.DrawImage(lu.Image, opts)
		}
	}

}

func (s *Weapon) Draw(screen *ebiten.Image) {
	msg1 := fmt.Sprintf("%s's : %d. Cost : %.2f", s.Name, s.Quantity, s.Cost)
	msg2 := fmt.Sprintf("%s's total damage : %d", s.Name, s.Damage*s.Quantity)
	opts := &ebiten.DrawImageOptions{}
	if s.Name == "Dagger" {
		opts.GeoM.Scale(0.04, 0.04)
		text.Draw(screen, msg1, TtfFont, int(s.X)+10, int(s.Y)-10, color.White)
		text.Draw(screen, msg2, TtfFont, int(s.X), int(s.Y), color.White)
	} else if s.Name == "Bow" {
		opts.GeoM.Scale(0.07, 0.07)
		text.Draw(screen, msg1, TtfFont, int(s.X)-10, int(s.Y)-10, color.White)
		text.Draw(screen, msg2, TtfFont, int(s.X)-20, int(s.Y), color.White)
	} else {
		opts.GeoM.Scale(0.05, 0.05)
		text.Draw(screen, msg1, TtfFont, int(s.X)+10, int(s.Y)-10, color.White)
		text.Draw(screen, msg2, TtfFont, int(s.X), int(s.Y), color.White)
	}

	opts.GeoM.Translate(s.X, s.Y)
	if s.IsClicked && time.Since(s.SpriteClickedStartTime) >= 10*time.Millisecond {
		op := &colorm.DrawImageOptions{}
		if s.Name == "Dagger" {
			op.GeoM.Scale(0.04, 0.04)
		} else {
			op.GeoM.Scale(0.05, 0.05)
		}
		op.GeoM.Translate(s.X, s.Y)
		cm := colorm.ColorM{}
		cm.Scale(1.0, 1.0, 1.0, 0.5)
		colorm.DrawImage(screen, s.Image, cm, op)
	} else {
		screen.DrawImage(s.Image, opts)
	}
}

func (s *PSprite) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(s.X, s.Y)
	if s.IsClicked && time.Since(s.SpriteClickedStartTime) >= 10*time.Millisecond {
		op := &colorm.DrawImageOptions{}
		op.GeoM.Translate(s.X, s.Y)
		cm := colorm.ColorM{}
		cm.Scale(1.0, 1.0, 1.0, 0.5)
		colorm.DrawImage(screen, s.Image, cm, op)
	} else {
		screen.DrawImage(s.Image, opts)
	}

	if CurrentHp > 0 {
		vector.DrawFilledRect(screen, float32(s.X-30), float32(s.Y+100), float32(BarWidth), float32(BarHeight), color.RGBA{255, 0, 0, 255}, true)
		vector.DrawFilledRect(screen, float32(s.X-30), float32(s.Y+100), float32(CurrentHp*BarWidth/MaxHP), float32(BarHeight), color.RGBA{0, 255, 0, 255}, true)
		hpmsg := fmt.Sprintf("HP: %d / %d", CurrentHp, MaxHP)
		text.Draw(screen, hpmsg, TtfFont, int(s.X), int(s.Y+115), color.White)
	} else {
		text.Draw(screen, "you won!!!", TtfFont, ScreenWidth/2, ScreenHeight/2, color.RGBA{0, 255, 0, 255})
	}
	// Draw the HP bar with the current health

	// vector.DrawFilledCircle(screen, float32(s.X+s.Width/2), float32(s.Y+s.Height/2), float32(5), color.RGBA{111, 2, 2, 2}, true)

	msg := fmt.Sprintf("gold: %d", Gold)
	text.Draw(screen, msg, TtfFont, 12, 36, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
