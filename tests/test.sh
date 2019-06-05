#!/bin/bash

printf "######## Initiating tests for the Bitnami TUI ########\n\n"

for test in exit.sh simpleCommand.sh submenu.sh arguments.sh; do
    printf "### Testing ${test}\n\n"
    ./$test
    echo "Exit code: $?"
done
