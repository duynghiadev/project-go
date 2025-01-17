Đây là cấu trúc Docker cho một ứng dụng bao gồm ba phần: frontend (Next.js), backend (Go), và cơ sở dữ liệu (PostgreSQL). Mỗi phần đều có Dockerfile riêng biệt để xây dựng và chạy trong các container riêng biệt. Dưới đây là phân tích chi tiết về từng phần.

### 1. **Dockerfile cho Go (Backend)**

- **Base image** : Sử dụng `golang:1.23.4-alpine3.20` là phiên bản Go với hệ điều hành Alpine, giúp giảm kích thước ảnh Docker.
- **WORKDIR** : Đặt thư mục làm việc trong container là `/app`. Mọi lệnh tiếp theo sẽ được thực hiện trong thư mục này.
- **COPY** : Sao chép tất cả các file từ thư mục hiện tại trên máy chủ vào thư mục `/app` trong container.
- **RUN go get** : Cài đặt tất cả dependencies Go cho dự án.
- **RUN go build** : Biên dịch mã nguồn Go và tạo file thực thi `todo`.
- **EXPOSE** : Mở cổng 8000 để container có thể nhận và gửi dữ liệu.
- **CMD** : Lệnh để chạy ứng dụng Go, khởi động file thực thi `todo`.

### 2. **Dockerfile cho Next.js (Frontend)**

Dockerfile này sử dụng multi-stage builds để giảm kích thước ảnh Docker và tối ưu hóa quá trình build:

- **Base image** : Sử dụng `node:20-alpine` làm base image cho tất cả các giai đoạn, giúp xây dựng môi trường cho ứng dụng Node.js.
- **Stage 1: `deps` (Install dependencies)** :
- Cài đặt một số package cần thiết như `libc6-compat`.
- Sao chép file `package.json` vào container và cài đặt các dependencies của frontend.
- **Stage 2: `builder` (Build the application)** :
- Sao chép thư mục `node_modules` từ giai đoạn `deps`.
- Sao chép toàn bộ mã nguồn vào container và chạy lệnh `npm run build` để biên dịch ứng dụng Next.js cho môi trường sản xuất.
- **Stage 3: `runner` (Production image)** :
- Đây là image cuối cùng để chạy ứng dụng Next.js trong môi trường sản xuất.
- Thiết lập các biến môi trường như `NODE_ENV=production`, chỉ định cổng `PORT=3000`, và cho phép ứng dụng lắng nghe trên tất cả các IP với `HOSTNAME=0.0.0.0`.
- Tạo user và group `nextjs` để chạy ứng dụng với quyền hạn hạn chế.
- Sao chép các file cần thiết từ giai đoạn `builder` vào thư mục thích hợp trong container.
- Đặt user `nextjs` để chạy ứng dụng với quyền hạn hạn chế.

### 3. **docker-compose.yml**

File `docker-compose.yml` giúp bạn quản lý các container cho frontend, backend, và cơ sở dữ liệu trong một môi trường Docker dễ dàng. Các dịch vụ được định nghĩa như sau:

- **nextapp (Frontend)** :
- `build`: Chỉ định context (thư mục chứa mã nguồn frontend) và Dockerfile (`next.dockerfile`) để xây dựng image.
- `ports`: Mở cổng 3000 của container và ánh xạ với cổng 3000 trên máy chủ để truy cập ứng dụng Next.js.
- `depends_on`: Chỉ định rằng frontend cần phải đợi backend (Go) khởi động trước khi bắt đầu.
- **golang (Backend)** :
- `build`: Chỉ định context (thư mục chứa mã nguồn backend) và Dockerfile (`go.dockerfile`) để xây dựng image.
- `environment`: Cấu hình biến môi trường cho PostgreSQL.
- `ports`: Mở cổng 8000 của container và ánh xạ với cổng 8000 trên máy chủ để truy cập ứng dụng Go.
- `depends_on`: Chỉ định rằng backend cần phải đợi cơ sở dữ liệu (PostgreSQL) khởi động trước khi bắt đầu.
- **db (PostgreSQL)** :
- `image`: Sử dụng image PostgreSQL với phiên bản `16.3-alpine3.20`.
- `environment`: Cấu hình mật khẩu, tên người dùng và tên cơ sở dữ liệu cho PostgreSQL.
- `ports`: Mở cổng 5432 của container và ánh xạ với cổng 5432 trên máy chủ để truy cập cơ sở dữ liệu.
- `volumes`: Đảm bảo dữ liệu của PostgreSQL được lưu trữ trong volume Docker (`todopostgres`) để duy trì dữ liệu giữa các lần khởi động lại container.

### 4. **Tổng kết**

- **Frontend (Next.js)** và **Backend (Go)** được xây dựng và chạy trong các container riêng biệt, giúp bạn quản lý và triển khai chúng độc lập.
- **Docker Compose** kết hợp tất cả các container lại với nhau, đảm bảo rằng frontend có thể giao tiếp với backend và backend có thể kết nối với cơ sở dữ liệu PostgreSQL.
- **Multi-stage builds** trong Dockerfile cho Next.js giúp giảm kích thước ảnh Docker và chỉ chứa những gì cần thiết cho môi trường sản xuất.

Cách cấu hình này giúp bạn dễ dàng triển khai ứng dụng với các phần tách biệt, tối ưu hóa cho môi trường sản xuất và giúp quản lý các dịch vụ trong một môi trường containerized.
