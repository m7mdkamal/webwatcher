package watcher

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/m7mdkamal/webwatcher/model"
)

// RedditWatcher responsible to get the data from reddit
type RedditWatcher struct {
	subreddit    string `yaml:"subreddit"`
	filterRegexp string
}

// NewRedditWatcher to create @RedditWatcher object
func NewRedditWatcher(sub string, filter string) *RedditWatcher {
	return &RedditWatcher{subreddit: sub, filterRegexp: GenerateRegexp(filter)}
}

// Run the watcher
func (w RedditWatcher) Run() []model.Result {
	var results []model.Result
	resp, err := w.getResponse()
	defer resp.Body.Close()
	if err != nil {
		// do what you have to do defer resp.Body.Close()
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {

	}
	redditResp := RedditResponse{}
	err = json.Unmarshal(body, &redditResp)
	if err != nil {
		panic(err)
	}
	results = append(results, w.filter(redditResp)...)
	// return filtered data
	return results
}

func (w RedditWatcher) getResponse() (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", w.getUrl(), nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "webwatcher:v1")
	return client.Do(req)
}

func (w RedditWatcher) filter(resp RedditResponse) []model.Result {

	var results []model.Result

	for _, post := range resp.Data.Children {
		if ContainKeywords(w.filterRegexp, post.Data.Title, post.Data.Selftext) {
			result := model.Result{}
			result.Title = post.Data.Title
			result.Content = post.Data.Selftext
			result.URL = post.Data.URL
			result.Time = time.Unix(int64(post.Data.CreatedUtc), 0).Format("2006-01-02 15:04:05")
			results = append(results, result)
		}
	}

	//todo:: filter by last checked
	// we need database
	return results
}

func (w RedditWatcher) getUrl() string {
	return fmt.Sprintf("https://www.reddit.com/r/%s/new.json", w.subreddit)
}

type RedditResponse struct {
	Kind string `json:"kind"`
	Data struct {
		After           string `json:"after"`
		Dist            int    `json:"dist"`
		Modhash         string `json:"modhash"`
		WhitelistStatus string `json:"whitelist_status"`
		Children        []struct {
			Kind string `json:"kind"`
			Data struct {
				SubredditID         string        `json:"subreddit_id"`
				ApprovedAtUtc       interface{}   `json:"approved_at_utc"`
				SendReplies         bool          `json:"send_replies"`
				ModReasonBy         interface{}   `json:"mod_reason_by"`
				BannedBy            interface{}   `json:"banned_by"`
				NumReports          interface{}   `json:"num_reports"`
				RemovalReason       interface{}   `json:"removal_reason"`
				Subreddit           string        `json:"subreddit"`
				SelftextHTML        interface{}   `json:"selftext_html"`
				Selftext            string        `json:"selftext"`
				Likes               interface{}   `json:"likes"`
				SuggestedSort       interface{}   `json:"suggested_sort"`
				UserReports         []interface{} `json:"user_reports"`
				SecureMedia         interface{}   `json:"secure_media"`
				IsRedditMediaDomain bool          `json:"is_reddit_media_domain"`
				Saved               bool          `json:"saved"`
				ID                  string        `json:"id"`
				BannedAtUtc         interface{}   `json:"banned_at_utc"`
				ModReasonTitle      interface{}   `json:"mod_reason_title"`
				ViewCount           interface{}   `json:"view_count"`
				Archived            bool          `json:"archived"`
				Clicked             bool          `json:"clicked"`
				NoFollow            bool          `json:"no_follow"`
				Title               string        `json:"title"`
				NumCrossposts       int           `json:"num_crossposts"`
				LinkFlairText       interface{}   `json:"link_flair_text"`
				ModReports          []interface{} `json:"mod_reports"`
				CanModPost          bool          `json:"can_mod_post"`
				IsCrosspostable     bool          `json:"is_crosspostable"`
				Pinned              bool          `json:"pinned"`
				Score               int           `json:"score"`
				ApprovedBy          interface{}   `json:"approved_by"`
				Over18              bool          `json:"over_18"`
				ReportReasons       interface{}   `json:"report_reasons"`
				Domain              string        `json:"domain"`
				Hidden              bool          `json:"hidden"`
				Thumbnail           string        `json:"thumbnail"`
				Edited              interface{}   `json:"edited"`
				LinkFlairCSSClass   interface{}   `json:"link_flair_css_class"`
				AuthorFlairCSSClass interface{}   `json:"author_flair_css_class"`
				ContestMode         bool          `json:"contest_mode"`
				Gilded              int           `json:"gilded"`
				Downs               int           `json:"downs"`
				BrandSafe           bool          `json:"brand_safe"`
				SecureMediaEmbed    struct {
				} `json:"secure_media_embed"`
				MediaEmbed struct {
				} `json:"media_embed"`
				AuthorFlairText       interface{} `json:"author_flair_text"`
				Stickied              bool        `json:"stickied"`
				CanGild               bool        `json:"can_gild"`
				IsSelf                bool        `json:"is_self"`
				ParentWhitelistStatus string      `json:"parent_whitelist_status"`
				Name                  string      `json:"name"`
				Spoiler               bool        `json:"spoiler"`
				Permalink             string      `json:"permalink"`
				SubredditType         string      `json:"subreddit_type"`
				Locked                bool        `json:"locked"`
				HideScore             bool        `json:"hide_score"`
				Created               float64     `json:"created"`
				URL                   string      `json:"url"`
				WhitelistStatus       string      `json:"whitelist_status"`
				Quarantine            bool        `json:"quarantine"`
				Author                string      `json:"author"`
				CreatedUtc            float64     `json:"created_utc"`
				SubredditNamePrefixed string      `json:"subreddit_name_prefixed"`
				Ups                   int         `json:"ups"`
				Media                 interface{} `json:"media"`
				NumComments           int         `json:"num_comments"`
				Visited               bool        `json:"visited"`
				ModNote               interface{} `json:"mod_note"`
				IsVideo               bool        `json:"is_video"`
				Distinguished         interface{} `json:"distinguished"`
			} `json:"data"`
		} `json:"children"`
		Before interface{} `json:"before"`
	} `json:"data"`
}
