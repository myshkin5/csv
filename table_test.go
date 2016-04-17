package csv_test

import (
	"io"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/myshkin5/csv"
)

var _ = Describe("Table", func() {
	var (
		table *csv.Table
		err   error
	)

	Describe("happy path", func() {
		BeforeEach(func() {
			table, err = csv.New([]string{"col1", "col2"},
				[][]string{{"col1", "col2"}, {"val1", "val2"}, {"val3", "val4"}})
			Expect(err).NotTo(HaveOccurred())
		})

		It("reads values in order", func() {
			Expect(table.Next()).To(Succeed())
			Expect(table.Value("col1")).To(Equal("val1"))
			Expect(table.Value("col2")).To(Equal("val2"))
		})

		It("reads the second record", func() {
			Expect(table.Next()).To(Succeed())
			Expect(table.Next()).To(Succeed())
			Expect(table.Value("col1")).To(Equal("val3"))
			Expect(table.Value("col2")).To(Equal("val4"))
		})

		It("returns EOF when there are no more records", func() {
			Expect(table.Next()).To(Succeed())
			Expect(table.Next()).To(Succeed())
			Expect(table.Next()).To(Equal(io.EOF))
		})
	})

	Describe("sad path", func() {
		Describe("column doesn't exist", func() {
			BeforeEach(func() {
				_, err = csv.New([]string{"col1", "col2"}, [][]string{{"col1", "col3"}})
				Expect(err).To(HaveOccurred())
			})

			It("returns an error", func() {
				Expect(err.Error()).To(Equal("Column col2 not found in header record"))
			})
		})
	})
})
