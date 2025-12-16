FROM rabbitmq:3.13-management
RUN rabbitmq-plugins enable rabbitmq_stomp

# docker run -d --name rabbitmq_stomp -p 5672:5672 -p 15672:15672 -p 61613:61613 rabbitmq-stomp
