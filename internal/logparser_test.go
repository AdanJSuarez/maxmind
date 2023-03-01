package internal

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TSLogParser struct{ suite.Suite }

func TestRunTSAccount(t *testing.T) {
	suite.Run(t, new(TSLogParser))
}

func (ts *TSLogParser) BeforeTest(_, _ string) {
}
