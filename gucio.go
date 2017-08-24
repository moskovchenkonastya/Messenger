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



type tParamsLoginStruct struct {
	username     string
	password string
}

var data tParamsLoginStruct

func layoutLogin(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	// fmt.Print("inir")

	if v, err := g.SetView("Authrorization", maxX/2 - 20, maxY/2 - 5, maxX/2 + 20, maxY/2 + 4); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.Editable = false

		v.Title = "Authrorization"

	}

	if v, err := g.SetView("username", maxX/2-18, maxY/2 - 4, maxX/2 + 18, maxY/2 - 2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.Frame = true
		v.Editable = true
		v.Wrap = true
		v.Title = "Username"
		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorBlack
	}

	if v, err := g.SetView("password", maxX/2-18,  maxY/2 , maxX/2 + 18, maxY/2 + 2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.Frame = true
		v.Editable = true
		v.Wrap = true
		v.Title = "Password"

		v.SelBgColor = gocui.ColorGreen
		v.SelFgColor = gocui.ColorGreen
	}


	if v, err := g.SetView("but1", maxX/2 - 20,  maxY/2 + 6, maxX/2 + 20, maxY/2 + 11); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Highlight = true
		v.SelBgColor = gocui.ColorBlue
		v.SelFgColor = gocui.ColorBlack
		fmt.Fprintln(v, "Login")
		fmt.Fprintln(v, "Register")
		fmt.Fprintln(v, "Forgot password")
		fmt.Fprintln(v, "Exit")
	}

	return nil
}

func layoutLoginNextView(g *gocui.Gui, v *gocui.View) error {

	_, err := g.SetCurrentView("password")

	return err
}

func quit(g *gocui.Gui, v *gocui.View) error {

	if v.Name() == "password"{
		data.password = v.Buffer()
	}

	return gocui.ErrQuit
}

func readUsername(g *gocui.Gui, v *gocui.View) error {

	if v.Name() == "username" {
		data.username = v.Buffer()
	}

	return nil
}

func delViews (g *gocui.Gui, v *gocui.View, s string) error {

	if err := g.DeleteView(s); err != nil {
		return err
	}
	fmt.Print("\n del - ", s)
	// g.Close()

	return nil

}



func login(g *gocui.Gui, v *gocui.View){

	go delViews(g,v,"but1")
	go delViews(g,v,"Authrorization")
	go delViews(g,v,"username")
	go delViews(g,v,"password")
	fmt.Print("login")
}


func getMethod (g *gocui.Gui, v *gocui.View) error {

	var l string
	var err error

	if _, err := g.SetCurrentView(v.Name()); err != nil {
		return err
	}

	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		l = ""
	}

	if l == "Login" {
		login(g, v)
	}
	/*
	maxX, maxY := g.Size()
	if v, err := g.SetView("msg", maxX/2-10, maxY/2, maxX/2+10, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, l)
	}*/
	return nil
}

func keybindings(g *gocui.Gui) error {

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlD, gocui.ModNone, readUsername); err != nil {
		return err
	}

	if err := g.SetKeybinding("but1", gocui.MouseLeft, gocui.ModNone, getMethod); err != nil {
		return err
	}

	for _, n := range []string{"username", "password"} {
		if err := g.SetKeybinding(n, gocui.MouseLeft, gocui.ModNone, showMsg); err != nil {
			return err
		}
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
		log.Printf("error cursor %s",l)
	}

	/*if v.Name() == "username"{
		_, cy := v.Cursor()
		username, _ = v.Line(cy)
		v.Clear()
		fmt.Print("Username:", username)

	}

	if v.Name() == "password"{
		//_, cy := v.Cursor()
		password = v.Buffer()
		fmt.Print("pass:", password)
	}*/


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

	//g.Close()

	/*
	fmt.Print("Username:", data.username)
	fmt.Printf("\n")
	fmt.Print("Password:", data.password)
	*/


}
