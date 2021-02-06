package ginkgo_demo_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGinkgoDemo(t *testing.T) {
	/*
	一个Ginkgo测试调用Ginkgo的Fail函数发出失败信号，
	使用RegisterFailHandler将这个函数传给Gomega，这是Ginkgo和Gomega唯一的连接点
	 */
	RegisterFailHandler(Fail)
	/*
	告诉Ginkgo开始这个测试套件，如果任意Specs（说明）失败了，Ginkgo会自动使testing.T失败
	 */
	RunSpecs(t, "GinkgoDemo Suite")
}

// 在任何spec之前运行
var _ = BeforeSuite(func() {})

// 在任何spec之后运行.
// 当使用ctrl c 打断测试时，该函数也会运行
// 通过传递带Done参数的函数，可以异步运行BeforeSuite和AfterSuite
var _ = AfterSuite(func() {})
