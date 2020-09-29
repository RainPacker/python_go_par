import sqlite3


def main():
    conn = sqlite3.connect("test.db")
    cur = conn.cursor()

    sql = " select *  from user "
    row = cur.execute(sql)
    for item in row:
        print(item)

    conn.commit()
    cur.close()
    conn.close()


if __name__ == '__main__':
    main()
