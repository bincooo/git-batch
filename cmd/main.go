package main

import (
	"fmt"
	batch "git-batch"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"log/slog"
	"os"
	"time"
)

func main() {
	rootCmd := NewGitCommand()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func NewGitCommand() *cobra.Command {
	opts := newGitOptions()
	cmd := &cobra.Command{
		Use:   "git-batch",
		Short: "git-batch is a tool about sync by git",
		Run: func(cmd *cobra.Command, args []string) {
			run(opts)
		},
	}

	addFlag(cmd, opts)
	return cmd
}

func addFlag(cmd *cobra.Command, opts *batch.GitOptions) {
	cmd.Flags().StringVarP(&opts.Directory, "path", "p", ".", "Git repository path")
	cmd.Flags().StringVarP(&opts.CommitName, "name", "n", "git-batch", "Commit name")
	cmd.Flags().StringVarP(&opts.CommitEmail, "email", "e", "", "Commit email")
	cmd.Flags().DurationVar(&opts.CommitInterval, "commit", time.Minute, "Commit interval")
	cmd.Flags().DurationVar(&opts.PushInterval, "push", time.Hour, "Push interval")
}

func newGitOptions() *batch.GitOptions {
	ops := &batch.GitOptions{}
	return ops
}

func run(opt *batch.GitOptions) {
	stop := make(chan int)
	tc := time.NewTicker(opt.CommitInterval)
	tp := time.NewTicker(opt.PushInterval)

	go func() {
		batch.AutoCommit(opt)
		for {
			select {
			case <-tc.C:
				slog.Info("try to commit")
				batch.AutoCommit(opt)
			}
		}
	}()

	go func() {
		batch.AutoPush(opt)
		time.Sleep(time.Second * 10)
		for {
			select {
			case <-tp.C:
				glog.V(4).Info("try to push")
				batch.AutoPush(opt)
			}
		}
	}()

	<-stop
}
