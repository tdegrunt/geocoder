package geocoder_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGeocoder(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Geocoder Suite")
}
