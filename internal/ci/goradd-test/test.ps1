# Windows version of test.sh
# This script is an aid to running the browser test as a github action. It is designed to be run from the goradd-test directory.
echo "*** building main"
go build goradd-test
echo "*** starting server"
Start-Job -Name "testserver" -ScriptBlock { ./goradd-test.exe }
Start-Sleep -Seconds 5
Start-Process "Chrome" "http://localhost:8000/goradd/Test.g?all=1 --headless --remote-debugging-port=9222"
Wait-Job -Name "testserver"
Receive-Job -Name "testserver"