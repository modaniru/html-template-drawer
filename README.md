## Установка своих one page страниц
Приложение имеет предустановленные странички по следующим url
* localhost/first
* localhost/second
* localhost/third
### html 
Разметка находится в template директории. Каждая страничка находится в отдельной папке для более удобной иерархии. html файлы должны иметь уникальное название.
### Статические файлы
Стили, картинки, скрипты, шрифты находятся в папке static. Они могут быть подключены к любым html файлам. Н.р. first.html может подключить себе static/style/second/style.css
### Роутинг html страниц
Для того, чтобы страница была доступна, нужно отредактировать internal/router.go. Можно редактировать следующий строчки в методе GetRouter()
~~~go
    // routing
	r.router.GET("/first", r.LoadHtmlPage("first.html")) // будет доступен по localhost/first, выдаст страничку first.html
	r.router.GET("/second", r.LoadHtmlPage("second.html"))
	// located in template/third/third.html but we must write just a file name
	r.router.GET("/third", r.LoadHtmlPage("third.html"))
~~~
Важно то, что в методе LoadHtmlPage мы вводим не относительный путь до файла, а ЕГО название
## Локальный запуск
Нужен установленный docker и docker compose
~~~bash
docker compose up
~~~
сервер будет развернут локально на 80 порту
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
