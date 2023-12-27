FROM rabbitmq:3.13-rc-management-alpine

COPY rabbitmq.conf /etc/rabbitmq/

ENV RABBITMQ_NODENAME=rabbit@localhost

RUN chown rabbitmq:rabbitmq /etc/rabbitmq/rabbitmq.conf
# Define environment variables.
ENV RABBITMQ_USER bushu
ENV RABBITMQ_PASSWORD yahweh
ENV RABBITMQ_PID_FILE /var/lib/rabbitmq/mnesia/rabbitmq

ADD rabbuser.sh /rabbuser.sh
RUN rabbitmq-plugins enable rabbitmq_management

RUN chmod +x rabbuser.sh

# Define default command
CMD ["/rabbuser.sh"]
EXPOSE 5672
EXPOSE 15672

