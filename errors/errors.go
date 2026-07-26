package errors

import "errors"

var ErrorNF = errors.New("target element not found in container")
var ErrorEmpty = errors.New("container is empty")
var ErrorOverflow = errors.New("container is too large to add new element")
