package fruit_test

import (
	"github.com/jwfriese/omgfruitapi/fruit"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FruitSource", func() {
	var subject fruit.FruitSource

	BeforeEach(func() {
		fruitFileNames := []string{
			"images/apple.png",
			"images/pineapple.png",
			"images/bananas.png",
			"images/eggplant.png",
		}
		subject = fruit.NewFruitSource(fruitFileNames)
	})

	It("returns fruit data in a cycle", func() {
		fruitName, fruitDescription, fruitImageData := subject.GetNextFruit()
		Expect(fruitName).To(Equal("Apple"))
		Expect(fruitDescription).To(Equal("omg, an apple"))
		Expect(fruitImageData).ToNot(BeNil())

		fruitName, fruitDescription, fruitImageData = subject.GetNextFruit()
		Expect(fruitName).To(Equal("Pineapple"))
		Expect(fruitDescription).To(Equal("dude look, a pineapple"))
		Expect(fruitImageData).ToNot(BeNil())

		fruitName, fruitDescription, fruitImageData = subject.GetNextFruit()
		Expect(fruitName).To(Equal("Banana"))
		Expect(fruitDescription).To(Equal("these bananas are bananas"))
		Expect(fruitImageData).ToNot(BeNil())

		fruitName, fruitDescription, fruitImageData = subject.GetNextFruit()
		Expect(fruitName).To(Equal("Eggplant"))
		Expect(fruitDescription).To(Equal("\"WTF man that's a fruit?\" Yeah, it is look it up"))
		Expect(fruitImageData).ToNot(BeNil())

		// Cycles back to the beginning
		fruitName, fruitDescription, fruitImageData = subject.GetNextFruit()
		Expect(fruitName).To(Equal("Apple"))
		Expect(fruitDescription).To(Equal("omg, an apple"))
		Expect(fruitImageData).ToNot(BeNil())
	})
})
