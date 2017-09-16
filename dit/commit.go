package dit

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	. "github.com/zddhub/dit/utils"
)

type commit struct {
	object
	tree      *tree
	parent    *commit
	message   string
	author    userInfo
	committer userInfo
}

type userInfo struct {
	name      string
	email     string
	timestamp time.Time
}

var user userInfo

func init() {
	// will read from config file
	user = userInfo{"zdd", "zddhub@gmail.com", time.Now()}
}

func (c commit) Content() []byte {
	var b bytes.Buffer
	var content string
	content += fmt.Sprintf("tree %s\n", c.tree.Sha1String())
	content += fmt.Sprintf("author %s\n", c.author)
	content += fmt.Sprintf("committer %s\n", c.committer)
	content += fmt.Sprintf("\n%s\n", c.message)
	LogI.Println(content)

	b.Write([]byte(content))
	return b.Bytes()
}

func (user userInfo) String() string {
	zone := strings.Split(user.timestamp.Local().String(), " ")[2]
	return fmt.Sprintf("%s <%s> %d %s", user.name, user.email, user.timestamp.Unix(), zone)
}

func (repo *repository) Commit(message string) {
	tree := repo.NewTree()
	tree.Write(tree.Content())

	user.timestamp = time.Now()
	commit := &commit{object{Type: "commit"}, tree, nil, message, user, user}
	commit.Write(commit.Content())
}
