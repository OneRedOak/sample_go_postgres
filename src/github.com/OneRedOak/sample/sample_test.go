package sample_test

import (
	. "github.com/OneRedOak/sample"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sample", func() {
    Describe("Saving a name in db", func() {
        Context("and then returning it back", func() {
            It("should result to 'Daniel'", func() {
                Expect(GetResult()).To(Equal("Daniel"))
            })
        })
    })
})
