#!/usr/bin/env expect

set timeout 10

if { [llength $argv] == 0 } {
    spawn go run ../app.go
} else {
    spawn [lindex $argv 0]
}

expect {
    timeout { puts "timed out when opening the application"; exit 1 }
    "Set up Let's Encrypt"
}
send -- "\[B"

send -- "\r"
expect {
    timeout { puts "timed out when waiting for the email information"; exit 1 }
    "email"
}

send -- "\r"
expect -exact "\[42;159H\[42;159H"

send -- "a\r"
expect {
    timeout { puts "timed out when waiting for the domain information"; exit 1 }
    "domain"
}

send -- "\r"
expect -exact "\[42;159H\[42;159H"

send -- "b\r"
expect -exact "\[42;159H\[42;159H"

send -- ""
expect {
    timeout { puts "timed out when moving to the main menu"; exit 1 }
    "Set up Let's Encrypt"
}

send -- ""
expect eof
