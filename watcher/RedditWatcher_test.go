package watcher

/*
func TestGetUrl(t *testing.T) {
	watcher := NewRedditWatcher("java", NewTask("java", "", 30))
	url := watcher.getUrl()
	if url != "https://www.reddit.com/r/java/new.json" {
		t.Error("Expected https://www.reddit.com/r/java/new.json, got ", url)
	}
}

func TestGetResponse(t *testing.T) {
	watcher := NewRedditWatcher("java", NewTask("java", "", 30))
	res, err := watcher.getResponse()
	defer res.Body.Close()
	if err != nil {
		t.Error(err)
	}
	//body, err := ioutil.ReadAll(res.Body)
	//t.Fatal(string(body))
	if res.StatusCode != 200 {
		t.Error("Expected 200 status code, got ", res.StatusCode)
	}

}

func TestRun(t *testing.T) {
	// watcher := NewRedditWatcher("java", NewTask("java", "", 30))
	// reddit := watcher.Run()
	// for _, d := range reddit.Data.Children {
	// 	t.Fatal(d.Data.Title)
	// }
}
*/
