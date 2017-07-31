package soupchain

import "github.com/anaskhan96/soup"

type errchain struct {
	err error
}

func (c *errchain) Attrs() (map[string]string, error) {
	return nil, c.err
}

func (c *errchain) Find(s ...string) Chainable {
	return c
}

func (c *errchain) Text() (string, error) {
	return "", c.err
}

func (c *errchain) FindAll(...string) ([]Chainable, error) {
	return nil, c.err
}

func (c *errchain) Result() (soup.Root, error) {
	return soup.Root{}, c.err
}
