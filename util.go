package goft

import (
	"slices"
	"strings"
)

type Link struct {
	First, Prev, Last, Next string
}

func ParseLink(rawLink string) Link {
	if rawLink == "" {
		return Link{}
	}

	rels := strings.Split(rawLink, ", ")

	var l Link
	for _, s := range rels {
		pair := strings.Split(s, "; rel=")
		link := strings.Trim(pair[0], "<>")
		link, _ = strings.CutPrefix(link, BaseUrlVersioned)
		rel := strings.Trim(pair[1], "\"")

		switch rel {
		case "first":
			l.First = link
		case "prev":
			l.Prev = link
		case "last":
			l.Last = link
		case "next":
			l.Next = link
		}
	}

	return l
}

func HasStudAchievement(achievements []Achievement) bool {
	return slices.ContainsFunc(achievements, func(e Achievement) bool {
		return (e.ID == 1 || e.ID == 218)
	})
}

func HasMemberCursus(cursusStudents []CursusUser) bool {
	return slices.ContainsFunc(cursusStudents, func(e CursusUser) bool {
		return strings.ToLower(e.Grade) == "member"
	})
}
