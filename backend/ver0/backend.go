package backend

import (
	"encoding/json"
	"fmt"
	"net/http"
	"appengine"
	"appengine/urlfetch"
)

//Response is a slice because the reddit api starts as an array with two objects
type Response []struct {
	Data struct {
		Children []struct {
			Comments struct {
				Body string `json:"body"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

func init() {
	http.HandleFunc("/api/getcomments", getComments)
	// http.HandleFunc("/api/test", testHiddenComment)
}

func fetchComments(c appengine.Context, url string) ([]string, error) {
	var comments []string

	client := urlfetch.Client(c)
	res, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error Get: %v", err)
	}

	content := new(Response)
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&content)
	if err != nil {
		return nil, fmt.Errorf("Error decoding: %v", err)
	}

	for _, el := range (*content)[1].Data.Children {
		comments = append(comments, el.Comments.Body)
	}

	return comments, nil

}

func fetchHiddenComment(c appengine.Context, url string) (string, error) {
	comment := ""
	client := urlfetch.Client(c)
	res, err := client.Get(url)
	if err != nil {
		return "", fmt.Errorf("Error Get: %v", err)
	}

	content := new(Response)
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&content)
	if err != nil {
		return "", fmt.Errorf("Error decodign: %v", err)
	}

	comment = (*content)[1].Data.Children[0].Comments.Body
	return comment, nil
}

func getComments(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	url := "https://www.reddit.com/r/AskReddit/comments/2pq2iz/reddit_what_song_consistently_gives_you_the_chills.json?limit=500"

	comments, err := fetchComments(c, url)
	if err != nil {
		http.Error(w, "Reddit API not working", http.StatusInternalServerError)
		c.Errorf("Fetching comments: %v", err)
	}

	for _, el := range comments {
		fmt.Fprint(w, el, "\n\n")
	}
}

