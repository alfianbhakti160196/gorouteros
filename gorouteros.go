package gorouteros

import (
	"fmt"
	"net/http"
	"strings"
)

type Session struct {
	Username string
	Password string
	IP       string
}

func Get(session Session, url string) *http.Request {
	req, err := http.NewRequest("GET", "", strings.NewReader(session.IP+url))
	if err != nil {
		fmt.Println(err)
		return req
	}

	req.SetBasicAuth(session.Username, session.Password)

	return req
}
