
#!/bin/bash
SC_VERSION=2.12
KF_VERSION=2.2.0
FILE=./binary/kafka_$SC_VERSION-$KF_VERSION.tgz
if test -f "$FILE"; then
    echo "$FILE exists."
else
    mkdir -p binary
    cd binary
    wget https://archive.apache.org/dist/kafka/$KF_VERSION/kafka_$SC_VERSION-$KF_VERSION.tgz
fi
