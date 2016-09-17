package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"

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

func main() {
	url := flag.String("url", "rss.php", "the url to the target rss feed")
	flag.Parse()

	rssFile, err := ioutil.ReadFile(*url)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	rssData, err := rss.Decode(rssFile)
	if err != nil {
		fmt.Println("Error decoding rss feed:", err)
	}

	fmt.Println(RssFeed(rssData))

	newRssFile, err := rss.Encode(rssData)
	if err != nil {
		fmt.Println("Error encoding rss feed:", err)
	}

	err = ioutil.WriteFile("test.php", newRssFile, 0644)
}

// TODO
// - convert to cui format (move around, launch browser, some pretty styling)
// - fetch feed from ze interwebz (instead of good old file)
// - file with a list of feeds (urls basically that are then fetched, decoded, and listed)
// - Add a heirarchy of Top(options, manage(add/remove feed), and feeds)/Feeds(based on input)/Items(fetched per feed)
// - possibly convert to client/server style with a DB in the bg that keeps track of read/unread, and managed feeds (futuristic)
