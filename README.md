** TODO List API (Go)**

Простой REST-сервис TODO-лист, написанный на Go.

Позволяет:

✔ Создавать задачи
✔ Получать задачу по названию
✔ Получать все задачи
✔ Фильтровать по параметру completed
✔ Отмечать задачу как выполненную
✔ Удалять задачи

📌 Реализовано с использованием:

net/http — стандартный HTTP-сервер Go

github.com/gorilla/mux — роутер для REST-маршрутов

Чёткая архитектура: business logic → handlers → HTTP server

🚀 API Endpoints
Метод	      Путь	                 Описание
POST	      /tasks	               Создать задачу
GET	        /tasks/{title}	       Получить задачу по названию
GET	        /tasks	               Получить все задачи
GET	        /tasks?completed=true	 Получить только завершённые
PATCH	      /tasks/{title}	       Отметить задачу как выполненную
DELETE	    /tasks	               Удалить задачу
