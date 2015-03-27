package gozoo_test

import (
	. "github.com/roxtar/gozoo"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Gozoo", func() {

	It("initializes zookeeper locally", func() {
		z := NewClient()
		err := z.Init("localhost:2181", 1000)
		Expect(err).ToNot(HaveOccurred())
		err = z.Close()
		Expect(err).ToNot(HaveOccurred())
	})

	It("creates a new znode, gets its value and deletes it", func() {
		z := NewClient()
		err := z.Init("localhost:2181", 1000)
		Expect(err).ToNot(HaveOccurred())

		value := []byte("this is a test of create, get and delete")
		path := "/gozoo_create_get_delete_test"
		_, err = z.Create(path, value)
		defer func() {
			z.Delete(path)
			z.Close()
		}()
		Expect(err).ToNot(HaveOccurred())

		actualValue, err := z.Get(path)
		Expect(err).ToNot(HaveOccurred())
		Expect(actualValue).To(Equal(value))

	})

	FIt("creates a new znode and sets it to a different value", func(done Done) {
		defer close(done)
		z := NewClient()
		wasCalled := make(chan string, 1)
		z.Callback = func(zooType int, zooState int, path string) {
			wasCalled <- path
		}

		err := z.Init("localhost:2181", 1000)
		Expect(err).ToNot(HaveOccurred())

		value := []byte("this is a test of create and set")
		path := "/gozoo_create_set_test"
		_, err = z.Create(path, value)
		defer func() {
			z.Delete(path)
			z.Close()
		}()
		Expect(err).ToNot(HaveOccurred())
		Eventually(wasCalled).Should(Receive())

		actualValue, err := z.Get(path)
		Expect(err).ToNot(HaveOccurred())
		Expect(actualValue).To(Equal(value))

		newValue := []byte("this is the modified value")
		err = z.Set(path, newValue)
		Expect(err).ToNot(HaveOccurred())
		Eventually(wasCalled).Should(Receive(Equal(path)))

		actualValue, err = z.Get(path)
		Expect(err).ToNot(HaveOccurred())
		Expect(actualValue).To(Equal(newValue))

	})
})
