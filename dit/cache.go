package dit

import (
	"encoding/json"
	. "github.com/zddhub/dit/utils"
	"sort"
)

type entries []*object
type extensions interface{}

type cache struct {
	Signature  string
	Version    int
	Entries    entries
	Extensions extensions
}

type cachedTreeEntry struct {
	Dirname       string
	Sha1          string
	ObjectsCount  int
	SubTreesCount int
}

type cachedTree struct {
	Signature string
	Entries   []*cachedTreeEntry
}

func (c *cache) loadCache(filename string) {
	buffer, err := ReadFile(filename)
	if err != nil {
		LogD.Println(err)
		c.Signature = "DIRC"
		c.Version = 0
		return
	}
	err = json.Unmarshal(buffer, &c)
	if err != nil {
		LogD.Println(err)
		c.Signature = "DIRC"
		c.Version = 0
		return
	}
}

func (c *cache) storeCache(filename string) error {
	c.sort()
	buffer, err := json.MarshalIndent(*c, "", "  ")

	if err != nil {
		LogD.Println(err)
		return err
	}

	LogI.Printf("%s\n", buffer)
	return WriteFile(filename, buffer, 0644)
}

func (c *cache) sort() {
	if len(c.Entries) < 2 {
		return
	}
	sort.Sort(c.Entries)
}

func (e entries) Len() int           { return len(e) }
func (e entries) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
func (e entries) Less(i, j int) bool { return e[i].Path < e[j].Path }
