package cpuid

import (
	"testing"
)

func TestMidr(t *testing.T) {
	dumpMidr()
}

func TestProcFeatures(t *testing.T) {
	dumpProcFeatures()
}

func TestInstAttributess(t *testing.T) {
	dumpInstAttributes()
}
