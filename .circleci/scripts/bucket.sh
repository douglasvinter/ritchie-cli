#!/bin/sh

if expr "$VERSION" : '^[0-9]\+.0.0-qa' >/dev/null; then
  echo "ritchie-13528094685555"
elif expr "$VERSION" : '^[0-9]\+' >/dev/null; then
  echo "ritchie-cli-bucket152849730126474"
elif expr "$VERSION" : '.*\.legacy\..*' >/dev/null; then
  echo "ritchie-cli-bucket152849730126474"
elif expr "$VERSION" : '^[0-9]\+' >/dev/null; then
  echo "ritchie-cli-bucket152849730126474"
else
  echo ""
fi