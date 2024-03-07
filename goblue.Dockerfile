FROM gpac/ubuntu:latest

RUN apt install -y libc6 libc-bin

RUN apt -y update && apt -y upgrade

RUN apt -y install build-essential pkg-config g++ git cmake yasm

RUN apt install build-essential pkg-config git

WORKDIR /playground/

COPY docs /playground/

COPY main /playground/

COPY goBluev2.db /playground/

COPY .env /playground/

RUN chmod +x main

CMD ["./main","run"]
 
