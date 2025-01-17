Các api được gọi ở trong folder `golang-todoapp-next` của ứng dụng postman

---

# Golang Fullstack To-Do List App Tutorial with Next.js, PostgreSQL, and Docker

In this comprehensive tutorial, we will build a fullstack To-Do List application using modern technologies: Next.js for the frontend, Golang for the backend, PostgreSQL for the database, and Docker for containerization. Follow along to learn how to integrate these powerful tools into a seamless, scalable application. Perfect for developers looking to enhance their fullstack skills with cutting-edge technologies.

---

### Hướng dẫn chạy ứng dụng với Docker

#### Yêu cầu

- **Docker** : Cần cài đặt Docker trên máy. Có thể tải và cài đặt Docker tại [https://www.docker.com/products/docker-desktop/]().

#### Các bước thực hiện

1. **Tải PostgreSQL từ Docker Hub (nếu cần)**
   Nếu chưa có PostgreSQL trên máy qua Docker, bạn có thể tải về bằng lệnh sau:

```bash
docker pull postgres:16.3-alpine3.20
```

2. **Cấu trúc Docker Compose**
   File `docker-compose.yaml` đã được thiết lập sẵn để:

   - Chạy PostgreSQL thông qua Docker mà không cần cài đặt trực tiếp.
   - Chạy Backend (Golang) và Frontend (Next.js) trong các container Docker.

   Cấu trúc chính của tệp bao gồm:

   - Service `nextapp` cho Frontend.
   - Service `golang` cho Backend.
   - Service `db` cho PostgreSQL.

3. **Chạy ứng dụng**

   Sử dụng lệnh sau để khởi chạy ứng dụng:

```bash
   docker compose up --build -d
```

Trong đó:

- `--build`: Build lại container từ các Dockerfile.
- `-d`: Chạy Docker ở chế độ nền.

4. **Truy cập ứng dụng**

   - Frontend chạy tại: [http://localhost:3000]().
   - Backend có thể kiểm tra qua: [http://localhost:8000](http://localhost:8000).
   - PostgreSQL hoạt động trên cổng `5432`.

5. **Dừng và xóa container**

   - Để dừng ứng dụng:
     docker compose down
   - Để xóa dữ liệu không cần thiết, bao gồm cả volume và image:

   ```bash
     docker system prune -a
   ```

6. **Kiểm tra log khi gặp lỗi**
   Nếu gặp lỗi khi chạy, có thể kiểm tra log của từng container bằng lệnh:

```bash
  docker compose logs <tên_container>
```

Ví dụ:

```bash
docker compose logs nextapp
docker compose logs golang
docker compose logs db
```

---

#### Sau khi bạn chạy lệnh:

```
docker compose up --build -d
```

Các container sẽ được khởi chạy, và bạn có thể mở trình duyệt tại [http://localhost:3000](http://localhost:3000) để xem giao diện của ứng dụng.

#### Lý do:

- Service nextapp (ứng dụng Frontend) được cấu hình trong file `docker-compose.yaml` với phần:

```
ports:
  - 3000:3000
```

Điều này ánh xạ cổng 3000 của container nextapp sang cổng `3000` của máy chủ host.

Kết quả là khi truy cập [http://localhost:3000](http://localhost:3000), bạn sẽ thấy giao diện của ứng dụng Next.js.

#### Lưu ý:

- Nếu giao diện không hiển thị hoặc gặp lỗi, hãy kiểm tra log của container bằng lệnh:

```
docker compose logs nextapp
```

- Đảm bảo không có ứng dụng nào khác trên máy sử dụng cổng 3000 trước khi chạy lệnh. Nếu có, bạn cần thay đổi cổng trong file `docker-compose.yaml`.
