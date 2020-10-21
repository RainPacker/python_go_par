import _thread
import time


def print_time(thread_name, delay):
    count = 0
    while count < 0:
        time.sleep(delay)
        count += 1
        print(count)
        print("%s:%s" % (thread_name, time.cntime(time.time)))

try:
    _thread.start_new_thread(print_time, ("Thread1",2,))
    _thread.start_new_thread(print_time, ("Thread2",4,))
except:
    print("无法启动线程。。。")
while 1:
    pass        