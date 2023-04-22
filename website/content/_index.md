---
title: About API Brew
linkTitle: About
menu: { main: { weight: 10 } }
---

{{% blocks/cover title="About API Brew" height="auto" %}}

Apibrew is an open-source tool that allows you to quickly and easily create CRUD APIs for your data. With Apibrew, you
define your schema in a declarative way, and the tool generates the corresponding APIs for you, minimizing the amount of
coding you have to do.

Apibrew is designed to be low-code, which means that you can create APIs without writing a lot of code. It is also
database agnostic and supports various database platforms, including PostgreSQL, MongoDB, MySQL, Redis, and more.
Apibrew supports multiple databases, so you can define multiple databases and perform operations on them.

{{% /blocks/cover %}}

{{% blocks/section color="white" %}}

## Quick Example

country.yml
```yaml
type: resource
name: country
properties:
  - name: name # name of the property
    type: STRING # type of the property
    length: 255
    required: true
    unique: true
  - name: description # name of the property
    type: STRING # type of the property
    length: 255
```

```bash
dhctl apply -f country.yml
```
So you are ready, you have fully established Rest API for country resource

```javascript
axios.post('http://localhost:9009/country', {
  name: 'Country1',
  description: 'Sample Country 1'
})
```

You can build entire application with resources and references between them (like relations in relational databases)

Everything can be written by resources, not?
**Let's extend our country resource**

```javascript
countryExtension.onCreate(async (country) => {
        country.description = country.description + ' - Extended'

        return country
    })
```
So we have extended our country resource with modification description on create

So, with **API Brew**, you can create your application with resources and you can customize behavior of your resources with extensions


{{% /blocks/section %}}
