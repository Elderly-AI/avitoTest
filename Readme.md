## AvitoTech trainee

## Принятые решения 
- Метод получения списка номеров отелей как я понял должен отдавать все номера без пагинация,
а сортировку сделал с помощью ручки ?sort="по чему сортируем", т.е. запретил сортировку сразу по двум полям.

## Запуск проекта 
make start (локальная постгреха должна быть остановлена "sudo service postgresql stop")

## Запуск тестов
make tests (покрытие около 66%)

## документация 
[Swagger-doc](http://localhost:9000/docs) (Все данные возвращаются внутри тега "data")

## p.s.
env файл закинул в репу чтобы удобно было запускать, так то его понятно там быть не должно)