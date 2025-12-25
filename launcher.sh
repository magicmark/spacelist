#!/bin/bash
WINDOW_ID=$(uuidgen)
open -na ghostty.app --args -e bash -c "\
  aerospace layout floating && \
  ~/apps/spacelist/sl --id=\"$WINDOW_ID\" \
"

success=false
for i in {1..5}; do
  output=$(./center.swift "spacelist-$WINDOW_ID" 2>&1)
  if [ $? -eq 0 ]; then
    success=true
    break
  fi
  sleep 0.01
done

if [ "$success" = false ]; then
  echo "$output"
fi
