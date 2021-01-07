#!/bin/bash

if [[ -z $INPUT_ACCESS_TOKEN ]]; then
	echo "No access token specified. Could not continue."
	exit 1
fi
