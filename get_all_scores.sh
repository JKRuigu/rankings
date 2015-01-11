for school in $(cat $1);
do
  echo "$school";
  phantomjs ex.js $school >> results.txt
done
