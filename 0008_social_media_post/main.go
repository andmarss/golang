package main

import (
	"fmt"
	"time"
)

type MoodState int

const (
	MoodStateNeutral MoodState = iota
	MoodStateHappy
	MoodStateSad
	MoodStateAngry
	MoodStateHopeful
	MoodStateThrilled
	MoodStateBored
	MoodStateShy
	MoodStateComical
	MoodStateOnCloudNine
)

type AuditableContent struct {
	TimeCreated  time.Time
	TimeModified time.Time
	CreatedBy    string
	modifiedBy   string
}

type Post struct {
	AuditableContent
	Caption      string
	MessageBody  string
	URL          string
	ImageURI     string
	ThumbnailURI string
	Keywords     []string
	AuthorMood   MoodState
}

var Moods map[string]MoodState

func init() {
	Moods = map[string]MoodState{"netural": MoodStateNeutral, "happy": MoodStateHappy, "sad": MoodStateSad, "angry": MoodStateAngry, "hopeful": MoodStateHopeful, "thrilled": MoodStateThrilled, "bored": MoodStateBored, "shy": MoodStateShy, "comical": MoodStateComical, "cloudnine": MoodStateOnCloudNine}
}

func NewPost(username string, mood MoodState, caption string, messageBody string, url string, imageURI string, thumbnailURI string, keywords []string) *Post {
	auditableContent := AuditableContent{CreatedBy: username, TimeCreated: time.Now()}
	return &Post{Caption: caption, MessageBody: messageBody, URL: url, ImageURI: imageURI, ThumbnailURI: thumbnailURI, AuthorMood: mood, Keywords: keywords, AuditableContent: auditableContent}
}

func main() {
	myPost := NewPost("Mars", Moods["thrilled"], "Go is awesome", "Check out", "https//golang.org", "", "", []string{"go", "golang", "programming language"})

	fmt.Printf("Mypost: %+v\n", myPost)
}
