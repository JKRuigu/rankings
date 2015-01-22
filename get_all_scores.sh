for school in $(cat $1);
do
  echo "$school"; # so we can know it's running
  phantomjs ex.js $school >> results.txt
done
