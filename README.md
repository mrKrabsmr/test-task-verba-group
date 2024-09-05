API Server

GET /api/v1/tasks<br>
GET /api/v1/tasks/{id}<br>
POST /api/v1/tasks<br>
PUT /api/v1/tasks/{id}<br>
DELETE /api/v1/tasks/{id}<br>

RUNNING STEPS:
1. create .env file and write there `ADDRESS` and `DBADDRESS`. EXAMPLE:
   ![image](https://github.com/user-attachments/assets/7cb49168-9e6e-4fba-ac04-ce059632da28)
2. run postgresql. EXAMPLE with using docker:
   ![image](https://github.com/user-attachments/assets/64832b29-47cd-45ad-a5bc-66288beda61b)
3. run golang application:
   
   go run cmd/main.go
   
     flags:
   
       --init (use for init db and create table tasks)
   
       --debug (use for debug mode)
   
       -v %d (use for set version, default 1)
   

