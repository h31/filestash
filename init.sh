#!/bin/sh

if [ ! -d "$SNAP_DATA/data" ]
then
echo "Copying default config..."
cp -r $SNAP/bin/data/ $SNAP_DATA/
fi
