#!/bin/bash

my_data=$(curl -s https://learn.zone01oujda.ma/assets/superhero/all.json)

result=$(echo "${my_data}" | jq '.[] | select(.id == 70) .name')

echo "${result}"