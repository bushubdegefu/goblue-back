FROM yugabytedb/yugabyte:2.19.3.0-b140

RUN echo "something" > myfile.txt

RUN dnf -y  update

RUN dnf install chrony -y

CMD bin/yugabyted start && tail -f myfile.txt 
