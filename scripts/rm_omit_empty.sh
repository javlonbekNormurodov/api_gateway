#!/bin/bash
CURRENT_DIR=$1
for service_folder in $(find ${CURRENT_DIR}/genproto/* -type d); do
    for x in $(find $service_folder/* -type f); do
        cd $service_folder && ls *.pb.go | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}'
    done
done
