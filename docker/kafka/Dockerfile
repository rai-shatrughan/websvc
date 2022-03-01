FROM alpine

ARG SC_VERSION=2.12
ARG KF_VERSION=2.2.0
ARG KAFKA_DIR=kafka_${SC_VERSION}-${KF_VERSION}
ARG KAFKA_BINARY=${KAFKA_DIR}.tgz

RUN apk --update add bash git wget ca-certificates sudo openssh rsync openjdk8 && \
  rm -rf /var/cache/apk/* && \
  rm -rf /opt  && \
  mkdir -p /opt 

ENV JAVA_HOME /usr/lib/jvm/java-1.8-openjdk
ENV PATH $JAVA_HOME/bin:$PATH

VOLUME /tmp

RUN mkdir -p /tmp/zookeeper


ADD binary/${KAFKA_BINARY} /opt/
ADD conf/server.properties /opt/${KAFKA_DIR}/config/server.properties
ADD conf/zookeeper.properties /opt/${KAFKA_DIR}/config/zookeeper.properties

ENV PATH=$PATH:/opt/${KAFKA_DIR}/bin
ENV KAFKA_HOME=/opt/${KAFKA_DIR}

WORKDIR /opt/${KAFKA_DIR}