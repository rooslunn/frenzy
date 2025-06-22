#!/bin/bash

mask="*.srt"
counter=1

for file in $mask; do
  echo "Processing file $counter: $file"
  # Your renaming logic here, potentially using $counter in the new name
  new_name="${counter}.png"
  echo "Renaming '$file' to '$new_name'"
  mv "$file" "$new_name" # Uncomment to actually rename
  ((counter++))
done
