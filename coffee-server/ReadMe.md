clone the project

Run> make stop_containers to stop any running containers

Run> make create_container to create the container

Run> make create_db to run the DB

Run>  psql -h localhost -p 5436 -U ${USER} -d ${DB_NAME} to connect to postgres using psql
