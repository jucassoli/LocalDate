# localdate for Go

[![Go Reference](https://pkg.go.dev/badge/golang.org/x/reference.svg)](https://github.com/jucassoli/localdate)
[![Build Status](https://travis-ci.org/golang/localdate.svg?branch=master)](https://travis-ci.org/jucassoli/localdate)

localdate package contains a date and time helper for Go language

## Installation

~~~~
go get github.com/jucassoli/localdate
~~~~

Or you can manually git clone the repository to
`$(go env GOPATH)/src/github.com/jucassoli/localdate`.

See pkg.go.dev for further documentation and examples.

* [pkg.go.dev/jucassoli/x/oauth2](https://pkg.go.dev/golang.org/x/localdate)
* [pkg.go.dev/jucassoli/x/oauth2/google](https://pkg.go.dev/golang.org/x/localdate/google)

## Policy for new packages

We no longer accept new provider-specific packages in this repo if all
they do is add a single endpoint variable. If you just want to add a
single endpoint, add it to the
[pkg.go.dev/jucassoli/x/oauth2/endpoints](https://pkg.go.dev/golang.org/x/oauth2/endpoints)
package.

## Report Issues / Send Patches

This repository uses Gerrit for code changes. To learn how to submit changes to
this repository, see https://golang.org/doc/contribute.html.

The main issue tracker for the oauth2 repository is located at
https://github.com/jucassoli/localdate/issues.