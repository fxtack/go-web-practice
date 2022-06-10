# Executable file authorization, and run.
# You need to note that the binary executable file is used in this run. If there is no such file in the root directory of the project, it will not run
cd ../../cmd

if [ -x GoWebPractic-linux ]
  then
    echo "start server by script"
  else
    chmod a+x ./GoWebPractic-linux
fi

./GoWebPractic-linux
