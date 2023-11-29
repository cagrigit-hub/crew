package main

func worker(url string, ch chan []string, fetcher Fetcher) {
	urls, err := fetcher.Fetch(url)
	if err != nil {
		ch <- []string{}
	} else {
		ch <- urls
	}
}

func master(ch chan []string, fetcher Fetcher) {
	n := 1
	fetched := make(map[string]bool)
	for urls := range ch {
		for _, url := range urls {
			if !fetched[url] {
				fetched[url] = true
				n++
				go worker(url, ch, fetcher)
			}
		}
		n--
		if n == 0 {
			break
		}
	}
}

func ConcurrentChannel(url string, fetcher Fetcher) {
	ch := make(chan []string)
	go func() {
		ch <- []string{url}
	}()
	master(ch, fetcher)
}

type Fetcher interface {
	Fetch(url string) (urls []string, err error)
}

type FakeFetcher struct{}

func (f *FakeFetcher) Fetch(url string) (urls []string, err error) {
	return []string{
		"http://golang.org/",
		"http://golang.org/pkg/",
		"http://golang.org/cmd/",
		"http://google.com/",
	}, nil
}
