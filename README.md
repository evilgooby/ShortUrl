# short_url_service
short_url_service: Сокращение URL-ссылок — это программа, написанная на языке Go, которая позволяет пользователям создавать короткие URL-ссылки для длинных URL-ссылок. Программа использует базу данных PostgreSQL для хранения оригинальных URL-ссылок и их коротких версий.

## Основные функции: 
Сокращение URL-ссылок: Пользователи могут ввести длинную URL-ссылку и получить короткую версию для использования в социальных сетях, электронных письмах или других местах.

Аудит: Программа ведет журнал всех операций, связанных с сокращением URL-ссылок. 

Тестирование: Программа содержит комплексные тесты, обеспечивающие стабильность и надежность.

## Установка:

```
git clone https://github.com/evilgooby/short_url_service.git
```

## Использование: 
Запустите программу:
```
sudo docker compose build
sudo docker compose up
```
Введите длинную URL-ссылку, которую вы хотите сократить.

Программа сгенерирует короткую версию URL-ссылки.

Пример запроса: 
```
curl -X POST -H "Content-Type: application/json" -d '{"long_url": "https://rutube.ru/video/c4e6290ace7933b84b73810771a6f31e/"}' http://localhost:8080/CreateShortUrl
```
Пример ответа:
```JSON
{
  "short_url":"http://localhost:8080/XkXzo63P"
}
```
Пример запроса:
```
curl -X GET -H "Content-Type: application/json" -d '{"short_url":"http://localhost:8080/XkXzo63P"}' http://localhost:8080/GetLongUrl
```
Пример ответа:
```JSON
{
  "long_url":"https://rutube.ru/video/c4e6290ace7933b84b73810771a6f31e/"
}
```
