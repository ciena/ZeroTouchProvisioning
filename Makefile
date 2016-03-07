build:
	GOOS=linux GOARCH=amd64 go build -o switchGo

image:
	sudo docker build -t cord/fabricdhcpharvester .

run1:
	sudo docker-compose -f harvest-compose-1.yml up

run2:
	sudo docker-compose -f harvest-compose-2.yml up

runquiet1:
	sudo docker-compose -f harvest-compose-1.yml up -d

runquiet2:
	sudo docker-compose -f harvest-compose-2.yml up -d