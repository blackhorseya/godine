## 0.4.6 (2024-07-06)

## 0.4.5 (2024-07-06)

### Feat

- improve service initialization and error handling (#71)
- improve server shutdown logging for multiple adapters
- implement graceful shutdown for server
- improve service initialization and error handling
- refactor command execution logic in main loop
- introduce new service listing functionality

### Refactor

- refactor signal handling and remove otelx references
- refactor and clean up code in multiple adapter packages
- improve error handling and signal handling
- refactor service commands in `start.go` and `cmdx.go`

## 0.4.4 (2024-07-06)

### Feat

- refactor configuration package in `configx` module
- add Configuration Field to Injector Struct (#68)
- improve error handling and dynamic config file path logic
- refactor configuration initialization across multiple files
- add Configuration Field to Injector Struct
- improve transaction handling and error logging (#67)
- improve transaction handling and error logging

### Fix

- refactor application initialization for configuration injection
- improve error handling and rollback in mariadb functions

### Refactor

- refactor codebase for improved efficiency
- refactor code to remove unnecessary variables
- refactor code to include `app *configx.Application` parameter
- refactor Configuration struct properties (#70)
- refactor Configuration struct properties
- update HTTP client functions to accept config parameter
- refactor configuration handling in external tests
- refactor configuration handling in wire setup
- refactor initialization process to use viper parameter
- refactor initialization process for better parameter passing
- refactor application configuration handling
- refactor initialization and imports in test files
- refactor configuration handling in cmdx package

## 0.4.3 (2024-07-06)

### Feat

- refactor snowflake node ID generation

### Refactor

- update order ID values throughout API tests

## 0.4.2 (2024-07-06)

### Feat

- refactor Order struct to include BigIntID field (#66)
- refactor Order struct to include BigIntID field
- implement snowflake node creation and ID generation
- add new field `NewID` to Order struct
- update `orderId` value in API endpoints

### Refactor

- refactor order model and error logging
- refactor order ID generation across files
- set default value for order.NewID when not provided

## 0.4.1 (2024-07-05)

### Feat

- update GORM implementation in Order struct (#63)
- update GORM implementation in Order struct
- implement database integration with MariaDB
- improve error handling and tracing throughout codebase
- consolidate default limit and offset handling
- improve error handling and testing for order functions
- improve logging functionality across the codebase
- improve transaction handling for order creation
- update order creation process in domain package
- refactor mariadb module for dependency injection (#60)
- implement logging and tracing in `mariadb.go`
- refactor test setup and configuration handling
- refactor MariaDB external tester configuration
- refactor mariadb module for dependency injection
- optimize MySQL client connection settings
- add support for MariaDB integration
- configure `mariadb` service and volume (#58)
- configure `mariadb` service and volume
- consolidate MySQL configurations in app.go

### Fix

- refactor return statements in BeforeSave and AfterFind
- improve error handling and import statements in Mariadb files

### Refactor

- optimize database queries in mariadb.go
- refactor struct tags for Order and OrderItem models (#62)
- improve error handling and context in List function
- refactor struct tags for Order and OrderItem models
- refactor order repository to use gorm instead of sqlx

## 0.4.0 (2024-07-04)

### Feat

- refactor database schema and model initialization

### Refactor

- update database column types to BIGINT

## 0.3.7 (2024-07-04)

### Feat

- update order and user information in API endpoints

### Perf

- optimize MongoDB query performance

## 0.3.6 (2024-07-03)

### Feat

- improve restaurant data handling and caching

### Refactor

- convert string IDs to ObjectIDs for MongoDB queries (#56)
- refactor code to use ObjectID for filter operations
- refactor UnmarshalBSON functions in multiple models
- convert string IDs to ObjectIDs for MongoDB queries
- refactor ID generation and imports in MongoDB model files

## 0.3.5 (2024-07-03)

### Feat

- implement caching functionality for restaurant data in Redis (#55)
- implement caching functionality for restaurant data in Redis
- integrate Redis service and configurations across services (#54)
- integrate Redis service and configurations across services
- refactor Redis configuration and dependencies (#52)
- refactor Redis configuration and dependencies

### Refactor

- refactor error handling and logging logic across multiple files
- update IDs in test and request files

## 0.3.4 (2024-07-02)

### Feat

- improve network performance and reliability
- optimize connection pool configuration in NewClientWithDSN function
- improve API test suite and assertions

## 0.3.3 (2024-07-02)

### Feat

- optimize connection pool configuration in NewClientWithDSN function

## 0.3.2 (2024-07-02)

### Refactor

- update data type of X-Total-Count header in API endpoints

## 0.3.1 (2024-06-29)

## 0.3.0 (2024-06-27)

### Feat

- improve error handling and data parsing in API requests
- refactor error handling and implement user list functionality in HTTP client (#47)
- refactor error handling and status code checks
- improve error handling and response parsing in DeleteUser method
- improve user update functionality with tracing and error handling
- refactor error handling and implement user list functionality in HTTP client

### Fix

- improve error handling in logistics HTTP client (#48)
- improve error handling in logistics HTTP client

### Refactor

- refactor notification service methods
- update error handling in logistics HTTP client files

## 0.2.3 (2024-06-27)

### Feat

- implement new CRUD endpoints and Swagger definitions (#46)
- refactor URL path concatenation using constant variable
- improve error handling and user status functionality
- update user information with error handling and return updated info
- implement new CRUD endpoints and Swagger definitions
- improve error handling and logging in UpdateMenuItem function (#45)
- improve error handling and logging in UpdateMenuItem function
- improve error handling and HTTP request in UpdateRestaurant function (#44)
- refactor restaurant HTTP client for error handling
- implement error handling and HTTP logic for DeleteRestaurant method
- improve error handling and HTTP request in UpdateRestaurant function
- refactor update and delete operations for restaurant model (#43)
- refactor update and delete operations for restaurant model
- improve restaurant entity management and error handling (#42)
- refactor restaurant API delete functionality
- improve restaurant entity management and error handling
- implement PUT and DELETE functionality for items API (#41)
- implement PUT and DELETE functionality for items API
- add restaurant api router
- improve error handling and logging in menu HTTP client (#40)
- improve error handling and logging in menu HTTP client
- implement new delivery status changed handler in biz directory
- consolidate logistics handling functionality (#39)
- refactor delivery status handling in logistics module
- consolidate logistics handling functionality
- implement dependency injection using Wire in handler package (#38)
- implement dependency injection using Wire in handler package
- improve logging in KafkaEventBus
- introduce a new function for creating writers with specified topics
- consolidate topic name usage in Kafka event bus
- integrate MQX event bus into logistics module
- introduce new event handling functionality (#37)
- integrate Kafka event bus into transport and update methods
- implement in-memory event bus with registration and publishing capabilities
- implement MQX transport for event bus integration
- introduce new event handling functionality
- consolidate event handling functionality
- refactor delivery status change methods across services
- implement Kafka message delivery for status updates
- update Kafka configuration and dependencies across files
- integrate Kafka configurations for new broker service

### Refactor

- improve context handling in user deletion process
- refactor user model references in http_client.go
- refactor order service methods and interfaces
- standardize error handling with `errorx` package
- standardize span naming conventions across files
- refactor restaurant API URL construction in http client
- standardize field name for page size across restaurant domain
- optimize imports across the codebase
- update `IsOpen` field in restaurant structs
- refactor event bus handling and logging in Kafka integration
- refactor event handling in logistics domain
- refactor event handling in delivery system
- refactor event handling to use topic instead of event type
- refactor error handling and logging in logistics and event bus modules
- refactor restaurant ID field in PostPayload struct
- refactor order handling interfaces
- refactor Kafka topic handling throughout the codebase

## 0.2.2 (2024-06-26)

### Feat

- refactor delivery status handling in logistics service
- update PATCH route and Swagger documentation in logistics API
- improve state transition handling in delivery process (#35)
- improve state transition handling in delivery process
- refactor default status handling in deliveries
- refactor model package and update dependencies
- improve error handling and notifications
- improve error handling and order retrieval in API functions
- refactor API handling for PatchWithStatus functionality
- update order status endpoint and documentation

### Refactor

- refactor delivery status handling in `delivery_state.go` files
- refactor delivery status handling in model logic

## 0.2.1 (2024-06-26)

### Feat

- refactor notification creation in order process
- refactor notification HTTP client functionality in Go codebase (#33)
- improve error handling and response parsing
- improve notification handling and error checking
- update notification API functionality (#32)
- update notification API functionality
- refactor notification HTTP client functionality in Go codebase
- implement RESTful API routing for notifications (#31)
- introduce new RESTful notifications API framework
- implement RESTful API routing for notifications
- implement error handling and result count logic across files
- refactor dependency management and notifications in MongoDB integration
- integrate new notification features across files
- introduce notifications functionality into the repo package (#28)
- implement MongoDB notification repository functionality
- introduce notifications functionality into the repo package
- create new notification adapter for wirex integration (#27)
- integrate new `notify_restful` service into compose configuration
- implement real-time notification feature with `notify` module
- integrate new NotifyRestful application configuration
- create new notification adapter for wirex integration
- consolidate notification business logic into separate package

### Fix

- improve error handling in GetByID function

### Refactor

- update naming convention for notification services
- refactor notification handling to use model.Notification struct
- refactor imports for notification module
- refactor UpdateStatus function and improve error handling
- remove references to notification biz from order adapter

## 0.2.0 (2024-06-25)

### Feat

- integrate OTel collector, Jaeger, and Prometheus services (#24)
- integrate OTel collector, Jaeger, and Prometheus services
- implement NewDelivery method and dependencies (#22)
- improve order creation and update functionality
- refactor order model to include delivery ID
- implement NewDelivery method and dependencies
- integrate logistics business logic into order module (#21)
- refactor error handling and response parsing in GetDelivery method
- consolidate dependencies and implement new method
- integrate logistics business logic into order module
- improve error handling and imports across files (#20)
- improve error handling and context in GetByID function
- improve error handling and documentation in Handle function
- improve error handling and imports across files
- introduce new deliveries API handlers and documentation (#19)
- improve delivery management API endpoints
- introduce new deliveries API handlers and documentation
- refactor filter conditions for driver_id in List function
- improve delivery service functionality (#18)
- improve delivery tracking functionality
- improve delivery service functionality
- implement distributed tracing with otelx package
- add dependencies and implement Create method in mongodb.go (#17)
- refactor MongoDB deliveries Delete method
- implement timeout and filter options in MongoDB list function
- add dependencies and implement Create method in mongodb.go
- introduce RESTful API for logistics service

### Fix

- improve database update operations and error handling
- refactor error handling in MongoDB operations

### Refactor

- improve code quality and readability in order creation process
- refactor code for improved efficiency
- refactor UpdateDeliveryStatus function

## 0.1.3 (2024-06-25)

### Feat

- consolidate new files for logistics/restful and wirex (#15)
- consolidate new files for logistics/restful and wirex
- integrate new LogisticsRestful application into Configuration struct
- consolidate Bazel build configuration for wirex module
- consolidate logistics domain files into biz directory (#14)
- integrate dependency injection with Wire framework
- consolidate logistics domain files into biz directory
- consolidate logistics business entity files
- consolidate delivery module files into separate directory
- create new structs for delivery entities and addresses (#11)
- create new structs for delivery entities and addresses

### Refactor

- refactor delivery methods in `ILogisticsBiz` interface (#13)
- refactor delivery methods in `ILogisticsBiz` interface

## 0.1.2 (2024-06-25)

### Feat

- implement error handling and list orders functionality (#10)
- refactor GetByID function for improved error handling and order information retrieval
- implement error handling and list orders functionality
- refactor API endpoints and Swagger documentation (#9)
- refactor API endpoints and Swagger documentation
- refactor tracing and error handling in order.go (#8)
- refactor tracing and error handling in order.go
- improve error handling and listing logic in functions (#6)
- improve error handling and listing logic in functions
- refactor order status handling and transitions (#5)
- enhance JSON serialization for Order struct
- implement Unmarshal methods for Order struct
- refactor order status handling and transitions
- refactor MongoDB integration in order service
- refactor order creation process (#4)
- enhance payload handling for orders API requests
- refactor order creation process
- update order API payloads and documentation
- refactor order creation process in `order.go`
- generate order items based on menu items and options
- consolidate restaurant menu endpoint logic (#2)
- consolidate restaurant menu endpoint logic
- refactor order business logic across multiple files (#1)
- refactor menuService implementation in NewOrderBiz function
- refactor order business logic across multiple files
- improve error handling and menu retrieval in API
- refactor naming conventions in codebase
- add user level field to API documentation
- refactor order creation process
- refactor HTTP client methods in biz library
- update user struct to include new level field
- update user retrieval functionality and Swagger documentation
- add RESTful API for user management in adapter user module
- integrate user-restful adapter into project
- consolidate user restful API and wirex implementations
- refactor user management and add `IsActive` field
- implement user-based filtering for update and delete operations
- refactor MongoDB list function for better error handling
- refactor user creation logic and data types
- implement Bazel build system and user repository interface
- refactor user methods to accept ID as string
- integrate RESTful order API versioning into system
- refactor order business logic and dependencies
- refactor error handling and API requests in restaurant HTTP client
- integrate UUID library for API endpoint retrieval
- update restaurant library with new HTTP client functions
- consolidate order domain and repository entity files

### Fix

- improve error handling and query options in user functions
- improve error handling for HTTP requests

### Refactor

- refactor return type of ListOrders to use pointers
- refactor order listing functions and structs (#7)
- refactor order listing functions and structs
- consolidate duplicate code in list functions
- refactor UpdateStatus method across multiple files
- refactor order state handling and add pending state functionality
- refactor API pagination parameters
- refactor parameter naming in CreateOrder function
- update data types from uuid.UUID to string across files (#3)
- update data types from uuid.UUID to string across files
- refactor order item initialization in CreateOrder function
- refactor user business logic across modules
- refactor user model and remove aggregate functionality
- refactor restaurant HTTP client functionality
- refactor data types and variable names in order models

## 0.1.1 (2024-06-13)

### Feat

- introduce Swagger documentation for order restful API
- refactor order wirex adapter configuration and build process
- add new POST method and documentation for menu items
- refactor API endpoints and update Swagger documentation
- create endpoint for posting restaurants and update swagger docs
- integrate v1 restful API for restaurant handling
- implement v1 restful API endpoints with Go_library targets
- improve error handling and update menu logic
- improve error handling and traceability in menu functionality
- refactor menuBiz struct and NewMenuBiz function
- refactor restaurant domain logic and models
- integrate MongoDB options for restaurant listing
- improve error handling and add timeouts to methods
- consolidate MongoDB package functionality
- create initial files for restaurant package and repo
- consolidate new functionality for the biz domain
- implement order-related methods and build file for `biz` package
- consolidate new files into biz directory
- implement CRUD operations for user management in IUserBiz interface
- consolidate business logic and structures
- consolidate new files for `model` package
- introduce new files to the `model` package
- implement restaurant model with necessary files and functions
- refactor health check endpoint and Swagger definitions
- refactor server start command using Cobra
- update Swagger documentation and add Bazel files for restaurant API
- implement dependency injection in restaurant adapter
- consolidate configuration files in app/infra/configx
- consolidate package structure into `pkg/cmdx` folder
- create new `adapterx` package with BUILD.bazel and adapterx.go files
- implement logging package with options and initialization
- implement custom error handling with errorx package

### Refactor

- refactor menu API and business logic
- improve error handling and logging for restaurant operations
- refactor restaurant aggregation logic in model files
- refactor struct field names and introduce new module in Bazel build

## 0.1.0 (2024-06-10)

### Feat

- consolidate netx package structure
- introduce new cmd target and go_binary in Bazel files
- update workspace definitions for new project integration
