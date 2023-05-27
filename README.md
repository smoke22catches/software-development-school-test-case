# Software Development School Test Case

## Як запустити проект?

1. Зібрати контейнер:

```bash
docker build -t test-case .
```

2. Запустити контейнер, при цьому відкривши порт 5000 для прослуховування:

```bash
docker run -p 5000:5000 test-case
```
