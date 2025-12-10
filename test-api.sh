#!/bin/bash

BASE_URL="http://localhost:8080"

echo "=== Testing Go JWT Tasks API ==="
echo ""

# Test health check
echo "1. Testing Health Check..."
curl -s $BASE_URL/health | jq .
echo -e "\n"

# Register a new user
echo "2. Registering a new user..."
REGISTER_RESPONSE=$(curl -s -X POST $BASE_URL/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "email": "test@example.com",
    "password": "password123"
  }')

echo $REGISTER_RESPONSE | jq .
TOKEN=$(echo $REGISTER_RESPONSE | jq -r '.token')
echo "Token: $TOKEN"
echo -e "\n"

# Login
echo "3. Logging in..."
LOGIN_RESPONSE=$(curl -s -X POST $BASE_URL/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123"
  }')

echo $LOGIN_RESPONSE | jq .
TOKEN=$(echo $LOGIN_RESPONSE | jq -r '.token')
echo "Token: $TOKEN"
echo -e "\n"

# Get profile
echo "4. Getting user profile..."
curl -s -X GET $BASE_URL/api/v1/profile \
  -H "Authorization: Bearer $TOKEN" | jq .
echo -e "\n"

# Create a task
echo "5. Creating a task..."
TASK_RESPONSE=$(curl -s -X POST $BASE_URL/api/v1/tasks \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "title": "Complete Go project",
    "description": "Finish the JWT Tasks API"
  }')

echo $TASK_RESPONSE | jq .
TASK_ID=$(echo $TASK_RESPONSE | jq -r '.id')
echo "Task ID: $TASK_ID"
echo -e "\n"

# Get all tasks
echo "6. Getting all tasks..."
curl -s -X GET $BASE_URL/api/v1/tasks \
  -H "Authorization: Bearer $TOKEN" | jq .
echo -e "\n"

# Get task by ID
echo "7. Getting task by ID..."
curl -s -X GET $BASE_URL/api/v1/tasks/$TASK_ID \
  -H "Authorization: Bearer $TOKEN" | jq .
echo -e "\n"

# Update task
echo "8. Updating task..."
curl -s -X PUT $BASE_URL/api/v1/tasks/$TASK_ID \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "title": "Updated Task Title",
    "completed": true
  }' | jq .
echo -e "\n"

# Get updated task
echo "9. Getting updated task..."
curl -s -X GET $BASE_URL/api/v1/tasks/$TASK_ID \
  -H "Authorization: Bearer $TOKEN" | jq .
echo -e "\n"

# Delete task
echo "10. Deleting task..."
curl -s -X DELETE $BASE_URL/api/v1/tasks/$TASK_ID \
  -H "Authorization: Bearer $TOKEN" | jq .
echo -e "\n"

# Verify task deleted
echo "11. Verifying task deleted..."
curl -s -X GET $BASE_URL/api/v1/tasks \
  -H "Authorization: Bearer $TOKEN" | jq .
echo -e "\n"

echo "=== Tests Complete ==="
