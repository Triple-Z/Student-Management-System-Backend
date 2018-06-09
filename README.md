# Student Management

<!-- TOC -->

- [Student Management](#student-management)
    - [Quick Start](#quick-start)
    - [API Reference](#api-reference)
        - [Students](#students)
            - [获取所有学生信息](#获取所有学生信息)
                - [Response](#response)
            - [新建学生](#新建学生)
                - [Parameters](#parameters)
                - [Response](#response-1)
            - [获取某个学生信息](#获取某个学生信息)
                - [Response](#response-2)
            - [更新某个学生信息](#更新某个学生信息)
                - [Parameters](#parameters-1)
                - [Response](#response-3)
            - [删除某个学生](#删除某个学生)
                - [Response](#response-4)
        - [Departments](#departments)
            - [获取所有学院信息](#获取所有学院信息)
                - [Response](#response-5)
            - [新建学院](#新建学院)
                - [Parameters](#parameters-2)
                - [Response](#response-6)
            - [获取某个学院信息](#获取某个学院信息)
                - [Response](#response-7)
            - [更新某个学院信息](#更新某个学院信息)
                - [Parameters](#parameters-3)
                - [Response](#response-8)
            - [删除某个学院](#删除某个学院)
                - [Response](#response-9)

<!-- /TOC -->

## Quick Start

```bash
$ git clone https://github.com/Triple-Z/Student-Management-System-Backend.git
$ cd Student-Management-System-Backend/
$ go run main.go
```

## API Reference

API Root: http://localhost:3000

### Students

#### 获取所有学生信息

```api
GET /api/v1/students
```

##### Response

```text
Status: 200 OK
```

```json
{
    "count": 3,
    "data": [
        {
            "Id": {
                "Int64": 3,
                "Valid": true
            },
            "Number": {
                "String": "161540121",
                "Valid": true
            },
            "Name": {
                "String": "赵祯真",
                "Valid": true
            },
            "Department_id": {
                "Int64": 1,
                "Valid": true
            },
            "Department_name": {
                "String": "计算机科学与技术学院",
                "Valid": true
            },
            "Create_date": "2018-06-09T15:01:04Z",
            "Last_updated": "2018-06-09T15:00:58Z"
        },
        {
            "Id": {
                "Int64": 7,
                "Valid": true
            },
            "Number": {
                "String": "161540110",
                "Valid": true
            },
            "Name": {
                "String": "lewis",
                "Valid": true
            },
            "Department_id": {
                "Int64": 1,
                "Valid": true
            },
            "Department_name": {
                "String": "计算机科学与技术学院",
                "Valid": true
            },
            "Create_date": "2018-06-09T15:01:04Z",
            "Last_updated": "2018-06-09T19:32:49Z"
        },
        {
            "Id": {
                "Int64": 10,
                "Valid": true
            },
            "Number": {
                "String": "161540119",
                "Valid": true
            },
            "Name": {
                "String": "斳森",
                "Valid": true
            },
            "Department_id": {
                "Int64": 1,
                "Valid": true
            },
            "Department_name": {
                "String": "计算机科学与技术学院",
                "Valid": true
            },
            "Create_date": "2018-06-09T15:01:04Z",
            "Last_updated": "2018-06-09T15:00:58Z"
        }
    ],
    "status": "ok"
}
```

#### 新建学生

```api
POST /api/v1/students
```

##### Parameters

| Name | Type | Description |
|:----: | :----: | :----: |
| `number` | `string` | 学号|
| `name` |`string` | 姓名|
| `department_id`| `int` | 学院 ID |\

##### Response

```text
Status: 200 OK
```

```json
{
    "new_id": 13,
    "status": "ok"
}
```

#### 获取某个学生信息

```api
GET /api/v1/students/:id
```

##### Response

```text
Status: 200 OK
```

```json
{
    "data": {
        "Id": {
            "Int64": 3,
            "Valid": true
        },
        "Number": {
            "String": "161540121",
            "Valid": true
        },
        "Name": {
            "String": "赵祯真",
            "Valid": true
        },
        "Department_id": {
            "Int64": 1,
            "Valid": true
        },
        "Department_name": {
            "String": "计算机科学与技术学院",
            "Valid": true
        },
        "Create_date": "2018-06-09T15:01:04Z",
        "Last_updated": "2018-06-09T15:00:58Z"
    },
    "status": "ok"
}
```

#### 更新某个学生信息

```api
PUT /api/v1/students/:id
```

##### Parameters

| Name |Type| Description |
|:----: | :----: | :----: |
| `number` | `string` | 学号|
| `name` | `string` | 姓名|
| `department_id` |`int` | 学院 ID|

##### Response

```text
Status: 200 OK
```

```json
{
    "status": "ok",
    "updated_id": "13"
}
```

#### 删除某个学生

```api
DELETE /api/v1/students/:id
```

##### Response

删除成功返回 `204` ：

```text
Status: 204 No Content
```

删除失败则返回 `404` ：

```text
Status: 404 Not Found
```

```json
{
    "message": "Cannot find this student by id: 13",
    "rows_affected": null,
    "status": "failed"
}
```

### Departments

#### 获取所有学院信息

```api
GET /api/v1/departments
```

##### Response

```text
Status: 200 OK
```

```json
{
    "count": 3,
    "data": [
        {
            "Id": {
                "Int64": 1,
                "Valid": true
            },
            "Name": {
                "String": "计算机科学与技术学院",
                "Valid": true
            },
            "Create_date": "2018-06-09T14:47:53Z",
            "Last_updated": "2018-06-09T14:58:24Z"
        },
        {
            "Id": {
                "Int64": 2,
                "Valid": true
            },
            "Name": {
                "String": "航空宇航学院",
                "Valid": true
            },
            "Create_date": "2018-06-09T14:56:12Z",
            "Last_updated": "2018-06-09T14:56:12Z"
        },
        {
            "Id": {
                "Int64": 3,
                "Valid": true
            },
            "Name": {
                "String": "自动化学院",
                "Valid": true
            },
            "Create_date": "2018-06-09T18:46:52Z",
            "Last_updated": "2018-06-09T20:18:24Z"
        }
    ],
    "status": "ok"
}
```

#### 新建学院

```api
POST /api/v1/departments
```

##### Parameters

| Name | Type | Description |
| :---: | :---: | :---: |
| `name` | `string` | 学院名 |

##### Response

```text
Status: 200 OK
```

```json
{
    "new_id": 5,
    "status": "ok"
}
```

#### 获取某个学院信息

```api
GET /api/v1/departments/:id
```

##### Response

```text
Status: 200 OK
```

```json
{
    "data": {
        "Id": {
            "Int64": 5,
            "Valid": true
        },
        "Name": {
            "String": "人工智能学院",
            "Valid": true
        },
        "Create_date": "2018-06-09T21:16:15Z",
        "Last_updated": "2018-06-09T21:16:15Z"
    },
    "status": "ok"
}
```

#### 更新某个学院信息

```api
PUT /api/v1/departments/:id
```

##### Parameters

| Name | Type | Description |
| :---: | :---: | :---: |
| `name` | `string` | 学院名 |

##### Response

```text
Status: 200 OK
```

```json
{
    "status": "ok",
    "updated_id": "5"
}
```

#### 删除某个学院

```api
DELETE /api/v1/departments/:id
```

##### Response

删除成功返回 `204` ：

```text
Status: 204 No Content
```

删除失败则返回 `404` ：

```text
Status: 404 Not Found
```

```json
{
    "message": "Cannot find this department by id: 5",
    "rows_affected": null,
    "status": "failed"
}
```