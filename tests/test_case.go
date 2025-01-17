package tests

import (
	"github.com/goravel/framework/testing"

	"github.com/startup-of-zero-reais/zoo-api/bootstrap"
)

func init() {
	bootstrap.Boot()
}

type TestCase struct {
	testing.TestCase
}
