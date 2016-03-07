build:
	GOOS=linux GOARCH=amd64 go build -o switchGo

image:
	sudo docker build -t cord/fabricdhcpharvester .

run:
	sudo docker-compose -f harvest-compose.yml up

runquiet:
	sudo docker-compose -f harvest-compose.yml up -d
