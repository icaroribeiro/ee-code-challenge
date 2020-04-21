# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# It starts from the node base image.
FROM node

# Add maintainer info.
LABEL maintainer="Ícaro Ribeiro <icaroribeiro@hotmail.com>"

# Set the current working directory inside the container.
WORKDIR /app

# Copy the source from the current directory to the working directory inside the container.
COPY . .

RUN npm install

# Command to start the application.
CMD ["npm", "start"]