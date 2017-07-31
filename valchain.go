package soupchain

import (
	"errors"
	"fmt"
	"strings"

	"github.com/anaskhan96/soup"
)

type Chainable interface {
	Find(...string) Chainable
	FindAll(...string) ([]Chainable, error)
	// FindNextSibling() struct{}
	// FindNextElementSibling() struct{}
	// FindPrevSibling() struct{}
	// FindPrevElementSibling() struct{}
	Attrs() (map[string]string, error)
	Text() (string, error)
	Result() (soup.Root, error)
}

func New(r soup.Root) Chainable {
	return &valchain{
		value: r,
	}
}

type valchain struct {
	value soup.Root
}

func (c *valchain) Find(s ...string) Chainable {
	val := c.value.Find(s...)
	if val.Pointer == nil {
		return &errchain{
			err: errors.New(fmt.Sprintf("error: could not find elms: %s", strings.Join(s, " "))),
		}
	}

	return &valchain{value: val}
}

func (c *valchain) Attrs() (map[string]string, error) {
	attrs := c.value.Attrs()
	if attrs == nil {
		attrs = map[string]string{}
	}

	return attrs, nil
}

func (c *valchain) Text() (string, error) {
	return c.value.Text(), nil
}

func (c *valchain) FindAll(s ...string) ([]Chainable, error) {
	vals := c.value.FindAll(s...)
	chainables := make([]Chainable, len(vals))
	for i, val := range vals {
		chainables[i] = &valchain{value: val}
	}

	return chainables, nil
}

func (c *valchain) Result() (soup.Root, error) {
	return c.value, nil
}
