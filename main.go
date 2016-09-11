package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/tingvarsson/rss"
)

func main() {
	url := flag.String("url", "rss.php", "the url to the target rss feed")
	flag.Parse()
	fmt.Println(*url)

	rssFile, err := ioutil.ReadFile(*url)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	rssData, err := rss.Decode(rssFile)
	if err != nil {
		fmt.Println("Error decoding rss feed:", err)
	}

	fmt.Println(rssData)

	newRssFile, err := rss.Encode(rssData)
	if err != nil {
		fmt.Println("Error encoding rss feed:", err)
	}

	err = ioutil.WriteFile("test.php", newRssFile, 0644)
}
