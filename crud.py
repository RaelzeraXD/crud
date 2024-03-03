#!/bin/python3
import mysql.connector

connect = mysql.connector.connect(
    host='localhost',
    user='root',
    password='',
    database='pydb'
)

cursor=connect.cursor()

class User:
    def __init__(self, name, age):
        self.name = name
        self.age = age

def Create(people):
    comand = f'INSERT INTO pytable VALUES ("{people.name}",{people.age})'
    cursor.execute(comand)
    connect.commit()
    print("successfully created")

def Read():
    comand = f'SELECT * FROM pytable'
    cursor.execute(comand)
    result = cursor.fetchall()
    print(result)

def Update(people):
    comand = f'UPDATE pytable SET age = 1000 WHERE name = "{people.name}"'
    cursor.execute(comand)
    connect.commit()
    print("sucessfully updated")

def Delete(people):
    comand = f'DELETE FROM pytable WHERE age = {people.age} and name = "{people.name}"'
    cursor.execute(comand)
    connect.commit()
    print("sucessfully deleted")

user=User("israel",22)
Create(user)
Read()

cursor.close()
connect.close()