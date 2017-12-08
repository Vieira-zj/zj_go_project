package tests_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TestDemo02", func() {
	Describe("Assert tests", func() {
		Context("Context: To and NotTo", func() {
			It("assert NotTo", func() {
				Expect(1).NotTo(Equal(2))
			})

			It("assert Zero", func() {
				Expect(0).To(BeZero())
			})
		})

		Context("Context: Should and ShouldNot", func() {
			It("assert Be", func() {
				Expect(true).Should(BeTrue())
			})
		})
	})
})
