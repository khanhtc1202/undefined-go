#!/usr/bin/env bash

docker run -v `pwd`:/app -it --privileged --rm golang:1.12-alpine /bin/sh
