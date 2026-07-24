package bunpaginate_test

import (
	"testing"

	"github.com/hanzo-fi/go-libs/testing/docker"
	"github.com/hanzo-fi/go-libs/testing/utils"

	"github.com/hanzo-fi/go-libs/testing/platform/pgtesting"

	"github.com/hanzo-fi/go-libs/logging"
)

var srv *pgtesting.PostgresServer

func TestMain(m *testing.M) {
	utils.WithTestMain(func(t *utils.TestingTForMain) int {
		srv = pgtesting.CreatePostgresServer(t, docker.NewPool(t, logging.Testing()))

		return m.Run()
	})
}
