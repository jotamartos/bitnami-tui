#!/usr/bin/env expect

set timeout 10

if { [llength $argv] == 0 } {
    spawn go run ../app.go
} else {
    spawn [lindex $argv 0]
}

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
