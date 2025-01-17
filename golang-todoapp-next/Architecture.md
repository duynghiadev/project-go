[video](https://www.youtube.com/watch?v=pTwtTtRcmLo)

---

## Flow app work

![1737014591048](image/structure/1737014591048.png)

---

**Kiến trúc dự án**

Dự án được triển khai trên nền tảng Docker và bao gồm ba thành phần chính:

1. **Frontend (Next.js)** :

- Đây là giao diện người dùng, được phát triển bằng Next.js.
- Thành phần này chạy trên một container riêng, sử dụng hình ảnh Docker được xây dựng từ file `next.dockerfile`.
- Giao diện này giao tiếp với backend (API) để hiển thị và quản lý dữ liệu.
- Cổng giao tiếp: **3000** (host).

2. **Backend (Golang)** :

- Đây là phần API xử lý logic của ứng dụng, được viết bằng ngôn ngữ Golang.
- Backend chạy trên một container riêng, sử dụng hình ảnh Docker được xây dựng từ file `go.dockerfile`.
- Thành phần này kết nối với cơ sở dữ liệu Postgres để truy xuất và lưu trữ dữ liệu.
- Cổng giao tiếp: **8000** (host).

3. **Database (Postgres)** :

- Cơ sở dữ liệu được triển khai bằng Postgres thông qua Docker Hub.
- Cơ sở dữ liệu này được sử dụng để lưu trữ dữ liệu cần thiết của ứng dụng (ví dụ: thông tin người dùng, bài viết, hoặc dữ liệu todo).
- Container Postgres được ánh xạ với volume để đảm bảo dữ liệu không bị mất khi container dừng hoặc bị xóa.
- Cổng giao tiếp: **5432** (host).

**Docker Image và Container:**

- Mỗi thành phần được đóng gói trong một **Docker Image** riêng biệt và được chạy trên các container tương ứng.
- Các container này tương tác với nhau thông qua mạng nội bộ Docker.

**Quy trình hoạt động:**

- Khi bạn chạy lệnh `docker compose up --build -d`, Docker sẽ:
  - Xây dựng và khởi động container cho từng thành phần (Frontend, Backend, Database).
  - Ánh xạ các cổng của container với cổng trên máy chủ host.
- Sau khi khởi động:
  - Bạn có thể truy cập giao diện ứng dụng tại [http://localhost:3000](http://localhost:3000).
  - Backend xử lý yêu cầu từ Frontend và giao tiếp với Database để trả về dữ liệu cần thiết.
