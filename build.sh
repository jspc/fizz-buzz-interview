#!/usr/bin/env bash

GOOS=linux go build -v -o linux/fizz-buzz-interview
strip linux/fizz-buzz-interview
