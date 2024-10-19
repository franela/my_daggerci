package main

import (
	"context"
	"dagger/my-daggerci-app/internal/dagger"
)

type MyDaggerciApp struct{}

func (m *MyDaggerciApp) Dispatch(ctx context.Context) *dagger.File {
	goApp := dag.Gha().Pipeline("go-app").
		OnChanges([]string{"**"}).
		OnPullRequest([]dagger.GhaAction{dagger.Opened, dagger.Synchronize}).
		Module("go-app").
		Call("test")

	nodeApp := dag.Gha().Pipeline("node-app").
		OnChanges([]string{"**"}).
		OnPullRequest([]dagger.GhaAction{dagger.Opened, dagger.Synchronize}).
		Module("node-app").
		Call("test")

	return dag.Gha().Pipelines([]*dagger.GhaPipeline{goApp, nodeApp})
}
