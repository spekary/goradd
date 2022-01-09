#starts chrome  in headless mode at the test web page, after waiting for the server to start
timeout /nobreak 10
start /b chrome --headless --remote-debugging-port=9222 http://localhost:8000/goradd/Test.g?all=1
