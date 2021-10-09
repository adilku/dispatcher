# dispatcher
**Пример рассылки**

```
 //Сборка
 go mod download -x
 go build -o dispatcher cmd/dispatcher/main.go
 
 //Запуск
 //Будем брать данные из data/example.json
 ./dispatcher data/example.json
 //данные расслаются параллельно в три канала с сохранением порядка.
 //при этом при ошибке отправки перепотправки не будет.
 ```
