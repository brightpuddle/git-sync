package main

import "errors"

// Repo is a local git repository
type Repo struct {
	id      string
	path    string
	remotes []string
}

// Git runs a git command against this repo
func (r Repo) Git(args ...string) (string, error) {
	return git(r.path, args...)
}

// NewRepo creates a new repo from a local path
func NewRepo(path string) (Repo, error) {
	cfg, err := git(path, "config", "--local", "-l")
	if err != nil {
		return Repo{}, err
	}
	remotes := parseRemotes(cfg)
	if len(remotes) == 0 {
		return Repo{}, errors.New("no remotes")
	}
	id := parseID(remotes[0])
	return Repo{
		id:      id,
		path:    path,
		remotes: remotes,
	}, nil
}

// Pull performs a git pull on this repo
func (repo Repo) Pull() (string, error) {
	return repo.Git("pull")
}

// Push performs a git push on this repo
func (repo Repo) Push() (string, error) {
	return repo.Git("push")
}
