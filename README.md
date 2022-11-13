# Simple calculator
## Describe
Project based on grpc framework and protobuf.
## Installation
Clone

    `git clone https://github.com/r3isenfe1d/calculator.git`

## Usage
Run in shell:

`make server p=[port] op=[option]` - start server as a background process to listen port and do operation;

`make client ip=[address] a=[first value] b=[second value]` - send request from client ot server;

`make stop` - close server;
## Logs
Server logs contain in `.log/server.log` file.
