# exchange-parser-server

## Navigation

- [Description](#desc)
- [Stack](#stack)
- [Backend Development](#launch)
- [Documentation](#docs)


<a name="desc"></a>

## Description
Parser P2P orders and SPOT pairs from crypto exchanges.
Using fiber and Cache Redis.
Basic routes developed


- Fiat: [RUB, TRY, USD]
- Exchanges: [Binance, Bybit, Huobi, Okx]


<a name="stack"></a>

## Stack

- Backend:
  <a href=https://github.com/gofiber/fiber ><img src=https://gofiber.io/assets/images/logo.svg width="100" height="30">
  [![Redis](https://img.shields.io/badge/redis-%23DD0031.svg?style=for-the-badge&logo=redis&logoColor=white)](https://redis.io/)
- Utils: [json-iterator](https://github.com/json-iterator/go), [req](https://github.com/imroc/req), [validation](https://github.com/go-playground/validator), [redis client](https://github.com/go-redis/redis)


<a name="launch"></a>

## Start application
Start Docker Compose:
1. Clone this repo:
2. Edit .env.example to .env
3. Set your settings on .env file
4. Start docker-compose file:
bash
docker-compose up -d

<a name="docs"></a>

## Docs

Upgrade swagger documentation

bash
swag init -g app/cmd -o app/docs