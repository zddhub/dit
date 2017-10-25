package dit

import (
	"encoding/json"
	"sort"

	utils "github.com/zddhub/dit/utils"
)

type entries []*object
type extensions interface{}

type cache struct {
	Signature  string
	Version    int
	Entries    entries
	Extensions extensions // cachedTree
}

func (c *cache) loadCache(filename string) {
	buffer, err := utils.ReadFile(filename)
	if err != nil {
		utils.LogD.Println(err)
		c.Signature = "DIRC"
		c.Version = 0
		return
	}
	err = json.Unmarshal(buffer, &c)
	if err != nil {
		utils.LogD.Println(err)
		c.Signature = "DIRC"
		c.Version = 0
		return
	}
}

func (c *cache) storeCache(filename string) error {
	sort.Sort(c.Entries)
	buffer, err := json.MarshalIndent(*c, "", "  ")

	if err != nil {
		utils.LogD.Println(err)
		return err
	}

	utils.LogI.Printf("%s\n", buffer)
	return utils.WriteFile(filename, buffer, 0644)
}

func (c *cache) addEntry(obj *object) {
	if yes, i := c.include(obj); yes {
		c.Entries[i] = obj
	} else {
		c.Entries = append(c.Entries, obj)
	}
}

func (c cache) include(obj *object) (bool, int) {
	for i, entry := range c.Entries {
		if obj.Sha1 == entry.Sha1 {
			return true, i
		}
	}
	return false, 0
}

func (e entries) Len() int           { return len(e) }
func (e entries) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
func (e entries) Less(i, j int) bool { return e[i].Path < e[j].Path }
