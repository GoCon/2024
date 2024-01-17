FROM node:20-slim as base

RUN npm install -g pnpm@8

WORKDIR /app

COPY package.json pnpm-lock.yaml ./
RUN pnpm install --frozen-lockfile

CMD [ "pnpm", "start", "--host" ]
