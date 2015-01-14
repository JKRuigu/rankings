#!/usr/bin/env bash

## generate school ids
#echo "Generating school index numbers"
#seq -w 10000000 99999999 > schools.txt

## split into smaller files for parallelization
#echo "splitting up index numbers"
#split -d -l 10000000 schools.txt split-

## run in parallel
## max of 200 concurrent jobs
#echo "Running jobs in parallel"
#ls split-* | parallel --gnu -j200 'sh get_all_scores.sh {}'

# removed middle men
seq -w 10000000 99985821 |sort -R  | parallel --gnu -j150 'phantomjs ex.js {} >> results.txt'
