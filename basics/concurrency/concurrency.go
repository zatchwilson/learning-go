package concurrency

type WebsiteChecker func(string) bool
type result struct { // New type that is associated with the return value of the WebsiteChecker
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result) // Make result channel

	for _, url := range urls {
		go func() {
			resultChannel <- result{url, wc(url)} // this is sending a 'result' struct for each call to wc to the resultChannel with a send statment
		}()
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
