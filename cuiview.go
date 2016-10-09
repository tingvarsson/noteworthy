package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
	"github.com/tingvarsson/rss"
)

type CuiView struct {
	gui   *gocui.Gui
	feed  *rss.TopElement
	index int
}

func NewCuiView(f *rss.TopElement) *CuiView {
	g := gocui.NewGui()
	if err := g.Init(); err != nil {
		log.Panicln(err)
	}

	initLayout(g)
	initKeybindings(g)

	return &CuiView{gui: g, feed: f, index: 0}
}

func (cv *CuiView) MainLoop() {
	if err := cv.gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func (cv *CuiView) Close() {
	cv.gui.Close()
}

func initLayout(g *gocui.Gui) {
	g.SetLayout(layoutHandler)
}

func initKeybindings(g *gocui.Gui) {
	for _, kb := range keybindingList {
		if err := g.SetKeybinding(kb.viewName, kb.key, kb.mod, kb.handler); err != nil {
			log.Panicln(err)
		}
	}
}

var keybindingList = []struct {
	viewName string
	key      interface{}
	mod      gocui.Modifier
	handler  gocui.KeybindingHandler
}{
	{"", gocui.KeyCtrlC, gocui.ModNone, quit},
	{"", gocui.KeyArrowUp, gocui.ModNone, prevItem},
	{"", gocui.KeyArrowDown, gocui.ModNone, nextItem},
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func prevItem(g *gocui.Gui, v *gocui.View) error {
	fmt.Fprintf(v, "%d", SelectedItem)
	SelectedItem--
	return nil
}

func nextItem(g *gocui.Gui, v *gocui.View) error {
	fmt.Fprintf(v, "%d", SelectedItem)
	SelectedItem++
	return nil
}
