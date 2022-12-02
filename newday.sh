#!/bin/bash

# A script to create folder structure for every new day
set -a
. .env
set +a

fetch_input() {
  echo "Fetching inputs"
  curl -b "session=$SESSION"  "https://adventofcode.com/2022/day/$1/input" > $2
}

if [[ $# != 1 ]]
then
  echo "Enter day no. as argument"
  exit 1
fi

day=$@
day=${day#0}  # making 01 into 1
dir="$(printf 'day%02d' $day)"  # 1 => day01; 25 => day25

if [[ -d $dir ]]
then
  echo "the folder for day $day already exists"
  exit 1
fi

input="./$dir/input.txt"

cp -r day00 $dir
fetch_input $day $input
cd $dir
