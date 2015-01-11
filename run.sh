# generate school ids
echo "Generating school index numbers"
seq -w 1 99999999 > schools.txt

# split into smaller files for parallelization
echo "splitting up index numbers"
splt -d -l 50 schools.txt split-

# run in parallel
# max of 200 concurrent jobs
echo "Running jobs in parallel"
ls split-* | parallel --gnu -j200 'sh get_all_scores.sh {}'
