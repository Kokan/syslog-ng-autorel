package gitminer

import (
	"fmt"
	"testing"

	"github.com/darxtrix/syslog-ng-autorel/internal/pkg/goroutinepool"
)

func TestGetMergeRequest(t *testing.T) {
	pool := goroutinepool.NewGoRoutinePool(2, 10, 1)
	gm, err := GetMiner("../../../temp/syslog-ng", "balabit", "syslog-ng", "<token>", "./temp", pool)
	if err != nil {
		panic(err)
	}
	firstCommit := "7be16513a3722488f5e3224a39f7076e6167f72b"
	lastCommit := "82a7a012353143314d8482b7f249e56367a4da59"

	mergeRequests, err := gm.GetMergeRequests(firstCommit, lastCommit)
	if err != nil {
		panic(err)
	}
	panic(err)
	fmt.Println(len(mergeRequests))
}
