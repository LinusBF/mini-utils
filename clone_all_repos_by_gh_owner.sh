#!/bin/bash
set -e

REPO_OWNERS=("Owner1" "Owner2" "Owner3")

mkdir -p dump
cd dump

for owner in ${REPO_OWNERS[@]}; do
    echo "Starting to clone all repos for owner $owner"
    mkdir $owner
    cd $owner

    for repo in $(gh repo list $owner -L 1000 -q '.[] | .sshUrl' --json sshUrl | xargs -L1 echo); do
        git clone $repo
    done

    echo "Cloned all repos for $owner"
    cd ..
done
