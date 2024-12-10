#!/bin/bash

# Builtin to easily create new cmdr scripts

if [ -z "$1" ]; then 
  echo "No argument passed. Please provide a name for the script you'd like to create."
  exit 1
fi

if find "$CMDR_SCRIPTS_DIR/$1" > /dev/null 2>&1; then
  echo "A script named $1 already exists"
  exit 1
fi 

touch "$CMDR_SCRIPTS_DIR/$1"

if [ -z "$EDITOR" ]; then
  echo "No editor configured."
  exit 1  
fi 

"$EDITOR" "$CMDR_SCRIPTS_DIR/$1"
