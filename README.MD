#В процессе допиливания - редактирование из клиента

# Спринты 11-12. Второй выпускной проект  "Менеджер паролей GophKeeper" (но получился go-keeper)
## Как запускать
### Предварительно создаем базу данных postgres и пользователя для нашего приложения.  
Где вы возьмете сервер Postgress думаю придумаете сами.
Краткая настройки БД:
```bash
sudo su postgres
psql
```
В CLI psql:
```postgres
create database go_keeper;
create user go_keeper with encrypted password 'go_keeper';
grant all on database go_keeper to go_keeper;
grant all on schema public to go_keeper;
```
### Сервер:
```bash
git clone git@github.com:KznRkjp/go-keeper.git
cd go-keeper/bin/server/linux/
chmod +x server
./server -d "host=localhost port=5432 user=go_keeper password=go_keeper dbname=go_keeper sslmode=disable" 
```
По умолчанию сервер стартует на ```http://localhost:4443``` (задается через ```-a localhost:4443```)
Так же можно включить отладочные сообщения ```-debug true```
На этом вкладку терминала с сервером можно оставить в покое.
### Клиент
```bash
cd go-keeper/bin/client/linux/
chmod +x client
./client
```
По умолчанию клиент подключается к  ```http://localhost:4443``` (задается через параметр  ```-a localhost:4443```)
