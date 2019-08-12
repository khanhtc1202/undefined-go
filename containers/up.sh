#!/usr/bin/env bash

docker run -v `pwd`:/app -it --user root --rm golang:1.12-alpine /bin/sh
