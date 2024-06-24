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
