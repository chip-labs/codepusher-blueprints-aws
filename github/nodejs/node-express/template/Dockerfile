FROM node:22-alpine AS base

FROM base AS builder

WORKDIR /usr/src/app

COPY package.json package-lock.json ./

RUN npm install --only=production

FROM base AS final

WORKDIR /usr/src/app

COPY --from=builder /usr/src/app/node_modules ./node_modules

COPY . .

RUN npm run build

EXPOSE 3000

ENTRYPOINT ["node", "dist/server.js"]
