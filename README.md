# Crud
a simple crud with mysql in go and python
## Python
first we need to inicialize the database so in the cli follow this steps to create a simple database
* sudo mysql -e "CREATE DATABASE banco; USE banco; CREATE TABLE tabela (id INT AUTO_INCREMENT PRIMARY KEY, chave VARCHAR(255), valor INT);"
* exit
### how to use?
* git clone https://github.com/RaelzeraXD/crud
* cd crud
* ./crud.py
### problems with mysql connection? try this!
1) This will make root password accessible
* sudo mysql -e "USE mysql; UPDATE user SET plugin='mysql_native_password' WHERE User='root';"
* exit
* sudo systemctl restart mysql.service
2) Case you dont know the password, reset with this (chane the field to your password and edit the file crud.py to it)
* mysql -uroot -e "ALTER USER 'root'@'localhost' IDENTIFIED BY 'PUT YOUR NEW PASSWORD HERE'; FLUSH PRIVILEGES;"
* sudo systemctl restart mysql.service
## Go
