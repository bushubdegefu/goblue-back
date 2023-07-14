FROM node:20-alpine3.17

RUN apk add npm

RUN npm install npm@latest

RUN npm i newman

RUN npm i newman-reporter-junitfullreport

RUN npm --version

WORKDIR /playground/

COPY docs /playground/

COPY main /playground/

COPY blue_test.json /playground/

COPY test_env.json /playground/

COPY goBlue.db /playground/

COPY .env /playground/

RUN chmod +x main

RUN ./main run >> applog.log &


RUN newman run blue_test.json -e test_env.json -r junit --reporter-junit-export result.xml