#!/usr/bin/env expect

set timeout 10

spawn go run ../app.go

expect {
    timeout { puts "timed out when opening the application"; exit 1 }
    "Manage the services"
}
send -- "\[B"
send -- "\[B"

send -- "\r"
expect {
    timeout { puts "timed out when moving to a submenu"; exit 1 }
    "Get status of the services"
}

send -- "\r"
expect {
    timeout { puts "timed out when getting the status of the services"; exit 1 }
    "This is"
}

send -- ""
expect {
    timeout { puts "timed out when returning to the submenu"; exit 1 }
    "Start"
}

send -- "\[B"

send -- "\r"
expect {
    timeout { puts "timed out when starting the services"; exit 1 }
    "Information"
}

send -- ""
expect {
    timeout { puts "timed out when returning to the submenu (2)"; exit 1 }
    "Stop the services"
}

send -- ""
expect {
    timeout { puts "timed out when moving to the main menu"; exit 1 }
    "Manage the services"
}

send -- ""
expect eof
