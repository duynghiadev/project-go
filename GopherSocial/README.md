### [Tutorial: Complete Backend Engineering Course](https://www.youtube.com/watch?v=h3fqD6IprIA&list=PLYEESps429vrFV0yiN_MCaDPhnYb0qRxK&index=6)

---

### What tech are we using, pretty simple:

- Go 1.22
- Docker

- Postgres running on Docker
- Swagger for docs

- Golang migrate for migrations

### Testing rate limiter

```bash
 npx autocannon -r 22 -d 1 -c 1 --renderStatusCodes http://localhost:8080/v1/health
```
