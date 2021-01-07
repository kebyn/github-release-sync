package main

import (
	"log"
	"sync"

	"github.com/kebyn/github-release-sync/github/tags"
)

func main() {
	urls := []string{
		"https://api.github.com/repos/golang/go/tags",
	}

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)

		go func(url string, wg *sync.WaitGroup) {

			defer wg.Done()
			latestVersion, name := tags.TheLatestVersion(url)
			log.Printf("%v %v", latestVersion, name)

			err := DownloadFile(name, latestVersion)
			if err != nil {
				log.Panicf("%v", err)
			}

			log.Printf("download %v done.", name)

		}(url, &wg)

	}
	wg.Wait()

}
