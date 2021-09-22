FROM mongo:5.0.3

COPY init.sh .
COPY mongo.conf /etc/mongo/mongo.conf

RUN mongod -f /etc/mongo/mongo.conf
RUN mongo < init.sh
