# Implementation of sending email through sendgrid using AWS Lambda and API Gateway











# Setting up RDS and connecting it with Database client locally
1. Go to AWS Console
2. Select RDS from AWS service list
3. Click on <b>"Create database"</b>
4. Select Standard create
5. In engine option select <b>"PostgresSQL"</b>
6. In Templates select <b>"Free tier"</b>
7. Set a name as your DB instance identifier say <b>"dbtest</b>
8. Set master username and master password
9. In connectivity we can proceed with "Don't connect to an EC2 compute resource
10. Set network type acordingly, but I preferred IPv4
11. <b>Now comes a very important part:</b> "Public access"
12. Down in additional configuration confirm the Database port as "5432"
13. In Database authentication: select <b>"Password and IAM database authentication"</b>
14. Down further in Aditional configuration, we will have <b>database options</b> select the database name here, I entered "sendgriddbtest"
15. Futher we will options for Encryption, backup, retention period, we can select all those according to the needs.
16. <b> Finally click on</b> : "Create database"

<b> Wait for a few minutes for it to set up until the status shows as "Available"</b>

#### After the status of the database is Available, we also want to connect it to Database client locally.
##### To connect to database client locally:
1. Click on the database created in the AWS RDS console
2. In under connectivity & security tab, we will see "Endpoint", for connecting to Database client that is our DB_HOST
3. The port number shown is our DB_PORT
4. Down under Security group rules, click on the inbound type security group and wait for it to open, then under Inbound rules, click on "Edit inbound rules"
5. Then click on Add rule and select source type as "PostgreSQL" and select source as Any IPv4. This will allow Database client on local machine to make a connection request to the RDS service on AWS.
6. Then under Manage IAM roles select the IAM role we created for this purpose of giving access of RDS to Lambda
7. Then on features, choose a feature to add, here in our case it is Lambda.
8. Now start the VS Code and download a database client, in my case I downloaded an extension named "Database Client".
9. Click on the extension and the click on "Create Connection".
10. In server type select PostgresSQL.
11. In the host section enter the endpoint on the RDS database created
12. In username enter the master username that we set during RDS Databse creation.
13. In password enter the master password.
14. We can also select the database name (optional), since my DB_NAME was sendgriddbtest, i entered the same.
15. Finally click on <b>"Connect"</b>
16. Now for saving the response in the database, we are required to create a table named sendgrid_response.
    ```
    CREATE TABLE sendgrid_response (
    id SERIAL PRIMARY KEY,
    status_code INTEGER,
    body TEXT,
    headers JSONB
    );
    ```

![VSCode_Database_Connection](https://github.com/Amarjit0511/go-task-sheet/assets/54772122/90cd76a4-e163-481e-bac3-15349cf64c7d)


#### Now that we have everything in place, we should test run the application
1. Go to AWS Lambda console
2. Go to the Test section
3. In Test event section click on Create new event
4. Down in event json, provide the following
```
{
  "resource": "/send-email",
  "path": "/send-email",
  "httpMethod": "POST",
  "headers": {
    "Content-Type": "application/json"
  },
  "body": "{\"name\": \"John Doe\", \"email\": \"johndoe@example.com\"}"
}
```
5. Then click on "Test"

![Execution_log_of_test](https://github.com/Amarjit0511/go-task-sheet/assets/54772122/6e399b2d-7417-441c-af2b-0b5cd2b5bc81)
