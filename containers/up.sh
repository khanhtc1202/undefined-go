#!/usr/bin/env bash

docker run -v `pwd`:/app -it --rm golang:1.12 /bin/bash
