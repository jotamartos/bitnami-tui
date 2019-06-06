#!/bin/bash

EXIT=0
printf "######## Initiating tests for the Bitnami TUI ########\n\n"

for test in exit.sh simpleCommand.sh submenu.sh arguments.sh; do
    printf "### Testing ${test}\n\n"
    ./$test $1
    EXIT_TMP=$?
    echo "Exit code: $EXIT_TMP"
    if [ $EXIT_TMP -ne 0 ]; then EXIT=1; fi
done

exit $EXIT
