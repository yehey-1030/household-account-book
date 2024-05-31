package main

import (
	"fmt"
	"github.com/yehey-1030/household-account-book/go/cmd/household-account-book/bootstrap"
	"strconv"
	"time"
)

var (
	Version   string
	BuildAt   string
	GoVersion string
)

func main() {
	// Build console message
	fmt.Println(fmt.Sprintf("household-account-book-server: now is %s", time.Now()))
	seconds, _ := strconv.ParseInt(BuildAt, 10, 64)
	buildAt := time.Unix(seconds, 0)
	var startupMessage = fmt.Sprintf("household-account-book-server %s (build at %s, by %s) starts at %s", Version, buildAt.UTC().Format(time.RFC3339), GoVersion, time.Now().UTC().Format(time.RFC3339))
	fmt.Println(startupMessage)

	bootstrap.Init()
	bootstrap.Run(startupMessage, Version)
}
