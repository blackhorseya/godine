# GoDine

[![Go](https://github.com/blackhorseya/godine/workflows/Go/badge.svg)](https://github.com/blackhorseya/godine/actions?query=workflow:"Go")
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=blackhorseya_godine&metric=coverage)](https://sonarcloud.io/summary/new_code?id=blackhorseya_godine)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=blackhorseya_godine&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=blackhorseya_godine)
[![GitHub release](https://img.shields.io/github/release/blackhorseya/godine?include_prereleases=&sort=semver&color=blue)](https://github.com/blackhorseya/godine/releases/)
![GitHub License](https://img.shields.io/github/license/blackhorseya/godine)

GoDine is an online food ordering system designed using Golang and following the principles of Domain-Driven Design (
DDD). The system includes multiple business domains such as User Management, Restaurant Management, Order Management,
Payment Management, and Notification Management. Each domain has clear boundaries and interacts with other domains to
provide a seamless online dining experience.

## Project Name Explanation

The name GoDine is derived from combining "Golang" and "Dine", representing an online food ordering system built using
Golang. This name is concise and representative, conveying the core functionality of the project, which is to provide
efficient online dining services using Golang technology.

## Features

- **User Management**: User registration, login, profile updates, account deletion.
- **Restaurant Management**: Restaurant information management, menu management.
- **Order Management**: Order creation, order status updates, order queries.
- **Payment Management**: Handling order payments, payment record queries.
- **Notification Management**: Sending notifications to users (e.g., order status updates).

## Technical Details

- **Programming Language**: Golang
- **Architecture Pattern**: Domain-Driven Design (DDD)
- **Major Modules**:
    - **User Management**: Responsible for user registration, login, profile management, etc.
    - **Restaurant Management**: Responsible for restaurant information and menu management.
    - **Order Management**: Handles order creation, management, and tracking.
    - **Payment Management**: Processes payments and manages payment records.
    - **Notification Management**: Sends notifications regarding order status changes to users.

## System Architecture Diagram

```mermaid
graph TD
    UserManagement[用戶管理]
    RestaurantManagement[餐廳管理]
    OrderManagement[訂單管理]
    PaymentManagement[支付管理]
    NotificationManagement[通知管理]
    LogisticsManagement[物流管理]

    UserManagement -->|確認用戶身份| OrderManagement
    RestaurantManagement -->|確認餐廳和菜單可用性| OrderManagement
    OrderManagement -->|確認訂單和金額| PaymentManagement
    PaymentManagement -->|通知支付結果| OrderManagement
    OrderManagement -->|通知訂單狀態變更| NotificationManagement
    OrderManagement -->|安排配送| LogisticsManagement
    LogisticsManagement -->|更新配送狀態| OrderManagement
    LogisticsManagement -->|通知配送狀態變更| NotificationManagement

    subgraph UserManagement[用戶管理]
        UM[User Service]
        UM --> U[User]
    end

    subgraph RestaurantManagement[餐廳管理]
        RM[Restaurant Service]
        M[Menu Service]
        RM --> R[Restaurant]
        M --> MI[MenuItem]
    end

    subgraph OrderManagement[訂單管理]
        OM[Order Service]
        OM --> O[Order]
        O --> OI[OrderItem]
    end

    subgraph PaymentManagement[支付管理]
        PM[Payment Service]
        PM --> PR[PaymentRecord]
    end

    subgraph NotificationManagement[通知管理]
        NM[Notification Service]
        NM --> N[Notification]
    end

    subgraph LogisticsManagement[物流管理]


        LM[Logistics Service]
        LM --> D[Delivery]
        D --> DS[DeliveryStatus]
    end
```

## Installation

1. **Clone the repository**:
    ```bash
    git clone https://github.com/blackhorseya/godine.git
    ```
2. **Navigate to the project directory**:
    ```bash
    cd godine
    ```
3. **Install dependencies**:
    ```bash
    go mod tidy
    ```

## Usage

1. **Run the service**:
    ```bash
    go run .
    ```
2. **Access the service**:
    - API documentation is available at `http://localhost:8080/swagger/index.html`.

## Contribution

We welcome contributions to this project. Please follow these steps:

1. **Fork the repository**
2. **Create your feature branch**:
    ```bash
    git checkout -b feature-branch
    ```
3. **Commit your changes**:
    ```bash
    git commit -am 'Add some feature'
    ```
4. **Push to the branch**:
    ```bash
    git push origin feature-branch
    ```
5. **Create a Pull Request**

## License

This project is licensed under the GPL-3.0 License. See the [LICENSE](LICENSE) file for details.

```

This README file provides a comprehensive overview of the GoDine project, including its name explanation, features, technical details, system architecture diagram, installation instructions, usage guidelines, contribution steps, and licensing information.