#!/bin/bash
o=$1
s=$2


echo "Building EdgeAgent..."
if [ $o = "windows" ];
  then
  if [ $s = "32" ];
    then
      export GOARCH=386
  elif [ $s = "64" ];
    then
      export GOARCH=amd64
  fi
elif [ $o = "linux" ];then
  if [ $s = "32" ];
    then
      export GOARCH=386
  elif [ $s = "64" ];
    then
      export GOARCH=amd64
  fi
fi

export GOOS=$o

echo $GOOS
echo $GOARCH
#go env

if [ $o = "windows" ];then
    go build -o EdgeAgent_$o$s.exe
elif [ $o = "linux" ];then
    go build -o EdgeAgent_$o$s
fi

