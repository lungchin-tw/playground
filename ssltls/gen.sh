#! /usr/bin/bash

echo '$(dirname $0):' $(dirname $0)
echo '$(basename $0):' $(basename $0)
echo '$(pwd):' $(pwd)
pushd $(dirname $0)

set -x

openssl version

rm ./*.pem


# 1.Generate CA's private key and self-signed certificated

# -x509              = Output a certificate instead of a request
# -newkey rsa:4096   = Create a private key with RSA 4096-bit key
# -days 365          = Specify the number of days that the certificate is valid for
# -keyout ca-key.pem = Write the created private key to ca-key.pem
# -out ca-cert.pem   = Write the certificate to ca-cert.pem

# Interactive Mode
# openssl req -x509 -newkey rsa:4096 -days 365 -keyout ca-key.pem -out ca-cert.pem

# Non-Interactive Mode
# -nodes = No pass phrase needed(when you want it becomes easier for development environment)
openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.pem -out ca-cert.pem \
-subj "/C=TW/ST=Some-State/L=Taipei/O=Amber/OU=Amber Taiwan/CN=*.jacky.amber.com.tw/emailAddress=jacky.chen@amberstudio.com"

# Pass phrase = Used to encrypt the private key before writing it to the file.
# Identity Information Part
#   - Country Name, for example: TW
#   - State or Province Name, for example: Taipei
#   - Locality Name, for example: Taipei
#   - Organization Name, for example: Tech School
#   - Organizational Unit Name, for example: Education
# Common Name = domain name, for example: *.techschool.guru
# Email Adress


# 2. Generate a server's private key and certificate signing request(CSR)
# remove -x509, -days 365
openssl req -newkey rsa:4096 -nodes -keyout server-key.pem -out server-req.pem \
-subj "/C=TW/ST=Some-State/L=Taipei/O=AmberStudio/OU=AmberBackend/CN=*.jacky.amberbackend.com.tw/emailAddress=jacky.chen@amberstudio.com"


# 3. Use CA's self-signed certificate private key to sign server's CSR and get back the signed certificate
openssl x509 -req -in server-req.pem -days 365 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem -extfile server-ext.cnf