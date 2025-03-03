package goft

import (
	"fmt"
	"time"
)

func (e Event) Duration() time.Duration {
	return e.EndAt.Sub(e.BeginAt)
}

func (e Event) Dates() string {
	end := "15h04"
	if e.BeginAt.YearDay() != e.EndAt.YearDay() {
		end = "Mon 02 15h04"
	}
	if e.BeginAt.Month() != e.EndAt.Month() {
		end = "Mon 02 Jan 15h04"
	}

	start := "Mon 02 Jan 15h04"
	if e.BeginAt.Year() != e.EndAt.Year() {
		start = "Mon 02 Jan 15h04 06"
		end = "Mon 02 Jan 15h04 06"
	}

	return fmt.Sprintf(
		"📅 %s → %s ⏳ %s",
		e.BeginAt.Local().Format(start),
		e.EndAt.Local().Format(end),
		e.Duration(),
	)
}

func (e Event) String() string {
	return fmt.Sprintf(`%d ● %s
%s
%s 📍 %s 
❌ %dh 👥 %d/%d
%s`,
		e.ID, e.Name,
		e.Dates(),
		e.Kind, e.Location,
		e.ProhibitionOfCancellation, e.NbrSubscribers, e.MaxPeople,
		e.Description,
	)

}

// Short does the same as String, but with less information
func (e Event) Short() string {
	return fmt.Sprintf(`%d ● %s
%s
%s 📍 %s 
❌ %dh 👥 %d/%d`,
		e.ID, e.Name,
		e.Dates(),
		e.Kind, e.Location,
		e.ProhibitionOfCancellation, e.NbrSubscribers, e.MaxPeople,
	)
}

func (e Event) ShortInfo() string {
	return fmt.Sprintf(`%s
%s 📍 %s 
❌ %dh 👥 %d/%d`,
		e.Dates(),
		e.Kind, e.Location,
		e.ProhibitionOfCancellation, e.NbrSubscribers, e.MaxPeople,
	)
}

func (e Event) Title() string {
	return fmt.Sprintf("%d ● %s", e.ID, e.Name)
}
