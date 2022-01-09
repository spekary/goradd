# Windows version of test.sh
# This script is an aid to running the browser test on travis. It is designed to be run from the goradd-test directory.
echo "*** building main"
go build goradd-test
echo "*** starting chrome"
start /b test-helper.cmd
echo "*** starting server"
start /wait goradd-test
