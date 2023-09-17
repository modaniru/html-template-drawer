## Развертывание на удаленном сервере
### ssl и dns
Чтобы развернуть приложение с работующим ssl, нужно сначала привязать dns сервера к домену.
### Убедитесь, что на сервере установлен docker и docker compose
install docker on ubuntu
~~~bash
apt install docker
~~~
install docker compose on ubuntu: [тык](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-compose-on-ubuntu-22-04)
### Клонирование репозитория
~~~bash
git clone https://github.com/modaniru/html-template-drawer
cd html-template-drawer
~~~
#### Редактирование Caddyfile
Чтобы развернуть с работующим ssl, нужно отредактированить Caddyfile следующим образом:
~~~bash
мой-домен.ру

reverse_proxy html-template
~~~
### Запуск
~~~bash
docker compose up
~~~
Приложение запустится на 80 порту (http) и 443 порту (https)
