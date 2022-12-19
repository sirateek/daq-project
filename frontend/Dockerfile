FROM node:16-alpine AS builder

# make the 'app' folder the current working directory
WORKDIR /app

# copy both 'package.json' and 'package-lock.json' (if available)
COPY package*.json ./
COPY yarn.lock ./

# install project dependencies
RUN yarn install

# copy project files and folders to the current working directory (i.e. 'app' folder)
COPY . .

# build app for production with minification
RUN yarn build


FROM nginx AS runner

COPY --from=builder /app/dist/. /usr/share/nginx/html/.
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
