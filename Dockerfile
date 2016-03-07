FROM iron/base

RUN apk update && apk upgrade \
  && apk add python \
  && rm -rf /var/cache/apk/*

ADD dhcpharvester.py /dhcpharvester.py
ADD switchGo /switchGo
ADD main.go /main.go
ADD ofdpa-i.12.1.1_12.1.1+accton1.7-1_amd64.deb /ofdpa-i.12.1.1_12.1.1+accton1.7-1_amd64.deb

ENTRYPOINT [ "python", "/dhcpharvester.py" ]
