import mysql.connector as m
import pandas as pd

def populate_users(df, mysql_db) :
    cursor = mysql_db.cursor()

    sql_cmd = "INSERT INTO Users(user_id, username, hashed_password, email, first_name, last_name, party_id) VALUES(%s, %s, %s, %s, %s, %s, %s)"

    for _, row in df.iterrows() :
        cursor.execute(sql_cmd, row.values.tolist())

    mysql_db.commit()

def populate_parties(df, mysql_db) :
    cursor = mysql_db.cursor()

    sql_cmd = "INSERT INTO Party(party_id, party) VALUES(%s, %s)"

    for _, row in df.iterrows() :
        cursor.execute(sql_cmd, row.values.tolist())

    mysql_db.commit()

def main() :
    user_file  = "data/user_data.csv"
    party_file = "data/party_data.csv"
    df_users   = pd.read_csv(user_file)
    df_party   = pd.read_csv(party_file)

    mysql_db = m.connect(
        host="mysql_db",
        user="root",
        password="secret",
        database="voting"
    )

    #populate_users(df, mysql_db)
    populate_parties(df_party, mysql_db)

main()
