import socket

# 使用 socket
s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
host = socket.gethostname()
port = 10086
s.connect((host, port))
print(s.recv(1024))
s.close()
