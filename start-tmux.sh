#!/bin/bash

# Ensure the 'printqueue' session exists
if ! tmux has-session -t printqueue 2>/dev/null; then
  # Create a new session named 'printqueue', detached
  cd /workspaces/printqueue
  tmux new-session -d -s printqueue
 fi
