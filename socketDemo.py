import socket


s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
host = socket.gethostname()
port = 10086
print(host)
s.bind((host, port))
s.listen(5)
print("server start....")
while True:
    c, addr = s.accept()
    print(c)
    msg = "'Hello python'"
    c.send(msg.encode("utf-8"))
    print(addr)
    c.close()
