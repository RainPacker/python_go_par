from pymysql import *

conn = connect(host='127.0.0.1', user="root", password="root", port=3306, database="wizarpos_international", charset="utf8")

conn.cursor()

try:
    cur = conn.cursor()
    sql = "select * from sys_m_user"
    cur.execute(sql)
    result = cur.fetchall()
    for item in result:
        print(item)
    print(result)
except Exception as result:
    print(result)
finally:
    cur.close()
    conn.close()
