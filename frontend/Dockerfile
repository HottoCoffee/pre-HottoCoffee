FROM node:18-alpine AS base

# 1. Install dependencies only when needed
FROM base AS deps
# Check https://github.com/nodejs/docker-node/tree/b4117f9333da4138b03a546ec926ef50a31506c3#nodealpine to understand why libc6-compat might be needed.
RUN apk add --no-cache libc6-compat

WORKDIR /app

# Install dependencies based on the preferred package manager
COPY frontend/package.json ./
COPY frontend/package-lock.json ./
COPY frontend/.npmrc ./
RUN npm install --legacy-peer-deps

# 2. Rebuild the source code only when needed
FROM base AS builder
WORKDIR /app/frontend
COPY --from=deps /app/node_modules ./node_modules
COPY ./frontend /app/frontend
COPY ./swagger /app/swagger
# This will do the trick, use the corresponding env file for each environment.
ENV NODE_ENV=production
COPY frontend/.env.docker .env.production

RUN npm run generate
RUN npm run build

# 3. Production image, copy all the files and run next
FROM nginx:latest

COPY --from=builder /app/frontend/out /usr/share/nginx/html
COPY ./frontend/docker/default.conf /etc/nginx/conf.d/default.conf

EXPOSE 80
RUN service nginx start
