# Build this image based on the specified image
FROM postgres:15.4


# Will execute any command in a shell inside the container environment
RUN apt-get update
RUN apt-get install -y procps vim


# This instruction is executed when a container starts
# There is only one "CMD" instruction in a Dockerfile
# CMD [ "npm", "start" ]
