Что нужно сделать
Напишите HTTP-сервис, который принимает входящие соединения с JSON-данными и обрабатывает их следующим образом: 

    1. Сделайте обработчик создания пользователя. У пользователя должны быть следующие поля: имя, возраст и массив друзей. Пользователя необходимо сохранять в мапу. Пример запроса: 

POST /create HTTP/1.1
Content-Type: application/json; charset=utf-8
Host: localhost:8080
{"name":"some name","age":"24","friends":[]}
Данный запрос должен возвращать ID пользователя и статус 201.



    2. Сделайте обработчик, который делает друзей из двух пользователей. Например, если мы создали двух пользователей и нам вернулись их ID, то в запросе мы можем указать ID пользователя, который инициировал запрос на дружбу, и ID пользователя, который примет инициатора в друзья. Пример запроса:

    POST /make_friends HTTP/1.1
Content-Type: application/json; charset=utf-8
Host: localhost:8080
{"source_id":"1","target_id":"2"}
Данный запрос должен возвращать статус 200 и сообщение «username_1 и username_2 теперь друзья».



    3. Сделайте обработчик, который удаляет пользователя. Данный обработчик принимает ID пользователя и удаляет его из хранилища, а также стирает его из массива friends у всех его друзей. Пример запроса:

    DELETE /user HTTP/1.1
Content-Type: application/json; charset=utf-8
Host: localhost:8080
{"target_id":"1"}
Данный запрос должен возвращать 200 и имя удалённого пользователя.



    4. Сделайте обработчик, который возвращает всех друзей пользователя. Пример запроса:

    GET /friends/user_id HTTP/1.1
Host: localhost:8080
Connection: close
После /friends/ указывается id пользователя, друзей которого мы хотим увидеть.



    5. Сделайте обработчик, который обновляет возраст пользователя. Пример запроса:

PUT /user_id HTTP/1.1
Content-Type: application/json; charset=utf-8
Host: localhost:8080
{"new age":"28"}
Запрос должен возвращать 200 и сообщение «возраст пользователя успешно обновлён».
