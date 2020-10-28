import mysql.connector as m
import pandas as pd

def populate_users(df, mysql_db) :
    cursor = mysql_db.cursor()

    sql_cmd = "INSERT INTO Users(user_id, username, hashed_password, email, first_name, last_name, party_id) VALUES(%s, %s, %s, %s, %s, %s, %s)"

    for _, row in df.iterrows() :
        cursor.execute(sql_cmd, row.values.tolist())

    mysql_db.commit()


def main() :
    csv_file = "data/user_data.csv"
    df       = pd.read_csv(csv_file)

    mysql_db = m.connect(
        host="localhost",
        user="root",
        password="secret",
        database="voting"
    )

    populate_users(df, mysql_db)

main()
