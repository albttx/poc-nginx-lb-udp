#!/bin/bash

ping -c 1 $1 >2

if [ "$?" == "0" ]; then
    echo -n "success"
else
    echo -n "fail"
fi
