package main

import (
	//"fmt"
	//"io"
	//"io/ioutil"
	"log"
	//"strings"

	"github.com/jroimartin/gocui"
	"fmt"
)

var username, password string

type tParamsLoginStruct struct {
	username     string
	password string
}

func layoutLogin(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	// fmt.Print("inir")

	if v, err := g.SetView("login", maxX/2 - 20, maxY/2 - 5, maxX/2 + 20, maxY/2 + 5); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.Editable = false

		v.Title = "Authrorization"

	}

	if v, err := g.SetView("username", maxX/2-18, maxY/2 - 4, maxX/2 + 18, maxY/2-2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.Frame = true
		v.Editable = true
		v.Wrap = true
		v.Title = "Login"
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
	}

	if v, err := g.SetView("password", maxX/2-18, maxY/2 , maxX/2 + 18, maxY/2 + 2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.Frame = true
		v.Editable = true
		v.Wrap = true
		v.Title = "Password"

		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
	}

	return nil
}

func layoutLoginNextView(g *gocui.Gui, v *gocui.View) error {

	_, err := g.SetCurrentView("password")

	return err
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
func funcUsername(g *gocui.Gui, v *gocui.View) error {
	if v.Name() == "username"{
		username = v.ViewBuffer()
		v.Clear()
		//fmt.Printf("vgbhnj")

	}

	if v.Name() == "password"{
		password = v.Word()
		v.Clear()
	}
	return gocui.ErrQuit
}



func keybindings(g *gocui.Gui) error {

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}

	for _, n := range []string{"username", "password"} {
		if err := g.SetKeybinding(n, gocui.MouseLeft, gocui.ModNone, showMsg); err != nil {
			return err
		}
	}

	if err := g.SetKeybinding("", gocui.KeyEsc, gocui.ModNone, funcUsername); err != nil {
		return err
	}

	return nil
}


func showMsg(g *gocui.Gui, v *gocui.View) error {
	//fmt.Println(v.Name())
	var l string
	var err error


	if _, err := g.SetCurrentView(v.Name()); err != nil {
		return err
	}

	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		l = ""
	}
	if err == nil {
		log.Printf("%s",l)
	}



	return nil
}


func main() {

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Cursor = true
	g.Mouse = true

	g.SetManagerFunc(layoutLogin)

	if err := keybindings(g); err != nil {
		log.Panicln(err)
	}


	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}

	g.Close()

	fmt.Printf("Username:", username)
	fmt.Printf("\n")
	fmt.Printf("Password:", password)


}
