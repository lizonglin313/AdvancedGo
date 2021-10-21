package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

//const REQUEST_URL = "https://api.github.com/users/"
const REQUEST_URL = "https://api.github.com/search/users?q="

type User struct {
	TotalCount        int  `json:"total_count"`
	IncompleteResults bool `json:"incomplete_results"`
	Items             []*Item
}

type Item struct {
	Login             string  `json:"login"`
	Id                int     `json:"id"`
	NodeId            string  `json:"node_id"`
	AvatarUrl         string  `json:"avatar_url"`
	GravatarId        string  `json:"gravatar_id"`
	Url               string  `json:"url"`
	HtmlUrl           string  `json:"html_url"`
	FollowersUrl      string  `json:"followers_url"`
	FollowingUrl      string  `json:"following_url"`
	GistsUrl          string  `json:"gists_url"`
	StarredUrl        string  `json:"starred_url"`
	SubscriptionsUrl  string  `json:"subscriptions_url"`
	OrganizationsUrl  string  `json:"organizations_url"`
	ReposUrl          string  `json:"repos_url"`
	EventsUrl         string  `json:"events_url"`
	ReceivedEventsUrl string  `json:"received_events_url"`
	Type              string  `json:"type"`
	SiteAdmin         bool    `json:"site_admin"`
	Score             float64 `json:"score"`
}

// {
// 	"total_count": 1,
// 	"incomplete_results": false,
// 	"items": [
// 	  {
// 		"login": "lizonglin313",
// 		"id": 40293922,
// 		"node_id": "MDQ6VXNlcjQwMjkzOTIy",
// 		"avatar_url": "https://avatars.githubusercontent.com/u/40293922?v=4",
// 		"gravatar_id": "",
// 		"url": "https://api.github.com/users/lizonglin313",
// 		"html_url": "https://github.com/lizonglin313",
// 		"followers_url": "https://api.github.com/users/lizonglin313/followers",
// 		"following_url": "https://api.github.com/users/lizonglin313/following{/other_user}",
// 		"gists_url": "https://api.github.com/users/lizonglin313/gists{/gist_id}",
// 		"starred_url": "https://api.github.com/users/lizonglin313/starred{/owner}{/repo}",
// 		"subscriptions_url": "https://api.github.com/users/lizonglin313/subscriptions",
// 		"organizations_url": "https://api.github.com/users/lizonglin313/orgs",
// 		"repos_url": "https://api.github.com/users/lizonglin313/repos",
// 		"events_url": "https://api.github.com/users/lizonglin313/events{/privacy}",
// 		"received_events_url": "https://api.github.com/users/lizonglin313/received_events",
// 		"type": "User",
// 		"site_admin": false,
// 		"score": 1.0
// 	  }
// 	]
//   }

func Get(url string) (*User, error) {
	user := &User{}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error of get url: %s", err)
		return nil, err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(user); err != nil {
		fmt.Printf("error of decode respon body: %s", err)
		return nil, err
	}
	return user, nil
}

// type Ts struct {
// 	Name    string   `json:"name"`
// 	Age     int      `json:"age"`
// 	Hobbies []*hobby `json:"hobbies"`
// }

// type hobby struct {
// 	Hname string `json:"hname"`
// }

// func testDecode() {
// 	str := `{
// 		"name":"lzl",
// 		"age":12,
// 		"hobbies":[

// 		]
// 	}`
// }

func main() {

	userName := os.Args[1]
	url := REQUEST_URL + userName
	fmt.Println(url)
	user, err := Get(url)
	if err != nil {
		fmt.Printf("error: %s", err)
	}
	fmt.Println(user)
	fmt.Println(user.Items[0])

	log.Printf("A log")
}
