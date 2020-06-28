# rest-controll-interface
a module for controlling a rasbery PI using HTTP REST
Until I get bluetooth working, a way to comunicate over the local network instructions from a client to a server with some servo motors attached

## Example usage

Call with the horizontal and vertical coordinates you want it to turn to
`curl -X POST localhost:8080/rotate -d '{"Horizontal":120, "Vertical": 110}`

