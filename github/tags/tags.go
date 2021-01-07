package tags

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Commit struct {
	Sha string `json:"sha"`
	URL string `json:"url"`
}
type Tag struct {
	Name       string `json:"name"`
	ZipballURL string `json:"zipball_url"`
	TarballURL string `json:"tarball_url"`
	Commit     Commit `json:"commit"`
	NodeID     string `json:"node_id"`
}

func TheLatestVersion(url string) (latestVersion, name string) {

	var client = &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		log.Panicf("%v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicf("%v", err)
	}

	tags := []Tag{}
	json.Unmarshal(body, &tags)

	latestVersion = tags[0].TarballURL
	name = fmt.Sprintf("%s.tar.gz", tags[0].Name)
	return latestVersion, name
}
