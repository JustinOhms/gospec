// Copyright © 2009 Esko Luontola <www.orfjackal.net>
// This software is released under the Apache License 2.0.
// The license text is at http://www.apache.org/licenses/LICENSE-2.0

package gospec

import (
	"testing"
)


func Test__When_a_spec_contains_passing_asserts__Then_the_spec_passes(t *testing.T) {
	results := resultsOfSpec(func(c *Context) {
		c.Then(1).Should.Equal(1)
	})
	assertEquals(1, results.PassCount(), t)
	assertEquals(0, results.FailCount(), t)
}

func Test__When_a_spec_contains_failing_asserts__Then_the_spec_fails(t *testing.T) {
	results := resultsOfSpec(func(c *Context) {
		c.Then(1).Should.Equal(2)
	})
	assertEquals(0, results.PassCount(), t)
	assertEquals(1, results.FailCount(), t)
}

func Test__When_a_passing_spec_has_children__Then_the_children_are_executed(t *testing.T) {
	results := resultsOfSpec(func(c *Context) {
		c.Then(1).Should.Equal(1)
		c.Specify("Child", func() {
		})
	})
	assertEquals(2, results.TotalCount(), t)
}

func Test__When_a_failing_spec_has_children__Then_the_children_are_not_executed(t *testing.T) {
	results := resultsOfSpec(func(c *Context) {
		c.Then(1).Should.Equal(2)
		c.Specify("Child", func() {
		})
	})
	assertEquals(1, results.TotalCount(), t)
}


func resultsOfSpec(spec func(*Context)) *ResultCollector {
	runner := NewRunner()
	runner.AddSpec("RootSpec", spec)
	runner.Run()
	return runner.compileResults()
}

