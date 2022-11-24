<a id="readme-top"></a>

<br />
<div align="center">
  <a href="https://github.com/ssergomol/E-Wallet">
</a>

<h3 align="center">E-Wallet</h3>
</div>





<!-- TABLE OF CONTENTS -->
<details>
  <summary>Содержание</summary>
  <ol>
    <li>
      <a href="#описание-проекта">Описание проекта</a>
      <ul>
        <li><a href="#технологии">Технологии</a></li>
      </ul>
    </li>
    <li><a href="#возможности-сервиса">Возможности сервиса</a></li>
    <li>
      <a href="#установка-и-запуск-сервера">Установка и запуск сервера</a>
    </li>
    <li><a href="#контакты">Контакты</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->
<!-- ## About the project -->
## Описание проекта

Сервис для управления балансом пользователей, который позволяет осуществлять основные денежные операции
<!-- Service for managing users' balance which allows the money operations such as crediting funds, debiting funds, transferring funds from user to user as well as  obtaining the user's balance -->

<!-- TECHNOLOGIES -->
### Технологии
<!-- ### Technologies -->

* [![GoLang-logo]][GoLang-url]
* [![PostgreSQL-logo]][PostgreSQL-url]
* [![Docker-logo]][Docker-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- Features -->
<!-- ## Features and implementations -->
## Возможности сервиса

* Сервис предоставляет HTTP API в формате JSON
* Зачисление средств на баланс
* Cписание денег c баланса
* Получение баланса пользователя

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ## Installation -->
## Установка и запуск сервера
1. Установите Docker и docker-compose для своей ОС https://docs.docker.com/compose/install/
2. Склонируйте репозиторий и перейдите в корень проекта
```
git clone https://github.com/ssergomol/E-Wallet.git && cd ./E-Wallet
```
3. Выполните
```sh
docker-compose up
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Примеры запросов и возможные ответы
1. Метод зачисление средств на баланс  
Принимает ID пользователя, ID операции (0 в данном случае) и сумму зачисления

Запрос:

```
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"user_id":0, "service_id": 0, "price": "100.00"}' \
  http://localhost:8080/balance
```

Тело ответа (история запросов):

```
[
    {
        "user_id": 0,
        "service_id":0,
        "price":"100.00",
        "description":"Replenish balance"
    }
]
```

2. Метод Cписания денег c баланса  
Принимает ID пользователя, ID услуги (1 в данном случае) и сумму списания

Запрос:
```
curl --header "Content-Type: application/json" \
  --request POST \
  --data \
  '{ "id": 0, "user_id": 0, "service_id": 1, "price":"10.00"}' \
  http://localhost:8080/accounts
```

Тело ответа (история запросов):
```
[
    {
        "user_id": 0,
        "service_id":0,
        "price":"100.00",
        "description":"Replenish balance"
    },
    {
       "user_id": 0,
        "service_id":1,
        "price":"10.00",
        "description":"Spend funds" 
    }
]
```

3. Метод получения баланса пользователя  
Принимает id пользователя
Запрос:
```
curl --header "Content-Type: application/json" \
  --request GET \
  --data \
  '{ "id": 0 }' \
  http://localhost:8080/balance
```

Тело ответа (в случае успеха возвращает id пользователя и его баланс):
```
{ 
    "id": 0,
    "sum": "1000.00"
}
```

<!-- CONTRIBUTING -->
<!-- ## Contributing

If you have any intentions that would make this project better, fork the repo and create pull request

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p> -->



<!-- CONTACT -->
<!-- ## Contact -->
## Контакты

Сергей Молчанов - @ssergomol - ssergomoll@gmail.com

<p align="right">(<a href="#readme-top">back to top</a>)</p>

[React-logo]: https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB
[React-url]: https://reactjs.org/
[GoLang-url]: https://go.dev
[GoLang-logo]: https://img.shields.io/badge/GoLang-ffffff?style=for-the-badge&logo=Go&logoColor=7bccec
[product-screenshot]: images/home_page.png
[PostgreSQL-url]: https://www.postgresql.org/
[PostgreSQL-logo]: https://img.shields.io/badge/PostgreSQL-ffffff?style=for-the-badge&logo=PostgreSQL&logoColor=008bb9
[JavaScript-url]: https://javascript.com
[JavaScript-logo]: https://img.shields.io/badge/JavaScript-323330?style=for-the-badge&logo=javascript&logoColor=f0db4f
[Docker-logo]: https://img.shields.io/badge/Docker-ffffff?style=for-the-badge&logo=docker&logoColor=0db7ed
[Docker-url]: https://www.docker.com/
