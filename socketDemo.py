import socket


s = socket.socket()
host = socket.getHostname()
port = 12345
print(host)
s.bind(host,port)
