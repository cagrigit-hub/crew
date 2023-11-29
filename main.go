package main

// DEV ENV.
func main() {
	fFetcher := FakeFetcher{}
	ConcurrentChannel("http://golang.org/", &fFetcher)
	// test
}
