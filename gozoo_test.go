package gozoo_test

import (
	"fmt"

	. "github.com/roxtar/gozoo"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Gozoo", func() {

	Context("Initializes", func() {
		It("initializes zookeeper locally", func() {
			z := NewClient()
			err := z.Init("localhost:2181", 1000)
			Expect(err).ToNot(HaveOccurred())
			err = z.Close()
			Expect(err).ToNot(HaveOccurred())
		})

		It("initialization of zookeeper fails for unreachable server", func() {
			z := NewClient()
			err := z.Init("unreachable:2181", 1000)
			Expect(err).To(HaveOccurred())
			z.Close()
		})
	})

	Context("Creates", func() {
		var z *ZooClient
		BeforeEach(func() {
			z = NewClient()
			z.Init("localhost:2181", 1000)
		})

		It("creates a new znode, gets its value and deletes it", func() {
			value := []byte("this is a test of create, get and delete")
			path := "/gozoo_create_get_delete_test"
			_, err := z.Create(path, value)
			defer func() {
				z.Delete(path)
				z.Close()
			}()
			Expect(err).ToNot(HaveOccurred())

			actualValue, err := z.Get(path)
			Expect(err).ToNot(HaveOccurred())
			Expect(actualValue).To(Equal(value))

		})

		It("creates a new znode and sets it to a different value", func(done Done) {
			defer close(done)
			wasCalled := make(chan string, 1)
			z.Callback = func(zooType ZookeeperEvent, zooState ZookeeperState, path string) {
				wasCalled <- path
			}

			value := []byte("this is a test of create and set")
			path := "/gozoo_create_set_test"
			_, err := z.Create(path, value)
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

		It("can create a node with a nil value", func() {
			path := "/gozoo_create_null"
			_, err := z.Create(path, nil)
			Expect(err).ToNot(HaveOccurred())
			defer func() {
				z.Delete(path)
				z.Close()
			}()
		})
	})

	Context("Sets and Gets", func() {
		var z *ZooClient
		BeforeEach(func() {
			z = NewClient()
			z.Init("localhost:2181", 1000)
		})

		It("can set a node with a nil value", func() {
			value := []byte("this is a test of create and set")
			path := "/gozoo_create_null_set"
			_, err := z.Create(path, value)
			defer func() {
				z.Delete(path)
				z.Close()
			}()
			err = z.Set(path, nil)
			Expect(err).ToNot(HaveOccurred())

			actualValue, err := z.Get(path)
			Expect(err).ToNot(HaveOccurred())
			Expect(actualValue).To(Equal([]byte{}))
		})

		It("returns all child nodes", func() {
			path := "/gozoo_get_children"
			z.Create(path, nil)
			defer z.Delete(path)

			child1 := fmt.Sprintf("%s/child1", path)
			z.Create(child1, nil)
			defer z.Delete(child1)

			child2 := fmt.Sprintf("%s/child2", path)
			z.Create(child2, nil)
			defer z.Delete(child2)

			children, err := z.GetChildren(path)
			Expect(err).ToNot(HaveOccurred())
			Expect(children).To(ConsistOf([]string{"child1", "child2"}))
		})
	})
})
