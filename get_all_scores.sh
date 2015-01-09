for school in $(cat sorted_schools.txt);
do
  echo "$school";
  phantomjs ex.js $school >> results.txt
done
