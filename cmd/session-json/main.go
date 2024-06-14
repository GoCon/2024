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

type SessionMap struct {
	No    int
	Slide string
}

func output(sessions Sessions, w io.Writer) error {
	m := map[string]SessionMap{
		"649102": {No: 1, Slide: "https://speakerdeck.com/utgwkk/go-conference-2024"},
		"644546": {No: 2, Slide: "https://audience.ahaslides.com/cl965inb88/review?lookback-tab=slides"},
		"648209": {No: 3, Slide: "https://speakerdeck.com/k1low/go-conference-2024"},
		"649848": {No: 4, Slide: "https://speakerdeck.com/masumomo/custom-logging-with-slog-making-logging-fun-again"},
		"646485": {No: 5, Slide: "https://github.com/97vaibhav/Go-Conference-2024-Tokyo"},
		"637496": {No: 6, Slide: "https://speakerdeck.com/yamatoya/go1-dot-21karadao-ru-sareta-go-toolchainnoshi-zu-miwomarututojie-shuo/"},
		"649626": {No: 7, Slide: "https://docs.google.com/presentation/d/e/2PACX-1vSWVLveC-AdAoBebAVx3lU4C8bUSBN5_vx-1x_4AWgwddild-kkEqypIp0ox8gGA32SMPu3xchsqGDw/pub?start=false&loop=false&delayms=3000&slide=id.g2e3b19212c6_0_0"},
		"649675": {No: 8, Slide: "https://speakerdeck.com/shohata/gonolanguage-server-protocolshi-zhuang-gopls-nozi-dong-bu-wan-noshi-zu-miwoxue-bu"},
		"649710": {No: 9, Slide: "https://nymphium.github.io/pdf/gocon2024.html"},
		"646478": {No: 10, Slide: "https://speakerdeck.com/replu/understanding-the-swisstable-being-considered-to-improve-map-performance"},
		"648864": {No: 11, Slide: "https://speakerdeck.com/biwashi/go-conference-2024-automating-feature-flag-instrumentation-by-constructing-go-ast-from-unified-diff-format"},
		"649858": {No: 12, Slide: "https://speakerdeck.com/convto/understanding-gob-encoding"},
		"649860": {No: 13, Slide: "https://docs.google.com/presentation/d/1X5dXShWTmjhQbXH7vXHnTLJ5Tca7QhU4Pq1YugGiIHs/edit#slide=id.p"},
		"647542": {No: 14, Slide: "https://speakerdeck.com/qualiarts/xiang-jie-fixing-for-loops-in-go-1-dot-22-zi-zuo-linterwogolangci-linthekontoribiyutositahua"},
		"632985": {No: 15, Slide: "https://speakerdeck.com/kuro_kurorrr/deep-dive-into-deadcode"},
		"649096": {No: 16, Slide: "https://speakerdeck.com/shinnosuke_kishida/go-get-dekao-lu-siteiru-huairusisutemunoju-dong-nituite"},
		"648614": {No: 17, Slide: "https://note.com/thousanda/n/n2b1c43b1814b"},
		"649907": {No: 18, Slide: ""},
		"649897": {No: 19, Slide: "https://speakerdeck.com/abekoh/table-driven-testing-nifu-rarenai-gonotesutopatan"},
		"648689": {No: 20, Slide: "https://speakerdeck.com/bmf_san/zi-zuo-httprutakaraxin-siiservemuxhe"},
		"647566": {No: 21, Slide: "https://speakerdeck.com/uhzz/zi-dong-sheng-cheng-saretahttpendopointogotonikasutamumidoruueawocha-ru-sitaihua"},
		"649845": {No: 22, Slide: "https://speakerdeck.com/bellwood4486/goconference2024lt-proxy"},
		"701753": {No: 23, Slide: "https://speakerdeck.com/uchihara/delish-kitchenniokerumasutadetakiyatusiyuzhan-lue-tosonoli-shi-de-bian-qian"},
	}
	for i, s := range sessions.Sessions {
		if id, ok := m[s.ID]; ok {
			sessions.Sessions[i].ID = fmt.Sprintf("%d", id.No)
			sessions.Sessions[i].Slide = id.Slide
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
		fmt.Printf("    \"slide\": %q,\n", s.Slide)

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
	Slide            string           `json:"slide"`
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
