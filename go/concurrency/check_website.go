package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)
	for _, url := range urls {
		go func() {
			// results[url] = wc(url)
			resultChannel <- result{url, wc(url)} // store 'result' struct to channel
		}()
	}
	for i := 0; i < len(urls); i++ {
		r := <-resultChannel // retrieve 'result' struct and store into 'r'
		results[r.string] = r.bool
	}
	return results
}
