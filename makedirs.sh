for((i=20;i>=1;i--))
do
	mkdir -p $i
	let ns=10*$i
	echo ""
	tail -n $ns makefiles.sh | head -n 10|awk -v dirname=$i 'BEGIN{filename=$1;}{cmd="mv "$filename" ./"dirname"/";system(cmd)}'
#	tail -n $ns makefiles.sh | head -n 10|awk -v dirname=$i '{print dirname}'
done

