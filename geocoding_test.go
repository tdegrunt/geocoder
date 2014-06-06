package geocoder_test

import (
	. "github.com/tdegrunt/geocoder"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"os"
)

var (
	GM_CLIENT = os.Getenv("GM_CLIENT")
	GM_PK     = os.Getenv("GM_PK")
)

var _ = Describe("Geocoding", func() {
	var (
		gc  *Geocoder
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

		gc.SetClient(GM_CLIENT)
		gc.SetPrivateKey(GM_PK)

		res, err := gc.Geocode("Rijksmuseum")
		Expect(err).NotTo(HaveOccurred())
		Expect(res.Status).To(Equal("OK"))

		bm, err := res.BestMatch()
		Expect(err).NotTo(HaveOccurred())
		Expect(bm.Geometry.Location.Latitude).To(Equal(52.3599976))
		Expect(bm.Geometry.Location.Longitude).To(Equal(4.8852188))
	})
	It("should geocode 'Margaretha van Borsselenlaan 36\n1181 DA Amstelveen\nThe Netherlands' with Enterprise key/client", func() {
		gc.SetClient(GM_CLIENT)
		gc.SetPrivateKey(GM_PK)

		res, err := gc.Geocode("Margaretha van Borsselenlaan 36\n1181 DA Amstelveen\nThe Netherlands")
		Expect(err).NotTo(HaveOccurred())
		Expect(res.Status).To(Equal("OK"))

		bm, err := res.BestMatch()
		Expect(err).NotTo(HaveOccurred())
		Expect(bm.Geometry.Location.Latitude).To(Equal(52.3190839))
		Expect(bm.Geometry.Location.Longitude).To(Equal(4.863637))
	})
})
