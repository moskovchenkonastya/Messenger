# Messenger
* Authorization
* Post a user with a username
* If the user does not already exist, give the password
* If the user is present, request a password
* List of users + avatar
* Message history
* Authorization list (list of friends)

Требования:
- Клиент-сервер (клиентом может быть браузер, телнет или отдельная go-программа), есть REST API
- Есть параллельная обработка: горутины (> 1 и с лимитом) или  каналы
- Считывание файлов в формате init, yaml, json, xml
- Хранение в СУБД не менее 3 связанных сущностей

