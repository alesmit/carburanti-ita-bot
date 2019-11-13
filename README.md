# Carburanti ITA Bot

[![Telegram](https://img.shields.io/badge/telegram-%40FuelMasterBot-%230088cc)](https://t.me/FuelMasterBot)

This is the source code of the Telegram bot [@FuelMasterBot](https://t.me/FuelMasterBot). The bot will receive your position and send a list of the three closest gas stations along with prices information.

Please note that this bot only works in Italy. Information about prices and gas stations are provided by the italian Ministry of Economic Development, which updates datasets everyday at 8:00 AM GMT+1. You can find more about it [here](https://www.mise.gov.it/index.php/it/open-data/elenco-dataset/2032336-carburanti-prezzi-praticati-e-anagrafica-degli-impianti).

This bot is deployed on Heroku and it uses the free plan. Since it doesn't use any database to cache datasets, the first request of the day after 8 AM takes a while to download the most fresh CSV. The CSV is then parsed on each request; this makes the response time quite bad, yeah, but this is what we can do with a free instance. I made this bot mainly because I was learning Go and I wanted to build something with it.

## Credits

- The icon in the botpic is Fuel by [Jason Grube](https://thenounproject.com/grubedoo/) from the Noun Project
- Thanks [@neonima](https://github.com/neonima) for helping me take the first steps with Go

## Contributions

Contributions are welcome, so feel free to fork the repo and open a PR.
