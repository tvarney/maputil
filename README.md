# maputil
Utility functions for manipulating JSON style maps in Go

[![Godoc](https://godoc.org/github.com/tvarney/maputil?status.svg)](https://godoc.org/github.com/tvarney/maputil)
[![Go Report Card](https://goreportcard.com/badge/github.com/tvarney/maputil)](https://goreportcard.com/report/github.com/tvarney/maputil)

## Overview

Map manipulation in go can be significantly verbose; normally this isn't a big
deal, but sometimes code needs to load JSON or YAML files into a map and
perform a variety of operations on it. This package provides convenience
functions for doing so.
