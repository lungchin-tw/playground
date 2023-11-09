# Build this image based on the specified image
FROM node:20-alpine

# COPY files in folder src to /app/ in the image
COPY src /app/

# Set the working directory for all following commands
WORKDIR /app

# Will execute any command in a shell inside the container environment
RUN npm install


# This instruction is executed when a container starts
# There is only one "CMD" instruction in a Dockerfile
CMD [ "npm", "start" ]
