package backend

import (
"fmt"
"net/http"
"encoding/json"

"appengine"
"appengine/urlfetch"
)

//Response is a slice because the reddit api starts as an array with two objects
type Response []struct {
	Data struct {
		Children []struct {
			Comments struct {
					Body string `json:"body"`
				}`json:"data"`
		} `json:"children"`
	} `json:"data"`
}

func init() {
	http.HandleFunc("/api/getcomments", getComments)
}

func getComments(w http.ResponseWriter, r *http.Request) {
	var comments []string
	url := "https://www.reddit.com/r/videos/comments/3vgdsb/recruitment_2016.json"

	//starting newContext and getting data from the url
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)
	resp, err := client.Get(url)
	if err != nil { fmt.Fprint(w, "Error client.Get(): ", err) }
	
	//creating new response object
	re := new(Response)
	fmt.Fprint(w, "Response struct: ", re, "\n")

	//decoding the data to json into our Response array
	errTwo := json.NewDecoder(resp.Body).Decode(&re)
	if errTwo != nil {  fmt.Fprint(w, "Error decoding: ", errTwo, "\n") }

	for _, el := range (*re)[1].Data.Children {
		comments = append(comments, el.Comments.Body)
	}

	for _, el := range comments {
		fmt.Fprint(w, el, "\n|||||||||||||||||||||||||||||||||||||||||||\n")
	}
}

