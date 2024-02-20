#!/bin/python3
import mysql.connector

connect = mysql.connector.connect(
    host='localhost',
    user='root',
    password='',
    database='banco'
)

cursor=connect.cursor()

#CREATE
chave = "it works!"
valor = 10
comand = f'INSERT INTO tabela (chave,valor) VALUES ("{chave}",{valor})'
cursor.execute(comand)
connect.commit()

#READ
comand = 'SELECT * FROM tabela'
cursor.execute(comand)
result = cursor.fetchall()
print(result)

#UPDATE
comand = 'UPDATE tabela SET chave = "this was changed by the update " WHERE valor = 10'
cursor.execute(comand)
connect.commit()

#DELETE
new_value = 99
comand = f'DELETE FROM tabela WHERE valor = {new_value}'
cursor.execute(comand)
connect.commit()

cursor.close()
connect.close()