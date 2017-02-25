#!/usr/bin/env bash

# cd docs
# shopt -s extglob
# rm -rf -- !(CNAME)
# shopt -u extglob
# cd ..

rm -rf docs/*
hugo -d docs
