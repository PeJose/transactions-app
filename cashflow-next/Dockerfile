FROM node:23-alpine

RUN corepack enable

WORKDIR /app

COPY ./cashflow-next/package.json ./cashflow-next/pnpm-lock.yaml ./

RUN pnpm install

CMD ["pnpm", "dev"]