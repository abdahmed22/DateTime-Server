# DateTime-Server-Abdelrahman-Mahmoud

## Introduction

A server is a computer system capable of delivering web content to end users over the internet via a web browser. In this project we deliver the current date and time to the user using two golang packages net/http and gin. 

## Feature /datetime 

This end point returns a message that has the current date and time as a JSON for simplicity and easier exchange of data between different technologies.

   ```JavaScript
{"message": "2024-07-07 16:08"}
   ```

## Setup

1. Clone the Repository to a directory of your choice.
2. Make sure you have go version 1.22.4 installed on your device.
3. Open the directory in a code editor of your choice.
4. Navigate to cmd/ginserver or cmd/httpserver directory
5. Build the project using:
   ```console
   user@user-VirtualBox:~$ go build main.go
   ```
6. Run the project using:
   ```console
   user@user-VirtualBox:~$ ./main
   ```
7. Open a web browser of your choice.
8. Use http://localhost:8090/datetime for http server or http://localhost:8085/datetime for gin server to perform a request.
9. Shutdown the server using ctrl + c inside the terminal.

## Tests

