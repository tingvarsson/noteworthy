package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jroimartin/gocui"

	"github.com/tingvarsson/rss"
)

type RssFeed rss.TopElement

func (r RssFeed) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Title:%s (%s)\n", r.Channel.Title, r.Channel.Link))
	for _, i := range r.Channel.Items {
		buffer.WriteString(fmt.Sprintf("-- [%s] %s (%s)\n%s\n\n", i.PubDate, i.Title, i.Link, i.Description))
	}
	return buffer.String()
}

var rssData rss.TopElement
var SelectedItem = 0
var ListViews []string
var ContentView *gocui.View

func main() {
	url := flag.String("url", "rss.php", "the url to the target rss feed")
	flag.Parse()

	rssFile, err := ioutil.ReadFile(*url)
	if err != nil {
		log.Panicln(err)
		return
	}

	rssData, err = rss.Decode(rssFile)
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println(RssFeed(rssData))

	cv := NewCuiView(&rssData)
	defer cv.Close()
	cv.MainLoop()
}

func feedLayout(g *gocui.Gui) error {
	if err := createListViews(g); err != nil {
		return err
	}

	if err := createContentView(g); err != nil {
		return err
	}

	populateList(g)
	populateContent(g)

	return nil
}

func createListViews(g *gocui.Gui) error {
	ListViews = []string{}
	maxX, maxY := g.Size()
	for i := 0; i <= getMaxNoListItems(maxY); i++ {
		viewName := fmt.Sprintf("ListItem%d", i)
		x0, x1, y0, y1 := listPosition(i, maxX, maxY)
		if _, err := g.SetView(viewName, x0, y0, x1, y1); err != nil {
			if err != gocui.ErrUnknownView {
				return err
			}
		}
		ListViews = append(ListViews, viewName)
	}
	return nil
}

func createContentView(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	x0, x1, y0, y1 := contentPosition(maxX, maxY)
	v, err := g.SetView("Content", x0, y0, x1, y1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}
	ContentView = v
	return nil
}

func populateList(g *gocui.Gui) {
	adjustedItemIndex := getAdjustedItemIndex(len(ListViews))
	for i, n := range ListViews {
		v, err := g.View(n)
		if err != nil {
			log.Panicln(err)
		}
		v.Clear()
		fmt.Fprintf(v, "%s", rssData.Channel.Items[i+adjustedItemIndex].Title)
		if i == 0 {
			g.SetCurrentView(v.Name())
		}
	}
}

func getAdjustedItemIndex(l int) int {
	return SelectedItem
}

func populateContent(g *gocui.Gui) {
	ContentView.Clear()
	fmt.Fprintln(ContentView, rssData.Channel.Items[SelectedItem].Description)
}

func listPosition(i int, maxX int, maxY int) (int, int, int, int) {
	return 0, getListWidth(maxX), i * ListHeight, i*ListHeight + ListHeight
}

func contentPosition(maxX, maxY int) (int, int, int, int) {
	return getListWidth(maxX), maxX - 1, 0, maxY - 1
}

const ListWidthRatio = 0.4
const ListHeight = 10

func getMaxNoListItems(y int) int {
	return y / ListHeight
}

func getListWidth(x int) int {
	return int(float32(x) * ListWidthRatio)
}

// TODO
// - convert to cui format (move around, launch browser, some pretty styling)
// - fetch feed from ze interwebz (instead of good old file)
// - file with a list of feeds (urls basically that are then fetched, decoded, and listed)
// - Add a heirarchy of Top(options, manage(add/remove feed), and feeds)/Feeds(based on input)/Items(fetched per feed)
// - possibly convert to client/server style with a DB in the bg that keeps track of read/unread, and managed feeds (futuristic)
