# Build the front-end
# This is a multi-stage build, so we need to specify the stage name
FROM node:16.13.0-alpine3.14
# This is the stage name
WORKDIR /usr/src/app

# Copy the package.json and package-lock.json files
COPY package.json ./
# Install the dependencies
RUN yarn

# Copy the rest of the files
COPY . .
# Build the front-end
RUN yarn build