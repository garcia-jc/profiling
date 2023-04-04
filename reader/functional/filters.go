package functional

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

func usesMasterBranch(item Item) bool {
	return strings.Contains(item.Payload.Ref, "master")
}

func countActors(actors map[string]bool, in Item) map[string]bool {
	actors[in.Actor.Login] = true
	return actors
}

func ReadItems(fn string) []Item {
	f, err := os.Open(fn)
	check(err)
	items := make([]Item, 0)
	err = json.NewDecoder(f).Decode(&items)
	check(err)
	check(f.Close())
	return items
}

func isPublic(i Item) bool {
	return i.Public
}

func Work(items []Item, w io.Writer) int {
	fmt.Fprintln(w, "total items", len(items))
	wg := new(sync.WaitGroup)
	publicRepos := filter(items, isPublic)
	fmt.Fprintln(w, "public repos", len(publicRepos))
	func() {
		wg.Add(1)
		defer wg.Done()
		mainBranch := filter(publicRepos, usesMasterBranch)
		fmt.Fprintln(w, "public repos with main branch", len(mainBranch))
	}()
	func() {
		wg.Add(1)
		defer wg.Done()
		initial := make(map[string]bool)
		distinctAuthors := reduce(items, initial, countActors)
		fmt.Fprintln(w, len(distinctAuthors), "distinct authors contribute to the public repos")
	}()
	func() {
		wg.Add(1)
		defer wg.Done()
		private := func(in Item) bool {
			return !in.Public
		}
		privateRepos := filter(items, private)
		fmt.Fprintln(w, "private repos", len(privateRepos))
	}()
	func() {
		wg.Add(1)
		defer wg.Done()
		gitRefs := forEach(items, func(in Item) string {
			return in.Payload.RefType
		})
		refTypes := make(map[string]struct{})
		refTypes = reduce(gitRefs, refTypes, groupByReferenceType)
		fmt.Fprintf(w, "over %d distinct ref types across %d repos", len(refTypes), len(items))
	}()
	func() {
		wg.Add(1)
		defer wg.Done()
		payloads := forEach(publicRepos, func(i Item) Payload {
			return i.Payload
		})
		notMaster := filter(payloads, func(p Payload) bool {
			return p.MasterBranch != "master"
		})
		fmt.Fprintln(w, "repos with a main branch not called ´master´", len(notMaster))
	}()
	wg.Wait()
	return len(items)
}

func groupByReferenceType(initial map[string]struct{}, refType string) map[string]struct{} {
	initial[refType] = struct{}{}
	return initial
}
