#!/bin/bash

# Builtin to easily edit cmdr scripts

if [ -z "$1" ]; then 
  echo "No argument passed. Please provide the name of the script you'd like to edit."
  exit 1
fi

if ! find "$CMDR_SCRIPTS_DIR/$1" > /dev/null 2>&1; then
  echo "No script named $1 exists"
  exit 1
fi 

if [ -z "$EDITOR" ]; then
  echo "No editor configured."
  exit 1  
fi 

"$EDITOR" "$CMDR_SCRIPTS_DIR/$1"
