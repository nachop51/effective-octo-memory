# Testing Guide

This directory contains all tests for the Effective Octo Memory REST API project, organized separately from the source code for better maintainability and clarity.

## Test Structure

```
tests/
├── unit/                    # Unit tests (isolated, fast)
│   ├── users/              # User domain unit tests
│   │   └── service_test.go
│   └── accounts/           # Account domain unit tests
├── integration/            # Integration tests (with database)
│   └── users_integration_test.go
├── mocks/                  # Mock implementations
│   └── user_repository_mock.go
├── helpers/                # Test utilities and helpers
│   └── test_helpers.go
├── config/                 # Test configuration
│   └── test_config.go
└── README.md              # This file
```

## Test Types

### Unit Tests (`./unit/`)
- **Purpose**: Test individual components in isolation
- **Scope**: Single functions, methods, or small units of code
- **Dependencies**: Use mocks for external dependencies
- **Speed**: Fast (< 1 second per test)
- **Database**: No database required

### Integration Tests (`./integration/`)
- **Purpose**: Test complete workflows and API endpoints
- **Scope**: Full HTTP request/response cycles
- **Dependencies**: Real database, full application stack
- **Speed**: Slower (few seconds per test)
- **Database**: Requires test database

## Running Tests

### Prerequisites
- Go 1.24.4 or higher
- PostgreSQL (for integration tests)

### Quick Commands

```bash
# Run all tests
make test

# Run only unit tests (fast)
make test-unit

# Run only integration tests
make test-integration

# Run tests with coverage
make test-coverage

# Run tests with coverage in console
make test-coverage-console

# Watch for changes and run tests
make test-watch
```

### Manual Commands

```bash
# Unit tests
go test -v ./tests/unit/...

# Integration tests
go test -v ./tests/integration/...

# All tests with coverage
go test -v -coverprofile=coverage.out ./tests/...
go tool cover -html=coverage.out -o coverage.html
```

## Test Database Setup

For integration tests, you need a test database. Set the following environment variables:

```bash
export TEST_DATABASE_URL="postgres://user:password@localhost:5432/test_db?sslmode=disable"
```

Or use individual variables:
```bash
export TEST_DB_HOST="localhost"
export TEST_DB_USERNAME="user"
export TEST_DB_PASSWORD="password"
export TEST_DB_NAME="test_db"
```

### Docker Setup (Recommended)

```bash
# Start PostgreSQL for testing
docker run --name postgres-test -e POSTGRES_PASSWORD=password -e POSTGRES_USER=user -e POSTGRES_DB=test_db -p 5432:5432 -d postgres:15

# Run tests
make test-integration
```

## Writing Tests

### Unit Test Example

```go
func TestUserService_CreateUser_Success(t *testing.T) {
    // Arrange
    mockRepo := new(mocks.MockUserRepository)
    service := users.NewUserService(mockRepo, []byte("test-secret"))
    
    userBody := helpers.CreateUserBody()
    mockRepo.On("CreateUser", mock.AnythingOfType("*users.User")).Return(nil)
    
    // Act
    result, err := service.CreateUser(userBody)
    
    // Assert
    assert.NoError(t, err)
    assert.NotNil(t, result)
    mockRepo.AssertExpectations(t)
}
```

### Integration Test Example

```go
func (suite *UserIntegrationTestSuite) TestCreateUser_Success() {
    // Arrange
    userBody := helpers.CreateUserBody()
    jsonBody, _ := json.Marshal(userBody)
    
    // Act
    resp, err := http.Post(
        fmt.Sprintf("%s/users", suite.server.URL),
        "application/json",
        bytes.NewBuffer(jsonBody),
    )
    
    // Assert
    suite.Require().NoError(err)
    suite.Equal(http.StatusCreated, resp.StatusCode)
}
```

## Test Utilities

### Helpers (`./helpers/`)
- `CreateTestUser()`: Creates a test user with default values
- `CreateUserBody()`: Creates a test request body
- `TestJWTSecret()`: Returns consistent JWT secret for testing

### Mocks (`./mocks/`)
- `MockUserRepository`: Mock implementation of UserRepository interface
- Uses `testify/mock` for behavior verification

## Best Practices

### Unit Tests
1. **Test one thing at a time**
2. **Use descriptive test names** (Given_When_Then pattern)
3. **Follow AAA pattern** (Arrange, Act, Assert)
4. **Mock external dependencies**
5. **Test both success and error cases**

### Integration Tests
1. **Use test suites** for setup/teardown
2. **Clean database** between tests
3. **Test complete workflows**
4. **Use real HTTP requests**
5. **Verify database state**

### General Guidelines
1. **Keep tests independent** - No test should depend on another
2. **Use meaningful assertions** - Check specific values, not just existence
3. **Test error cases** - Don't just test happy paths
4. **Use test helpers** - Reduce duplication with helper functions
5. **Mock external services** - Don't make real API calls in tests

## Test Coverage

Aim for:
- **Unit tests**: 80%+ coverage
- **Integration tests**: Cover all API endpoints
- **Critical paths**: 100% coverage for authentication, validation, etc.

```bash
# Check current coverage
make test-coverage-console
```

## Continuous Integration

Tests are automatically run on:
- Pull requests
- Main branch pushes
- Release builds

Make sure all tests pass before merging code.

## Troubleshooting

### Common Issues

1. **Database connection errors**
   - Check if PostgreSQL is running
   - Verify connection string
   - Ensure test database exists

2. **Test failures after code changes**
   - Update mocks if interfaces changed
   - Check test data setup
   - Verify assertions are still valid

3. **Slow tests**
   - Use unit tests for fast feedback
   - Optimize database queries
   - Consider test parallelization

### Getting Help

1. Check test output for specific error messages
2. Use `go test -v` for verbose output
3. Run individual test files: `go test -v ./tests/unit/users/service_test.go`
4. Check mock expectations with `mockRepo.AssertExpectations(t)`

## Adding New Tests

1. **For new domain**: Create `tests/unit/[domain]/` directory
2. **For new service**: Add `[service]_test.go` file
3. **For new endpoint**: Add integration test case
4. **For new interface**: Create mock in `tests/mocks/`

Remember: Good tests are an investment in code quality and development speed!