#!/bin/bash

# COmpare average_color results with results of imagemagick's convert utility on images in a given directory.

if [ $# -eq 0 ]
then
  echo "Usage: $0 <directory with images>"
  exit 1
fi

MATCHED=0
MISMATCHED=0

for FILE in $(find "$1" -iname '*.png' -o -iname '*.jpg' -o -iname '*.jpeg' -o -iname '*.gif')
do
  RES1=$(average_color -rgb "$FILE" | tail -1 | sed -E 's/rgba?\(([^)]+)\)/\1/')
  # rgb to rgba
  if [ $(awk -F, '{print NF-1}' <<<"$RES1") -eq 2 ]
  then
    RES1="$RES1,255"
  fi
  RES2=$(convert "$FILE" -scale 1x1 -format "%[fx:round(255*r)],%[fx:round(255*g)],%[fx:round(255*b)],%[fx:round(255*a)]" info:)
  if [ "$RES1" != "$RES2" ]
  then
    echo "Mismatch: $FILE: $RES1 != $RES2"
    (( MISMATCHED++ ))
  else
    (( MATCHED++ ))
  fi
done

echo "$MATCHED matched, $MISMATCHED mismatched"
