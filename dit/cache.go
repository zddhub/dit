package dit

import (
	"encoding/json"
	. "github.com/zddhub/dit/utils"
)

type cache struct {
	Signature string
	Version   int
	Entries   []*object
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
	buffer, err := json.MarshalIndent(*c, "", "  ")
	if err != nil {
		LogD.Println(err)
		return err
	}

	LogI.Printf("%s\n", buffer)
	return WriteFile(filename, buffer, 0644)
}
