# regen [![Build Status](https://travis-ci.org/zach-klippenstein/regen.svg?branch=master)](https://travis-ci.org/zach-klippenstein/regen)

Simple command-line utility for generating strings from regular expression patterns.
Uses [goregen](https://github.com/zach-klippenstein/goregen).

## Installation
```bash
go get github.com/zach-klippenstein/regen
```

## Usage

Regen takes a regular expression as an argument, and prints the generated strings to standard out.

```bash
regen [args] pattern
```
where `pattern` is a regular expression. Run with `-h` for more information on arguments.

For example,
```bash
regen '[a-f0-9]{8}(-[a-f0-9]{4}){3}-[a-f0-9]{12}'
```
will generate a UUID (e.g. `bc8c3f7d-1176-ae94-492a-2dab8fc06640`).
