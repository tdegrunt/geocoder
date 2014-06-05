package geocoder_test

import (
	. "github.com/tdegrunt/geocoder"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Geocoding", func() {
	var (
		gc *Geocoder
		err error
	)
	BeforeEach(func() {
		gc, err = NewGeocoder()
		Expect(err).NotTo(HaveOccurred())
	})
	It("should geocode 'Rijksmuseum'", func() {
		res, err := gc.Geocode("Rijksmuseum")
		Expect(err).NotTo(HaveOccurred())
		Expect(res.Status).To(Equal("OK"))
		
		bm, err := res.BestMatch()
		Expect(err).NotTo(HaveOccurred())
		Expect(bm.Geometry.Location.Latitude).To(Equal(52.3599976))
		Expect(bm.Geometry.Location.Longitude).To(Equal(4.8852188))
	})
	It("should geocode 'Rijksmuseum' with Enterprise key/client", func() {
		gc.SetClient("yourclientidhere")
		gc.SetPrivateKey("yourprivatekeyhere")
		
		res, err := gc.Geocode("Rijksmuseum")
		Expect(err).NotTo(HaveOccurred())
		Expect(res.Status).To(Equal("OK"))
		
		bm, err := res.BestMatch()
		Expect(err).NotTo(HaveOccurred())
		Expect(bm.Geometry.Location.Latitude).To(Equal(52.3599976))
		Expect(bm.Geometry.Location.Longitude).To(Equal(4.8852188))
	})
})
