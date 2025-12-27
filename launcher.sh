#!/bin/bash
# https://stackoverflow.com/a/246128/4396258
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
WINDOW_ID=$(uuidgen)

open -na ghostty.app --args -e bash -c "\
  aerospace layout floating && \
  spacelist --id=\"$WINDOW_ID\" \
"

success=false
for i in {1..5}; do
  output=$("${DIR}/center.swift" "spacelist-$WINDOW_ID" 2>&1)
  if [ $? -eq 0 ]; then
    success=true
    break
  fi
  sleep 0.01
done

if [ "$success" = false ]; then
  echo "$output"
fi
