package dit

import (
	"fmt"
	"strings"
	"time"
)

type userInfo struct {
	name      string
	email     string
	timestamp time.Time
}

func (user userInfo) String() string {
	zone := strings.Split(user.timestamp.Local().String(), " ")[2]
	return fmt.Sprintf("%s <%s> %d %s", user.name, user.email, user.timestamp.Unix(), zone)
}

var user *userInfo

func init() {
	// will read from config file
	user = &userInfo{"zdd", "zddhub@gmail.com", time.Now()}
}
