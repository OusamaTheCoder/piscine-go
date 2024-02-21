#!/bin/bash

my_data=$(curl -s https://learn.zone01oujda.ma/assets/superhero/all.json)

result=$(echo "${my_data}" | jq --arg id "$HERO_ID" '.[] | select(.id == ($id | tonumber))' | jq '.connections.relatives')

f_result=$(echo "${result}" | sed 's/"//g')

echo "${f_result}"