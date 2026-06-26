#!/bin/bash

trap "linedriver -ctx $ctx close" INT # clean up browser when <Ctrl-C> is encountered

geckodriver > /dev/null 2>&1 &
killall firefox > /dev/null 2>&1

ctx=$(linedriver) # init webdriver and hold the browser context

linedriver -ctx "$ctx" navigate "https://www.github.careers/careers-home/jobs?page=1&locations=,,United%20States&categories=Engineering&tags6=Yes" # open UR:

sleep 5s # wait for the network to settle

linedriver -ctx "$ctx" url   # get URL
linedriver -ctx "$ctx" title # get title
for i in {0..9}; do 
	printf "<!DOCTYPE html>\n%s" "$(linedriver -ctx "$ctx" src | jq -r .value | perl -MHTML::Entities -MEncode -pe 'BEGIN { binmode(STDIN, ":utf8"); binmode(STDOUT, ":utf8"); } s/\\n//g; s/\\//g; s/&#39;/"/g; decode_entities($_);')" | xsel -a -t html -x "//*[@id='mat-expansion-panel-header-$i']/span[1]/mat-panel-title/p[2]" - | cut -d " " -f 4
done
# s/\\//g
linedriver -ctx $ctx close

killall geckodriver
# firefox https://gemini.google.com > /dev/null 2>&1 &
