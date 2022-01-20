go run wordlehack.go | tee score
awk '{ print $2 " " $1 }' scores|sort -n|head

python3 sub.py raise < La.json > La_no_raise.json
go run wordlehack.go La_no_raise.json |tee scores_no_raise
awk '{ print $2 " " $1 }' scores_no_raise |sort -n|head
