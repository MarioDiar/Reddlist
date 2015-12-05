package backend

import (
"fmt"
"net/http"
"encoding/json"

"appengine"
"appengine/urlfetch"
)

//struct of reddit json api
//Respons is a slice because for some reason, the reddit api starts as an array
type Response []struct {
	Parent struct {
		Data struct {
			Children []struct {
				Com Comment
			}
		}
	}
}
//struct of comment
type Comment struct {
	Name string `json:"body"`
}

func init() {
	http.HandleFunc("/api/getcomments", getComments)
}

func getComments(w http.ResponseWriter, r *http.Request) {

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

	fmt.Fprint(w, "Response struct: ", re)
}

