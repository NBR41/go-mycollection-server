#!/bin/sh


# the for loop is in case the database container is not yet
# ready, we give it some times before retrying
while [ 1 ]; do
    echo "exit" | nc mysql 3306 1>/dev/null 2>&1
    if [ $? != 0 ]; then
        echo "no database"
        sleep 5
        continue
    fi

    /go/bin/go-myinventory-server
done
