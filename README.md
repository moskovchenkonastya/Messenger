## Сonsol Messenger

   * Авторизация
   * Постом создать пользователя с логином
   * Если пользователя еще нет, выдать пароль
   * Если пользователь есть, запросить пароль
   * Список пользователей + аватарка
   * История сообщений
   * Список авторизаций
     
# Требования:
- Клиент-сервер (клиентом может быть браузер, телнет или отдельная go-программа), есть REST API
- Есть параллельная обработка: горутины (> 1 и с лимитом) или  каналы
- Считывание файлов в формате init, yaml, json, xml
- Хранение в СУБД не менее 3 связанных сущностей
 
 # Реализация: 
 ![Модель](https://github.com/moskovchenkonastya/Messenger/blob/master/Screen%20Shot%202017-09-06%20at%2014.48.42.png)
 
 # Скриншоты приложения
 ![View Login](https://github.com/moskovchenkonastya/Messenger/blob/master/Screen%20Shot%202017-09-06%20at%2014.48.55.png)
 ![View main](https://github.com/moskovchenkonastya/Messenger/blob/master/Screen%20Shot%202017-09-06%20at%2014.49.03.png)
 ![View profile](https://github.com/moskovchenkonastya/Messenger/blob/master/Screen%20Shot%202017-09-06%20at%2014.49.11.png)
