#!/bin/bash
gnome-terminal \
  --geometry=110x40 \
  --window \
  --title="GoSnake" --profile="Monokai" --command="bash -c 'go run .;'"
