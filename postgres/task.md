# Creating a new user

<b> First go to the main user </b>
```
sudo -u postgres psql
```
<b> We will enter postgres shell for user postgres </b>

<b> Now we will create a new user </b>
```
create user amarjit with password amarjit;
```

# Creating a new database 

<b> Now we will create a database with user - amarjit as the ownwer of the database </b>
```
create database amarjitdb ownwer amarjit;
```

# Creating a table employees with columns Name, Email_ID, Phone_No
```
sudo su amarjit
```
```
psql -U amarjit -d amarjitdb
```
```
create table employees ( Name varchar(100), Email_ID varchar(100), Phone_No varchar(20));
```
![Screenshot 2023-07-20 at 9 23 04 AM](https://github.com/Amarjit0511/go-task-sheet/assets/54772122/d8fa3f96-334b-4b61-98e9-3e37582b6673)
