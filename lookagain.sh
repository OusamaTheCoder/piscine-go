#!/bin/bash
find . -type f -name "*.sh" | sort -r | sed 's/\.\///;s/\.sh$//;s/.*\///' s