#!/bin/bash

# request GET to home
curl http://localhost:8080/

# request POST to tasks
curl -X POST -d '{"id":3, "task": "task1", "isCompleted": false}' http://localhost:8080/tasks

# request GET to tasks
curl -X GET -d '{"id":3}' http://localhost:8080/tasks

# request PUT to tasks
curl -X PUT -d '{"task":"homework,hurry","isCompleted":true}' http://localhost:8080/tasks/3

# request DELETE to tasks
curl -X DELETE http://localhost:8080/tasks/3