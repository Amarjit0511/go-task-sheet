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

