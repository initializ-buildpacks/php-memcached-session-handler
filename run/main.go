package main

import (
	"os"

	phpmemcachedhandler "github.com/initializ-buildpacks/php-memcached-session-handler"
	"github.com/paketo-buildpacks/packit/v2"
	"github.com/paketo-buildpacks/packit/v2/scribe"
	"github.com/paketo-buildpacks/packit/v2/servicebindings"
)

func main() {
	logEmitter := scribe.NewEmitter(os.Stdout).WithLevel(os.Getenv("BP_LOG_LEVEL"))
	serviceResolver := servicebindings.NewResolver()

	packit.Run(
		phpmemcachedhandler.Detect(
			serviceResolver,
		),
		phpmemcachedhandler.Build(
			phpmemcachedhandler.NewMemcachedConfigParser(),
			serviceResolver,
			phpmemcachedhandler.NewMemcachedConfigWriter(logEmitter),
			logEmitter,
		),
	)
}
