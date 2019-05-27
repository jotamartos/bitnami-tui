#!/usr/bin/env expect

cd tests
spawn go run ../app.go

expect {
    timeout { puts "timed out when opening the application"; exit 1 }
    "Manage the services"
}
send -- ""
expect eof
