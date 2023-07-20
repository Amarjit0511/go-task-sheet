# Installing PostgresSQL Locally 
[Download link for MacOS and Windows](https://www.enterprisedb.com/downloads/postgres-postgresql-downloads)
[For Linux](https://www.postgresql.org/download/linux/ubuntu/)

<b> For Mac and Windows, all the packages for pg admin and postgres shell comes in a bundle </b>

### Downloaidng pgAdmin for Linux:
<b> Follow the instructions given at </b>
[pgAdmin](https://www.pgadmin.org/download/pgadmin-4-apt/)

#### Step 1:
```
curl -fsS https://www.pgadmin.org/static/packages_pgadmin_org.pub | sudo gpg --dearmor -o /usr/share/keyrings/packages-pgadmin-org.gpg
```
<b>Here we will be prompted for our system's password</b>

#### Step 2:
```
sudo sh -c 'echo "deb [signed-by=/usr/share/keyrings/packages-pgadmin-org.gpg] https://ftp.postgresql.org/pub/pgadmin/pgadmin4/apt/$(lsb_release -cs) pgadmin4 main" > /etc/apt/sources.list.d/pgadmin4.list && apt update'
```

#### Step 3:
```
sudo apt install pgadmin4
```

#### Step 4:
```
sudo apt install pgadmin4-desktop
```

#### Step 5:
```
sudo apt install pgadmin4-web
```
#### Step 6:
```
sudo /usr/pgadmin4/bin/setup-web.sh
```
<b> Here it will ask for email id and a new password </b>
<b> After this pgadmin web can be started at [http://127.0.0.1/pgadmin4](http://127.0.0.1/pgadmin4) with the above provided credentials</b>

<b>Alternatively to start postgres</b>
```
sudo su postgres
```
<b> Here we will be prompted a terminal like below:</b>
```
postgres=#
```
<b> Now we can perform various operations like creating a new user, creating database, etc</b>


<b>After the application is installed</b>
<b>1. Click on Add Server</b>
<b>2. Add the required confiuration - user_name, password, port, etc and connect</b>

