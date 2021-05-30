package pkg

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Pkg.Sum", func() {
	Context("Sum case1", func() {
		It("Should correct", func() {
			res := Add(1, 2)
			Expect(res).Should(BeEquivalentTo(3))
		})
	})

	Context("Sum case2", func() {
		PIt("Should correct", func() {
			res := Add(-1, 2)
			Expect(res).Should(BeEquivalentTo(1))
		})
	})

	Context("Sum case3", func() {
		It("Should correct", func() {
			res := Add(-1, -2)
			Expect(res).Should(BeEquivalentTo(0))
		})
	})

})
