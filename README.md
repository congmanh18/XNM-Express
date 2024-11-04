# XNM-Express

XNM-Express is a web application that allows users to manage their express deliveries.

## Features

- [ ] User Authentication
- [ ] User Dashboard
- [ ] Admin Dashboard
- [ ] Tracking
- [ ] Reporting

## Technologies

- [ ] Golang: Programming language
- [ ] Docker: Containerization
- [ ] Redis: In-memory data structure store
- [ ] PostgreSQL: Object-relational database


## Overview of the services

1. **Bảng điều khiển dành cho quản trị viên**: Quản lý các chức năng quản trị, cho phép quản trị viên giám sát và quản lý hoạt động của hệ thống.
2. **Nhật ký kiểm tra**: Theo dõi và ghi lại tất cả các hành động và sự kiện quan trọng trong hệ thống để giải trình và khắc phục sự cố.
3. **Xác thực**: Xử lý xác thực người dùng, ủy quyền và quản lý phiên để bảo mật ứng dụng.
4. **Theo dõi giao hàng**: Quản lý theo dõi và cập nhật theo thời gian thực cho các lần giao hàng đang diễn ra.
5. **Hàng tồn kho**: Kiểm soát và giám sát lượng mặt hàng có sẵn để giao.
6. **Thông báo**: Quản lý việc gửi thông báo đến người dùng, bao gồm cập nhật trạng thái đơn hàng và tin nhắn khuyến mãi.
7. **Quản lý đơn hàng**: Chịu trách nhiệm xử lý và quản lý đơn hàng, từ khi tạo đến khi hoàn thành.
8. **Thanh toán**: Quản lý các giao dịch thanh toán và tích hợp với các cổng thanh toán.
9. **Báo cáo**: Cung cấp dữ liệu và báo cáo phân tích để theo dõi hiệu suất và ra quyết định.
10. **Tìm kiếm**: Kích hoạt chức năng tìm kiếm trên các đơn hàng, khoảng không quảng cáo và người dùng.
11. **Hồ sơ người dùng**: Quản lý thông tin tài khoản người dùng và tùy chọn.

## Architecture

Dự án này được xây dựng dưới dạng một tập hợp các dịch vụ độc lập giao tiếp qua HTTP/gRPC.

### Công nghệ
- **Golang**: Ngôn ngữ chính cho tất cả dịch vụ.
- **Docker**: Container hóa các dịch vụ.
- **gRPC/HTTP**: Giao tiếp giữa các dịch vụ.

## Getting Started

1. Clone the repository.
2. Set up the necessary environment variables.
3. Run each service individually or use Docker to start the entire stack.

## Contributing

Please follow the instructions in [CONTRIBUTING.md](CONTRIBUTING.md).

## License

This project is licensed under the MIT License.