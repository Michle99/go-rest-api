# Docker – Create MySQL Container

# Install MySQL Server using Docker
docker pull mysql/mysql-server:8.0

# Verify MySQL Image:
docker images

# Create MySQL Docker Container
Following is the command to create MySQL docker container.

1.  --name=mysql-docker – To Specify Container name.
2.  -p 3306:3306 – Port Binding -p <host_port>:<container_port> , if would like to access form different port (ex: 3309) then command would be -p 3309:3306
3.  -e MYSQL_ROOT_HOST='%' – to specify host to access MySQL DB for root user. '%' is to allow to access from any other containers. You can change this to specific host.
4.  -e MYSQL_ROOT_PASSWORD=password – To set root user password.
5.  -e MYSQL_DATABASE=docker_db – Name of the Database to be created.
6.  -d mysql/mysql-server:8.0 – To Specify image name (which is created in previous step), -d means containers starts in detached mode.

command:
docker run --name=mysql-docker -p 3306:3306 -e MYSQL_ROOT_HOST='%' -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=docker_db -d mysql/mysql-server:8.0

# Verify Created Container:
docker ps -a 

# Stop MySQL Container:
docker stop mysql-docker

# Start MySQL Container:
docker start mysql-docker

# Connecting MySQL server from within container
You can run MySQL client using command: 
``` 
docker exec -it mysql-docker mysql -uroot -p 
```
It asks user password, enter password that used to create container.

# Show the Database names to verify Setup :
Issue mysql command: 
```
show databases;
```

# Setup MySQL server and create a container using Docker-Compose
Make sure you already install Docker-Compose. Before starting Docker Compose remove previous setup,
issue the following commands:
1.   ``` docker stop mysql-docker```  – Stop container
2.     ```docker container rm mysql-docker``` – Remove container
3.     ```docker image rm mysql/mysql-server:8.0``` – Remove Image

# Create a Directory
$ mkdir mysql-db

# Create a docker compose .yml file
$ ```cd mysql-db``` – Navigate to directory which is created in previous step

$ ```touch docker-compose.yml``` and $ ```sudo nano docker-compose.yml``` – Add the following lines and save the file.

```
version: '3'

services:

  mysql-docker:
    image: mysql/mysql-server:8.0
    environment:
      MYSQL_ROOT_HOST: "%"
      MYSQL_ROOT_PASSWORD: password
      MYSQL_DATABASE: docker_db
    ports:
      - "3306:3306"
```

Once you created docker-compose.yml run the following command in the same directory where .yml is created.
``` $ docker-compose up ```

# Verify docker container
You can verify created container and images using $ docker-compose ps and $ docker-compose images