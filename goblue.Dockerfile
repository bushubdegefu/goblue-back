FROM ubuntu:latest

RUN apt install -y libc6 libc-bin

WORKDIR /playground/

COPY docs /playground/

COPY main /playground/

COPY goBluev2.db /playground/

COPY .env /playground/

RUN chmod +x main

CMD ["./main","run"]
 
