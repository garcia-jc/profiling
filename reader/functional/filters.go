package functional

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
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

func Work(items []Item, w io.Writer) (int, int) {
	fmt.Fprintln(w, "total items", len(items))
	publicRepos := filter(items, isPublic)
	fmt.Fprintln(w, "public repos", len(publicRepos))
	mainBranch := filter(publicRepos, usesMasterBranch)
	fmt.Fprintln(w, "public repos with main branch", len(mainBranch))
	in := make(map[string]bool)
	distinctAuthors := reduce(items, in, countActors)
	fmt.Fprintln(w, len(distinctAuthors), "distinct authors contribute to the public repos")
	private := func(in Item) bool {
		return !in.Public
	}
	privateRepos := filter(items, private)
	fmt.Fprintln(w, "private repos", len(privateRepos))
	gitRefs := forEach(items, func(in Item) string {
		return in.Payload.RefType
	})
	refTypes := make(map[string]struct{})
	refTypes = reduce(gitRefs, refTypes, func(initial map[string]struct{}, refType string) map[string]struct{} {
		initial[refType] = struct{}{}
		return initial
	})
	fmt.Fprintf(w, "over %d distinct ref types across %d repos", len(refTypes), len(items))
	return len(publicRepos), len(privateRepos)
}
