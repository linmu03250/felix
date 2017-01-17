// Copyright (c) 2016-2017 Tigera, Inc. All rights reserved.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package set_test

import (
	"github.com/projectcalico/felix/go/felix/set"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Set", func() {
	var s set.Set
	BeforeEach(func() {
		s = set.New()
	})

<<<<<<< 369494803c8730eca18e569b1b3edfdcf440f4eb
	It("should be empty", func() {
		Expect(s.Len()).To(BeZero())
	})
	It("should iterate over no items", func() {
		called := false
		s.Iter(func(item interface{}) error {
			called = true
			return nil
		})
		Expect(called).To(BeFalse())
	})
	It("should do nothing on clear", func() {
		s.Clear()
		Expect(s.Len()).To(BeZero())
=======
	Describe("Set created by FromArray", func() {
		BeforeEach(func() {
			s = set.FromArray([]int{1, 2})
		})
		It("should contain 1", func() {
			Expect(s.Contains(1)).To(BeTrue())
		})
		It("should contain 2", func() {
			Expect(s.Contains(2)).To(BeTrue())
		})
		It("should not contain 3", func() {
			Expect(s.Contains(3)).To(BeFalse())
		})
>>>>>>> Add set.FromArray() and Set.AddAll() functions.
	})

	Describe("after adding 1 and 2", func() {
		BeforeEach(func() {
			s.Add(1)
			s.Add(2)
			s.Add(2) // Duplicate should have no effect
		})
		It("should contain 1", func() {
			Expect(s.Contains(1)).To(BeTrue())
		})
		It("should contain 2", func() {
			Expect(s.Contains(2)).To(BeTrue())
		})
		It("should not contain 3", func() {
			Expect(s.Contains(3)).To(BeFalse())
		})
		It("should iterate over 1 and 2 in some order", func() {
			seen1 := false
			seen2 := false
			s.Iter(func(item interface{}) error {
				if item.(int) == 1 {
					Expect(seen1).To(BeFalse())
					seen1 = true
				} else if item.(int) == 2 {
					Expect(seen2).To(BeFalse())
					seen2 = true
				} else {
					Fail("Unexpected item")
				}
				return nil
			})
			Expect(seen1).To(BeTrue())
			Expect(seen2).To(BeTrue())
		})
		It("should allow remove during iteration", func() {
			s.Iter(func(item interface{}) error {
				if item.(int) == 1 {
					return set.RemoveItem
				}
				return nil
			})
			Expect(s.Contains(1)).To(BeFalse())
			Expect(s.Contains(2)).To(BeTrue())
		})
		It("should support stopping iteration", func() {
			iterationStarted := false
			s.Iter(func(item interface{}) error {
				if iterationStarted {
					Fail("Iteration continued after stop")
				}
				iterationStarted = true
				return set.StopIteration
			})
			Expect(s.Contains(1)).To(BeTrue())
			Expect(s.Contains(2)).To(BeTrue())
		})
		It("can copy a Set", func() {
			c := s.Copy()
			Expect(c.Len()).To(Equal(s.Len()))
			Expect(c).NotTo(BeIdenticalTo(s)) // Check they're not the same object.
			Expect(c).To(Equal(s))            // DeepEquals, will check the contents.
		})
		It("should correctly determine set equality", func() {
			c := s.Copy()
			Expect(c.Equals(s)).To(BeTrue())
			Expect(s.Equals(c)).To(BeTrue())
			c.Add(3)
			Expect(c.Equals(s)).To(BeFalse())
			Expect(s.Equals(c)).To(BeFalse())
			c.Discard(2)
			Expect(c.Equals(s)).To(BeFalse())
			Expect(s.Equals(c)).To(BeFalse())
			c.Add(2)
			c.Discard(3)
			Expect(c.Equals(s)).To(BeTrue())
			Expect(s.Equals(c)).To(BeTrue())
		})

		Describe("after removing 2", func() {
			BeforeEach(func() {
				s.Discard(2)
			})
			It("should contain 1", func() {
				Expect(s.Contains(1)).To(BeTrue())
			})
			It("should not contain 2", func() {
				Expect(s.Contains(2)).To(BeFalse())
			})
			It("should not contain 3", func() {
				Expect(s.Contains(3)).To(BeFalse())
			})
		})
<<<<<<< 369494803c8730eca18e569b1b3edfdcf440f4eb

		Describe("after Clear()", func() {
			BeforeEach(func() {
				s.Clear()
			})
			It("should be empty", func() {
				Expect(s.Len()).To(BeZero())
=======
		Describe("after using AddAll to add 2, 3, 4", func() {
			BeforeEach(func() {
				s.AddAll([]int{2, 3, 4})
			})
			It("should contain 1", func() {
				Expect(s.Contains(1)).To(BeTrue())
			})
			It("should contain 2", func() {
				Expect(s.Contains(2)).To(BeTrue())
			})
			It("should contain 3", func() {
				Expect(s.Contains(3)).To(BeTrue())
			})
			It("should contain 4", func() {
				Expect(s.Contains(4)).To(BeTrue())
>>>>>>> Add set.FromArray() and Set.AddAll() functions.
			})
		})
	})
})

var _ = Describe("EmptySet", func() {
	var empty set.Set
	BeforeEach(func() {
		empty = set.Empty()
	})
	It("has length 0", func() {
		Expect(empty.Len()).To(Equal(0))
	})
	It("should panic on add", func() {
		Expect(func() { empty.Add("foo") }).To(Panic())
	})
	It("should ignore discard", func() {
		Expect(func() { empty.Discard("foo") }).NotTo(Panic())
	})
	It("should iterate 0 times", func() {
		empty.Iter(func(item interface{}) error {
			Fail("Iterated > 0 times")
			return nil
		})
	})
})
