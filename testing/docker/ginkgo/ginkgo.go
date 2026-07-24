package ginkgo

import (
	"github.com/hanzo-fi/go-libs/v2/logging"
	"github.com/hanzo-fi/go-libs/v2/testing/docker"
	. "github.com/onsi/ginkgo/v2"
)

var pool = new(docker.Pool)

func ActualDockerPool() *docker.Pool {
	return pool
}

func WithNewDockerPool(logger logging.Logger, fn func()) bool {
	return Context("With docker pool", func() {
		BeforeEach(func() {
			*pool = *docker.NewPool(GinkgoT(), logger)
		})
		fn()
	})
}
