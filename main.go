//+build !mobilebind

package main

import (
	"image/color"
	"log"
	"strconv"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

type DefaultScene struct{}

type MyLabel struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

type Guy struct {
	ecs.BasicEntity

	common.RenderComponent
	common.SpaceComponent
	common.MouseComponent
	common.CollisionComponent
}

var (
	fnt   *common.Font
	red   *common.Texture
	green *common.Texture
	blue  *common.Texture
)

func (*DefaultScene) Preload() {
	err := engo.Files.Load("guy.png")
	if err != nil {
		log.Println(err)
	}
	err = engo.Files.Load("blue.png")
	if err != nil {
		log.Println(err)
	}
	err = engo.Files.Load("green.png")
	if err != nil {
		log.Println(err)
	}
	err = engo.Files.Load("Roboto-Regular.ttf")
	if err != nil {
		panic(err)
	}
}

func (*DefaultScene) Setup(u engo.Updater) {
	w, _ := u.(*ecs.World)
	common.SetBackground(color.White)

	w.AddSystem(&common.RenderSystem{})
	w.AddSystem(&common.MouseSystem{})
	w.AddSystem(&ControlSystem{})
	w.AddSystem(&ClickSystem{})
	w.AddSystem(&common.CollisionSystem{})

	// These are not required, but allow you to move / rotate and still see that it works
	w.AddSystem(&common.MouseZoomer{-0.125})
	w.AddSystem(common.NewKeyboardScroller(500, engo.DefaultHorizontalAxis, engo.DefaultVerticalAxis))
	w.AddSystem(&common.MouseRotator{RotationSpeed: 0.125})

	//text
	fnt = &common.Font{
		URL:  "Roboto-Regular.ttf",
		FG:   color.Black,
		Size: 64,
	}
	err := fnt.CreatePreloaded()
	if err != nil {
		panic(err)
	}

	label1 := MyLabel{BasicEntity: ecs.NewBasic()}
	label1.RenderComponent.Drawable = common.Text{
		Font: fnt,
		Text: "Hello world !",
	}
	label1.SetShader(common.HUDShader)

	label2 := MyLabel{BasicEntity: ecs.NewBasic()}
	label2.RenderComponent.Drawable = common.Text{
		Font:          fnt,
		Text:          "blo",
		LineSpacing:   0.5,
		LetterSpacing: 0.15,
	}
	label2.SpaceComponent.Position = engo.Point{500, 0}
	label2.SetShader(common.HUDShader)

	label3 := MyLabel{BasicEntity: ecs.NewBasic()}
	label3.RenderComponent.Drawable = common.Text{
		Font:          fnt,
		Text:          "Hey",
		LineSpacing:   0.5,
		LetterSpacing: 0.15,
	}
	label3.SpaceComponent.Position = engo.Point{0, 150}
	label3.SetShader(common.HUDShader)

	label4 := MyLabel{BasicEntity: ecs.NewBasic()}
	label4.RenderComponent.Drawable = common.Text{
		Font:          fnt,
		Text:          "howdy",
		LineSpacing:   0.5,
		LetterSpacing: 0.15,
	}
	label4.SpaceComponent.Position = engo.Point{500, 150}
	label4.SetShader(common.HUDShader)

	label5 := MyLabel{BasicEntity: ecs.NewBasic()}
	label5.RenderComponent.Drawable = common.Text{
		Font:          fnt,
		Text:          "MX: 0",
		LineSpacing:   0.5,
		LetterSpacing: 0.15,
	}
	label5.SpaceComponent.Position = engo.Point{0, 300}
	label5.SetShader(common.HUDShader)

	label6 := MyLabel{BasicEntity: ecs.NewBasic()}
	label6.RenderComponent.Drawable = common.Text{
		Font:          fnt,
		Text:          "MY: 0",
		LineSpacing:   0.5,
		LetterSpacing: 0.15,
	}
	label6.SpaceComponent.Position = engo.Point{500, 300}
	label6.SetShader(common.HUDShader)

	label7 := MyLabel{BasicEntity: ecs.NewBasic()}
	label7.RenderComponent.Drawable = common.Text{
		Font:          fnt,
		Text:          "DX: 0",
		LineSpacing:   0.5,
		LetterSpacing: 0.15,
	}
	label7.SpaceComponent.Position = engo.Point{0, 450}
	label7.SetShader(common.HUDShader)

	label8 := MyLabel{BasicEntity: ecs.NewBasic()}
	label8.RenderComponent.Drawable = common.Text{
		Font:          fnt,
		Text:          "DY: 0",
		LineSpacing:   0.5,
		LetterSpacing: 0.15,
	}
	label8.SpaceComponent.Position = engo.Point{500, 450}
	label8.SetShader(common.HUDShader)

	// Add our text to appropriate systems
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&label1.BasicEntity, &label1.RenderComponent, &label1.SpaceComponent)
			sys.Add(&label2.BasicEntity, &label2.RenderComponent, &label2.SpaceComponent)
			sys.Add(&label3.BasicEntity, &label3.RenderComponent, &label3.SpaceComponent)
			sys.Add(&label4.BasicEntity, &label4.RenderComponent, &label4.SpaceComponent)
			sys.Add(&label5.BasicEntity, &label5.RenderComponent, &label5.SpaceComponent)
			sys.Add(&label6.BasicEntity, &label6.RenderComponent, &label6.SpaceComponent)
			sys.Add(&label7.BasicEntity, &label7.RenderComponent, &label7.SpaceComponent)
			sys.Add(&label8.BasicEntity, &label8.RenderComponent, &label8.SpaceComponent)
		case *ControlSystem:
			sys.Add(&label1.BasicEntity, &label1.RenderComponent, &label1.SpaceComponent, "first")
			sys.Add(&label2.BasicEntity, &label2.RenderComponent, &label2.SpaceComponent, "second")
			sys.Add(&label3.BasicEntity, &label3.RenderComponent, &label3.SpaceComponent, "third")
			sys.Add(&label4.BasicEntity, &label4.RenderComponent, &label4.SpaceComponent, "fourth")
			sys.Add(&label5.BasicEntity, &label5.RenderComponent, &label5.SpaceComponent, "fifth")
			sys.Add(&label6.BasicEntity, &label6.RenderComponent, &label6.SpaceComponent, "sixth")
			sys.Add(&label7.BasicEntity, &label7.RenderComponent, &label7.SpaceComponent, "seventh")
			sys.Add(&label8.BasicEntity, &label8.RenderComponent, &label8.SpaceComponent, "eighth")
		}
	}

	//Guy icon
	// Retrieve a texture
	red, err = common.LoadedSprite("guy.png")
	if err != nil {
		log.Println(err)
	}

	green, err = common.LoadedSprite("green.png")
	if err != nil {
		log.Println(err)
	}

	blue, err = common.LoadedSprite("blue.png")
	if err != nil {
		log.Println(err)
	}

	// Create an entity
	guy := Guy{BasicEntity: ecs.NewBasic()}

	// Initialize the components, set scale to 8x
	guy.RenderComponent = common.RenderComponent{
		Drawable: red,
		Scale:    engo.Point{8, 8},
	}
	guy.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{0, 0},
		Width:    red.Width() * guy.RenderComponent.Scale.X,
		Height:   red.Height() * guy.RenderComponent.Scale.Y,
	}
	guy.CollisionComponent = common.CollisionComponent{
		Main:  1,
		Group: 1,
	}

	// Create an entity
	guy2 := Guy{BasicEntity: ecs.NewBasic()}

	// Initialize the components, set scale to 8x
	guy2.RenderComponent = common.RenderComponent{
		Drawable: red,
		Scale:    engo.Point{8, 8},
	}
	guy2.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{0, 0},
		Width:    red.Width() * guy2.RenderComponent.Scale.X,
		Height:   red.Height() * guy2.RenderComponent.Scale.Y,
	}

	// Create an entity
	guy3 := Guy{BasicEntity: ecs.NewBasic()}

	// Initialize the components, set scale to 8x
	guy3.RenderComponent = common.RenderComponent{
		Drawable: red,
		Scale:    engo.Point{8, 8},
	}
	guy3.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{300, 300},
		Width:    red.Width() * guy3.RenderComponent.Scale.X,
		Height:   red.Height() * guy3.RenderComponent.Scale.Y,
	}
	guy3.CollisionComponent = common.CollisionComponent{
		Main:  1,
		Group: 1,
	}

	// Add it to appropriate systems
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			sys.Add(&guy.BasicEntity, &guy.RenderComponent, &guy.SpaceComponent)
			sys.Add(&guy2.BasicEntity, &guy2.RenderComponent, &guy2.SpaceComponent)
			sys.Add(&guy3.BasicEntity, &guy3.RenderComponent, &guy3.SpaceComponent)
		case *common.MouseSystem:
			sys.Add(&guy.BasicEntity, &guy.MouseComponent, &guy.SpaceComponent, &guy.RenderComponent)
			sys.Add(&guy2.BasicEntity, &guy2.MouseComponent, &guy2.SpaceComponent, &guy2.RenderComponent)
			sys.Add(&guy3.BasicEntity, &guy3.MouseComponent, &guy3.SpaceComponent, &guy3.RenderComponent)
		case *ClickSystem:
			sys.Add(&guy.BasicEntity, &guy.RenderComponent, &guy.SpaceComponent, &guy.MouseComponent, &guy.CollisionComponent, "not", "red")
			sys.Add(&guy2.BasicEntity, &guy2.RenderComponent, &guy2.SpaceComponent, &guy2.MouseComponent, &guy.CollisionComponent, "clicky", "red")
			sys.Add(&guy3.BasicEntity, &guy3.RenderComponent, &guy3.SpaceComponent, &guy3.MouseComponent, &guy.CollisionComponent, "colidey", "red")
		case *common.CollisionSystem:
			sys.Add(&guy.BasicEntity, &guy.CollisionComponent, &guy.SpaceComponent)
			sys.Add(&guy3.BasicEntity, &guy3.CollisionComponent, &guy3.SpaceComponent)
		}
	}
}

func (*DefaultScene) Type() string { return "GameWorld" }

type controlEntity struct {
	*ecs.BasicEntity
	*common.RenderComponent
	*common.SpaceComponent
	whichone string
}

type ControlSystem struct {
	entities []controlEntity

	MouseX, MouseY float32

	WindowHeight, WindowWidth float32

	CanvasHeight, CanvasWidth float32
}

func (c *ControlSystem) Add(basic *ecs.BasicEntity, render *common.RenderComponent, space *common.SpaceComponent, id string) {
	c.entities = append(c.entities, controlEntity{basic, render, space, id})
}

func (c *ControlSystem) Remove(basic ecs.BasicEntity) {
	delete := -1
	for index, e := range c.entities {
		if e.BasicEntity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		c.entities = append(c.entities[:delete], c.entities[delete+1:]...)
	}
}

func (c *ControlSystem) Update(float32) {
	for _, e := range c.entities {
		switch e.whichone {
		case "first":
			if c.WindowWidth != engo.WindowWidth() {
				c.WindowWidth = engo.WindowWidth()
				txt := "WW: " + strconv.FormatFloat(float64(c.WindowWidth), 'f', -1, 32)
				e.RenderComponent.Drawable.Close()
				e.RenderComponent.Drawable = common.Text{
					Font: fnt,
					Text: txt,
				}
			}
		case "second":
			if c.WindowHeight != engo.WindowHeight() {
				c.WindowHeight = engo.WindowHeight()
				txt := "WH: " + strconv.FormatFloat(float64(c.WindowHeight), 'f', -1, 32)
				e.RenderComponent.Drawable.Close()
				e.RenderComponent.Drawable = common.Text{
					Font: fnt,
					Text: txt,
				}
			}
		case "third":
			if c.CanvasWidth != engo.CanvasWidth() {
				c.CanvasWidth = engo.CanvasWidth()
				txt := "CW: " + strconv.FormatFloat(float64(c.CanvasWidth), 'f', -1, 32)
				e.RenderComponent.Drawable.Close()
				e.RenderComponent.Drawable = common.Text{
					Font: fnt,
					Text: txt,
				}
			}
		case "fourth":
			if c.CanvasHeight != engo.CanvasHeight() {
				c.CanvasHeight = engo.CanvasHeight()
				txt := "CH: " + strconv.FormatFloat(float64(c.CanvasHeight), 'f', -1, 32)
				e.RenderComponent.Drawable.Close()
				e.RenderComponent.Drawable = common.Text{
					Font: fnt,
					Text: txt,
				}
			}
		case "fifth":
			if engo.Input.Mouse.Action == engo.Press && c.MouseX != engo.Input.Mouse.X {
				c.MouseX = engo.Input.Mouse.X
				txt := "MX: " + strconv.FormatFloat(float64(c.MouseX), 'f', -1, 32)
				e.RenderComponent.Drawable.Close()
				e.RenderComponent.Drawable = common.Text{
					Font: fnt,
					Text: txt,
				}
			}
		case "sixth":
			if engo.Input.Mouse.Action == engo.Press && c.MouseY != engo.Input.Mouse.Y {
				c.MouseY = engo.Input.Mouse.Y
				txt := "MY: " + strconv.FormatFloat(float64(c.MouseY), 'f', -1, 32)
				e.RenderComponent.Drawable.Close()
				e.RenderComponent.Drawable = common.Text{
					Font: fnt,
					Text: txt,
				}
			}
		case "seventh":
			txt := "DX: " + strconv.FormatFloat(float64(engo.Input.Mouse.X), 'f', -1, 32)
			e.RenderComponent.Drawable.Close()
			e.RenderComponent.Drawable = common.Text{
				Font: fnt,
				Text: txt,
			}
		case "eighth":
			txt := "DY: " + strconv.FormatFloat(float64(engo.Input.Mouse.Y), 'f', -1, 32)
			e.RenderComponent.Drawable.Close()
			e.RenderComponent.Drawable = common.Text{
				Font: fnt,
				Text: txt,
			}
		}
	}
}

type clickEntity struct {
	*ecs.BasicEntity
	*common.RenderComponent
	*common.SpaceComponent
	*common.MouseComponent
	*common.CollisionComponent

	Name  string
	Color string
}

type ClickSystem struct {
	entities []clickEntity
	camera   *common.CameraSystem
}

func (c *ClickSystem) New(w *ecs.World) {
	for _, system := range w.Systems() {
		switch sys := system.(type) {
		case *common.CameraSystem:
			c.camera = sys
		}
	}
}

func (c *ClickSystem) Add(basic *ecs.BasicEntity, render *common.RenderComponent, space *common.SpaceComponent, mouse *common.MouseComponent, coll *common.CollisionComponent, name, color string) {
	c.entities = append(c.entities, clickEntity{basic, render, space, mouse, coll, name, color})
}

func (c *ClickSystem) Remove(basic ecs.BasicEntity) {
	delete := -1
	for index, e := range c.entities {
		if e.BasicEntity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		c.entities = append(c.entities[:delete], c.entities[delete+1:]...)
	}
}

func (c *ClickSystem) Update(float32) {
	for i := 0; i < len(c.entities); i++ {
		if c.entities[i].Name == "clicky" {
			if c.entities[i].Clicked {
				if c.entities[i].Color == "red" {
					c.entities[i].Color = "green"
					c.entities[i].RenderComponent.Drawable = green
				} else if c.entities[i].Color == "green" {
					c.entities[i].Color = "blue"
					c.entities[i].RenderComponent.Drawable = blue
				} else {
					c.entities[i].Color = "red"
					c.entities[i].RenderComponent.Drawable = red
				}
			}
		} else if c.entities[i].Name == "colidey" {
			if c.entities[i].Collides == 1 {
				if c.entities[i].Color == "red" {
					c.entities[i].Color = "green"
					c.entities[i].RenderComponent.Drawable = green
				} else if c.entities[i].Color == "green" {
					c.entities[i].Color = "blue"
					c.entities[i].RenderComponent.Drawable = blue
				} else {
					c.entities[i].Color = "red"
					c.entities[i].RenderComponent.Drawable = red
				}
			}
		} else {
			c.entities[i].Position.X = ((engo.Input.Mouse.X * c.camera.Z() * engo.GameWidth() / engo.WindowWidth()) + (c.camera.X()-(engo.GameWidth()/2)*c.camera.Z())/engo.GetGlobalScale().X)
			c.entities[i].Position.Y = ((engo.Input.Mouse.Y * c.camera.Z() * engo.GameHeight() / engo.WindowHeight()) + (c.camera.Y()-(engo.GameHeight()/2)*c.camera.Z())/engo.GetGlobalScale().Y)
		}
	}
}

func main() {
	opts := engo.RunOptions{
		Title:         "Mouse Demo",
		Width:         1024,
		Height:        640,
		ScaleOnResize: true,
	}
	engo.Run(opts, &DefaultScene{})
}
