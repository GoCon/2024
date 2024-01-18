FROM node:20-slim as base

RUN apt-get update && \
  apt-get install -y --no-install-recommends git && \
  npm install -g pnpm@8 && \
  apt-get clean && \
  rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY package.json pnpm-lock.yaml ./
RUN pnpm install --frozen-lockfile

CMD [ "pnpm", "start", "--host" ]
