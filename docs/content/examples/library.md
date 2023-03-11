Creating Library Backend
====

### Overview

Let's assume that we want to crate a library backend. We will store Books and Authors information

Author

| Property   | Type    | Flags    |
|------------|---------|----------|
| Full Name  | String  | required |
| Country    | String  |          |
| BestSeller | Boolean |          |

Book

| Property     | Type      | Flags              |
|--------------|-----------|--------------------|
| Name         | String    | required           |
| ISBN         | String    | required           |
| Author       | Reference | [Author], required |
| Release Year | INT32     |                    |

### Creation of Resources

Now let's create resources for given structures

First we need to prepare yaml file and describe those resources

```yml
name: Author
sourceConfig:
    dataSource: default
    entity: author
properties:
    - name: fullName
      type: STRING
    - name: country
      type: STRING
    - name: bestSeller
      type: BOOLEAN
---
name: Book
sourceConfig:
  dataSource: default
  entity: author
properties:
  - name: fullName
    type: STRING
  - name: country
    type: STRING
  - name: bestSeller
    type: BOOLEAN
```

### Now let's create resources from yaml definition file
```shell
dhctl aply -f resources.yml
```

