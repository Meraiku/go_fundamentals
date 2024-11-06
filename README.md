# Задания по Go

Я совместил задания по Go fund, Go tests, DevOps.

## Реализовано

- Unit тесы
- Integration тесты
- Запуск в контейнере
- CI pipeline

## Запуск

### Тестового сервиса

```bash
make run          #Локально
make docker       #В контейнере
```

- Запускает HTTP сервер с 2-умя ручками
	- `/` - Выводит `Hello World`
	- `/search` - Своего рода прокси на поисковик google. Возвращает только HTML страницу :)
- Требуется для Integration тестов

### Unit тестов

```bash
make unit
make cover
```

### Integration тестов

```bash
make integration
```
