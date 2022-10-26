NUMBER=$(grep "CLUE" crimescene | grep Annabel people | sed '179!d' streets/Buckingham_Place | sed 's/^.*#//')
echo $NUMBER
cat interviews/interview-"$NUMBER"
echo $MAIN_SUSPECT