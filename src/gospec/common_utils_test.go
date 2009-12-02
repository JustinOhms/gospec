// Copyright © 2009 Esko Luontola <www.orfjackal.net>
// This software is released under the Apache License 2.0.
// The license text is at http://www.apache.org/licenses/LICENSE-2.0

package gospec

import (
	"exp/iterable";
	"fmt";
	"testing";
)


// Generic test utilities

var testSpy = "";

func resetTestSpy() {
	testSpy = "";
}

func assertTestSpyHas(expected string, t *testing.T) {
	assertEquals(expected, testSpy, t);
}

func assertEquals(expected interface{}, actual interface{}, t *testing.T) {
	if expected != actual {
		t.Error(fmt.Sprintf("Expected '%v' but was '%v'", expected, actual));
	}
}


// GoSpec specific test utilites

func runSpec(name string, closure func(*Context), context *Context) *taskResult {
	resetTestSpy();
	r := NewSpecRunner();
	return r.execute(name, closure, context);
}

func countSpecNames(specs iterable.Iterable) map[string]int {
	results := make(map[string]int);
	for v := range specs.Iter() {
		spec := v.(*spec);
		key := spec.name;
		if _, preset := results[key]; !preset {
			results[key] = 0;
		}
		results[key]++;
	}
	return results
}


// Test dummies for all test cases

func DummySpecWithNoChildren(c *Context) {
	testSpy += "root";
}

func DummySpecWithOneChild(c *Context) {
	testSpy += "root";
	c.Specify("Child A", func() {
		testSpy += ",a";
	});
}

func DummySpecWithTwoChildren(c *Context) {
	testSpy += "root";
	c.Specify("Child A", func() {
		testSpy += ",a";
	});
	c.Specify("Child B", func() {
		testSpy += ",b";
	});
}

func DummySpecWithNestedChildren(c *Context) {
	testSpy += "root";
	c.Specify("Child A", func() {
		testSpy += ",a";
		c.Specify("Child AA", func() {
			testSpy += ",aa";
		});
	});
}

func DummySpecWithMultipleNestedChildren(c *Context) {
	testSpy += "root";
	c.Specify("Child A", func() {
		testSpy += ",a";
		c.Specify("Child AA", func() {
			testSpy += ",aa";
		});
		c.Specify("Child AB", func() {
			testSpy += ",ab";
		});
	});
	c.Specify("Child B", func() {
		testSpy += ",b";
		c.Specify("Child BA", func() {
			testSpy += ",ba";
		});
		c.Specify("Child BB", func() {
			testSpy += ",bb";
		});
		c.Specify("Child BC", func() {
			testSpy += ",bc";
		});
	});
}

