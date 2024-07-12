package postgres

import "errors"

var (
	errNoSubs = errors.New("there is no subscribers in database")
)
