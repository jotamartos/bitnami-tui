#!/usr/bin/env expect

set timeout 10

cd tests
spawn go run ../app.go

expect {
    timeout { puts "timed out when opening the application"; exit 1 }
    "Manage the services"
}
send -- "\[B"
send -- "\[A"

send -- "\r"
expect {
    timeout { puts "timed out when running command"; exit 1 }
    "Success"
}

send -- ""
expect {
    timeout { puts "timed out when moving to the main menu"; exit 1 }
    "Manage the services"
}

send -- ""
expect eof
