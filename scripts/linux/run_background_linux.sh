# Executable file authorization, and run in the background.
cd ../../cmd

if [ -x GoWebPractic-linux ]
  then
    echo "start server by script"
  else
    chmod a+x GoWebPractic-linux
fi

nohup ./GoWebPractic-linux > server.out &


