package main

// curl https://sessionize.com/api/v2/3xsxwh64/view/All > ./cmd/session-json/all.json
// go run ./cmd/session-json > ./src/content/sessions/data.json

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	b, err := os.ReadFile("./cmd/session-json/all.json")
	if err != nil {
		return err
	}

	sessions := Sessions{}
	err = json.Unmarshal(b, &sessions)
	if err != nil {
		return err
	}

	w, err := os.Create("out.txt")
	if err != nil {
		return err
	}
	defer w.Close()

	err = output(sessions, w)
	if err != nil {
		return err
	}

	return nil
}

func output(sessions Sessions, w io.Writer) error {
	m := map[string]int{
		"649102": 1,
		"644546": 2,
		"648209": 3,
		"649848": 4,
		"646485": 5,
		"637496": 6,
		"649626": 7,
		"649675": 8,
		"649710": 9,
		"646478": 10,
		"648864": 11,
		"649858": 12,
		"649860": 13,
		"647542": 14,
		"632985": 15,
		"649096": 16,
		"648614": 17,
		"649907": 18,
		"649897": 19,
		"648689": 20,
		"647566": 21,
		"649845": 22,
		"701753": 23,
	}
	for i, s := range sessions.Sessions {
		if id, ok := m[s.ID]; ok {
			sessions.Sessions[i].ID = fmt.Sprintf("%d", id)
		}
	}

	sort.Slice(sessions.Sessions, func(i, j int) bool {
		iid, _ := strconv.ParseInt(sessions.Sessions[i].ID, 10, 0)
		jid, _ := strconv.ParseInt(sessions.Sessions[j].ID, 10, 0)
		return iid < jid
	})

	fmt.Printf("[\n")
	for i, s := range sessions.Sessions {
		fmt.Printf("  {\n")
		id := 0
		switch s.ID {
		case "644546":
			id = 1
		}
		if id > 0 {
			fmt.Printf("    \"id\": %d,\n", id)
		} else {
			fmt.Printf("    \"id\": %s,\n", s.ID)
		}

		typ := ""
		for _, cate := range s.CategoryItems {
			switch cate {
			case 248141:
				typ = "long"
			case 248142:
				typ = "short"
			case 248143:
				typ = "challenge"
			case 248144:
				typ = "lt"
			}
		}
		fmt.Printf("    \"type\": %q,\n", typ)

		level := ""
		for _, cate := range s.CategoryItems {
			switch cate {
			case 248145:
				level = "all"
			case 248146:
				level = "beginner"
			case 248147:
				level = "intermediate"
			case 248148:
				level = "advanced"
			}
		}
		fmt.Printf("    \"level\": %q,\n", level)
		roomID := ""
		switch s.RoomID {
		case 44252:
			roomID = "room-1"
		default:
			roomID = "room-2"
		}
		fmt.Printf("    \"track\": %q,\n", roomID)

		fmt.Printf("    \"startsAt\": %q,\n", s.StartsAt.Format(time.RFC3339))
		fmt.Printf("    \"endsAt\": %q,\n", s.EndsAt.Format(time.RFC3339))

		fmt.Printf("    \"title\": %q,\n", s.Title)
		fmt.Printf("    \"description\": %q,\n", s.Description)

		fmt.Printf("    \"speaker\": {\n")
		speaker := Speaker{}
		for _, sp := range sessions.Speakers {
			for _, sp2 := range s.Speakers {
				if sp.ID == sp2 {
					speaker = sp
				}
			}
		}
		fmt.Printf("      \"avatar\": %q,\n", speaker.ProfilePicture)
		fmt.Printf("      \"name\": %q,\n", fmt.Sprintf("%s %s", speaker.FirstName, speaker.LastName))
		fmt.Printf("      \"company\": %q,\n", speaker.TagLine)
		fmt.Printf("      \"bio\": %q,\n", speaker.Bio)
		twitter := Link{}
		for _, t := range speaker.Links {
			if t.LinkType == "Twitter" {
				twitter = t
			}
		}
		if twitter.URL == "" {
			fmt.Printf("      \"twitter\": null\n")
		} else {
			fmt.Printf("      \"twitter\": %q\n", twitter.URL)
		}
		fmt.Printf("    }\n")

		if i == len(sessions.Sessions)-1 {
			fmt.Printf("  }\n")
		} else {
			fmt.Printf("  },\n")
		}
	}
	fmt.Printf("]\n")
	return nil
}

// JSONの構造体定義
type Sessions struct {
	Sessions []Session `json:"sessions"`
	Speakers []Speaker `json:"speakers"`
}

type Session struct {
	ID               string           `json:"id"`
	Title            string           `json:"title"`
	Description      string           `json:"description"`
	StartsAt         time.Time        `json:"startsAt"`
	EndsAt           time.Time        `json:"endsAt"`
	IsServiceSession bool             `json:"isServiceSession"`
	IsPlenumSession  bool             `json:"isPlenumSession"`
	Speakers         []string         `json:"speakers"`
	CategoryItems    []int            `json:"categoryItems"`
	QuestionAnswers  []QuestionAnswer `json:"questionAnswers"`
	RoomID           int              `json:"roomId"`
	LiveURL          *string          `json:"liveUrl"`
	RecordingURL     *string          `json:"recordingUrl"`
	Status           string           `json:"status"`
	IsInformed       bool             `json:"isInformed"`
	IsConfirmed      bool             `json:"isConfirmed"`
}

type Speaker struct {
	ID              string           `json:"id"`
	FirstName       string           `json:"firstName"`
	LastName        string           `json:"lastName"`
	Bio             string           `json:"bio"`
	TagLine         string           `json:"tagLine"`
	ProfilePicture  string           `json:"profilePicture"`
	IsTopSpeaker    bool             `json:"isTopSpeaker"`
	Links           []Link           `json:"links"`
	Sessions        []int            `json:"sessions"`
	FullName        string           `json:"fullName"`
	CategoryItems   []int            `json:"categoryItems"`
	QuestionAnswers []QuestionAnswer `json:"questionAnswers"`
}

type Link struct {
	Title    string `json:"title"`
	URL      string `json:"url"`
	LinkType string `json:"linkType"`
}

type QuestionAnswer struct {
	QuestionID  int    `json:"questionId"`
	AnswerValue string `json:"answerValue"`
}
