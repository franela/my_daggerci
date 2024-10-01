// A generated module for MyDaggerciApp functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/my-daggerci-app/internal/dagger"
	"fmt"
)

type MyDaggerciApp struct{}

func (m *MyDaggerciApp) Dispatch(ctx context.Context, eventTrigger *dagger.File) error {

	c, err := eventTrigger.Contents(ctx)

	fmt.Println("LALALALA")
	fmt.Println(c, err)

	return dag.Gha(eventTrigger).WithPipeline("test").
		WithRunsOn("dagger-2c").
		WithOnPullRequest([]dagger.GhaAction{dagger.Opened, dagger.Synchronize}).
		WithModule("go-app").
		WithOnChanges([]string{"**"}).
		Call(ctx, "test")
}

// Returns a container that echoes whatever string argument is provided
func (m *MyDaggerciApp) Test(ctx context.Context,
	// +defaultPath="/"
	src *dagger.Directory,
) (string, error) {
	return dag.Container().From("golang:alpine").WithMountedDirectory("/app", src).
		WithWorkdir("/app").
		WithExec([]string{"go", "test", "-v", "./..."}).Stdout(ctx)

}
