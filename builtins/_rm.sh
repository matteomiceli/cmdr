#!/bin/bash

# Builtin to easily delete cmdr scripts

if [ -z "$1" ]; then 
  echo "No argument passed. Please provide the name of a script you'd like to remove."
  exit 1
fi

if ! find "$CMDR_SCRIPTS_DIR/$1" > /dev/null 2>&1; then
  echo "No script named $1 exists"
  exit 1
fi 

rm "$CMDR_SCRIPTS_DIR/$1"
echo "Removed script $CMDR_SCRIPTS_DIR/$1"
