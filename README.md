# Voting Web Service
* We use docker to spin up the HTTP webserver and a MySQL database
* Golang microservice is a REST API which we use to facilitate communication between our backend and frontend
* Frontend is written in Vanilla JS

### Front End Setup
In the project directory, install `NodeJS`
* Follow https://nodejs.org/en/download/package-manager/ to install
* Run `npm install speakeasy qrcode` to install 2FA libraries 

### MySQL Server Setup ###
Run the docker compose file from the top directory and ensure that the init_db.sql file and mysql-compose.yaml files are present. Run `sudo docker-compose -f mysql-compose.yaml up --build`. Clean up by running with `sudo docker-compose down -v --remove-orphans`.

### Run Go API ###
* Navigate to cmd/voting_web_service/main.go
* Execute go run main.go

### HTTPS
To enable HTTPS, you'll need to set up local certificate.

Helpful links:
* https://www.freecodecamp.org/news/how-to-get-https-working-on-your-local-development-environment-in-5-minutes-7af615770eec/
* https://askubuntu.com/questions/73287/how-do-i-install-a-root-certificate
* https://www.digitalocean.com/community/tutorials/how-to-configure-nginx-with-ssl-as-a-reverse-proxy-for-jenkins


### Localhost
Go to `https://localhost:8880` to see application