#!/usr/bin/env expect

if { [llength $argv] == 0 } {
    spawn go run ../app.go
} else {
    spawn [lindex $argv 0]
}

expect {
    timeout { puts "timed out when opening the application"; exit 1 }
    "Manage the services"
}
send -- ""
expect eof
