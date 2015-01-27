# rankings
An attempt at ranking 2014 KCPE candidates using 2012 data

## installation

For the scraping

Install [golang](http://golang.org/)

Download the project

`go get -u github.com/librekcpe/rankings`

##dev run

`cd $GOPATH/src/github.com/librekcpe/rankings`

`go run main.go`

You should start seeing results on the console, you might want to set debug as true in the code



## production run

Compile to run

 `go install github.com/librekcpe/rankings`

 `rankings > /path/to/results.csv`

##sample result

The results should be a csv that looks like:

index,total,name,eng,kis,mat,sci,ssr,schoolName,gender,engGrade,kisGrade,matGrade,sciGrade,ssrGrade
01101304001,372,DEBORAH   MASUBO,83,79,72,68,70,CHOKE,F,A,A-,B+,B,B+

.
.
.

01101306001,269,MWASARU RACHEL SARU,54,66,51,53,45,WUMINGU,F,C,B,C,C,C
01101302001,309,MWAWUGHANGA WAKIO CAROLINE,71,67,44,58,69,KITUMBI,F,B+,B,C-,C+,B





