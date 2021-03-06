#!/bin/bash

# Copyright (C) 2016 Yoshiki Shibata. All rights reserved.

go vet
if [ $? != 0 ]
then
	exit 1
fi

go build -o surface
./surface -type=0 > surface0.svg
./surface -type=1 > surface1.svg
./surface -type=2 > surface2.svg
./surface -type=3 > surface3.svg

open surface0.svg
open surface1.svg
open surface2.svg
open surface3.svg

sleep 5
rm surface
rm surface0.svg surface1.svg surface2.svg surface3.svg
