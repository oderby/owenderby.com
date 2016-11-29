#!/usr/bin/env bash

cd docs
shopt -s extglob
rm -rf -- !(CNAME)
shopt -u extglob
cd ..

hugo -d docs
