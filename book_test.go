package ginkgo_demo_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Book", func() {
	var (
		a int
		b int
	)

	// BeforeEach在每次It前执行。
	// 在此重新初始化任何会被测试改变的变量，保证共享变量是最初的值
	BeforeEach(func() {
		a = 1
		b = 2
	})
	// BeforeEach之后、It之前运行
	JustBeforeEach(func() {})
	// AfterEach在每次It后执行
	AfterEach(func() {})

	Describe("a和b的算术运算", func() {
		Context("在进行加法时", func() {
			It("和应该是3", func() {
				By("开始测试")
				Expect(a + b).To(Equal(3))
			})
		})
		Context("在进行减法时", func() {
			It("差应该是-1", func() {
				By("开始测试")
				Expect(a - b).To(Equal(-1))
			})
		})
	})

	Describe("异步测试", func() {
		Context("goruntine在进行测试", func() {
			It("goruntine发生了panic", func(done Done) {
				By("开始测试")
				go func() {
					defer GinkgoRecover()
					By("goruntine开始运行")
					Expect(false).Should(BeTrue())
					// 异步测试，通过关闭done channel来通知Ginkgo测试已经完成
					close(done)
				}()
			})
		})
	})
})
