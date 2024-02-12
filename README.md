
# app
# Проект Максим

Микросервис

- api
    - chi, fiber, http, gin, gorilla
    - /import/<something>
    - /liveness
        - 200 OK
        - metadata
    - /readiness
        - 200 OK
        - service healthy
    - Получает на вход json файл
    - Парсит его
        - Если не смогла спарсить → пушит в кафку Error
        - Если ок → сохраняет на s3 и пушит FileExtractSuccess
- processor (deferred processing)
    - Загружает файл из S3, валидирует и складывает в postgres
- monitoring (monitor events from api and processor), not required
- postgres
- s3 (minio)

ETL (extract, transform, load) events:

- FileExtractError
- FileExtractSuccess
    - key:
    - value: filename, request_id,
- ValidationError

Формат данных (найдем примеры)

```
/cmd
	/api
		main.go
	/processor
		main.go
	/monitoring
/deployments
	/k8s
		ingress.yaml
		/api
			deployments.yaml
			service.yaml
			ingress.yaml
		/processor
		/monitoring
	/test
		.env.example
		Dockerfile
		docker-compose.yaml
			- api
			- processor
			- monitoring
			- zookeeper
			- kafka
			- postgres
			- prometheus
			- grafana
			- grafana-dashboard
	/development
		.env.example
		Dockerfile (CompileDaemon + ldflags(apiVersion))
		docker-compose.yaml
Dockerfile.api
Dockerfile.processor
Makefile
```

ldflags

```
--ldflags
go tool nm ./main | grep utils
go build -ldflags="-X 'main.branchName=$(git rev-parse --abbrev-ref HEAD)'" -o ./main main.go && ./main
```

CompileDaemon

https://github.com/githubnemo/CompileDaemon

Makefile

```
make:
	test // запускаем юнит тесты go test ./...
	itest-up // поднимает docker compose с окружением для интегр. тестов
	itest-migrate // apply миграций
	itest-rollback // rollback миграций
	itest-down // удаляет контейнеры и volumes
	dev-up
	dev-migrate
	dev-rollback
	dev-down
	build // собирает Docker образ from scratch
		/api
		/processor  // пока не нужно
		/monitoring // пока не нужно
```

1 этап

- Создать репозиторий +
- Выбрать http библиотеку fiber(возможности роутинга, активно поддерживаемая мидлвара для прометеуса, сама либа активно развивается, интересно попробовать)
- Написать Docker образ под dev окружение +
    - Запускаем через CompileDaemon (в докер образе нужно скачать go install compiledaemon) +
- Написать Makefile
    - dev-*
- Написать docker-compose
    - Сервисы
        - api

Для кафки библиотека sarama от IBM

/examples
    /data
        /json
        /sh скрипт отправка курлом данных в апи

/usecase, перенести логику, парсинг(дефер body)
JsonImporter
    newJsonImporter
    importCompany(ctx context, body []byte) error (etag)

1. Кладем на s3, получаем метаданные о файле (возвращает id)
   /companyName/files
2. Отправляем метаданные о файле в топик processedFile
   ключ - название компании, для каждой компании - партиция, 
   значение url на s3, это сообщение примет консюмер
   3. Создать клиент продюсер sarama, при инициализации юскейса прокинуть
   продюсера, юскейс сохраняет на с3 и пушит в топик processedFile

3. Приложение consumer читает, обрабатывает и парсит, пушит в топик(потом)

топик processedFile
кластер из 3 брокеров, репФактор 3, инсинкРеп 2, партиций 3

Тестовый файл с данными c.json(один элемент из companies.json)

//kafka connect, эластик

балансирование по партициям, хеш, раунд робин, настройка партиции
на консюмере(антипатерн)



todo:
логгер
врап ошибок
унести логин и пароль от минио в енв(лучше в секреты)
добавить автоматическое создание бакета для с3 при запуске приложения в докере