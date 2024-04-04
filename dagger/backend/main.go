// A generated module for Backend functions
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

import "runtime"

type Backend struct{}

// Build the backend binary
func (m *Backend) Build(
	// The backend source to build
	source *Directory,
	// The os to build for
	// +optional
	// +default="linux"
	os string,
	// The arch to build for
	// +optional
	arch string,
) *File {
	if arch == "" {
		arch = runtime.GOARCH
	}
	return dag.Container().From("golang:latest").
		WithMountedDirectory("/src", source).
		WithWorkdir("/src").
		WithEnvVariable("GOOS", os).
		WithEnvVariable("GOARCH", arch).
		WithExec([]string{"go", "build", "-o", "greetings"}).
		File("/src/greetings")
}

func (b *Backend) Run(source *Directory) *Service {
	bin := b.Build(source, "linux", runtime.GOARCH)

	return dag.Container().From("cgr.dev/chainguard/wolfi-base").
		WithFile("/bin/greetings", bin).
		WithExposedPort(8080).
		WithExec([]string{"/bin/greetings"}).AsService()
}
