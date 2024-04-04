// A generated module for KubesimplifyDaggerWorkshop functions
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

type KubesimplifyDaggerWorkshop struct{}

// Returns a container that echoes whatever string argument is provided
func (m *KubesimplifyDaggerWorkshop) Build(source *Directory) *Directory {
	backendBuild := dag.Backend().Build(source)
	frontendBuild := dag.Frontend().Build(source.Directory("website"))

	return dag.Directory().WithDirectory("/website", frontendBuild).WithFile("/greetings", backendBuild)
}

// Run the whole stack
func (m *KubesimplifyDaggerWorkshop) Run(source *Directory) *Service {
	backendService := dag.Backend().Run(source)
	frontendService := dag.Frontend().Run(source.Directory("website"))
	return dag.Proxy().
		WithService(backendService, "backend", 8080, 8080).
		WithService(frontendService, "frontend", 8081, 80).
		Service()
}
