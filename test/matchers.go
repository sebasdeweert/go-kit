package test

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/golang/mock/gomock"
)

type withElementsMatcher struct {
	expected interface{}
}

// WithElements returns a matcher that ensures a slice will have the expected
// elements, not considering their position.
func WithElements(expected interface{}) gomock.Matcher {
	return withElementsMatcher{
		expected,
	}
}

func (m withElementsMatcher) Matches(actual interface{}) bool {
	v1 := reflect.ValueOf(actual)

	if !v1.IsValid() {
		return false
	}

	if v1.Kind() != reflect.Slice {
		return false
	}

	v2 := reflect.ValueOf(m.expected)

	if !v2.IsValid() {
		return false
	}

	if v2.Kind() != reflect.Slice {
		return false
	}

	if v1.IsNil() && v2.IsNil() {
		return true
	}

	if v1.IsNil() && !v2.IsNil() {
		return false
	}

	if v2.IsNil() && !v1.IsNil() {
		return false
	}

	if v1.Len() != v2.Len() {
		return false
	}

	for i := 0; i < v1.Len(); i++ {
		match := false

		for j := 0; j < v2.Len(); j++ {
			if reflect.DeepEqual(v1.Index(i).Interface(), v2.Index(j).Interface()) {
				v2 = reflect.AppendSlice(v2.Slice(0, j), v2.Slice(j+1, v2.Len()))
				match = true

				break
			}
		}

		if !match {
			return false
		}
	}

	return true
}

func (m withElementsMatcher) String() string {
	return fmt.Sprintf("is slice with same elements as %v", m.expected)
}

var trimmer = regexp.MustCompile(`\n+\t*`)

type withRawStringMatcher struct {
	expected string
}

// WithRawString returns a matcher that ensures a string will have a certain
// content, trimming any leading new lines and tabs.
func WithRawString(expected string) gomock.Matcher {
	return withRawStringMatcher{
		expected: strings.TrimSpace(trimmer.ReplaceAllString(expected, " ")),
	}
}

func (m withRawStringMatcher) Matches(actual interface{}) bool {
	v1 := reflect.ValueOf(actual)

	if !v1.IsValid() || v1.Kind() != reflect.String {
		return false
	}

	return m.expected == strings.TrimSpace(trimmer.ReplaceAllString(actual.(string), " "))
}

func (m withRawStringMatcher) String() string {
	return fmt.Sprintf("is string with content %s", m.expected)
}
