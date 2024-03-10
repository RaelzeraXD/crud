#!/bin/python3
import mysql.connector 

mydb = mysql.connector.connect(
    host='localhost',
    user='root',
    password='',
    database='pydb'
)

mycursor=mydb.cursor()

def create(person_name, person_age):
    command = 'INSERT INTO pytable VALUES (%s, %s)'
    mycursor.execute(command, (person_name, person_age))
    mydb.commit()
    print("successfully created")

def read():
    command = 'SELECT * FROM pytable'
    mycursor.execute(command)
    result = mycursor.fetchall()
    print(result)

def update(person_name,new_age,):
    command = 'UPDATE pytable SET age = %s WHERE name = %s'
    mycursor.execute(command, (new_age, person_name))
    mydb.commit()
    print("sucessfully updated")

def delete(person_name, person_age):
    command = 'DELETE FROM pytable WHERE name = %s AND age = %s'
    mycursor.execute(command, (person_name, person_age))
    mydb.commit()
    print("sucessfully deleted")


create("joao",10)
read()

mycursor.close()
mydb.close()