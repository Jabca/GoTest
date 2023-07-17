# Stage 1: Build the React app
FROM node:14.17.6 as build-stage
WORKDIR /app
COPY frontend/package*.json ./
RUN npm install
COPY frontend ./
RUN npm run build

# Stage 2: Serve the built app using Nginx
FROM nginx:1.21.3-alpine as production-stage
COPY --from=build-stage /app/build /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
