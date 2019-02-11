. fd
echo 'tail -f out1.txt | grep ^[^2]'
g factomd --prefix="x" --count=1 --networkport="34341" --port="8092" --controlpanelport="9081" --logPort="6062" --db=Map --peers="127.0.0.1:37003" --network=LOCAL -blktime=30 -net=alot+ -startdelay=1 > out1.txt
