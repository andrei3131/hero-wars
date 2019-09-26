package player_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPlayerBuilder(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Player Builder Suite")
}
