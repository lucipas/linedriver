#!/bin/bash
trap "linedriver -ctx $ctx close" INT # clean up browser when <Ctrl-C> is encountered

geckodriver & 2> /dev/null

ctx=$(linedriver) # init webdriver and hold the browser context

linedriver -ctx $ctx navigate "file:///usr/share/doc/" # open UR:
sleep 10s # wait for the network to settle

linedriver -ctx $ctx url   # get URL
linedriver -ctx $ctx title # get title
linedriver -ctx $ctx src   # get src
linedriver -ctx $ctx close

killall geckodriver
