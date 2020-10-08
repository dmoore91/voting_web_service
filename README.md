# Voting Web Service
* We use docker to spin up the HTTP webserver and a MySQL database
* Golang microservice is a REST API which we use to facilitate communication between our backend and frontend
* Frontend is written in Vanilla JS

### MySQL Server Setup ###
Run the docker compose file from the top directory and ensure that the init_db.sql file and mysql-compose.yaml files are present. Run `sudo docker-compose -f mysql-compose.yaml up --build`. Clean up by running with `sudo docker-compose down -v --remove-orphans`.
