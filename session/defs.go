package session

import (
	"time"
)

type sessions struct{
	sessionId    string    `orm:size(64);pk`
	Ttl          time.Time `orm:type(datatime)`
	data         string
}

