#!/usr/bin/env expect

set timeout 10

spawn go run ../app.go

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
expect -exact "\[?25l\[?25l"

send -- "a\r"
expect {
    timeout { puts "timed out when waiting for the domain information"; exit 1 }
    "domain"
}

send -- "\r"
expect -exact "\[?25l\[?25l"

send -- "b\r"
expect -exact "\[?25l\[?25l"

send -- ""
expect {
    timeout { puts "timed out when moving to the main menu"; exit 1 }
    "Set up Let's Encrypt"
}

send -- ""
expect eof
