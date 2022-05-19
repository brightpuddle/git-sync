package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

// Src is a source folder, i.e. a collection of repositories
type Src struct {
	remotes map[string]set
	repos   []Repo
}

// NewSrc reads all repos in a root source folder
func NewSrc(root string) (Src, error) {
	src := Src{
		remotes: map[string]set{},
		repos:   []Repo{},
	}
	err := filepath.WalkDir(root,
		func(path string, info fs.DirEntry, err error) error {
			for _, dir := range exclude {
				if info.Name() == dir {
					return filepath.SkipDir
				}
			}
			if info.Name() == ".git" {
				repo, err := NewRepo(filepath.Dir(path))
				if err != nil {
					return filepath.SkipDir
				}
				if _, ok := src.remotes[repo.id]; !ok {
					src.remotes[repo.id] = set{}
				}
				for _, url := range repo.remotes {
					id := parseID(url)
					src.remotes[repo.id].add(id)
				}
				// TODO add submodules to src.remotes
				src.repos = append(src.repos, repo)
				return filepath.SkipDir
			}
			return nil
		})
	return src, err
}

// NormalizeRemotes adds any missing remotes for dual-remote repos
func (src Src) NormalizeRemotes() {
	for _, repo := range src.repos {
		if len(repo.remotes) == 0 {
			continue
		}
		allRemotes := src.remotes[repo.id]
		if len(repo.remotes) < len(allRemotes) {
			fmt.Println(repo.path, "missing remotes")
			// TODO add missing remote as second origin
		}
	}
	// TODO also check that all remotes are the same origin
	// This will ensure git push/pull will update all
	// TODO check that remotes are ssh and not https
}
