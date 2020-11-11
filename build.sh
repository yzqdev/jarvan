#!/usr/bin/env bash
if [ $# != 1 ] ; then
echo "USAGE: sh build.sh dev test or release"
exit 1;
fi

CHANNEL=$1;
# echo $CHANNEL;

if [ $CHANNEL = "dev" ]; then
    echo "=====copy channel/app.ini.dev to conf/app.ini====";
    cp ./channel/app.ini.dev ./conf/app.ini
fi;

if [ $CHANNEL = "test" ]; then
    echo "=====copy channel/app.ini.test to conf/app.ini====";
    cp ./channel/app.ini.test ./conf/app.ini
fi;

if [ $CHANNEL = "release" ]; then
    echo "=====copy channel/app.ini.release to conf/app.ini====";
    cp ./channel/app.ini.release ./conf/app.ini
fi;