#!/bin/sh

swag init

if git diff --name-only | grep -q "docs/"; then
  echo "Swagger is NOT up to date"
  return 1
else
  echo "Swagger is up to date"
  return 0
fi