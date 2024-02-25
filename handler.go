package batch

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"log/slog"
	"time"
)

type GitOptions struct {
	CommitInterval time.Duration
	PushInterval   time.Duration
	Directory      string
	CommitName     string
	CommitEmail    string
}

func AutoCommit(opt *GitOptions) {
	r, err := git.PlainOpen(opt.Directory)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	w, err := r.Worktree()
	if err != nil {
		slog.Error(err.Error())
		return
	}
	status, err := w.Status()
	if len(status) == 0 {
		return
	}

	for k, _ := range status {
		_, err = w.Add(k)
	}

	commit, err := w.Commit("update", &git.CommitOptions{
		Author: &object.Signature{
			Name:  opt.CommitName,
			Email: opt.CommitEmail,
			When:  time.Now(),
		},
	})

	if err != nil {
		slog.Error(err.Error())
		return
	}

	_, err = r.CommitObject(commit)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	slog.Info("commit success")
}

func AutoPush(opt *GitOptions) {
	r, err := git.PlainOpen(opt.Directory)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	err = r.Push(&git.PushOptions{})
	if err != nil {
		slog.Error(err.Error())
		return
	}
	slog.Info("push success")
}
