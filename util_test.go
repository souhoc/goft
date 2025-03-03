package goft_test

import (
	"testing"

	"github.com/souhoc/goft"
)

const rawLink = `<https://api.intra.42.fr/v2/users/shocquen/events?page=1&per_page=20>; rel="first", <https://api.intra.42.fr/v2/users/shocquen/events?page=1&per_page=20>; rel="prev", <https://api.intra.42.fr/v2/users/shocquen/events?page=5&per_page=20>; rel="last", <https://api.intra.42.fr/v2/users/shocquen/events?page=3&per_page=20>; rel="next"`

var goodLink = goft.Link{
	First: "/users/shocquen/events?page=1&per_page=20",
	Prev:  "/users/shocquen/events?page=1&per_page=20",
	Last:  "/users/shocquen/events?page=5&per_page=20",
	Next:  "/users/shocquen/events?page=3&per_page=20",
}

func TestParseLink(t *testing.T) {

	l := goft.ParseLink(rawLink)
	t.Log(l)
	if l != goodLink {
		t.Fail()
	}

}
