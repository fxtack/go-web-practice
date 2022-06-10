# This scripts to stop the running server
d=`ps -ef | grep GoWebPractic-linux | grep -v "grep" | awk '{print $2}'`
if [ -n "$d" ]
  then
    kill -9 $d
    echo "server stop"
  else
    echo "server not started"
fi


