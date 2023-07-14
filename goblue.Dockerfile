FROM ubuntu:latest

RUN apt install -y libc6 libc-bin

WORKDIR /playground/

COPY docs /playground/

COPY main /playground/

COPY goBlue.db /playground/

COPY .env /playground/

RUN chmod +x main

CMD ["./main","run"]
 
