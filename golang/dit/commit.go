package dit

import (
	"bytes"
	"fmt"

	utils "github.com/zddhub/dit/golang/utils"
)

type commit struct {
	object
	tree      *tree
	parent    *commit
	message   string
	author    *userInfo
	committer *userInfo
}

func (c commit) Content() []byte {
	var b bytes.Buffer
	var content string
	content += fmt.Sprintf("tree %s\n", c.tree.Sha1)
	if c.parent.Sha1 != "" {
		content += fmt.Sprintf("parent %s\n", c.parent.Sha1)
	}
	content += fmt.Sprintf("author %s\n", c.author)
	content += fmt.Sprintf("committer %s\n", c.committer)
	content += fmt.Sprintf("\n%s\n", c.message)
	utils.LogI.Println(content)

	b.Write([]byte(content))
	return b.Bytes()
}
