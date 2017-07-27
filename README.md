# Golang Project Sample - Fibonacci Calculator

[![GoDoc](https://godoc.org/github.com/sevlyar/fibonacci?status.svg)](https://godoc.org/github.com/sevlyar/fibonacci) [![Build Status](https://travis-ci.org/sevlyar/fibonacci.svg?branch=master)](https://travis-ci.org/sevlyar/fibonacci) [![codecov](https://codecov.io/gh/sevlyar/fibonacci/branch/master/graph/badge.svg)](https://codecov.io/gh/sevlyar/fibonacci) [![Go Report Card](https://goreportcard.com/badge/github.com/sevlyar/fibonacci)](https://goreportcard.com/report/github.com/sevlyar/fibonacci) [![Sourcegraph](https://sourcegraph.com/github.com/sevlyar/fibonacci/-/badge.svg)](https://sourcegraph.com/github.com/sevlyar/fibonacci?badge)

*The project is not ready for production. It is still experimental and subject to change.*

Fibonacci Calculator calculates n-th Fibonacci number. It computes Fibonacci numbers in time O(log(n))
and uses 'Rising a matrix to the power' method.

Algorithm is implemented in library but you also can use utility 'fib'
to calculate Fibonacci numbers using command-line.

## Key Features

* Fast computation algorithm in time O(log(n));
* Almost constant memory usage.

## Documentation

* [Library](https://godoc.org/github.com/sevlyar/fibonacci);
* [Utility](https://godoc.org/github.com/sevlyar/fibonacci/cmd/fib).

## Installation

You can download latest release of the calculator for Darwin_amd64 from [here](https://github.com/sevlyar/fibonacci/releases/latest).

You need [Go distribution](https://golang.org/doc/install) to install project from sources:

	go get github.com/sevlyar/fibonacci
	go install github.com/sevlyar/fibonacci/...

