#!/bin/sh
col=$((`cat COL.txt`))
test_process=$(ps aux | grep "[a]nchordav-linux-arm")
if [ -n "$test_process" ]
then
    result="File Anchor is Running.....  "
else
    result="File Anchor is not running   "
fi
eips $col 3 "$result"