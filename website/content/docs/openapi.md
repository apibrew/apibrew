---
title: Rest Api
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
highlight_theme: darkula
headingLevel: 2

---

<!-- Generator: Widdershins v4.0.1 -->

<h1 id="apibrew">API Brew v1.0</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

OpenApi 3.0 definition for API Brew Rest API

Email: <a href="mailto:talehsmail@gmail.com">Taleh Ibrahimli</a> Web: <a href="https://github.com/apibrew/apibrew">Taleh Ibrahimli</a> 
License: <a href="https://github.com/google/gnostic/blob/master/LICENSE">Apache License</a>

# Authentication

- HTTP Authentication, scheme: bearer 

<h1 id="apibrew-authentication">Authentication</h1>

Authentication Service is for authentication related operations

## authenticationRenewToken

<a id="opIdauthenticationRenewToken"></a>

> Code samples

```shell
# You can also use wget
curl -X PUT /authentication/token \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
PUT /authentication/token HTTP/1.1

Content-Type: application/json
Accept: application/json

```

```javascript
const inputBody = '{
  "token": "string",
  "term": "SHORT"
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('/authentication/token',
{
  method: 'PUT',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Content-Type' => 'application/json',
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.put '/authentication/token',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Content-Type': 'application/json',
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.put('/authentication/token', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Content-Type' => 'application/json',
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('PUT','/authentication/token', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/authentication/token");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("PUT");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Content-Type": []string{"application/json"},
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("PUT", "/authentication/token", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`PUT /authentication/token`

*Renew token*

Renew token with existing token

> Body parameter

```json
{
  "token": "string",
  "term": "SHORT"
}
```

<h3 id="authenticationrenewtoken-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[RenewTokenRequest](#schemarenewtokenrequest)|true|none|

> Example responses

> 200 Response

```json
{
  "token": {
    "term": "SHORT",
    "content": "string",
    "expiration": "2019-08-24T14:15:22Z"
  }
}
```

<h3 id="authenticationrenewtoken-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[RenewTokenResponse](#schemarenewtokenresponse)|
|default|Default|Default error response|[Status](#schemastatus)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

## authenticationAuthenticate

<a id="opIdauthenticationAuthenticate"></a>

> Code samples

```shell
# You can also use wget
curl -X POST /authentication/token \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
POST /authentication/token HTTP/1.1

Content-Type: application/json
Accept: application/json

```

```javascript
const inputBody = '{
  "username": "string",
  "password": "string",
  "term": "SHORT"
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('/authentication/token',
{
  method: 'POST',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Content-Type' => 'application/json',
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.post '/authentication/token',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Content-Type': 'application/json',
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.post('/authentication/token', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Content-Type' => 'application/json',
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('POST','/authentication/token', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/authentication/token");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("POST");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Content-Type": []string{"application/json"},
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("POST", "/authentication/token", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /authentication/token`

*Authentication with username/password*

Authentication with username/password and create new token

> Body parameter

```json
{
  "username": "string",
  "password": "string",
  "term": "SHORT"
}
```

<h3 id="authenticationauthenticate-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[AuthenticationRequest](#schemaauthenticationrequest)|true|none|

> Example responses

> 200 Response

```json
{
  "token": {
    "term": "SHORT",
    "content": "string",
    "expiration": "2019-08-24T14:15:22Z"
  }
}
```

<h3 id="authenticationauthenticate-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[AuthenticationResponse](#schemaauthenticationresponse)|
|default|Default|Default error response|[Status](#schemastatus)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

<h1 id="apibrew-datasource">DataSource</h1>

DataSource Service is for managing data sources

## DataSource_List

<a id="opIdDataSource_List"></a>

> Code samples

```shell
# You can also use wget
curl -X GET /system/data-sources \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
GET /system/data-sources HTTP/1.1

Accept: application/json

```

```javascript

const headers = {
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('/system/data-sources',
{
  method: 'GET',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.get '/system/data-sources',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.get('/system/data-sources', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('GET','/system/data-sources', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/system/data-sources");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("GET");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "/system/data-sources", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /system/data-sources`

*List*

<h3 id="datasource_list-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|token|query|string|false|none|

> Example responses

> 200 Response

```json
{
  "content": [
    {
      "id": "string",
      "backend": "POSTGRESQL",
      "name": "string",
      "description": "string",
      "postgresqlParams": {
        "username": "string",
        "password": "string",
        "host": "string",
        "port": 0,
        "dbName": "string",
        "defaultSchema": "string"
      },
      "mysqlParams": {
        "username": "string",
        "password": "string",
        "host": "string",
        "port": 0,
        "dbName": "string",
        "defaultSchema": "string"
      },
      "virtualParams": {
        "mode": "DISCARD"
      },
      "redisParams": {
        "addr": "string",
        "password": "string",
        "db": 0
      },
      "mongoParams": {
        "uri": "string",
        "dbName": "string"
      },
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}
```

<h3 id="datasource_list-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[ListDataSourceResponse](#schemalistdatasourceresponse)|
|default|Default|Default error response|[Status](#schemastatus)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

## DataSource_Update

<a id="opIdDataSource_Update"></a>

> Code samples

```shell
# You can also use wget
curl -X PUT /system/data-sources \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
PUT /system/data-sources HTTP/1.1

Content-Type: application/json
Accept: application/json

```

```javascript
const inputBody = '{
  "token": "string",
  "dataSources": [
    {
      "id": "string",
      "backend": "POSTGRESQL",
      "name": "string",
      "description": "string",
      "postgresqlParams": {
        "username": "string",
        "password": "string",
        "host": "string",
        "port": 0,
        "dbName": "string",
        "defaultSchema": "string"
      },
      "mysqlParams": {
        "username": "string",
        "password": "string",
        "host": "string",
        "port": 0,
        "dbName": "string",
        "defaultSchema": "string"
      },
      "virtualParams": {
        "mode": "DISCARD"
      },
      "redisParams": {
        "addr": "string",
        "password": "string",
        "db": 0
      },
      "mongoParams": {
        "uri": "string",
        "dbName": "string"
      },
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('/system/data-sources',
{
  method: 'PUT',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Content-Type' => 'application/json',
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.put '/system/data-sources',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Content-Type': 'application/json',
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.put('/system/data-sources', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Content-Type' => 'application/json',
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('PUT','/system/data-sources', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/system/data-sources");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("PUT");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Content-Type": []string{"application/json"},
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("PUT", "/system/data-sources", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`PUT /system/data-sources`

*Update*

> Body parameter

```json
{
  "token": "string",
  "dataSources": [
    {
      "id": "string",
      "backend": "POSTGRESQL",
      "name": "string",
      "description": "string",
      "postgresqlParams": {
        "username": "string",
        "password": "string",
        "host": "string",
        "port": 0,
        "dbName": "string",
        "defaultSchema": "string"
      },
      "mysqlParams": {
        "username": "string",
        "password": "string",
        "host": "string",
        "port": 0,
        "dbName": "string",
        "defaultSchema": "string"
      },
      "virtualParams": {
        "mode": "DISCARD"
      },
      "redisParams": {
        "addr": "string",
        "password": "string",
        "db": 0
      },
      "mongoParams": {
        "uri": "string",
        "dbName": "string"
      },
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}
```

<h3 id="datasource_update-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[UpdateDataSourceRequest](#schemaupdatedatasourcerequest)|true|none|

> Example responses

> 200 Response

```json
{
  "dataSources": [
    {
      "id": "string",
      "backend": "POSTGRESQL",
      "name": "string",
      "description": "string",
      "postgresqlParams": {
        "username": "string",
        "password": "string",
        "host": "string",
        "port": 0,
        "dbName": "string",
        "defaultSchema": "string"
      },
      "mysqlParams": {
        "username": "string",
        "password": "string",
        "host": "string",
        "port": 0,
        "dbName": "string",
        "defaultSchema": "string"
      },
      "virtualParams": {
        "mode": "DISCARD"
      },
      "redisParams": {
        "addr": "string",
        "password": "string",
        "db": 0
      },
      "mongoParams": {
        "uri": "string",
        "dbName": "string"
      },
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}
```

<h3 id="datasource_update-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[UpdateDataSourceResponse](#schemaupdatedatasourceresponse)|
|default|Default|Default error response|[Status](#schemastatus)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

## DataSource_Create

<a id="opIdDataSource_Create"></a>

> Code samples

```shell
# You can also use wget
curl -X POST /system/data-sources \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
POST /system/data-sources HTTP/1.1

Content-Type: application/json
Accept: application/json

```

```javascript
const inputBody = '{
  "token": "string",
  "dataSources": [
    {
      "id": "string",
      "backend": "POSTGRESQL",
      "name": "string",
      "description": "string",
      "postgresqlParams": {
        "username": "string",
        "password": "string",
        "host": "string",
        "port": 0,
        "dbName": "string",
        "defaultSchema": "string"
      },
      "mysqlParams": {
        "username": "string",
        "password": "string",
        "host": "string",
        "port": 0,
        "dbName": "string",
        "defaultSchema": "string"
      },
      "virtualParams": {
        "mode": "DISCARD"
      },
      "redisParams": {
        "addr": "string",
        "password": "string",
        "db": 0
      },
      "mongoParams": {
        "uri": "string",
        "dbName": "string"
      },
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('/system/data-sources',
{
  method: 'POST',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Content-Type' => 'application/json',
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.post '/system/data-sources',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Content-Type': 'application/json',
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.post('/system/data-sources', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Content-Type' => 'application/json',
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('POST','/system/data-sources', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/system/data-sources");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("POST");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Content-Type": []string{"application/json"},
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("POST", "/system/data-sources", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /system/data-sources`

*Create*

> Body parameter

```json
{
  "token": "string",
  "dataSources": [
    {
      "id": "string",
      "backend": "POSTGRESQL",
      "name": "string",
      "description": "string",
      "postgresqlParams": {
        "username": "string",
        "password": "string",
        "host": "string",
        "port": 0,
        "dbName": "string",
        "defaultSchema": "string"
      },
      "mysqlParams": {
        "username": "string",
        "password": "string",
        "host": "string",
        "port": 0,
        "dbName": "string",
        "defaultSchema": "string"
      },
      "virtualParams": {
        "mode": "DISCARD"
      },
      "redisParams": {
        "addr": "string",
        "password": "string",
        "db": 0
      },
      "mongoParams": {
        "uri": "string",
        "dbName": "string"
      },
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}
```

<h3 id="datasource_create-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[CreateDataSourceRequest](#schemacreatedatasourcerequest)|true|none|

> Example responses

> 200 Response

```json
{
  "dataSources": [
    {
      "id": "string",
      "backend": "POSTGRESQL",
      "name": "string",
      "description": "string",
      "postgresqlParams": {
        "username": "string",
        "password": "string",
        "host": "string",
        "port": 0,
        "dbName": "string",
        "defaultSchema": "string"
      },
      "mysqlParams": {
        "username": "string",
        "password": "string",
        "host": "string",
        "port": 0,
        "dbName": "string",
        "defaultSchema": "string"
      },
      "virtualParams": {
        "mode": "DISCARD"
      },
      "redisParams": {
        "addr": "string",
        "password": "string",
        "db": 0
      },
      "mongoParams": {
        "uri": "string",
        "dbName": "string"
      },
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}
```

<h3 id="datasource_create-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[CreateDataSourceResponse](#schemacreatedatasourceresponse)|
|default|Default|Default error response|[Status](#schemastatus)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

## DataSource_Delete

<a id="opIdDataSource_Delete"></a>

> Code samples

```shell
# You can also use wget
curl -X DELETE /system/data-sources \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
DELETE /system/data-sources HTTP/1.1

Content-Type: application/json
Accept: application/json

```

```javascript
const inputBody = '{
  "token": "string",
  "ids": [
    "string"
  ]
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('/system/data-sources',
{
  method: 'DELETE',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Content-Type' => 'application/json',
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.delete '/system/data-sources',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Content-Type': 'application/json',
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.delete('/system/data-sources', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Content-Type' => 'application/json',
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('DELETE','/system/data-sources', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/system/data-sources");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("DELETE");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Content-Type": []string{"application/json"},
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("DELETE", "/system/data-sources", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`DELETE /system/data-sources`

*Delete*

> Body parameter

```json
{
  "token": "string",
  "ids": [
    "string"
  ]
}
```

<h3 id="datasource_delete-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[DeleteDataSourceRequest](#schemadeletedatasourcerequest)|true|none|

> Example responses

> 200 Response

```json
{}
```

<h3 id="datasource_delete-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[DeleteDataSourceResponse](#schemadeletedatasourceresponse)|
|default|Default|Default error response|[Status](#schemastatus)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

## DataSource_Get

<a id="opIdDataSource_Get"></a>

> Code samples

```shell
# You can also use wget
curl -X GET /system/data-sources/{id} \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
GET /system/data-sources/{id} HTTP/1.1

Accept: application/json

```

```javascript

const headers = {
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('/system/data-sources/{id}',
{
  method: 'GET',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.get '/system/data-sources/{id}',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.get('/system/data-sources/{id}', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('GET','/system/data-sources/{id}', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/system/data-sources/{id}");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("GET");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "/system/data-sources/{id}", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /system/data-sources/{id}`

*Get*

<h3 id="datasource_get-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|id|path|string|true|none|
|token|query|string|false|none|

> Example responses

> 200 Response

```json
{
  "dataSource": {
    "id": "string",
    "backend": "POSTGRESQL",
    "name": "string",
    "description": "string",
    "postgresqlParams": {
      "username": "string",
      "password": "string",
      "host": "string",
      "port": 0,
      "dbName": "string",
      "defaultSchema": "string"
    },
    "mysqlParams": {
      "username": "string",
      "password": "string",
      "host": "string",
      "port": 0,
      "dbName": "string",
      "defaultSchema": "string"
    },
    "virtualParams": {
      "mode": "DISCARD"
    },
    "redisParams": {
      "addr": "string",
      "password": "string",
      "db": 0
    },
    "mongoParams": {
      "uri": "string",
      "dbName": "string"
    },
    "auditData": {
      "createdOn": "2019-08-24T14:15:22Z",
      "updatedOn": "2019-08-24T14:15:22Z",
      "createdBy": "string",
      "updatedBy": "string"
    },
    "version": 0
  }
}
```

<h3 id="datasource_get-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[GetDataSourceResponse](#schemagetdatasourceresponse)|
|default|Default|Default error response|[Status](#schemastatus)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

## DataSource_PrepareResourceFromEntity

<a id="opIdDataSource_PrepareResourceFromEntity"></a>

> Code samples

```shell
# You can also use wget
curl -X POST /system/data-sources/{id}/_prepare_entity \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
POST /system/data-sources/{id}/_prepare_entity HTTP/1.1

Content-Type: application/json
Accept: application/json

```

```javascript
const inputBody = '{
  "token": "string",
  "id": "string",
  "catalog": "string",
  "entity": "string"
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('/system/data-sources/{id}/_prepare_entity',
{
  method: 'POST',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Content-Type' => 'application/json',
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.post '/system/data-sources/{id}/_prepare_entity',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Content-Type': 'application/json',
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.post('/system/data-sources/{id}/_prepare_entity', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Content-Type' => 'application/json',
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('POST','/system/data-sources/{id}/_prepare_entity', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/system/data-sources/{id}/_prepare_entity");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("POST");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Content-Type": []string{"application/json"},
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("POST", "/system/data-sources/{id}/_prepare_entity", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /system/data-sources/{id}/_prepare_entity`

*PrepareResourceFromEntity*

> Body parameter

```json
{
  "token": "string",
  "id": "string",
  "catalog": "string",
  "entity": "string"
}
```

<h3 id="datasource_prepareresourcefromentity-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|id|path|string|true|none|
|body|body|[PrepareResourceFromEntityRequest](#schemaprepareresourcefromentityrequest)|true|none|

> Example responses

> 200 Response

```json
{
  "resource": {
    "id": "string",
    "name": "string",
    "namespace": "string",
    "sourceConfig": {
      "dataSource": "string",
      "catalog": "string",
      "entity": "string"
    },
    "properties": [
      {
        "id": "string",
        "name": "string",
        "type": "BOOL",
        "mapping": "string",
        "required": true,
        "primary": true,
        "length": 0,
        "unique": true,
        "immutable": true,
        "securityContext": {
          "constraints": [
            {
              "namespace": "string",
              "resource": "string",
              "property": "string",
              "before": "2019-08-24T14:15:22Z",
              "after": "2019-08-24T14:15:22Z",
              "principal": "string",
              "recordIds": [
                "string"
              ],
              "operation": "OPERATION_TYPE_READ",
              "permit": "PERMIT_TYPE_ALLOW"
            }
          ]
        },
        "defaultValue": null,
        "exampleValue": null,
        "enumValues": [
          null
        ],
        "reference": {
          "referencedResource": "string",
          "cascade": true
        },
        "properties": [
          {}
        ],
        "Item": {},
        "title": "string",
        "description": "string",
        "annotations": {
          "property1": "string",
          "property2": "string"
        }
      }
    ],
    "indexes": [
      {
        "properties": [
          {
            "name": "string",
            "order": "ORDER_UNKNOWN"
          }
        ],
        "indexType": "BTREE",
        "unique": true,
        "annotations": {
          "property1": "string",
          "property2": "string"
        }
      }
    ],
    "securityContext": {
      "constraints": [
        {
          "namespace": "string",
          "resource": "string",
          "property": "string",
          "before": "2019-08-24T14:15:22Z",
          "after": "2019-08-24T14:15:22Z",
          "principal": "string",
          "recordIds": [
            "string"
          ],
          "operation": "OPERATION_TYPE_READ",
          "permit": "PERMIT_TYPE_ALLOW"
        }
      ]
    },
    "virtual": true,
    "immutable": true,
    "abstract": true,
    "title": "string",
    "description": "string",
    "auditData": {
      "createdOn": "2019-08-24T14:15:22Z",
      "updatedOn": "2019-08-24T14:15:22Z",
      "createdBy": "string",
      "updatedBy": "string"
    },
    "version": 0,
    "annotations": {
      "property1": "string",
      "property2": "string"
    }
  }
}
```

<h3 id="datasource_prepareresourcefromentity-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[PrepareResourceFromEntityResponse](#schemaprepareresourcefromentityresponse)|
|default|Default|Default error response|[Status](#schemastatus)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

## DataSource_Status

<a id="opIdDataSource_Status"></a>

> Code samples

```shell
# You can also use wget
curl -X GET /system/data-sources/{id}/_status \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
GET /system/data-sources/{id}/_status HTTP/1.1

Accept: application/json

```

```javascript

const headers = {
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('/system/data-sources/{id}/_status',
{
  method: 'GET',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.get '/system/data-sources/{id}/_status',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.get('/system/data-sources/{id}/_status', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('GET','/system/data-sources/{id}/_status', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/system/data-sources/{id}/_status");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("GET");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "/system/data-sources/{id}/_status", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /system/data-sources/{id}/_status`

*Status*

<h3 id="datasource_status-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|id|path|string|true|none|
|token|query|string|false|none|

> Example responses

> 200 Response

```json
{
  "connectionAlreadyInitiated": true,
  "testConnection": true
}
```

<h3 id="datasource_status-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[StatusResponse](#schemastatusresponse)|
|default|Default|Default error response|[Status](#schemastatus)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

## DataSource_ListEntities

<a id="opIdDataSource_ListEntities"></a>

> Code samples

```shell
# You can also use wget
curl -X GET /system/data-sources/{id}/entities \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
GET /system/data-sources/{id}/entities HTTP/1.1

Accept: application/json

```

```javascript

const headers = {
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('/system/data-sources/{id}/entities',
{
  method: 'GET',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.get '/system/data-sources/{id}/entities',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.get('/system/data-sources/{id}/entities', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('GET','/system/data-sources/{id}/entities', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/system/data-sources/{id}/entities");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("GET");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "/system/data-sources/{id}/entities", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /system/data-sources/{id}/entities`

*ListEntities*

<h3 id="datasource_listentities-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|id|path|string|true|none|
|token|query|string|false|none|

> Example responses

> 200 Response

```json
{
  "catalogs": [
    {
      "name": "string",
      "entities": [
        {
          "name": "string",
          "readOnly": true
        }
      ]
    }
  ]
}
```

<h3 id="datasource_listentities-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[ListEntitiesResponse](#schemalistentitiesresponse)|
|default|Default|Default error response|[Status](#schemastatus)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

<h1 id="apibrew-extension">Extension</h1>

Extension Service is for managing extensions

## Extension_List

<a id="opIdExtension_List"></a>

> Code samples

```shell
# You can also use wget
curl -X GET /system/extensions \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
GET /system/extensions HTTP/1.1

Accept: application/json

```

```javascript

const headers = {
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('/system/extensions',
{
  method: 'GET',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.get '/system/extensions',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.get('/system/extensions', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('GET','/system/extensions', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/system/extensions");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("GET");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "/system/extensions", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /system/extensions`

*List*

<h3 id="extension_list-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|token|query|string|false|none|

> Example responses

> 200 Response

```json
{
  "content": [
    {
      "id": "string",
      "name": "string",
      "description": "string",
      "namespace": "string",
      "resource": "string",
      "before": {
        "all": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "create": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "update": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "delete": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "get": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "list": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "sync": true
      },
      "instead": {
        "all": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "create": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "update": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "delete": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "get": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "list": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "finalize": true
      },
      "after": {
        "all": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "create": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "update": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "delete": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "get": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "list": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "sync": true
      },
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}
```

<h3 id="extension_list-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[ListExtensionResponse](#schemalistextensionresponse)|
|default|Default|Default error response|[Status](#schemastatus)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

## Extension_Update

<a id="opIdExtension_Update"></a>

> Code samples

```shell
# You can also use wget
curl -X PUT /system/extensions \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
PUT /system/extensions HTTP/1.1

Content-Type: application/json
Accept: application/json

```

```javascript
const inputBody = '{
  "token": "string",
  "extensions": [
    {
      "id": "string",
      "name": "string",
      "description": "string",
      "namespace": "string",
      "resource": "string",
      "before": {
        "all": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "create": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "update": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "delete": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "get": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "list": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "sync": true
      },
      "instead": {
        "all": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "create": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "update": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "delete": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "get": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "list": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "finalize": true
      },
      "after": {
        "all": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "create": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "update": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "delete": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "get": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "list": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "sync": true
      },
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('/system/extensions',
{
  method: 'PUT',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Content-Type' => 'application/json',
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.put '/system/extensions',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Content-Type': 'application/json',
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.put('/system/extensions', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Content-Type' => 'application/json',
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('PUT','/system/extensions', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/system/extensions");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("PUT");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Content-Type": []string{"application/json"},
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("PUT", "/system/extensions", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`PUT /system/extensions`

*Update*

> Body parameter

```json
{
  "token": "string",
  "extensions": [
    {
      "id": "string",
      "name": "string",
      "description": "string",
      "namespace": "string",
      "resource": "string",
      "before": {
        "all": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "create": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "update": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "delete": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "get": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "list": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "sync": true
      },
      "instead": {
        "all": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "create": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "update": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "delete": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "get": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "list": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "finalize": true
      },
      "after": {
        "all": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "create": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "update": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "delete": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "get": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "list": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "sync": true
      },
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}
```

<h3 id="extension_update-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[UpdateExtensionRequest](#schemaupdateextensionrequest)|true|none|

> Example responses

> 200 Response

```json
{
  "extensions": [
    {
      "id": "string",
      "name": "string",
      "description": "string",
      "namespace": "string",
      "resource": "string",
      "before": {
        "all": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "create": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "update": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "delete": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "get": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "list": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "sync": true
      },
      "instead": {
        "all": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "create": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "update": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "delete": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "get": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "list": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "finalize": true
      },
      "after": {
        "all": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "create": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "update": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "delete": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "get": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "list": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "sync": true
      },
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}
```

<h3 id="extension_update-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[UpdateExtensionResponse](#schemaupdateextensionresponse)|
|default|Default|Default error response|[Status](#schemastatus)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

## Extension_Create

<a id="opIdExtension_Create"></a>

> Code samples

```shell
# You can also use wget
curl -X POST /system/extensions \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
POST /system/extensions HTTP/1.1

Content-Type: application/json
Accept: application/json

```

```javascript
const inputBody = '{
  "token": "string",
  "extensions": [
    {
      "id": "string",
      "name": "string",
      "description": "string",
      "namespace": "string",
      "resource": "string",
      "before": {
        "all": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "create": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "update": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "delete": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "get": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "list": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "sync": true
      },
      "instead": {
        "all": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "create": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "update": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "delete": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "get": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "list": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "finalize": true
      },
      "after": {
        "all": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "create": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "update": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "delete": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "get": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "list": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "sync": true
      },
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('/system/extensions',
{
  method: 'POST',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Content-Type' => 'application/json',
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.post '/system/extensions',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Content-Type': 'application/json',
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.post('/system/extensions', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Content-Type' => 'application/json',
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('POST','/system/extensions', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/system/extensions");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("POST");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Content-Type": []string{"application/json"},
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("POST", "/system/extensions", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /system/extensions`

*Create*

> Body parameter

```json
{
  "token": "string",
  "extensions": [
    {
      "id": "string",
      "name": "string",
      "description": "string",
      "namespace": "string",
      "resource": "string",
      "before": {
        "all": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "create": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "update": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "delete": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "get": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "list": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "sync": true
      },
      "instead": {
        "all": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "create": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "update": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "delete": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "get": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "list": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "finalize": true
      },
      "after": {
        "all": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "create": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "update": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "delete": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "get": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "list": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "sync": true
      },
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}
```

<h3 id="extension_create-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[CreateExtensionRequest](#schemacreateextensionrequest)|true|none|

> Example responses

> 200 Response

```json
{
  "extensions": [
    {
      "id": "string",
      "name": "string",
      "description": "string",
      "namespace": "string",
      "resource": "string",
      "before": {
        "all": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "create": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "update": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "delete": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "get": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "list": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "sync": true
      },
      "instead": {
        "all": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "create": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "update": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "delete": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "get": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "list": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "finalize": true
      },
      "after": {
        "all": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "create": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "update": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "delete": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "get": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "list": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "sync": true
      },
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}
```

<h3 id="extension_create-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[CreateExtensionResponse](#schemacreateextensionresponse)|
|default|Default|Default error response|[Status](#schemastatus)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

## Extension_Delete

<a id="opIdExtension_Delete"></a>

> Code samples

```shell
# You can also use wget
curl -X DELETE /system/extensions \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
DELETE /system/extensions HTTP/1.1

Content-Type: application/json
Accept: application/json

```

```javascript
const inputBody = '{
  "token": "string",
  "ids": [
    "string"
  ]
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('/system/extensions',
{
  method: 'DELETE',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Content-Type' => 'application/json',
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.delete '/system/extensions',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Content-Type': 'application/json',
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.delete('/system/extensions', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Content-Type' => 'application/json',
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('DELETE','/system/extensions', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/system/extensions");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("DELETE");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Content-Type": []string{"application/json"},
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("DELETE", "/system/extensions", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`DELETE /system/extensions`

*Delete*

> Body parameter

```json
{
  "token": "string",
  "ids": [
    "string"
  ]
}
```

<h3 id="extension_delete-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[DeleteExtensionRequest](#schemadeleteextensionrequest)|true|none|

> Example responses

> 200 Response

```json
{}
```

<h3 id="extension_delete-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[DeleteExtensionResponse](#schemadeleteextensionresponse)|
|default|Default|Default error response|[Status](#schemastatus)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

## Extension_Get

<a id="opIdExtension_Get"></a>

> Code samples

```shell
# You can also use wget
curl -X GET /system/extensions/{id} \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
GET /system/extensions/{id} HTTP/1.1

Accept: application/json

```

```javascript

const headers = {
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('/system/extensions/{id}',
{
  method: 'GET',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.get '/system/extensions/{id}',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.get('/system/extensions/{id}', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('GET','/system/extensions/{id}', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/system/extensions/{id}");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("GET");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "/system/extensions/{id}", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /system/extensions/{id}`

*Get*

<h3 id="extension_get-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|id|path|string|true|none|
|token|query|string|false|none|

> Example responses

> 200 Response

```json
{
  "extension": {
    "id": "string",
    "name": "string",
    "description": "string",
    "namespace": "string",
    "resource": "string",
    "before": {
      "all": {
        "functionCall": {
          "host": "string",
          "functionName": "string"
        },
        "httpCall": {
          "uri": "string",
          "method": "string"
        }
      },
      "create": {
        "functionCall": {
          "host": "string",
          "functionName": "string"
        },
        "httpCall": {
          "uri": "string",
          "method": "string"
        }
      },
      "update": {
        "functionCall": {
          "host": "string",
          "functionName": "string"
        },
        "httpCall": {
          "uri": "string",
          "method": "string"
        }
      },
      "delete": {
        "functionCall": {
          "host": "string",
          "functionName": "string"
        },
        "httpCall": {
          "uri": "string",
          "method": "string"
        }
      },
      "get": {
        "functionCall": {
          "host": "string",
          "functionName": "string"
        },
        "httpCall": {
          "uri": "string",
          "method": "string"
        }
      },
      "list": {
        "functionCall": {
          "host": "string",
          "functionName": "string"
        },
        "httpCall": {
          "uri": "string",
          "method": "string"
        }
      },
      "sync": true
    },
    "instead": {
      "all": {
        "functionCall": {
          "host": "string",
          "functionName": "string"
        },
        "httpCall": {
          "uri": "string",
          "method": "string"
        }
      },
      "create": {
        "functionCall": {
          "host": "string",
          "functionName": "string"
        },
        "httpCall": {
          "uri": "string",
          "method": "string"
        }
      },
      "update": {
        "functionCall": {
          "host": "string",
          "functionName": "string"
        },
        "httpCall": {
          "uri": "string",
          "method": "string"
        }
      },
      "delete": {
        "functionCall": {
          "host": "string",
          "functionName": "string"
        },
        "httpCall": {
          "uri": "string",
          "method": "string"
        }
      },
      "get": {
        "functionCall": {
          "host": "string",
          "functionName": "string"
        },
        "httpCall": {
          "uri": "string",
          "method": "string"
        }
      },
      "list": {
        "functionCall": {
          "host": "string",
          "functionName": "string"
        },
        "httpCall": {
          "uri": "string",
          "method": "string"
        }
      },
      "finalize": true
    },
    "after": {
      "all": {
        "functionCall": {
          "host": "string",
          "functionName": "string"
        },
        "httpCall": {
          "uri": "string",
          "method": "string"
        }
      },
      "create": {
        "functionCall": {
          "host": "string",
          "functionName": "string"
        },
        "httpCall": {
          "uri": "string",
          "method": "string"
        }
      },
      "update": {
        "functionCall": {
          "host": "string",
          "functionName": "string"
        },
        "httpCall": {
          "uri": "string",
          "method": "string"
        }
      },
      "delete": {
        "functionCall": {
          "host": "string",
          "functionName": "string"
        },
        "httpCall": {
          "uri": "string",
          "method": "string"
        }
      },
      "get": {
        "functionCall": {
          "host": "string",
          "functionName": "string"
        },
        "httpCall": {
          "uri": "string",
          "method": "string"
        }
      },
      "list": {
        "functionCall": {
          "host": "string",
          "functionName": "string"
        },
        "httpCall": {
          "uri": "string",
          "method": "string"
        }
      },
      "sync": true
    },
    "auditData": {
      "createdOn": "2019-08-24T14:15:22Z",
      "updatedOn": "2019-08-24T14:15:22Z",
      "createdBy": "string",
      "updatedBy": "string"
    },
    "version": 0
  }
}
```

<h3 id="extension_get-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[GetExtensionResponse](#schemagetextensionresponse)|
|default|Default|Default error response|[Status](#schemastatus)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

<h1 id="apibrew-namespace">Namespace</h1>

Namespace Service is for managing namespaces

## Namespace_List

<a id="opIdNamespace_List"></a>

> Code samples

```shell
# You can also use wget
curl -X GET /system/namespaces \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
GET /system/namespaces HTTP/1.1

Accept: application/json

```

```javascript

const headers = {
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('/system/namespaces',
{
  method: 'GET',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.get '/system/namespaces',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.get('/system/namespaces', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('GET','/system/namespaces', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/system/namespaces");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("GET");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "/system/namespaces", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /system/namespaces`

*List*

<h3 id="namespace_list-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|token|query|string|false|none|

> Example responses

> 200 Response

```json
{
  "content": [
    {
      "id": "string",
      "name": "string",
      "description": "string",
      "details": {},
      "securityContext": {
        "constraints": [
          {
            "namespace": "string",
            "resource": "string",
            "property": "string",
            "before": "2019-08-24T14:15:22Z",
            "after": "2019-08-24T14:15:22Z",
            "principal": "string",
            "recordIds": [
              "string"
            ],
            "operation": "OPERATION_TYPE_READ",
            "permit": "PERMIT_TYPE_ALLOW"
          }
        ]
      },
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}
```

<h3 id="namespace_list-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[ListNamespaceResponse](#schemalistnamespaceresponse)|
|default|Default|Default error response|[Status](#schemastatus)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

## Namespace_Update

<a id="opIdNamespace_Update"></a>

> Code samples

```shell
# You can also use wget
curl -X PUT /system/namespaces \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
PUT /system/namespaces HTTP/1.1

Content-Type: application/json
Accept: application/json

```

```javascript
const inputBody = '{
  "token": "string",
  "Namespaces": [
    {
      "id": "string",
      "name": "string",
      "description": "string",
      "details": {},
      "securityContext": {
        "constraints": [
          {
            "namespace": "string",
            "resource": "string",
            "property": "string",
            "before": "2019-08-24T14:15:22Z",
            "after": "2019-08-24T14:15:22Z",
            "principal": "string",
            "recordIds": [
              "string"
            ],
            "operation": "OPERATION_TYPE_READ",
            "permit": "PERMIT_TYPE_ALLOW"
          }
        ]
      },
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('/system/namespaces',
{
  method: 'PUT',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Content-Type' => 'application/json',
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.put '/system/namespaces',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Content-Type': 'application/json',
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.put('/system/namespaces', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Content-Type' => 'application/json',
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('PUT','/system/namespaces', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/system/namespaces");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("PUT");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Content-Type": []string{"application/json"},
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("PUT", "/system/namespaces", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`PUT /system/namespaces`

*Update*

> Body parameter

```json
{
  "token": "string",
  "Namespaces": [
    {
      "id": "string",
      "name": "string",
      "description": "string",
      "details": {},
      "securityContext": {
        "constraints": [
          {
            "namespace": "string",
            "resource": "string",
            "property": "string",
            "before": "2019-08-24T14:15:22Z",
            "after": "2019-08-24T14:15:22Z",
            "principal": "string",
            "recordIds": [
              "string"
            ],
            "operation": "OPERATION_TYPE_READ",
            "permit": "PERMIT_TYPE_ALLOW"
          }
        ]
      },
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}
```

<h3 id="namespace_update-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[UpdateNamespaceRequest](#schemaupdatenamespacerequest)|true|none|

> Example responses

> 200 Response

```json
{
  "Namespaces": [
    {
      "id": "string",
      "name": "string",
      "description": "string",
      "details": {},
      "securityContext": {
        "constraints": [
          {
            "namespace": "string",
            "resource": "string",
            "property": "string",
            "before": "2019-08-24T14:15:22Z",
            "after": "2019-08-24T14:15:22Z",
            "principal": "string",
            "recordIds": [
              "string"
            ],
            "operation": "OPERATION_TYPE_READ",
            "permit": "PERMIT_TYPE_ALLOW"
          }
        ]
      },
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}
```

<h3 id="namespace_update-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[UpdateNamespaceResponse](#schemaupdatenamespaceresponse)|
|default|Default|Default error response|[Status](#schemastatus)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

## Namespace_Create

<a id="opIdNamespace_Create"></a>

> Code samples

```shell
# You can also use wget
curl -X POST /system/namespaces \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
POST /system/namespaces HTTP/1.1

Content-Type: application/json
Accept: application/json

```

```javascript
const inputBody = '{
  "token": "string",
  "Namespaces": [
    {
      "id": "string",
      "name": "string",
      "description": "string",
      "details": {},
      "securityContext": {
        "constraints": [
          {
            "namespace": "string",
            "resource": "string",
            "property": "string",
            "before": "2019-08-24T14:15:22Z",
            "after": "2019-08-24T14:15:22Z",
            "principal": "string",
            "recordIds": [
              "string"
            ],
            "operation": "OPERATION_TYPE_READ",
            "permit": "PERMIT_TYPE_ALLOW"
          }
        ]
      },
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('/system/namespaces',
{
  method: 'POST',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Content-Type' => 'application/json',
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.post '/system/namespaces',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Content-Type': 'application/json',
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.post('/system/namespaces', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Content-Type' => 'application/json',
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('POST','/system/namespaces', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/system/namespaces");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("POST");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Content-Type": []string{"application/json"},
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("POST", "/system/namespaces", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /system/namespaces`

*Create*

> Body parameter

```json
{
  "token": "string",
  "Namespaces": [
    {
      "id": "string",
      "name": "string",
      "description": "string",
      "details": {},
      "securityContext": {
        "constraints": [
          {
            "namespace": "string",
            "resource": "string",
            "property": "string",
            "before": "2019-08-24T14:15:22Z",
            "after": "2019-08-24T14:15:22Z",
            "principal": "string",
            "recordIds": [
              "string"
            ],
            "operation": "OPERATION_TYPE_READ",
            "permit": "PERMIT_TYPE_ALLOW"
          }
        ]
      },
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}
```

<h3 id="namespace_create-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[CreateNamespaceRequest](#schemacreatenamespacerequest)|true|none|

> Example responses

> 200 Response

```json
{
  "Namespaces": [
    {
      "id": "string",
      "name": "string",
      "description": "string",
      "details": {},
      "securityContext": {
        "constraints": [
          {
            "namespace": "string",
            "resource": "string",
            "property": "string",
            "before": "2019-08-24T14:15:22Z",
            "after": "2019-08-24T14:15:22Z",
            "principal": "string",
            "recordIds": [
              "string"
            ],
            "operation": "OPERATION_TYPE_READ",
            "permit": "PERMIT_TYPE_ALLOW"
          }
        ]
      },
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}
```

<h3 id="namespace_create-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[CreateNamespaceResponse](#schemacreatenamespaceresponse)|
|default|Default|Default error response|[Status](#schemastatus)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

## Namespace_Delete

<a id="opIdNamespace_Delete"></a>

> Code samples

```shell
# You can also use wget
curl -X DELETE /system/namespaces \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
DELETE /system/namespaces HTTP/1.1

Content-Type: application/json
Accept: application/json

```

```javascript
const inputBody = '{
  "token": "string",
  "ids": [
    "string"
  ]
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('/system/namespaces',
{
  method: 'DELETE',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Content-Type' => 'application/json',
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.delete '/system/namespaces',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Content-Type': 'application/json',
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.delete('/system/namespaces', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Content-Type' => 'application/json',
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('DELETE','/system/namespaces', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/system/namespaces");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("DELETE");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Content-Type": []string{"application/json"},
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("DELETE", "/system/namespaces", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`DELETE /system/namespaces`

*Delete*

> Body parameter

```json
{
  "token": "string",
  "ids": [
    "string"
  ]
}
```

<h3 id="namespace_delete-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[DeleteNamespaceRequest](#schemadeletenamespacerequest)|true|none|

> Example responses

> 200 Response

```json
{
  "Namespaces": [
    {
      "id": "string",
      "name": "string",
      "description": "string",
      "details": {},
      "securityContext": {
        "constraints": [
          {
            "namespace": "string",
            "resource": "string",
            "property": "string",
            "before": "2019-08-24T14:15:22Z",
            "after": "2019-08-24T14:15:22Z",
            "principal": "string",
            "recordIds": [
              "string"
            ],
            "operation": "OPERATION_TYPE_READ",
            "permit": "PERMIT_TYPE_ALLOW"
          }
        ]
      },
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}
```

<h3 id="namespace_delete-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[DeleteNamespaceResponse](#schemadeletenamespaceresponse)|
|default|Default|Default error response|[Status](#schemastatus)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

## Namespace_Get

<a id="opIdNamespace_Get"></a>

> Code samples

```shell
# You can also use wget
curl -X GET /system/namespaces/{id} \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
GET /system/namespaces/{id} HTTP/1.1

Accept: application/json

```

```javascript

const headers = {
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('/system/namespaces/{id}',
{
  method: 'GET',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.get '/system/namespaces/{id}',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.get('/system/namespaces/{id}', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('GET','/system/namespaces/{id}', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/system/namespaces/{id}");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("GET");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "/system/namespaces/{id}", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /system/namespaces/{id}`

*Get*

<h3 id="namespace_get-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|id|path|string|true|none|
|token|query|string|false|none|

> Example responses

> 200 Response

```json
{
  "Namespace": {
    "id": "string",
    "name": "string",
    "description": "string",
    "details": {},
    "securityContext": {
      "constraints": [
        {
          "namespace": "string",
          "resource": "string",
          "property": "string",
          "before": "2019-08-24T14:15:22Z",
          "after": "2019-08-24T14:15:22Z",
          "principal": "string",
          "recordIds": [
            "string"
          ],
          "operation": "OPERATION_TYPE_READ",
          "permit": "PERMIT_TYPE_ALLOW"
        }
      ]
    },
    "auditData": {
      "createdOn": "2019-08-24T14:15:22Z",
      "updatedOn": "2019-08-24T14:15:22Z",
      "createdBy": "string",
      "updatedBy": "string"
    },
    "version": 0
  }
}
```

<h3 id="namespace_get-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[GetNamespaceResponse](#schemagetnamespaceresponse)|
|default|Default|Default error response|[Status](#schemastatus)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

<h1 id="apibrew-record">Record</h1>

Record service is an abstract service for records of all resources. You can do CRUD like operations with Record service

## Record_ReadStream

<a id="opIdRecord_ReadStream"></a>

> Code samples

```shell
# You can also use wget
curl -X POST /records/read_stream \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
POST /records/read_stream HTTP/1.1

Content-Type: application/json
Accept: application/json

```

```javascript
const inputBody = '{
  "token": "string",
  "namespace": "string",
  "resource": "string",
  "query": {
    "and": {
      "expressions": [
        {}
      ]
    },
    "or": {
      "expressions": [
        {}
      ]
    },
    "not": {},
    "equal": {
      "left": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "right": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    },
    "lessThan": {
      "left": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "right": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    },
    "greaterThan": {
      "left": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "right": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    },
    "lessThanOrEqual": {
      "left": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "right": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    },
    "greaterThanOrEqual": {
      "left": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "right": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    },
    "in": {
      "left": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "right": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    },
    "isNull": {
      "additionalProperties": [
        {
          "name": "string",
          "value": {
            "value": {
              "@type": "string"
            },
            "yaml": "string"
          }
        }
      ]
    },
    "regexMatch": {
      "pattern": "string",
      "expression": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    }
  },
  "limit": 0,
  "offset": 0,
  "useHistory": true,
  "resolveReferences": [
    "string"
  ],
  "useTransaction": true,
  "packRecords": true,
  "annotations": {
    "property1": "string",
    "property2": "string"
  }
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('/records/read_stream',
{
  method: 'POST',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Content-Type' => 'application/json',
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.post '/records/read_stream',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Content-Type': 'application/json',
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.post('/records/read_stream', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Content-Type' => 'application/json',
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('POST','/records/read_stream', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/records/read_stream");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("POST");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Content-Type": []string{"application/json"},
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("POST", "/records/read_stream", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /records/read_stream`

*ReadStream*

> Body parameter

```json
{
  "token": "string",
  "namespace": "string",
  "resource": "string",
  "query": {
    "and": {
      "expressions": [
        {}
      ]
    },
    "or": {
      "expressions": [
        {}
      ]
    },
    "not": {},
    "equal": {
      "left": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "right": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    },
    "lessThan": {
      "left": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "right": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    },
    "greaterThan": {
      "left": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "right": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    },
    "lessThanOrEqual": {
      "left": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "right": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    },
    "greaterThanOrEqual": {
      "left": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "right": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    },
    "in": {
      "left": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "right": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    },
    "isNull": {
      "additionalProperties": [
        {
          "name": "string",
          "value": {
            "value": {
              "@type": "string"
            },
            "yaml": "string"
          }
        }
      ]
    },
    "regexMatch": {
      "pattern": "string",
      "expression": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    }
  },
  "limit": 0,
  "offset": 0,
  "useHistory": true,
  "resolveReferences": [
    "string"
  ],
  "useTransaction": true,
  "packRecords": true,
  "annotations": {
    "property1": "string",
    "property2": "string"
  }
}
```

<h3 id="record_readstream-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[ReadStreamRequest](#schemareadstreamrequest)|true|none|

> Example responses

> 200 Response

```json
{
  "id": "string",
  "properties": {
    "property1": null,
    "property2": null
  },
  "propertiesPacked": [
    null
  ]
}
```

<h3 id="record_readstream-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[Record](#schemarecord)|
|default|Default|Default error response|[Status](#schemastatus)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

## Record_List

<a id="opIdRecord_List"></a>

> Code samples

```shell
# You can also use wget
curl -X GET /records/{namespace}/{resource} \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
GET /records/{namespace}/{resource} HTTP/1.1

Accept: application/json

```

```javascript

const headers = {
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('/records/{namespace}/{resource}',
{
  method: 'GET',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.get '/records/{namespace}/{resource}',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.get('/records/{namespace}/{resource}', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('GET','/records/{namespace}/{resource}', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/records/{namespace}/{resource}");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("GET");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "/records/{namespace}/{resource}", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /records/{namespace}/{resource}`

*List*

<h3 id="record_list-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|namespace|path|string|true|none|
|resource|path|string|true|none|
|token|query|string|false|none|
|limit|query|integer(uint32)|false|none|
|offset|query|integer(uint64)|false|none|
|useHistory|query|boolean|false|none|
|resolveReferences|query|array[string]|false|none|

> Example responses

> 200 Response

```json
{
  "total": 0,
  "content": [
    {
      "id": "string",
      "properties": {
        "property1": null,
        "property2": null
      },
      "propertiesPacked": [
        null
      ]
    }
  ]
}
```

<h3 id="record_list-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[ListRecordResponse](#schemalistrecordresponse)|
|default|Default|Default error response|[Status](#schemastatus)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

## Record_Create

<a id="opIdRecord_Create"></a>

> Code samples

```shell
# You can also use wget
curl -X POST /records/{namespace}/{resource}/_bulk \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
POST /records/{namespace}/{resource}/_bulk HTTP/1.1

Accept: application/json

```

```javascript

const headers = {
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('/records/{namespace}/{resource}/_bulk',
{
  method: 'POST',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.post '/records/{namespace}/{resource}/_bulk',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.post('/records/{namespace}/{resource}/_bulk', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('POST','/records/{namespace}/{resource}/_bulk', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/records/{namespace}/{resource}/_bulk");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("POST");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("POST", "/records/{namespace}/{resource}/_bulk", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /records/{namespace}/{resource}/_bulk`

*Create*

<h3 id="record_create-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|namespace|path|string|true|none|
|resource|path|string|true|none|
|token|query|string|false|none|
|record.id|query|string|false|none|
|record.propertiesPacked|query|array[any]|false|none|
|ignoreIfExists|query|boolean|false|none|

> Example responses

> 200 Response

```json
{
  "record": {
    "id": "string",
    "properties": {
      "property1": null,
      "property2": null
    },
    "propertiesPacked": [
      null
    ]
  },
  "records": [
    {
      "id": "string",
      "properties": {
        "property1": null,
        "property2": null
      },
      "propertiesPacked": [
        null
      ]
    }
  ],
  "inserted": [
    true
  ]
}
```

<h3 id="record_create-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[CreateRecordResponse](#schemacreaterecordresponse)|
|default|Default|Default error response|[Status](#schemastatus)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

## Record_Apply

<a id="opIdRecord_Apply"></a>

> Code samples

```shell
# You can also use wget
curl -X PATCH /records/{namespace}/{resource}/_bulk \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
PATCH /records/{namespace}/{resource}/_bulk HTTP/1.1

Accept: application/json

```

```javascript

const headers = {
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('/records/{namespace}/{resource}/_bulk',
{
  method: 'PATCH',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.patch '/records/{namespace}/{resource}/_bulk',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.patch('/records/{namespace}/{resource}/_bulk', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('PATCH','/records/{namespace}/{resource}/_bulk', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/records/{namespace}/{resource}/_bulk");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("PATCH");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("PATCH", "/records/{namespace}/{resource}/_bulk", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`PATCH /records/{namespace}/{resource}/_bulk`

*Apply*

<h3 id="record_apply-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|namespace|path|string|true|none|
|resource|path|string|true|none|
|token|query|string|false|none|
|record.id|query|string|false|none|
|record.propertiesPacked|query|array[any]|false|none|
|checkVersion|query|boolean|false|none|

> Example responses

> 200 Response

```json
{
  "record": {
    "id": "string",
    "properties": {
      "property1": null,
      "property2": null
    },
    "propertiesPacked": [
      null
    ]
  },
  "records": [
    {
      "id": "string",
      "properties": {
        "property1": null,
        "property2": null
      },
      "propertiesPacked": [
        null
      ]
    }
  ]
}
```

<h3 id="record_apply-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[ApplyRecordResponse](#schemaapplyrecordresponse)|
|default|Default|Default error response|[Status](#schemastatus)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

## Record_Update

<a id="opIdRecord_Update"></a>

> Code samples

```shell
# You can also use wget
curl -X PUT /records/{namespace}/{resource}/{id} \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
PUT /records/{namespace}/{resource}/{id} HTTP/1.1

Content-Type: application/json
Accept: application/json

```

```javascript
const inputBody = '{
  "key": "string",
  "value": null
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('/records/{namespace}/{resource}/{id}',
{
  method: 'PUT',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Content-Type' => 'application/json',
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.put '/records/{namespace}/{resource}/{id}',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Content-Type': 'application/json',
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.put('/records/{namespace}/{resource}/{id}', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Content-Type' => 'application/json',
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('PUT','/records/{namespace}/{resource}/{id}', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/records/{namespace}/{resource}/{id}");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("PUT");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Content-Type": []string{"application/json"},
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("PUT", "/records/{namespace}/{resource}/{id}", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`PUT /records/{namespace}/{resource}/{id}`

*Update*

> Body parameter

```json
{
  "key": "string",
  "value": null
}
```

<h3 id="record_update-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|namespace|path|string|true|none|
|resource|path|string|true|none|
|id|path|string|true|none|
|token|query|string|false|none|
|checkVersion|query|boolean|false|none|
|body|body|[UpdateRecordRequest_PropertiesEntry](#schemaupdaterecordrequest_propertiesentry)|true|none|

> Example responses

> 200 Response

```json
{
  "record": {
    "id": "string",
    "properties": {
      "property1": null,
      "property2": null
    },
    "propertiesPacked": [
      null
    ]
  },
  "records": [
    {
      "id": "string",
      "properties": {
        "property1": null,
        "property2": null
      },
      "propertiesPacked": [
        null
      ]
    }
  ]
}
```

<h3 id="record_update-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[UpdateRecordResponse](#schemaupdaterecordresponse)|
|default|Default|Default error response|[Status](#schemastatus)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

## Record_Delete

<a id="opIdRecord_Delete"></a>

> Code samples

```shell
# You can also use wget
curl -X DELETE /records/{namespace}/{resource}/{id} \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
DELETE /records/{namespace}/{resource}/{id} HTTP/1.1

Accept: application/json

```

```javascript

const headers = {
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('/records/{namespace}/{resource}/{id}',
{
  method: 'DELETE',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.delete '/records/{namespace}/{resource}/{id}',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.delete('/records/{namespace}/{resource}/{id}', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('DELETE','/records/{namespace}/{resource}/{id}', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/records/{namespace}/{resource}/{id}");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("DELETE");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("DELETE", "/records/{namespace}/{resource}/{id}", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`DELETE /records/{namespace}/{resource}/{id}`

*Delete*

<h3 id="record_delete-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|namespace|path|string|true|none|
|resource|path|string|true|none|
|id|path|string|true|none|
|token|query|string|false|none|

> Example responses

> 200 Response

```json
{}
```

<h3 id="record_delete-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[DeleteRecordResponse](#schemadeleterecordresponse)|
|default|Default|Default error response|[Status](#schemastatus)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

## Record_UpdateMulti

<a id="opIdRecord_UpdateMulti"></a>

> Code samples

```shell
# You can also use wget
curl -X POST /records/{namespace}/{resource}/_multi \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
POST /records/{namespace}/{resource}/_multi HTTP/1.1

Accept: application/json

```

```javascript

const headers = {
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('/records/{namespace}/{resource}/_multi',
{
  method: 'POST',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.post '/records/{namespace}/{resource}/_multi',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.post('/records/{namespace}/{resource}/_multi', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('POST','/records/{namespace}/{resource}/_multi', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/records/{namespace}/{resource}/_multi");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("POST");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("POST", "/records/{namespace}/{resource}/_multi", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /records/{namespace}/{resource}/_multi`

*UpdateMulti*

<h3 id="record_updatemulti-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|namespace|path|string|true|none|
|resource|path|string|true|none|
|token|query|string|false|none|
|query.not.not.equal.left.property|query|string|false|none|
|query.not.not.equal.left.value|query|[GoogleProtobufValue](#schemagoogleprotobufvalue)|false|none|
|query.not.not.equal.left.refValue.namespace|query|string|false|none|
|query.not.not.equal.left.refValue.resource|query|string|false|none|
|query.not.not.equal.right.property|query|string|false|none|
|query.not.not.equal.right.value|query|[GoogleProtobufValue](#schemagoogleprotobufvalue)|false|none|
|query.not.not.equal.right.refValue.namespace|query|string|false|none|
|query.not.not.equal.right.refValue.resource|query|string|false|none|
|query.not.not.regexMatch.pattern|query|string|false|none|
|query.not.regexMatch.pattern|query|string|false|none|

> Example responses

> 200 Response

```json
{
  "record": {
    "id": "string",
    "properties": {
      "property1": null,
      "property2": null
    },
    "propertiesPacked": [
      null
    ]
  },
  "records": [
    {
      "id": "string",
      "properties": {
        "property1": null,
        "property2": null
      },
      "propertiesPacked": [
        null
      ]
    }
  ]
}
```

<h3 id="record_updatemulti-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[UpdateMultiRecordResponse](#schemaupdatemultirecordresponse)|
|default|Default|Default error response|[Status](#schemastatus)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

## Record_Search

<a id="opIdRecord_Search"></a>

> Code samples

```shell
# You can also use wget
curl -X POST /records/{namespace}/{resource}/_search \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
POST /records/{namespace}/{resource}/_search HTTP/1.1

Content-Type: application/json
Accept: application/json

```

```javascript
const inputBody = '{
  "token": "string",
  "namespace": "string",
  "resource": "string",
  "query": {
    "and": {
      "expressions": [
        {}
      ]
    },
    "or": {
      "expressions": [
        {}
      ]
    },
    "not": {},
    "equal": {
      "left": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "right": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    },
    "lessThan": {
      "left": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "right": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    },
    "greaterThan": {
      "left": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "right": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    },
    "lessThanOrEqual": {
      "left": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "right": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    },
    "greaterThanOrEqual": {
      "left": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "right": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    },
    "in": {
      "left": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "right": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    },
    "isNull": {
      "additionalProperties": [
        {
          "name": "string",
          "value": {
            "value": {
              "@type": "string"
            },
            "yaml": "string"
          }
        }
      ]
    },
    "regexMatch": {
      "pattern": "string",
      "expression": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    }
  },
  "limit": 0,
  "offset": 0,
  "useHistory": true,
  "resolveReferences": [
    "string"
  ],
  "annotations": {
    "property1": "string",
    "property2": "string"
  }
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('/records/{namespace}/{resource}/_search',
{
  method: 'POST',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Content-Type' => 'application/json',
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.post '/records/{namespace}/{resource}/_search',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Content-Type': 'application/json',
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.post('/records/{namespace}/{resource}/_search', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Content-Type' => 'application/json',
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('POST','/records/{namespace}/{resource}/_search', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/records/{namespace}/{resource}/_search");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("POST");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Content-Type": []string{"application/json"},
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("POST", "/records/{namespace}/{resource}/_search", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /records/{namespace}/{resource}/_search`

*Search*

> Body parameter

```json
{
  "token": "string",
  "namespace": "string",
  "resource": "string",
  "query": {
    "and": {
      "expressions": [
        {}
      ]
    },
    "or": {
      "expressions": [
        {}
      ]
    },
    "not": {},
    "equal": {
      "left": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "right": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    },
    "lessThan": {
      "left": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "right": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    },
    "greaterThan": {
      "left": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "right": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    },
    "lessThanOrEqual": {
      "left": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "right": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    },
    "greaterThanOrEqual": {
      "left": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "right": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    },
    "in": {
      "left": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "right": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    },
    "isNull": {
      "additionalProperties": [
        {
          "name": "string",
          "value": {
            "value": {
              "@type": "string"
            },
            "yaml": "string"
          }
        }
      ]
    },
    "regexMatch": {
      "pattern": "string",
      "expression": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    }
  },
  "limit": 0,
  "offset": 0,
  "useHistory": true,
  "resolveReferences": [
    "string"
  ],
  "annotations": {
    "property1": "string",
    "property2": "string"
  }
}
```

<h3 id="record_search-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|namespace|path|string|true|none|
|resource|path|string|true|none|
|body|body|[SearchRecordRequest](#schemasearchrecordrequest)|true|none|

> Example responses

> 200 Response

```json
{
  "total": 0,
  "content": [
    {
      "id": "string",
      "properties": {
        "property1": null,
        "property2": null
      },
      "propertiesPacked": [
        null
      ]
    }
  ]
}
```

<h3 id="record_search-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[SearchRecordResponse](#schemasearchrecordresponse)|
|default|Default|Default error response|[Status](#schemastatus)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

## Record_Get

<a id="opIdRecord_Get"></a>

> Code samples

```shell
# You can also use wget
curl -X GET /records/{namespace}/{resource}/{id} \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
GET /records/{namespace}/{resource}/{id} HTTP/1.1

Accept: application/json

```

```javascript

const headers = {
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('/records/{namespace}/{resource}/{id}',
{
  method: 'GET',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.get '/records/{namespace}/{resource}/{id}',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.get('/records/{namespace}/{resource}/{id}', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('GET','/records/{namespace}/{resource}/{id}', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/records/{namespace}/{resource}/{id}");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("GET");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "/records/{namespace}/{resource}/{id}", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /records/{namespace}/{resource}/{id}`

*Get*

<h3 id="record_get-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|namespace|path|string|true|none|
|resource|path|string|true|none|
|id|path|string|true|none|
|token|query|string|false|none|

> Example responses

> 200 Response

```json
{
  "record": {
    "id": "string",
    "properties": {
      "property1": null,
      "property2": null
    },
    "propertiesPacked": [
      null
    ]
  }
}
```

<h3 id="record_get-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[GetRecordResponse](#schemagetrecordresponse)|
|default|Default|Default error response|[Status](#schemastatus)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

<h1 id="apibrew-resource">Resource</h1>

Resource service is for managing resources

## Resource_List

<a id="opIdResource_List"></a>

> Code samples

```shell
# You can also use wget
curl -X GET /system/resources \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
GET /system/resources HTTP/1.1

Accept: application/json

```

```javascript

const headers = {
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('/system/resources',
{
  method: 'GET',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.get '/system/resources',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.get('/system/resources', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('GET','/system/resources', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/system/resources");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("GET");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "/system/resources", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /system/resources`

*List*

<h3 id="resource_list-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|token|query|string|false|none|

> Example responses

> 200 Response

```json
{
  "resources": [
    {
      "id": "string",
      "name": "string",
      "namespace": "string",
      "sourceConfig": {
        "dataSource": "string",
        "catalog": "string",
        "entity": "string"
      },
      "properties": [
        {
          "id": "string",
          "name": "string",
          "type": "BOOL",
          "mapping": "string",
          "required": true,
          "primary": true,
          "length": 0,
          "unique": true,
          "immutable": true,
          "securityContext": {
            "constraints": [
              {
                "namespace": "string",
                "resource": "string",
                "property": "string",
                "before": "2019-08-24T14:15:22Z",
                "after": "2019-08-24T14:15:22Z",
                "principal": "string",
                "recordIds": [
                  "string"
                ],
                "operation": "OPERATION_TYPE_READ",
                "permit": "PERMIT_TYPE_ALLOW"
              }
            ]
          },
          "defaultValue": null,
          "exampleValue": null,
          "enumValues": [
            null
          ],
          "reference": {
            "referencedResource": "string",
            "cascade": true
          },
          "properties": [
            {}
          ],
          "Item": {},
          "title": "string",
          "description": "string",
          "annotations": {
            "property1": "string",
            "property2": "string"
          }
        }
      ],
      "indexes": [
        {
          "properties": [
            {
              "name": "string",
              "order": "ORDER_UNKNOWN"
            }
          ],
          "indexType": "BTREE",
          "unique": true,
          "annotations": {
            "property1": "string",
            "property2": "string"
          }
        }
      ],
      "securityContext": {
        "constraints": [
          {
            "namespace": "string",
            "resource": "string",
            "property": "string",
            "before": "2019-08-24T14:15:22Z",
            "after": "2019-08-24T14:15:22Z",
            "principal": "string",
            "recordIds": [
              "string"
            ],
            "operation": "OPERATION_TYPE_READ",
            "permit": "PERMIT_TYPE_ALLOW"
          }
        ]
      },
      "virtual": true,
      "immutable": true,
      "abstract": true,
      "title": "string",
      "description": "string",
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0,
      "annotations": {
        "property1": "string",
        "property2": "string"
      }
    }
  ]
}
```

<h3 id="resource_list-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[ListResourceResponse](#schemalistresourceresponse)|
|default|Default|Default error response|[Status](#schemastatus)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

## Resource_Update

<a id="opIdResource_Update"></a>

> Code samples

```shell
# You can also use wget
curl -X PUT /system/resources \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
PUT /system/resources HTTP/1.1

Content-Type: application/json
Accept: application/json

```

```javascript
const inputBody = '{
  "token": "string",
  "resources": [
    {
      "id": "string",
      "name": "string",
      "namespace": "string",
      "sourceConfig": {
        "dataSource": "string",
        "catalog": "string",
        "entity": "string"
      },
      "properties": [
        {
          "id": "string",
          "name": "string",
          "type": "BOOL",
          "mapping": "string",
          "required": true,
          "primary": true,
          "length": 0,
          "unique": true,
          "immutable": true,
          "securityContext": {
            "constraints": [
              {
                "namespace": "string",
                "resource": "string",
                "property": "string",
                "before": "2019-08-24T14:15:22Z",
                "after": "2019-08-24T14:15:22Z",
                "principal": "string",
                "recordIds": [
                  "string"
                ],
                "operation": "OPERATION_TYPE_READ",
                "permit": "PERMIT_TYPE_ALLOW"
              }
            ]
          },
          "defaultValue": null,
          "exampleValue": null,
          "enumValues": [
            null
          ],
          "reference": {
            "referencedResource": "string",
            "cascade": true
          },
          "properties": [
            {}
          ],
          "Item": {},
          "title": "string",
          "description": "string",
          "annotations": {
            "property1": "string",
            "property2": "string"
          }
        }
      ],
      "indexes": [
        {
          "properties": [
            {
              "name": "string",
              "order": "ORDER_UNKNOWN"
            }
          ],
          "indexType": "BTREE",
          "unique": true,
          "annotations": {
            "property1": "string",
            "property2": "string"
          }
        }
      ],
      "securityContext": {
        "constraints": [
          {
            "namespace": "string",
            "resource": "string",
            "property": "string",
            "before": "2019-08-24T14:15:22Z",
            "after": "2019-08-24T14:15:22Z",
            "principal": "string",
            "recordIds": [
              "string"
            ],
            "operation": "OPERATION_TYPE_READ",
            "permit": "PERMIT_TYPE_ALLOW"
          }
        ]
      },
      "virtual": true,
      "immutable": true,
      "abstract": true,
      "title": "string",
      "description": "string",
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0,
      "annotations": {
        "property1": "string",
        "property2": "string"
      }
    }
  ],
  "doMigration": true,
  "forceMigration": true,
  "annotations": {
    "property1": "string",
    "property2": "string"
  }
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('/system/resources',
{
  method: 'PUT',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Content-Type' => 'application/json',
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.put '/system/resources',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Content-Type': 'application/json',
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.put('/system/resources', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Content-Type' => 'application/json',
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('PUT','/system/resources', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/system/resources");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("PUT");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Content-Type": []string{"application/json"},
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("PUT", "/system/resources", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`PUT /system/resources`

*Update*

> Body parameter

```json
{
  "token": "string",
  "resources": [
    {
      "id": "string",
      "name": "string",
      "namespace": "string",
      "sourceConfig": {
        "dataSource": "string",
        "catalog": "string",
        "entity": "string"
      },
      "properties": [
        {
          "id": "string",
          "name": "string",
          "type": "BOOL",
          "mapping": "string",
          "required": true,
          "primary": true,
          "length": 0,
          "unique": true,
          "immutable": true,
          "securityContext": {
            "constraints": [
              {
                "namespace": "string",
                "resource": "string",
                "property": "string",
                "before": "2019-08-24T14:15:22Z",
                "after": "2019-08-24T14:15:22Z",
                "principal": "string",
                "recordIds": [
                  "string"
                ],
                "operation": "OPERATION_TYPE_READ",
                "permit": "PERMIT_TYPE_ALLOW"
              }
            ]
          },
          "defaultValue": null,
          "exampleValue": null,
          "enumValues": [
            null
          ],
          "reference": {
            "referencedResource": "string",
            "cascade": true
          },
          "properties": [
            {}
          ],
          "Item": {},
          "title": "string",
          "description": "string",
          "annotations": {
            "property1": "string",
            "property2": "string"
          }
        }
      ],
      "indexes": [
        {
          "properties": [
            {
              "name": "string",
              "order": "ORDER_UNKNOWN"
            }
          ],
          "indexType": "BTREE",
          "unique": true,
          "annotations": {
            "property1": "string",
            "property2": "string"
          }
        }
      ],
      "securityContext": {
        "constraints": [
          {
            "namespace": "string",
            "resource": "string",
            "property": "string",
            "before": "2019-08-24T14:15:22Z",
            "after": "2019-08-24T14:15:22Z",
            "principal": "string",
            "recordIds": [
              "string"
            ],
            "operation": "OPERATION_TYPE_READ",
            "permit": "PERMIT_TYPE_ALLOW"
          }
        ]
      },
      "virtual": true,
      "immutable": true,
      "abstract": true,
      "title": "string",
      "description": "string",
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0,
      "annotations": {
        "property1": "string",
        "property2": "string"
      }
    }
  ],
  "doMigration": true,
  "forceMigration": true,
  "annotations": {
    "property1": "string",
    "property2": "string"
  }
}
```

<h3 id="resource_update-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[UpdateResourceRequest](#schemaupdateresourcerequest)|true|none|

> Example responses

> 200 Response

```json
{
  "resources": [
    {
      "id": "string",
      "name": "string",
      "namespace": "string",
      "sourceConfig": {
        "dataSource": "string",
        "catalog": "string",
        "entity": "string"
      },
      "properties": [
        {
          "id": "string",
          "name": "string",
          "type": "BOOL",
          "mapping": "string",
          "required": true,
          "primary": true,
          "length": 0,
          "unique": true,
          "immutable": true,
          "securityContext": {
            "constraints": [
              {
                "namespace": "string",
                "resource": "string",
                "property": "string",
                "before": "2019-08-24T14:15:22Z",
                "after": "2019-08-24T14:15:22Z",
                "principal": "string",
                "recordIds": [
                  "string"
                ],
                "operation": "OPERATION_TYPE_READ",
                "permit": "PERMIT_TYPE_ALLOW"
              }
            ]
          },
          "defaultValue": null,
          "exampleValue": null,
          "enumValues": [
            null
          ],
          "reference": {
            "referencedResource": "string",
            "cascade": true
          },
          "properties": [
            {}
          ],
          "Item": {},
          "title": "string",
          "description": "string",
          "annotations": {
            "property1": "string",
            "property2": "string"
          }
        }
      ],
      "indexes": [
        {
          "properties": [
            {
              "name": "string",
              "order": "ORDER_UNKNOWN"
            }
          ],
          "indexType": "BTREE",
          "unique": true,
          "annotations": {
            "property1": "string",
            "property2": "string"
          }
        }
      ],
      "securityContext": {
        "constraints": [
          {
            "namespace": "string",
            "resource": "string",
            "property": "string",
            "before": "2019-08-24T14:15:22Z",
            "after": "2019-08-24T14:15:22Z",
            "principal": "string",
            "recordIds": [
              "string"
            ],
            "operation": "OPERATION_TYPE_READ",
            "permit": "PERMIT_TYPE_ALLOW"
          }
        ]
      },
      "virtual": true,
      "immutable": true,
      "abstract": true,
      "title": "string",
      "description": "string",
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0,
      "annotations": {
        "property1": "string",
        "property2": "string"
      }
    }
  ]
}
```

<h3 id="resource_update-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[UpdateResourceResponse](#schemaupdateresourceresponse)|
|default|Default|Default error response|[Status](#schemastatus)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

## Resource_Create

<a id="opIdResource_Create"></a>

> Code samples

```shell
# You can also use wget
curl -X POST /system/resources \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
POST /system/resources HTTP/1.1

Content-Type: application/json
Accept: application/json

```

```javascript
const inputBody = '{
  "token": "string",
  "resources": [
    {
      "id": "string",
      "name": "string",
      "namespace": "string",
      "sourceConfig": {
        "dataSource": "string",
        "catalog": "string",
        "entity": "string"
      },
      "properties": [
        {
          "id": "string",
          "name": "string",
          "type": "BOOL",
          "mapping": "string",
          "required": true,
          "primary": true,
          "length": 0,
          "unique": true,
          "immutable": true,
          "securityContext": {
            "constraints": [
              {
                "namespace": "string",
                "resource": "string",
                "property": "string",
                "before": "2019-08-24T14:15:22Z",
                "after": "2019-08-24T14:15:22Z",
                "principal": "string",
                "recordIds": [
                  "string"
                ],
                "operation": "OPERATION_TYPE_READ",
                "permit": "PERMIT_TYPE_ALLOW"
              }
            ]
          },
          "defaultValue": null,
          "exampleValue": null,
          "enumValues": [
            null
          ],
          "reference": {
            "referencedResource": "string",
            "cascade": true
          },
          "properties": [
            {}
          ],
          "Item": {},
          "title": "string",
          "description": "string",
          "annotations": {
            "property1": "string",
            "property2": "string"
          }
        }
      ],
      "indexes": [
        {
          "properties": [
            {
              "name": "string",
              "order": "ORDER_UNKNOWN"
            }
          ],
          "indexType": "BTREE",
          "unique": true,
          "annotations": {
            "property1": "string",
            "property2": "string"
          }
        }
      ],
      "securityContext": {
        "constraints": [
          {
            "namespace": "string",
            "resource": "string",
            "property": "string",
            "before": "2019-08-24T14:15:22Z",
            "after": "2019-08-24T14:15:22Z",
            "principal": "string",
            "recordIds": [
              "string"
            ],
            "operation": "OPERATION_TYPE_READ",
            "permit": "PERMIT_TYPE_ALLOW"
          }
        ]
      },
      "virtual": true,
      "immutable": true,
      "abstract": true,
      "title": "string",
      "description": "string",
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0,
      "annotations": {
        "property1": "string",
        "property2": "string"
      }
    }
  ],
  "doMigration": true,
  "forceMigration": true,
  "annotations": {
    "property1": "string",
    "property2": "string"
  }
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('/system/resources',
{
  method: 'POST',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Content-Type' => 'application/json',
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.post '/system/resources',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Content-Type': 'application/json',
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.post('/system/resources', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Content-Type' => 'application/json',
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('POST','/system/resources', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/system/resources");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("POST");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Content-Type": []string{"application/json"},
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("POST", "/system/resources", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /system/resources`

*Create*

> Body parameter

```json
{
  "token": "string",
  "resources": [
    {
      "id": "string",
      "name": "string",
      "namespace": "string",
      "sourceConfig": {
        "dataSource": "string",
        "catalog": "string",
        "entity": "string"
      },
      "properties": [
        {
          "id": "string",
          "name": "string",
          "type": "BOOL",
          "mapping": "string",
          "required": true,
          "primary": true,
          "length": 0,
          "unique": true,
          "immutable": true,
          "securityContext": {
            "constraints": [
              {
                "namespace": "string",
                "resource": "string",
                "property": "string",
                "before": "2019-08-24T14:15:22Z",
                "after": "2019-08-24T14:15:22Z",
                "principal": "string",
                "recordIds": [
                  "string"
                ],
                "operation": "OPERATION_TYPE_READ",
                "permit": "PERMIT_TYPE_ALLOW"
              }
            ]
          },
          "defaultValue": null,
          "exampleValue": null,
          "enumValues": [
            null
          ],
          "reference": {
            "referencedResource": "string",
            "cascade": true
          },
          "properties": [
            {}
          ],
          "Item": {},
          "title": "string",
          "description": "string",
          "annotations": {
            "property1": "string",
            "property2": "string"
          }
        }
      ],
      "indexes": [
        {
          "properties": [
            {
              "name": "string",
              "order": "ORDER_UNKNOWN"
            }
          ],
          "indexType": "BTREE",
          "unique": true,
          "annotations": {
            "property1": "string",
            "property2": "string"
          }
        }
      ],
      "securityContext": {
        "constraints": [
          {
            "namespace": "string",
            "resource": "string",
            "property": "string",
            "before": "2019-08-24T14:15:22Z",
            "after": "2019-08-24T14:15:22Z",
            "principal": "string",
            "recordIds": [
              "string"
            ],
            "operation": "OPERATION_TYPE_READ",
            "permit": "PERMIT_TYPE_ALLOW"
          }
        ]
      },
      "virtual": true,
      "immutable": true,
      "abstract": true,
      "title": "string",
      "description": "string",
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0,
      "annotations": {
        "property1": "string",
        "property2": "string"
      }
    }
  ],
  "doMigration": true,
  "forceMigration": true,
  "annotations": {
    "property1": "string",
    "property2": "string"
  }
}
```

<h3 id="resource_create-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[CreateResourceRequest](#schemacreateresourcerequest)|true|none|

> Example responses

> 200 Response

```json
{
  "resources": [
    {
      "id": "string",
      "name": "string",
      "namespace": "string",
      "sourceConfig": {
        "dataSource": "string",
        "catalog": "string",
        "entity": "string"
      },
      "properties": [
        {
          "id": "string",
          "name": "string",
          "type": "BOOL",
          "mapping": "string",
          "required": true,
          "primary": true,
          "length": 0,
          "unique": true,
          "immutable": true,
          "securityContext": {
            "constraints": [
              {
                "namespace": "string",
                "resource": "string",
                "property": "string",
                "before": "2019-08-24T14:15:22Z",
                "after": "2019-08-24T14:15:22Z",
                "principal": "string",
                "recordIds": [
                  "string"
                ],
                "operation": "OPERATION_TYPE_READ",
                "permit": "PERMIT_TYPE_ALLOW"
              }
            ]
          },
          "defaultValue": null,
          "exampleValue": null,
          "enumValues": [
            null
          ],
          "reference": {
            "referencedResource": "string",
            "cascade": true
          },
          "properties": [
            {}
          ],
          "Item": {},
          "title": "string",
          "description": "string",
          "annotations": {
            "property1": "string",
            "property2": "string"
          }
        }
      ],
      "indexes": [
        {
          "properties": [
            {
              "name": "string",
              "order": "ORDER_UNKNOWN"
            }
          ],
          "indexType": "BTREE",
          "unique": true,
          "annotations": {
            "property1": "string",
            "property2": "string"
          }
        }
      ],
      "securityContext": {
        "constraints": [
          {
            "namespace": "string",
            "resource": "string",
            "property": "string",
            "before": "2019-08-24T14:15:22Z",
            "after": "2019-08-24T14:15:22Z",
            "principal": "string",
            "recordIds": [
              "string"
            ],
            "operation": "OPERATION_TYPE_READ",
            "permit": "PERMIT_TYPE_ALLOW"
          }
        ]
      },
      "virtual": true,
      "immutable": true,
      "abstract": true,
      "title": "string",
      "description": "string",
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0,
      "annotations": {
        "property1": "string",
        "property2": "string"
      }
    }
  ]
}
```

<h3 id="resource_create-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[CreateResourceResponse](#schemacreateresourceresponse)|
|default|Default|Default error response|[Status](#schemastatus)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

## Resource_Delete

<a id="opIdResource_Delete"></a>

> Code samples

```shell
# You can also use wget
curl -X DELETE /system/resources \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
DELETE /system/resources HTTP/1.1

Content-Type: application/json
Accept: application/json

```

```javascript
const inputBody = '{
  "token": "string",
  "ids": [
    "string"
  ],
  "doMigration": true,
  "forceMigration": true,
  "annotations": {
    "property1": "string",
    "property2": "string"
  }
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('/system/resources',
{
  method: 'DELETE',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Content-Type' => 'application/json',
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.delete '/system/resources',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Content-Type': 'application/json',
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.delete('/system/resources', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Content-Type' => 'application/json',
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('DELETE','/system/resources', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/system/resources");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("DELETE");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Content-Type": []string{"application/json"},
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("DELETE", "/system/resources", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`DELETE /system/resources`

*Delete*

> Body parameter

```json
{
  "token": "string",
  "ids": [
    "string"
  ],
  "doMigration": true,
  "forceMigration": true,
  "annotations": {
    "property1": "string",
    "property2": "string"
  }
}
```

<h3 id="resource_delete-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[DeleteResourceRequest](#schemadeleteresourcerequest)|true|none|

> Example responses

> 200 Response

```json
{}
```

<h3 id="resource_delete-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[DeleteResourceResponse](#schemadeleteresourceresponse)|
|default|Default|Default error response|[Status](#schemastatus)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

## Resource_PrepareResourceMigrationPlan

<a id="opIdResource_PrepareResourceMigrationPlan"></a>

> Code samples

```shell
# You can also use wget
curl -X POST /system/resources/_migrate \
  -H 'Content-Type: application/json' \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
POST /system/resources/_migrate HTTP/1.1

Content-Type: application/json
Accept: application/json

```

```javascript
const inputBody = '{
  "token": "string",
  "prepareFromDataSource": true,
  "resources": [
    {
      "id": "string",
      "name": "string",
      "namespace": "string",
      "sourceConfig": {
        "dataSource": "string",
        "catalog": "string",
        "entity": "string"
      },
      "properties": [
        {
          "id": "string",
          "name": "string",
          "type": "BOOL",
          "mapping": "string",
          "required": true,
          "primary": true,
          "length": 0,
          "unique": true,
          "immutable": true,
          "securityContext": {
            "constraints": [
              {
                "namespace": "string",
                "resource": "string",
                "property": "string",
                "before": "2019-08-24T14:15:22Z",
                "after": "2019-08-24T14:15:22Z",
                "principal": "string",
                "recordIds": [
                  "string"
                ],
                "operation": "OPERATION_TYPE_READ",
                "permit": "PERMIT_TYPE_ALLOW"
              }
            ]
          },
          "defaultValue": null,
          "exampleValue": null,
          "enumValues": [
            null
          ],
          "reference": {
            "referencedResource": "string",
            "cascade": true
          },
          "properties": [
            {}
          ],
          "Item": {},
          "title": "string",
          "description": "string",
          "annotations": {
            "property1": "string",
            "property2": "string"
          }
        }
      ],
      "indexes": [
        {
          "properties": [
            {
              "name": "string",
              "order": "ORDER_UNKNOWN"
            }
          ],
          "indexType": "BTREE",
          "unique": true,
          "annotations": {
            "property1": "string",
            "property2": "string"
          }
        }
      ],
      "securityContext": {
        "constraints": [
          {
            "namespace": "string",
            "resource": "string",
            "property": "string",
            "before": "2019-08-24T14:15:22Z",
            "after": "2019-08-24T14:15:22Z",
            "principal": "string",
            "recordIds": [
              "string"
            ],
            "operation": "OPERATION_TYPE_READ",
            "permit": "PERMIT_TYPE_ALLOW"
          }
        ]
      },
      "virtual": true,
      "immutable": true,
      "abstract": true,
      "title": "string",
      "description": "string",
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0,
      "annotations": {
        "property1": "string",
        "property2": "string"
      }
    }
  ],
  "annotations": {
    "property1": "string",
    "property2": "string"
  }
}';
const headers = {
  'Content-Type':'application/json',
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('/system/resources/_migrate',
{
  method: 'POST',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Content-Type' => 'application/json',
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.post '/system/resources/_migrate',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Content-Type': 'application/json',
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.post('/system/resources/_migrate', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Content-Type' => 'application/json',
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('POST','/system/resources/_migrate', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/system/resources/_migrate");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("POST");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Content-Type": []string{"application/json"},
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("POST", "/system/resources/_migrate", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /system/resources/_migrate`

*PrepareResourceMigrationPlan*

> Body parameter

```json
{
  "token": "string",
  "prepareFromDataSource": true,
  "resources": [
    {
      "id": "string",
      "name": "string",
      "namespace": "string",
      "sourceConfig": {
        "dataSource": "string",
        "catalog": "string",
        "entity": "string"
      },
      "properties": [
        {
          "id": "string",
          "name": "string",
          "type": "BOOL",
          "mapping": "string",
          "required": true,
          "primary": true,
          "length": 0,
          "unique": true,
          "immutable": true,
          "securityContext": {
            "constraints": [
              {
                "namespace": "string",
                "resource": "string",
                "property": "string",
                "before": "2019-08-24T14:15:22Z",
                "after": "2019-08-24T14:15:22Z",
                "principal": "string",
                "recordIds": [
                  "string"
                ],
                "operation": "OPERATION_TYPE_READ",
                "permit": "PERMIT_TYPE_ALLOW"
              }
            ]
          },
          "defaultValue": null,
          "exampleValue": null,
          "enumValues": [
            null
          ],
          "reference": {
            "referencedResource": "string",
            "cascade": true
          },
          "properties": [
            {}
          ],
          "Item": {},
          "title": "string",
          "description": "string",
          "annotations": {
            "property1": "string",
            "property2": "string"
          }
        }
      ],
      "indexes": [
        {
          "properties": [
            {
              "name": "string",
              "order": "ORDER_UNKNOWN"
            }
          ],
          "indexType": "BTREE",
          "unique": true,
          "annotations": {
            "property1": "string",
            "property2": "string"
          }
        }
      ],
      "securityContext": {
        "constraints": [
          {
            "namespace": "string",
            "resource": "string",
            "property": "string",
            "before": "2019-08-24T14:15:22Z",
            "after": "2019-08-24T14:15:22Z",
            "principal": "string",
            "recordIds": [
              "string"
            ],
            "operation": "OPERATION_TYPE_READ",
            "permit": "PERMIT_TYPE_ALLOW"
          }
        ]
      },
      "virtual": true,
      "immutable": true,
      "abstract": true,
      "title": "string",
      "description": "string",
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0,
      "annotations": {
        "property1": "string",
        "property2": "string"
      }
    }
  ],
  "annotations": {
    "property1": "string",
    "property2": "string"
  }
}
```

<h3 id="resource_prepareresourcemigrationplan-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[PrepareResourceMigrationPlanRequest](#schemaprepareresourcemigrationplanrequest)|true|none|

> Example responses

> 200 Response

```json
{
  "plans": [
    {
      "existingResource": {
        "id": "string",
        "name": "string",
        "namespace": "string",
        "sourceConfig": {
          "dataSource": "string",
          "catalog": "string",
          "entity": "string"
        },
        "properties": [
          {
            "id": "string",
            "name": "string",
            "type": "BOOL",
            "mapping": "string",
            "required": true,
            "primary": true,
            "length": 0,
            "unique": true,
            "immutable": true,
            "securityContext": {
              "constraints": [
                {
                  "namespace": "string",
                  "resource": "string",
                  "property": "string",
                  "before": "2019-08-24T14:15:22Z",
                  "after": "2019-08-24T14:15:22Z",
                  "principal": "string",
                  "recordIds": [
                    "string"
                  ],
                  "operation": "OPERATION_TYPE_READ",
                  "permit": "PERMIT_TYPE_ALLOW"
                }
              ]
            },
            "defaultValue": null,
            "exampleValue": null,
            "enumValues": [
              null
            ],
            "reference": {
              "referencedResource": "string",
              "cascade": true
            },
            "properties": [
              {}
            ],
            "Item": {},
            "title": "string",
            "description": "string",
            "annotations": {
              "property1": "string",
              "property2": "string"
            }
          }
        ],
        "indexes": [
          {
            "properties": [
              {
                "name": "string",
                "order": "ORDER_UNKNOWN"
              }
            ],
            "indexType": "BTREE",
            "unique": true,
            "annotations": {
              "property1": "string",
              "property2": "string"
            }
          }
        ],
        "securityContext": {
          "constraints": [
            {
              "namespace": "string",
              "resource": "string",
              "property": "string",
              "before": "2019-08-24T14:15:22Z",
              "after": "2019-08-24T14:15:22Z",
              "principal": "string",
              "recordIds": [
                "string"
              ],
              "operation": "OPERATION_TYPE_READ",
              "permit": "PERMIT_TYPE_ALLOW"
            }
          ]
        },
        "virtual": true,
        "immutable": true,
        "abstract": true,
        "title": "string",
        "description": "string",
        "auditData": {
          "createdOn": "2019-08-24T14:15:22Z",
          "updatedOn": "2019-08-24T14:15:22Z",
          "createdBy": "string",
          "updatedBy": "string"
        },
        "version": 0,
        "annotations": {
          "property1": "string",
          "property2": "string"
        }
      },
      "currentResource": {
        "id": "string",
        "name": "string",
        "namespace": "string",
        "sourceConfig": {
          "dataSource": "string",
          "catalog": "string",
          "entity": "string"
        },
        "properties": [
          {
            "id": "string",
            "name": "string",
            "type": "BOOL",
            "mapping": "string",
            "required": true,
            "primary": true,
            "length": 0,
            "unique": true,
            "immutable": true,
            "securityContext": {
              "constraints": [
                {
                  "namespace": "string",
                  "resource": "string",
                  "property": "string",
                  "before": "2019-08-24T14:15:22Z",
                  "after": "2019-08-24T14:15:22Z",
                  "principal": "string",
                  "recordIds": [
                    "string"
                  ],
                  "operation": "OPERATION_TYPE_READ",
                  "permit": "PERMIT_TYPE_ALLOW"
                }
              ]
            },
            "defaultValue": null,
            "exampleValue": null,
            "enumValues": [
              null
            ],
            "reference": {
              "referencedResource": "string",
              "cascade": true
            },
            "properties": [
              {}
            ],
            "Item": {},
            "title": "string",
            "description": "string",
            "annotations": {
              "property1": "string",
              "property2": "string"
            }
          }
        ],
        "indexes": [
          {
            "properties": [
              {
                "name": "string",
                "order": "ORDER_UNKNOWN"
              }
            ],
            "indexType": "BTREE",
            "unique": true,
            "annotations": {
              "property1": "string",
              "property2": "string"
            }
          }
        ],
        "securityContext": {
          "constraints": [
            {
              "namespace": "string",
              "resource": "string",
              "property": "string",
              "before": "2019-08-24T14:15:22Z",
              "after": "2019-08-24T14:15:22Z",
              "principal": "string",
              "recordIds": [
                "string"
              ],
              "operation": "OPERATION_TYPE_READ",
              "permit": "PERMIT_TYPE_ALLOW"
            }
          ]
        },
        "virtual": true,
        "immutable": true,
        "abstract": true,
        "title": "string",
        "description": "string",
        "auditData": {
          "createdOn": "2019-08-24T14:15:22Z",
          "updatedOn": "2019-08-24T14:15:22Z",
          "createdBy": "string",
          "updatedBy": "string"
        },
        "version": 0,
        "annotations": {
          "property1": "string",
          "property2": "string"
        }
      },
      "steps": [
        {
          "createResource": {},
          "deleteResource": {},
          "updateResource": {
            "changedFields": [
              "string"
            ]
          },
          "createProperty": {
            "property": "string"
          },
          "deleteProperty": {
            "existingProperty": "string"
          },
          "updateProperty": {
            "existingProperty": "string",
            "property": "string",
            "changedFields": [
              "string"
            ]
          },
          "createIndex": {
            "index": 0
          },
          "deleteIndex": {
            "existingIndex": 0
          }
        }
      ]
    }
  ]
}
```

<h3 id="resource_prepareresourcemigrationplan-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[PrepareResourceMigrationPlanResponse](#schemaprepareresourcemigrationplanresponse)|
|default|Default|Default error response|[Status](#schemastatus)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

## Resource_Get

<a id="opIdResource_Get"></a>

> Code samples

```shell
# You can also use wget
curl -X GET /system/resources/{id} \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
GET /system/resources/{id} HTTP/1.1

Accept: application/json

```

```javascript

const headers = {
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('/system/resources/{id}',
{
  method: 'GET',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.get '/system/resources/{id}',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.get('/system/resources/{id}', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('GET','/system/resources/{id}', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/system/resources/{id}");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("GET");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "/system/resources/{id}", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /system/resources/{id}`

*Get*

<h3 id="resource_get-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|id|path|string|true|none|
|token|query|string|false|none|

> Example responses

> 200 Response

```json
{
  "resource": {
    "id": "string",
    "name": "string",
    "namespace": "string",
    "sourceConfig": {
      "dataSource": "string",
      "catalog": "string",
      "entity": "string"
    },
    "properties": [
      {
        "id": "string",
        "name": "string",
        "type": "BOOL",
        "mapping": "string",
        "required": true,
        "primary": true,
        "length": 0,
        "unique": true,
        "immutable": true,
        "securityContext": {
          "constraints": [
            {
              "namespace": "string",
              "resource": "string",
              "property": "string",
              "before": "2019-08-24T14:15:22Z",
              "after": "2019-08-24T14:15:22Z",
              "principal": "string",
              "recordIds": [
                "string"
              ],
              "operation": "OPERATION_TYPE_READ",
              "permit": "PERMIT_TYPE_ALLOW"
            }
          ]
        },
        "defaultValue": null,
        "exampleValue": null,
        "enumValues": [
          null
        ],
        "reference": {
          "referencedResource": "string",
          "cascade": true
        },
        "properties": [
          {}
        ],
        "Item": {},
        "title": "string",
        "description": "string",
        "annotations": {
          "property1": "string",
          "property2": "string"
        }
      }
    ],
    "indexes": [
      {
        "properties": [
          {
            "name": "string",
            "order": "ORDER_UNKNOWN"
          }
        ],
        "indexType": "BTREE",
        "unique": true,
        "annotations": {
          "property1": "string",
          "property2": "string"
        }
      }
    ],
    "securityContext": {
      "constraints": [
        {
          "namespace": "string",
          "resource": "string",
          "property": "string",
          "before": "2019-08-24T14:15:22Z",
          "after": "2019-08-24T14:15:22Z",
          "principal": "string",
          "recordIds": [
            "string"
          ],
          "operation": "OPERATION_TYPE_READ",
          "permit": "PERMIT_TYPE_ALLOW"
        }
      ]
    },
    "virtual": true,
    "immutable": true,
    "abstract": true,
    "title": "string",
    "description": "string",
    "auditData": {
      "createdOn": "2019-08-24T14:15:22Z",
      "updatedOn": "2019-08-24T14:15:22Z",
      "createdBy": "string",
      "updatedBy": "string"
    },
    "version": 0,
    "annotations": {
      "property1": "string",
      "property2": "string"
    }
  }
}
```

<h3 id="resource_get-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[GetResourceResponse](#schemagetresourceresponse)|
|default|Default|Default error response|[Status](#schemastatus)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

## Resource_GetByName

<a id="opIdResource_GetByName"></a>

> Code samples

```shell
# You can also use wget
curl -X GET /system/resources/{namespace}/{name} \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
GET /system/resources/{namespace}/{name} HTTP/1.1

Accept: application/json

```

```javascript

const headers = {
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('/system/resources/{namespace}/{name}',
{
  method: 'GET',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.get '/system/resources/{namespace}/{name}',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.get('/system/resources/{namespace}/{name}', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('GET','/system/resources/{namespace}/{name}', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/system/resources/{namespace}/{name}");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("GET");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "/system/resources/{namespace}/{name}", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /system/resources/{namespace}/{name}`

*GetByName*

<h3 id="resource_getbyname-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|namespace|path|string|true|none|
|name|path|string|true|none|
|token|query|string|false|none|

> Example responses

> 200 Response

```json
{
  "resource": {
    "id": "string",
    "name": "string",
    "namespace": "string",
    "sourceConfig": {
      "dataSource": "string",
      "catalog": "string",
      "entity": "string"
    },
    "properties": [
      {
        "id": "string",
        "name": "string",
        "type": "BOOL",
        "mapping": "string",
        "required": true,
        "primary": true,
        "length": 0,
        "unique": true,
        "immutable": true,
        "securityContext": {
          "constraints": [
            {
              "namespace": "string",
              "resource": "string",
              "property": "string",
              "before": "2019-08-24T14:15:22Z",
              "after": "2019-08-24T14:15:22Z",
              "principal": "string",
              "recordIds": [
                "string"
              ],
              "operation": "OPERATION_TYPE_READ",
              "permit": "PERMIT_TYPE_ALLOW"
            }
          ]
        },
        "defaultValue": null,
        "exampleValue": null,
        "enumValues": [
          null
        ],
        "reference": {
          "referencedResource": "string",
          "cascade": true
        },
        "properties": [
          {}
        ],
        "Item": {},
        "title": "string",
        "description": "string",
        "annotations": {
          "property1": "string",
          "property2": "string"
        }
      }
    ],
    "indexes": [
      {
        "properties": [
          {
            "name": "string",
            "order": "ORDER_UNKNOWN"
          }
        ],
        "indexType": "BTREE",
        "unique": true,
        "annotations": {
          "property1": "string",
          "property2": "string"
        }
      }
    ],
    "securityContext": {
      "constraints": [
        {
          "namespace": "string",
          "resource": "string",
          "property": "string",
          "before": "2019-08-24T14:15:22Z",
          "after": "2019-08-24T14:15:22Z",
          "principal": "string",
          "recordIds": [
            "string"
          ],
          "operation": "OPERATION_TYPE_READ",
          "permit": "PERMIT_TYPE_ALLOW"
        }
      ]
    },
    "virtual": true,
    "immutable": true,
    "abstract": true,
    "title": "string",
    "description": "string",
    "auditData": {
      "createdOn": "2019-08-24T14:15:22Z",
      "updatedOn": "2019-08-24T14:15:22Z",
      "createdBy": "string",
      "updatedBy": "string"
    },
    "version": 0,
    "annotations": {
      "property1": "string",
      "property2": "string"
    }
  }
}
```

<h3 id="resource_getbyname-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[GetResourceByNameResponse](#schemagetresourcebynameresponse)|
|default|Default|Default error response|[Status](#schemastatus)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

<h1 id="apibrew-user">User</h1>

User service is for managing users

## User_List

<a id="opIdUser_List"></a>

> Code samples

```shell
# You can also use wget
curl -X GET /system/users \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
GET /system/users HTTP/1.1

Accept: application/json

```

```javascript

const headers = {
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('/system/users',
{
  method: 'GET',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.get '/system/users',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.get('/system/users', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('GET','/system/users', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/system/users");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("GET");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "/system/users", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /system/users`

*List*

<h3 id="user_list-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|token|query|string|false|none|
|limit|query|integer(uint32)|false|none|
|offset|query|integer(uint64)|false|none|

> Example responses

> 200 Response

```json
{
  "content": [
    {
      "id": "string",
      "username": "string",
      "password": "string",
      "securityContext": {
        "constraints": [
          {
            "namespace": "string",
            "resource": "string",
            "property": "string",
            "before": "2019-08-24T14:15:22Z",
            "after": "2019-08-24T14:15:22Z",
            "principal": "string",
            "recordIds": [
              "string"
            ],
            "operation": "OPERATION_TYPE_READ",
            "permit": "PERMIT_TYPE_ALLOW"
          }
        ]
      },
      "details": {},
      "signKey": "string",
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}
```

<h3 id="user_list-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[ListUserResponse](#schemalistuserresponse)|
|default|Default|Default error response|[Status](#schemastatus)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

## User_Update

<a id="opIdUser_Update"></a>

> Code samples

```shell
# You can also use wget
curl -X PUT /system/users/_bulk \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
PUT /system/users/_bulk HTTP/1.1

Accept: application/json

```

```javascript

const headers = {
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('/system/users/_bulk',
{
  method: 'PUT',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.put '/system/users/_bulk',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.put('/system/users/_bulk', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('PUT','/system/users/_bulk', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/system/users/_bulk");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("PUT");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("PUT", "/system/users/_bulk", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`PUT /system/users/_bulk`

*Update*

<h3 id="user_update-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|token|query|string|false|none|
|user.id|query|string|false|none|
|user.username|query|string|false|none|
|user.password|query|string|false|none|
|user.signKey|query|string|false|none|
|user.auditData.createdOn|query|string(date-time)|false|none|
|user.auditData.updatedOn|query|string(date-time)|false|none|
|user.auditData.createdBy|query|string|false|none|
|user.auditData.updatedBy|query|string|false|none|
|user.version|query|integer(uint32)|false|none|

> Example responses

> 200 Response

```json
{
  "user": {
    "id": "string",
    "username": "string",
    "password": "string",
    "securityContext": {
      "constraints": [
        {
          "namespace": "string",
          "resource": "string",
          "property": "string",
          "before": "2019-08-24T14:15:22Z",
          "after": "2019-08-24T14:15:22Z",
          "principal": "string",
          "recordIds": [
            "string"
          ],
          "operation": "OPERATION_TYPE_READ",
          "permit": "PERMIT_TYPE_ALLOW"
        }
      ]
    },
    "details": {},
    "signKey": "string",
    "auditData": {
      "createdOn": "2019-08-24T14:15:22Z",
      "updatedOn": "2019-08-24T14:15:22Z",
      "createdBy": "string",
      "updatedBy": "string"
    },
    "version": 0
  },
  "users": [
    {
      "id": "string",
      "username": "string",
      "password": "string",
      "securityContext": {
        "constraints": [
          {
            "namespace": "string",
            "resource": "string",
            "property": "string",
            "before": "2019-08-24T14:15:22Z",
            "after": "2019-08-24T14:15:22Z",
            "principal": "string",
            "recordIds": [
              "string"
            ],
            "operation": "OPERATION_TYPE_READ",
            "permit": "PERMIT_TYPE_ALLOW"
          }
        ]
      },
      "details": {},
      "signKey": "string",
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}
```

<h3 id="user_update-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[UpdateUserResponse](#schemaupdateuserresponse)|
|default|Default|Default error response|[Status](#schemastatus)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

## User_Create

<a id="opIdUser_Create"></a>

> Code samples

```shell
# You can also use wget
curl -X POST /system/users/_bulk \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
POST /system/users/_bulk HTTP/1.1

Accept: application/json

```

```javascript

const headers = {
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('/system/users/_bulk',
{
  method: 'POST',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.post '/system/users/_bulk',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.post('/system/users/_bulk', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('POST','/system/users/_bulk', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/system/users/_bulk");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("POST");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("POST", "/system/users/_bulk", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /system/users/_bulk`

*Create*

<h3 id="user_create-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|token|query|string|false|none|
|user.id|query|string|false|none|
|user.username|query|string|false|none|
|user.password|query|string|false|none|
|user.signKey|query|string|false|none|
|user.auditData.createdOn|query|string(date-time)|false|none|
|user.auditData.updatedOn|query|string(date-time)|false|none|
|user.auditData.createdBy|query|string|false|none|
|user.auditData.updatedBy|query|string|false|none|
|user.version|query|integer(uint32)|false|none|

> Example responses

> 200 Response

```json
{
  "user": {
    "id": "string",
    "username": "string",
    "password": "string",
    "securityContext": {
      "constraints": [
        {
          "namespace": "string",
          "resource": "string",
          "property": "string",
          "before": "2019-08-24T14:15:22Z",
          "after": "2019-08-24T14:15:22Z",
          "principal": "string",
          "recordIds": [
            "string"
          ],
          "operation": "OPERATION_TYPE_READ",
          "permit": "PERMIT_TYPE_ALLOW"
        }
      ]
    },
    "details": {},
    "signKey": "string",
    "auditData": {
      "createdOn": "2019-08-24T14:15:22Z",
      "updatedOn": "2019-08-24T14:15:22Z",
      "createdBy": "string",
      "updatedBy": "string"
    },
    "version": 0
  },
  "users": [
    {
      "id": "string",
      "username": "string",
      "password": "string",
      "securityContext": {
        "constraints": [
          {
            "namespace": "string",
            "resource": "string",
            "property": "string",
            "before": "2019-08-24T14:15:22Z",
            "after": "2019-08-24T14:15:22Z",
            "principal": "string",
            "recordIds": [
              "string"
            ],
            "operation": "OPERATION_TYPE_READ",
            "permit": "PERMIT_TYPE_ALLOW"
          }
        ]
      },
      "details": {},
      "signKey": "string",
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}
```

<h3 id="user_create-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[CreateUserResponse](#schemacreateuserresponse)|
|default|Default|Default error response|[Status](#schemastatus)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

## User_Delete

<a id="opIdUser_Delete"></a>

> Code samples

```shell
# You can also use wget
curl -X DELETE /system/users/_bulk \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
DELETE /system/users/_bulk HTTP/1.1

Accept: application/json

```

```javascript

const headers = {
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('/system/users/_bulk',
{
  method: 'DELETE',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.delete '/system/users/_bulk',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.delete('/system/users/_bulk', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('DELETE','/system/users/_bulk', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/system/users/_bulk");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("DELETE");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("DELETE", "/system/users/_bulk", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`DELETE /system/users/_bulk`

*Delete*

<h3 id="user_delete-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|token|query|string|false|none|
|id|query|string|false|none|
|ids|query|array[string]|false|none|

> Example responses

> 200 Response

```json
{}
```

<h3 id="user_delete-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[DeleteUserResponse](#schemadeleteuserresponse)|
|default|Default|Default error response|[Status](#schemastatus)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

## User_Get

<a id="opIdUser_Get"></a>

> Code samples

```shell
# You can also use wget
curl -X GET /system/users/{id} \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
GET /system/users/{id} HTTP/1.1

Accept: application/json

```

```javascript

const headers = {
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('/system/users/{id}',
{
  method: 'GET',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.get '/system/users/{id}',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.get('/system/users/{id}', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('GET','/system/users/{id}', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/system/users/{id}");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("GET");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "/system/users/{id}", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /system/users/{id}`

*Get*

<h3 id="user_get-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|id|path|string|true|none|
|token|query|string|false|none|

> Example responses

> 200 Response

```json
{
  "user": {
    "id": "string",
    "username": "string",
    "password": "string",
    "securityContext": {
      "constraints": [
        {
          "namespace": "string",
          "resource": "string",
          "property": "string",
          "before": "2019-08-24T14:15:22Z",
          "after": "2019-08-24T14:15:22Z",
          "principal": "string",
          "recordIds": [
            "string"
          ],
          "operation": "OPERATION_TYPE_READ",
          "permit": "PERMIT_TYPE_ALLOW"
        }
      ]
    },
    "details": {},
    "signKey": "string",
    "auditData": {
      "createdOn": "2019-08-24T14:15:22Z",
      "updatedOn": "2019-08-24T14:15:22Z",
      "createdBy": "string",
      "updatedBy": "string"
    },
    "version": 0
  }
}
```

<h3 id="user_get-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[GetUserResponse](#schemagetuserresponse)|
|default|Default|Default error response|[Status](#schemastatus)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

<h1 id="apibrew-watch">Watch</h1>

Watch service watching operations on records

## Watch_Watch

<a id="opIdWatch_Watch"></a>

> Code samples

```shell
# You can also use wget
curl -X POST /system/watch \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
POST /system/watch HTTP/1.1

Accept: application/json

```

```javascript

const headers = {
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('/system/watch',
{
  method: 'POST',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.post '/system/watch',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.post('/system/watch', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('POST','/system/watch', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("/system/watch");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("POST");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("POST", "/system/watch", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /system/watch`

*Watch*

Sends a greeting

<h3 id="watch_watch-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|token|query|string|false|none|
|namespace|query|string|false|none|
|resource|query|string|false|none|
|query.not.not.equal.left.property|query|string|false|none|
|query.not.not.equal.left.value|query|[GoogleProtobufValue](#schemagoogleprotobufvalue)|false|none|
|query.not.not.equal.left.refValue.namespace|query|string|false|none|
|query.not.not.equal.left.refValue.resource|query|string|false|none|
|query.not.not.equal.right.property|query|string|false|none|
|query.not.not.equal.right.value|query|[GoogleProtobufValue](#schemagoogleprotobufvalue)|false|none|
|query.not.not.equal.right.refValue.namespace|query|string|false|none|
|query.not.not.equal.right.refValue.resource|query|string|false|none|
|query.not.not.regexMatch.pattern|query|string|false|none|
|query.not.regexMatch.pattern|query|string|false|none|
|events|query|array[string]|false|none|

#### Enumerated Values

|Parameter|Value|
|---|---|
|events|CREATE|
|events|UPDATE|
|events|DELETE|
|events|GET|
|events|LIST|

> Example responses

> 200 Response

```json
{
  "changes": {},
  "recordIds": [
    "string"
  ],
  "event": "CREATE",
  "eventOn": "2019-08-24T14:15:22Z"
}
```

<h3 id="watch_watch-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[WatchMessage](#schemawatchmessage)|
|default|Default|Default error response|[Status](#schemastatus)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearerAuth
</aside>

# Schemas

<h2 id="tocS_Any">Any</h2>
<!-- backwards compatibility -->
<a id="schemaany"></a>
<a id="schema_Any"></a>
<a id="tocSany"></a>
<a id="tocsany"></a>

```json
{
  "value": {
    "@type": "string"
  },
  "yaml": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|value|[GoogleProtobufAny](#schemagoogleprotobufany)|false|none|Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.|
|yaml|string|false|none|none|

<h2 id="tocS_ApplyRecordRequest_PropertiesEntry">ApplyRecordRequest_PropertiesEntry</h2>
<!-- backwards compatibility -->
<a id="schemaapplyrecordrequest_propertiesentry"></a>
<a id="schema_ApplyRecordRequest_PropertiesEntry"></a>
<a id="tocSapplyrecordrequest_propertiesentry"></a>
<a id="tocsapplyrecordrequest_propertiesentry"></a>

```json
{
  "key": "string",
  "value": null
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|key|string|false|none|none|
|value|[GoogleProtobufValue](#schemagoogleprotobufvalue)|false|none|Represents a dynamically typed value which can be either null, a number, a string, a boolean, a recursive struct value, or a list of values.|

<h2 id="tocS_ApplyRecordResponse">ApplyRecordResponse</h2>
<!-- backwards compatibility -->
<a id="schemaapplyrecordresponse"></a>
<a id="schema_ApplyRecordResponse"></a>
<a id="tocSapplyrecordresponse"></a>
<a id="tocsapplyrecordresponse"></a>

```json
{
  "record": {
    "id": "string",
    "properties": {
      "property1": null,
      "property2": null
    },
    "propertiesPacked": [
      null
    ]
  },
  "records": [
    {
      "id": "string",
      "properties": {
        "property1": null,
        "property2": null
      },
      "propertiesPacked": [
        null
      ]
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|record|[Record](#schemarecord)|false|none|none|
|records|[[Record](#schemarecord)]|false|none|none|

<h2 id="tocS_AuditData">AuditData</h2>
<!-- backwards compatibility -->
<a id="schemaauditdata"></a>
<a id="schema_AuditData"></a>
<a id="tocSauditdata"></a>
<a id="tocsauditdata"></a>

```json
{
  "createdOn": "2019-08-24T14:15:22Z",
  "updatedOn": "2019-08-24T14:15:22Z",
  "createdBy": "string",
  "updatedBy": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|createdOn|string(date-time)|false|none|none|
|updatedOn|string(date-time)|false|none|none|
|createdBy|string|false|none|none|
|updatedBy|string|false|none|none|

<h2 id="tocS_AuthenticationRequest">AuthenticationRequest</h2>
<!-- backwards compatibility -->
<a id="schemaauthenticationrequest"></a>
<a id="schema_AuthenticationRequest"></a>
<a id="tocSauthenticationrequest"></a>
<a id="tocsauthenticationrequest"></a>

```json
{
  "username": "string",
  "password": "string",
  "term": "SHORT"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|username|string|false|none|none|
|password|string|false|none|none|
|term|string(enum)|false|none|none|

#### Enumerated Values

|Property|Value|
|---|---|
|term|SHORT|
|term|MIDDLE|
|term|LONG|
|term|VERY_LONG|

<h2 id="tocS_AuthenticationResponse">AuthenticationResponse</h2>
<!-- backwards compatibility -->
<a id="schemaauthenticationresponse"></a>
<a id="schema_AuthenticationResponse"></a>
<a id="tocSauthenticationresponse"></a>
<a id="tocsauthenticationresponse"></a>

```json
{
  "token": {
    "term": "SHORT",
    "content": "string",
    "expiration": "2019-08-24T14:15:22Z"
  }
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|token|[Token](#schematoken)|false|none|none|

<h2 id="tocS_BooleanExpression">BooleanExpression</h2>
<!-- backwards compatibility -->
<a id="schemabooleanexpression"></a>
<a id="schema_BooleanExpression"></a>
<a id="tocSbooleanexpression"></a>
<a id="tocsbooleanexpression"></a>

```json
{
  "and": {
    "expressions": [
      {
        "and": {},
        "or": {},
        "not": {},
        "equal": {
          "left": {
            "additionalProperties": [
              {
                "name": "string",
                "value": {
                  "value": {
                    "@type": "string"
                  },
                  "yaml": "string"
                }
              }
            ]
          },
          "right": {
            "additionalProperties": [
              {
                "name": "string",
                "value": {
                  "value": {
                    "@type": "string"
                  },
                  "yaml": "string"
                }
              }
            ]
          }
        },
        "lessThan": {
          "left": {
            "additionalProperties": [
              {
                "name": "string",
                "value": {
                  "value": {
                    "@type": "string"
                  },
                  "yaml": "string"
                }
              }
            ]
          },
          "right": {
            "additionalProperties": [
              {
                "name": "string",
                "value": {
                  "value": {
                    "@type": "string"
                  },
                  "yaml": "string"
                }
              }
            ]
          }
        },
        "greaterThan": {
          "left": {
            "additionalProperties": [
              {
                "name": "string",
                "value": {
                  "value": {
                    "@type": "string"
                  },
                  "yaml": "string"
                }
              }
            ]
          },
          "right": {
            "additionalProperties": [
              {
                "name": "string",
                "value": {
                  "value": {
                    "@type": "string"
                  },
                  "yaml": "string"
                }
              }
            ]
          }
        },
        "lessThanOrEqual": {
          "left": {
            "additionalProperties": [
              {
                "name": "string",
                "value": {
                  "value": {
                    "@type": "string"
                  },
                  "yaml": "string"
                }
              }
            ]
          },
          "right": {
            "additionalProperties": [
              {
                "name": "string",
                "value": {
                  "value": {
                    "@type": "string"
                  },
                  "yaml": "string"
                }
              }
            ]
          }
        },
        "greaterThanOrEqual": {
          "left": {
            "additionalProperties": [
              {
                "name": "string",
                "value": {
                  "value": {
                    "@type": "string"
                  },
                  "yaml": "string"
                }
              }
            ]
          },
          "right": {
            "additionalProperties": [
              {
                "name": "string",
                "value": {
                  "value": {
                    "@type": "string"
                  },
                  "yaml": "string"
                }
              }
            ]
          }
        },
        "in": {
          "left": {
            "additionalProperties": [
              {
                "name": "string",
                "value": {
                  "value": {
                    "@type": "string"
                  },
                  "yaml": "string"
                }
              }
            ]
          },
          "right": {
            "additionalProperties": [
              {
                "name": "string",
                "value": {
                  "value": {
                    "@type": "string"
                  },
                  "yaml": "string"
                }
              }
            ]
          }
        },
        "isNull": {
          "additionalProperties": [
            {
              "name": "string",
              "value": {
                "value": {
                  "@type": "string"
                },
                "yaml": "string"
              }
            }
          ]
        },
        "regexMatch": {
          "pattern": "string",
          "expression": {
            "additionalProperties": [
              {
                "name": "string",
                "value": {
                  "value": {
                    "@type": "string"
                  },
                  "yaml": "string"
                }
              }
            ]
          }
        }
      }
    ]
  },
  "or": {
    "expressions": [
      {
        "and": {},
        "or": {},
        "not": {},
        "equal": {
          "left": {
            "additionalProperties": [
              {
                "name": "string",
                "value": {
                  "value": {
                    "@type": "string"
                  },
                  "yaml": "string"
                }
              }
            ]
          },
          "right": {
            "additionalProperties": [
              {
                "name": "string",
                "value": {
                  "value": {
                    "@type": "string"
                  },
                  "yaml": "string"
                }
              }
            ]
          }
        },
        "lessThan": {
          "left": {
            "additionalProperties": [
              {
                "name": "string",
                "value": {
                  "value": {
                    "@type": "string"
                  },
                  "yaml": "string"
                }
              }
            ]
          },
          "right": {
            "additionalProperties": [
              {
                "name": "string",
                "value": {
                  "value": {
                    "@type": "string"
                  },
                  "yaml": "string"
                }
              }
            ]
          }
        },
        "greaterThan": {
          "left": {
            "additionalProperties": [
              {
                "name": "string",
                "value": {
                  "value": {
                    "@type": "string"
                  },
                  "yaml": "string"
                }
              }
            ]
          },
          "right": {
            "additionalProperties": [
              {
                "name": "string",
                "value": {
                  "value": {
                    "@type": "string"
                  },
                  "yaml": "string"
                }
              }
            ]
          }
        },
        "lessThanOrEqual": {
          "left": {
            "additionalProperties": [
              {
                "name": "string",
                "value": {
                  "value": {
                    "@type": "string"
                  },
                  "yaml": "string"
                }
              }
            ]
          },
          "right": {
            "additionalProperties": [
              {
                "name": "string",
                "value": {
                  "value": {
                    "@type": "string"
                  },
                  "yaml": "string"
                }
              }
            ]
          }
        },
        "greaterThanOrEqual": {
          "left": {
            "additionalProperties": [
              {
                "name": "string",
                "value": {
                  "value": {
                    "@type": "string"
                  },
                  "yaml": "string"
                }
              }
            ]
          },
          "right": {
            "additionalProperties": [
              {
                "name": "string",
                "value": {
                  "value": {
                    "@type": "string"
                  },
                  "yaml": "string"
                }
              }
            ]
          }
        },
        "in": {
          "left": {
            "additionalProperties": [
              {
                "name": "string",
                "value": {
                  "value": {
                    "@type": "string"
                  },
                  "yaml": "string"
                }
              }
            ]
          },
          "right": {
            "additionalProperties": [
              {
                "name": "string",
                "value": {
                  "value": {
                    "@type": "string"
                  },
                  "yaml": "string"
                }
              }
            ]
          }
        },
        "isNull": {
          "additionalProperties": [
            {
              "name": "string",
              "value": {
                "value": {
                  "@type": "string"
                },
                "yaml": "string"
              }
            }
          ]
        },
        "regexMatch": {
          "pattern": "string",
          "expression": {
            "additionalProperties": [
              {
                "name": "string",
                "value": {
                  "value": {
                    "@type": "string"
                  },
                  "yaml": "string"
                }
              }
            ]
          }
        }
      }
    ]
  },
  "not": {
    "and": {
      "expressions": [
        {}
      ]
    },
    "or": {
      "expressions": [
        {}
      ]
    },
    "not": {},
    "equal": {
      "left": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "right": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    },
    "lessThan": {
      "left": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "right": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    },
    "greaterThan": {
      "left": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "right": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    },
    "lessThanOrEqual": {
      "left": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "right": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    },
    "greaterThanOrEqual": {
      "left": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "right": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    },
    "in": {
      "left": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "right": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    },
    "isNull": {
      "additionalProperties": [
        {
          "name": "string",
          "value": {
            "value": {
              "@type": "string"
            },
            "yaml": "string"
          }
        }
      ]
    },
    "regexMatch": {
      "pattern": "string",
      "expression": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    }
  },
  "equal": {
    "left": {
      "additionalProperties": [
        {
          "name": "string",
          "value": {
            "value": {
              "@type": "string"
            },
            "yaml": "string"
          }
        }
      ]
    },
    "right": {
      "additionalProperties": [
        {
          "name": "string",
          "value": {
            "value": {
              "@type": "string"
            },
            "yaml": "string"
          }
        }
      ]
    }
  },
  "lessThan": {
    "left": {
      "additionalProperties": [
        {
          "name": "string",
          "value": {
            "value": {
              "@type": "string"
            },
            "yaml": "string"
          }
        }
      ]
    },
    "right": {
      "additionalProperties": [
        {
          "name": "string",
          "value": {
            "value": {
              "@type": "string"
            },
            "yaml": "string"
          }
        }
      ]
    }
  },
  "greaterThan": {
    "left": {
      "additionalProperties": [
        {
          "name": "string",
          "value": {
            "value": {
              "@type": "string"
            },
            "yaml": "string"
          }
        }
      ]
    },
    "right": {
      "additionalProperties": [
        {
          "name": "string",
          "value": {
            "value": {
              "@type": "string"
            },
            "yaml": "string"
          }
        }
      ]
    }
  },
  "lessThanOrEqual": {
    "left": {
      "additionalProperties": [
        {
          "name": "string",
          "value": {
            "value": {
              "@type": "string"
            },
            "yaml": "string"
          }
        }
      ]
    },
    "right": {
      "additionalProperties": [
        {
          "name": "string",
          "value": {
            "value": {
              "@type": "string"
            },
            "yaml": "string"
          }
        }
      ]
    }
  },
  "greaterThanOrEqual": {
    "left": {
      "additionalProperties": [
        {
          "name": "string",
          "value": {
            "value": {
              "@type": "string"
            },
            "yaml": "string"
          }
        }
      ]
    },
    "right": {
      "additionalProperties": [
        {
          "name": "string",
          "value": {
            "value": {
              "@type": "string"
            },
            "yaml": "string"
          }
        }
      ]
    }
  },
  "in": {
    "left": {
      "additionalProperties": [
        {
          "name": "string",
          "value": {
            "value": {
              "@type": "string"
            },
            "yaml": "string"
          }
        }
      ]
    },
    "right": {
      "additionalProperties": [
        {
          "name": "string",
          "value": {
            "value": {
              "@type": "string"
            },
            "yaml": "string"
          }
        }
      ]
    }
  },
  "isNull": {
    "additionalProperties": [
      {
        "name": "string",
        "value": {
          "value": {
            "@type": "string"
          },
          "yaml": "string"
        }
      }
    ]
  },
  "regexMatch": {
    "pattern": "string",
    "expression": {
      "additionalProperties": [
        {
          "name": "string",
          "value": {
            "value": {
              "@type": "string"
            },
            "yaml": "string"
          }
        }
      ]
    }
  }
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|and|[CompoundBooleanExpression](#schemacompoundbooleanexpression)|false|none|none|
|or|[CompoundBooleanExpression](#schemacompoundbooleanexpression)|false|none|none|
|not|[BooleanExpression](#schemabooleanexpression)|false|none|none|
|equal|[PairExpression](#schemapairexpression)|false|none|none|
|lessThan|[PairExpression](#schemapairexpression)|false|none|none|
|greaterThan|[PairExpression](#schemapairexpression)|false|none|none|
|lessThanOrEqual|[PairExpression](#schemapairexpression)|false|none|none|
|greaterThanOrEqual|[PairExpression](#schemapairexpression)|false|none|none|
|in|[PairExpression](#schemapairexpression)|false|none|none|
|isNull|[Expression](#schemaexpression)|false|none|none|
|regexMatch|[RegexMatchExpression](#schemaregexmatchexpression)|false|none|none|

<h2 id="tocS_CompoundBooleanExpression">CompoundBooleanExpression</h2>
<!-- backwards compatibility -->
<a id="schemacompoundbooleanexpression"></a>
<a id="schema_CompoundBooleanExpression"></a>
<a id="tocScompoundbooleanexpression"></a>
<a id="tocscompoundbooleanexpression"></a>

```json
{
  "expressions": [
    {
      "and": {
        "expressions": []
      },
      "or": {
        "expressions": []
      },
      "not": {},
      "equal": {
        "left": {
          "additionalProperties": [
            {
              "name": "string",
              "value": {
                "value": {
                  "@type": "string"
                },
                "yaml": "string"
              }
            }
          ]
        },
        "right": {
          "additionalProperties": [
            {
              "name": "string",
              "value": {
                "value": {
                  "@type": "string"
                },
                "yaml": "string"
              }
            }
          ]
        }
      },
      "lessThan": {
        "left": {
          "additionalProperties": [
            {
              "name": "string",
              "value": {
                "value": {
                  "@type": "string"
                },
                "yaml": "string"
              }
            }
          ]
        },
        "right": {
          "additionalProperties": [
            {
              "name": "string",
              "value": {
                "value": {
                  "@type": "string"
                },
                "yaml": "string"
              }
            }
          ]
        }
      },
      "greaterThan": {
        "left": {
          "additionalProperties": [
            {
              "name": "string",
              "value": {
                "value": {
                  "@type": "string"
                },
                "yaml": "string"
              }
            }
          ]
        },
        "right": {
          "additionalProperties": [
            {
              "name": "string",
              "value": {
                "value": {
                  "@type": "string"
                },
                "yaml": "string"
              }
            }
          ]
        }
      },
      "lessThanOrEqual": {
        "left": {
          "additionalProperties": [
            {
              "name": "string",
              "value": {
                "value": {
                  "@type": "string"
                },
                "yaml": "string"
              }
            }
          ]
        },
        "right": {
          "additionalProperties": [
            {
              "name": "string",
              "value": {
                "value": {
                  "@type": "string"
                },
                "yaml": "string"
              }
            }
          ]
        }
      },
      "greaterThanOrEqual": {
        "left": {
          "additionalProperties": [
            {
              "name": "string",
              "value": {
                "value": {
                  "@type": "string"
                },
                "yaml": "string"
              }
            }
          ]
        },
        "right": {
          "additionalProperties": [
            {
              "name": "string",
              "value": {
                "value": {
                  "@type": "string"
                },
                "yaml": "string"
              }
            }
          ]
        }
      },
      "in": {
        "left": {
          "additionalProperties": [
            {
              "name": "string",
              "value": {
                "value": {
                  "@type": "string"
                },
                "yaml": "string"
              }
            }
          ]
        },
        "right": {
          "additionalProperties": [
            {
              "name": "string",
              "value": {
                "value": {
                  "@type": "string"
                },
                "yaml": "string"
              }
            }
          ]
        }
      },
      "isNull": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "regexMatch": {
        "pattern": "string",
        "expression": {
          "additionalProperties": [
            {
              "name": "string",
              "value": {
                "value": {
                  "@type": "string"
                },
                "yaml": "string"
              }
            }
          ]
        }
      }
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|expressions|[[BooleanExpression](#schemabooleanexpression)]|false|none|none|

<h2 id="tocS_CreateDataSourceRequest">CreateDataSourceRequest</h2>
<!-- backwards compatibility -->
<a id="schemacreatedatasourcerequest"></a>
<a id="schema_CreateDataSourceRequest"></a>
<a id="tocScreatedatasourcerequest"></a>
<a id="tocscreatedatasourcerequest"></a>

```json
{
  "token": "string",
  "dataSources": [
    {
      "id": "string",
      "backend": "POSTGRESQL",
      "name": "string",
      "description": "string",
      "postgresqlParams": {
        "username": "string",
        "password": "string",
        "host": "string",
        "port": 0,
        "dbName": "string",
        "defaultSchema": "string"
      },
      "mysqlParams": {
        "username": "string",
        "password": "string",
        "host": "string",
        "port": 0,
        "dbName": "string",
        "defaultSchema": "string"
      },
      "virtualParams": {
        "mode": "DISCARD"
      },
      "redisParams": {
        "addr": "string",
        "password": "string",
        "db": 0
      },
      "mongoParams": {
        "uri": "string",
        "dbName": "string"
      },
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|token|string|false|none|none|
|dataSources|[[DataSource](#schemadatasource)]|false|none|none|

<h2 id="tocS_CreateDataSourceResponse">CreateDataSourceResponse</h2>
<!-- backwards compatibility -->
<a id="schemacreatedatasourceresponse"></a>
<a id="schema_CreateDataSourceResponse"></a>
<a id="tocScreatedatasourceresponse"></a>
<a id="tocscreatedatasourceresponse"></a>

```json
{
  "dataSources": [
    {
      "id": "string",
      "backend": "POSTGRESQL",
      "name": "string",
      "description": "string",
      "postgresqlParams": {
        "username": "string",
        "password": "string",
        "host": "string",
        "port": 0,
        "dbName": "string",
        "defaultSchema": "string"
      },
      "mysqlParams": {
        "username": "string",
        "password": "string",
        "host": "string",
        "port": 0,
        "dbName": "string",
        "defaultSchema": "string"
      },
      "virtualParams": {
        "mode": "DISCARD"
      },
      "redisParams": {
        "addr": "string",
        "password": "string",
        "db": 0
      },
      "mongoParams": {
        "uri": "string",
        "dbName": "string"
      },
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|dataSources|[[DataSource](#schemadatasource)]|false|none|none|

<h2 id="tocS_CreateExtensionRequest">CreateExtensionRequest</h2>
<!-- backwards compatibility -->
<a id="schemacreateextensionrequest"></a>
<a id="schema_CreateExtensionRequest"></a>
<a id="tocScreateextensionrequest"></a>
<a id="tocscreateextensionrequest"></a>

```json
{
  "token": "string",
  "extensions": [
    {
      "id": "string",
      "name": "string",
      "description": "string",
      "namespace": "string",
      "resource": "string",
      "before": {
        "all": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "create": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "update": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "delete": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "get": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "list": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "sync": true
      },
      "instead": {
        "all": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "create": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "update": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "delete": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "get": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "list": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "finalize": true
      },
      "after": {
        "all": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "create": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "update": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "delete": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "get": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "list": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "sync": true
      },
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|token|string|false|none|none|
|extensions|[[Extension](#schemaextension)]|false|none|none|

<h2 id="tocS_CreateExtensionResponse">CreateExtensionResponse</h2>
<!-- backwards compatibility -->
<a id="schemacreateextensionresponse"></a>
<a id="schema_CreateExtensionResponse"></a>
<a id="tocScreateextensionresponse"></a>
<a id="tocscreateextensionresponse"></a>

```json
{
  "extensions": [
    {
      "id": "string",
      "name": "string",
      "description": "string",
      "namespace": "string",
      "resource": "string",
      "before": {
        "all": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "create": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "update": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "delete": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "get": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "list": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "sync": true
      },
      "instead": {
        "all": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "create": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "update": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "delete": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "get": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "list": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "finalize": true
      },
      "after": {
        "all": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "create": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "update": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "delete": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "get": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "list": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "sync": true
      },
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|extensions|[[Extension](#schemaextension)]|false|none|none|

<h2 id="tocS_CreateNamespaceRequest">CreateNamespaceRequest</h2>
<!-- backwards compatibility -->
<a id="schemacreatenamespacerequest"></a>
<a id="schema_CreateNamespaceRequest"></a>
<a id="tocScreatenamespacerequest"></a>
<a id="tocscreatenamespacerequest"></a>

```json
{
  "token": "string",
  "Namespaces": [
    {
      "id": "string",
      "name": "string",
      "description": "string",
      "details": {},
      "securityContext": {
        "constraints": [
          {
            "namespace": "string",
            "resource": "string",
            "property": "string",
            "before": "2019-08-24T14:15:22Z",
            "after": "2019-08-24T14:15:22Z",
            "principal": "string",
            "recordIds": [
              "string"
            ],
            "operation": "OPERATION_TYPE_READ",
            "permit": "PERMIT_TYPE_ALLOW"
          }
        ]
      },
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|token|string|false|none|none|
|Namespaces|[[Namespace](#schemanamespace)]|false|none|none|

<h2 id="tocS_CreateNamespaceResponse">CreateNamespaceResponse</h2>
<!-- backwards compatibility -->
<a id="schemacreatenamespaceresponse"></a>
<a id="schema_CreateNamespaceResponse"></a>
<a id="tocScreatenamespaceresponse"></a>
<a id="tocscreatenamespaceresponse"></a>

```json
{
  "Namespaces": [
    {
      "id": "string",
      "name": "string",
      "description": "string",
      "details": {},
      "securityContext": {
        "constraints": [
          {
            "namespace": "string",
            "resource": "string",
            "property": "string",
            "before": "2019-08-24T14:15:22Z",
            "after": "2019-08-24T14:15:22Z",
            "principal": "string",
            "recordIds": [
              "string"
            ],
            "operation": "OPERATION_TYPE_READ",
            "permit": "PERMIT_TYPE_ALLOW"
          }
        ]
      },
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|Namespaces|[[Namespace](#schemanamespace)]|false|none|none|

<h2 id="tocS_CreateRecordRequest_PropertiesEntry">CreateRecordRequest_PropertiesEntry</h2>
<!-- backwards compatibility -->
<a id="schemacreaterecordrequest_propertiesentry"></a>
<a id="schema_CreateRecordRequest_PropertiesEntry"></a>
<a id="tocScreaterecordrequest_propertiesentry"></a>
<a id="tocscreaterecordrequest_propertiesentry"></a>

```json
{
  "key": "string",
  "value": null
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|key|string|false|none|none|
|value|[GoogleProtobufValue](#schemagoogleprotobufvalue)|false|none|Represents a dynamically typed value which can be either null, a number, a string, a boolean, a recursive struct value, or a list of values.|

<h2 id="tocS_CreateRecordResponse">CreateRecordResponse</h2>
<!-- backwards compatibility -->
<a id="schemacreaterecordresponse"></a>
<a id="schema_CreateRecordResponse"></a>
<a id="tocScreaterecordresponse"></a>
<a id="tocscreaterecordresponse"></a>

```json
{
  "record": {
    "id": "string",
    "properties": {
      "property1": null,
      "property2": null
    },
    "propertiesPacked": [
      null
    ]
  },
  "records": [
    {
      "id": "string",
      "properties": {
        "property1": null,
        "property2": null
      },
      "propertiesPacked": [
        null
      ]
    }
  ],
  "inserted": [
    true
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|record|[Record](#schemarecord)|false|none|none|
|records|[[Record](#schemarecord)]|false|none|none|
|inserted|[boolean]|false|none|none|

<h2 id="tocS_CreateResourceRequest">CreateResourceRequest</h2>
<!-- backwards compatibility -->
<a id="schemacreateresourcerequest"></a>
<a id="schema_CreateResourceRequest"></a>
<a id="tocScreateresourcerequest"></a>
<a id="tocscreateresourcerequest"></a>

```json
{
  "token": "string",
  "resources": [
    {
      "id": "string",
      "name": "string",
      "namespace": "string",
      "sourceConfig": {
        "dataSource": "string",
        "catalog": "string",
        "entity": "string"
      },
      "properties": [
        {
          "id": "string",
          "name": "string",
          "type": "BOOL",
          "mapping": "string",
          "required": true,
          "primary": true,
          "length": 0,
          "unique": true,
          "immutable": true,
          "securityContext": {
            "constraints": [
              {
                "namespace": "string",
                "resource": "string",
                "property": "string",
                "before": "2019-08-24T14:15:22Z",
                "after": "2019-08-24T14:15:22Z",
                "principal": "string",
                "recordIds": [
                  "string"
                ],
                "operation": "OPERATION_TYPE_READ",
                "permit": "PERMIT_TYPE_ALLOW"
              }
            ]
          },
          "defaultValue": null,
          "exampleValue": null,
          "enumValues": [
            null
          ],
          "reference": {
            "referencedResource": "string",
            "cascade": true
          },
          "properties": [
            {}
          ],
          "Item": {},
          "title": "string",
          "description": "string",
          "annotations": {
            "property1": "string",
            "property2": "string"
          }
        }
      ],
      "indexes": [
        {
          "properties": [
            {
              "name": "string",
              "order": "ORDER_UNKNOWN"
            }
          ],
          "indexType": "BTREE",
          "unique": true,
          "annotations": {
            "property1": "string",
            "property2": "string"
          }
        }
      ],
      "securityContext": {
        "constraints": [
          {
            "namespace": "string",
            "resource": "string",
            "property": "string",
            "before": "2019-08-24T14:15:22Z",
            "after": "2019-08-24T14:15:22Z",
            "principal": "string",
            "recordIds": [
              "string"
            ],
            "operation": "OPERATION_TYPE_READ",
            "permit": "PERMIT_TYPE_ALLOW"
          }
        ]
      },
      "virtual": true,
      "immutable": true,
      "abstract": true,
      "title": "string",
      "description": "string",
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0,
      "annotations": {
        "property1": "string",
        "property2": "string"
      }
    }
  ],
  "doMigration": true,
  "forceMigration": true,
  "annotations": {
    "property1": "string",
    "property2": "string"
  }
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|token|string|false|none|none|
|resources|[[Resource](#schemaresource)]|false|none|none|
|doMigration|boolean|false|none|none|
|forceMigration|boolean|false|none|none|
|annotations|object|false|none|none|
|» **additionalProperties**|string|false|none|none|

<h2 id="tocS_CreateResourceResponse">CreateResourceResponse</h2>
<!-- backwards compatibility -->
<a id="schemacreateresourceresponse"></a>
<a id="schema_CreateResourceResponse"></a>
<a id="tocScreateresourceresponse"></a>
<a id="tocscreateresourceresponse"></a>

```json
{
  "resources": [
    {
      "id": "string",
      "name": "string",
      "namespace": "string",
      "sourceConfig": {
        "dataSource": "string",
        "catalog": "string",
        "entity": "string"
      },
      "properties": [
        {
          "id": "string",
          "name": "string",
          "type": "BOOL",
          "mapping": "string",
          "required": true,
          "primary": true,
          "length": 0,
          "unique": true,
          "immutable": true,
          "securityContext": {
            "constraints": [
              {
                "namespace": "string",
                "resource": "string",
                "property": "string",
                "before": "2019-08-24T14:15:22Z",
                "after": "2019-08-24T14:15:22Z",
                "principal": "string",
                "recordIds": [
                  "string"
                ],
                "operation": "OPERATION_TYPE_READ",
                "permit": "PERMIT_TYPE_ALLOW"
              }
            ]
          },
          "defaultValue": null,
          "exampleValue": null,
          "enumValues": [
            null
          ],
          "reference": {
            "referencedResource": "string",
            "cascade": true
          },
          "properties": [
            {}
          ],
          "Item": {},
          "title": "string",
          "description": "string",
          "annotations": {
            "property1": "string",
            "property2": "string"
          }
        }
      ],
      "indexes": [
        {
          "properties": [
            {
              "name": "string",
              "order": "ORDER_UNKNOWN"
            }
          ],
          "indexType": "BTREE",
          "unique": true,
          "annotations": {
            "property1": "string",
            "property2": "string"
          }
        }
      ],
      "securityContext": {
        "constraints": [
          {
            "namespace": "string",
            "resource": "string",
            "property": "string",
            "before": "2019-08-24T14:15:22Z",
            "after": "2019-08-24T14:15:22Z",
            "principal": "string",
            "recordIds": [
              "string"
            ],
            "operation": "OPERATION_TYPE_READ",
            "permit": "PERMIT_TYPE_ALLOW"
          }
        ]
      },
      "virtual": true,
      "immutable": true,
      "abstract": true,
      "title": "string",
      "description": "string",
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0,
      "annotations": {
        "property1": "string",
        "property2": "string"
      }
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|resources|[[Resource](#schemaresource)]|false|none|none|

<h2 id="tocS_CreateUserResponse">CreateUserResponse</h2>
<!-- backwards compatibility -->
<a id="schemacreateuserresponse"></a>
<a id="schema_CreateUserResponse"></a>
<a id="tocScreateuserresponse"></a>
<a id="tocscreateuserresponse"></a>

```json
{
  "user": {
    "id": "string",
    "username": "string",
    "password": "string",
    "securityContext": {
      "constraints": [
        {
          "namespace": "string",
          "resource": "string",
          "property": "string",
          "before": "2019-08-24T14:15:22Z",
          "after": "2019-08-24T14:15:22Z",
          "principal": "string",
          "recordIds": [
            "string"
          ],
          "operation": "OPERATION_TYPE_READ",
          "permit": "PERMIT_TYPE_ALLOW"
        }
      ]
    },
    "details": {},
    "signKey": "string",
    "auditData": {
      "createdOn": "2019-08-24T14:15:22Z",
      "updatedOn": "2019-08-24T14:15:22Z",
      "createdBy": "string",
      "updatedBy": "string"
    },
    "version": 0
  },
  "users": [
    {
      "id": "string",
      "username": "string",
      "password": "string",
      "securityContext": {
        "constraints": [
          {
            "namespace": "string",
            "resource": "string",
            "property": "string",
            "before": "2019-08-24T14:15:22Z",
            "after": "2019-08-24T14:15:22Z",
            "principal": "string",
            "recordIds": [
              "string"
            ],
            "operation": "OPERATION_TYPE_READ",
            "permit": "PERMIT_TYPE_ALLOW"
          }
        ]
      },
      "details": {},
      "signKey": "string",
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|user|[User](#schemauser)|false|none|none|
|users|[[User](#schemauser)]|false|none|none|

<h2 id="tocS_DataSource">DataSource</h2>
<!-- backwards compatibility -->
<a id="schemadatasource"></a>
<a id="schema_DataSource"></a>
<a id="tocSdatasource"></a>
<a id="tocsdatasource"></a>

```json
{
  "id": "string",
  "backend": "POSTGRESQL",
  "name": "string",
  "description": "string",
  "postgresqlParams": {
    "username": "string",
    "password": "string",
    "host": "string",
    "port": 0,
    "dbName": "string",
    "defaultSchema": "string"
  },
  "mysqlParams": {
    "username": "string",
    "password": "string",
    "host": "string",
    "port": 0,
    "dbName": "string",
    "defaultSchema": "string"
  },
  "virtualParams": {
    "mode": "DISCARD"
  },
  "redisParams": {
    "addr": "string",
    "password": "string",
    "db": 0
  },
  "mongoParams": {
    "uri": "string",
    "dbName": "string"
  },
  "auditData": {
    "createdOn": "2019-08-24T14:15:22Z",
    "updatedOn": "2019-08-24T14:15:22Z",
    "createdBy": "string",
    "updatedBy": "string"
  },
  "version": 0
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|id|string|false|none|none|
|backend|string(enum)|false|none|none|
|name|string|false|none|none|
|description|string|false|none|none|
|postgresqlParams|[PostgresqlParams](#schemapostgresqlparams)|false|none|none|
|mysqlParams|[MysqlParams](#schemamysqlparams)|false|none|none|
|virtualParams|[VirtualParams](#schemavirtualparams)|false|none|none|
|redisParams|[RedisParams](#schemaredisparams)|false|none|none|
|mongoParams|[MongoParams](#schemamongoparams)|false|none|none|
|auditData|[AuditData](#schemaauditdata)|false|none|none|
|version|integer(uint32)|false|none|none|

#### Enumerated Values

|Property|Value|
|---|---|
|backend|POSTGRESQL|
|backend|VIRTUAL|
|backend|MYSQL|
|backend|ORACLE|
|backend|MONGODB|
|backend|REDIS|

<h2 id="tocS_DataSourceCatalog">DataSourceCatalog</h2>
<!-- backwards compatibility -->
<a id="schemadatasourcecatalog"></a>
<a id="schema_DataSourceCatalog"></a>
<a id="tocSdatasourcecatalog"></a>
<a id="tocsdatasourcecatalog"></a>

```json
{
  "name": "string",
  "entities": [
    {
      "name": "string",
      "readOnly": true
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|name|string|false|none|none|
|entities|[[DataSourceEntity](#schemadatasourceentity)]|false|none|none|

<h2 id="tocS_DataSourceEntity">DataSourceEntity</h2>
<!-- backwards compatibility -->
<a id="schemadatasourceentity"></a>
<a id="schema_DataSourceEntity"></a>
<a id="tocSdatasourceentity"></a>
<a id="tocsdatasourceentity"></a>

```json
{
  "name": "string",
  "readOnly": true
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|name|string|false|none|none|
|readOnly|boolean|false|none|none|

<h2 id="tocS_DeleteDataSourceRequest">DeleteDataSourceRequest</h2>
<!-- backwards compatibility -->
<a id="schemadeletedatasourcerequest"></a>
<a id="schema_DeleteDataSourceRequest"></a>
<a id="tocSdeletedatasourcerequest"></a>
<a id="tocsdeletedatasourcerequest"></a>

```json
{
  "token": "string",
  "ids": [
    "string"
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|token|string|false|none|none|
|ids|[string]|false|none|none|

<h2 id="tocS_DeleteDataSourceResponse">DeleteDataSourceResponse</h2>
<!-- backwards compatibility -->
<a id="schemadeletedatasourceresponse"></a>
<a id="schema_DeleteDataSourceResponse"></a>
<a id="tocSdeletedatasourceresponse"></a>
<a id="tocsdeletedatasourceresponse"></a>

```json
{}

```

### Properties

*None*

<h2 id="tocS_DeleteExtensionRequest">DeleteExtensionRequest</h2>
<!-- backwards compatibility -->
<a id="schemadeleteextensionrequest"></a>
<a id="schema_DeleteExtensionRequest"></a>
<a id="tocSdeleteextensionrequest"></a>
<a id="tocsdeleteextensionrequest"></a>

```json
{
  "token": "string",
  "ids": [
    "string"
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|token|string|false|none|none|
|ids|[string]|false|none|none|

<h2 id="tocS_DeleteExtensionResponse">DeleteExtensionResponse</h2>
<!-- backwards compatibility -->
<a id="schemadeleteextensionresponse"></a>
<a id="schema_DeleteExtensionResponse"></a>
<a id="tocSdeleteextensionresponse"></a>
<a id="tocsdeleteextensionresponse"></a>

```json
{}

```

### Properties

*None*

<h2 id="tocS_DeleteNamespaceRequest">DeleteNamespaceRequest</h2>
<!-- backwards compatibility -->
<a id="schemadeletenamespacerequest"></a>
<a id="schema_DeleteNamespaceRequest"></a>
<a id="tocSdeletenamespacerequest"></a>
<a id="tocsdeletenamespacerequest"></a>

```json
{
  "token": "string",
  "ids": [
    "string"
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|token|string|false|none|none|
|ids|[string]|false|none|none|

<h2 id="tocS_DeleteNamespaceResponse">DeleteNamespaceResponse</h2>
<!-- backwards compatibility -->
<a id="schemadeletenamespaceresponse"></a>
<a id="schema_DeleteNamespaceResponse"></a>
<a id="tocSdeletenamespaceresponse"></a>
<a id="tocsdeletenamespaceresponse"></a>

```json
{
  "Namespaces": [
    {
      "id": "string",
      "name": "string",
      "description": "string",
      "details": {},
      "securityContext": {
        "constraints": [
          {
            "namespace": "string",
            "resource": "string",
            "property": "string",
            "before": "2019-08-24T14:15:22Z",
            "after": "2019-08-24T14:15:22Z",
            "principal": "string",
            "recordIds": [
              "string"
            ],
            "operation": "OPERATION_TYPE_READ",
            "permit": "PERMIT_TYPE_ALLOW"
          }
        ]
      },
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|Namespaces|[[Namespace](#schemanamespace)]|false|none|none|

<h2 id="tocS_DeleteRecordResponse">DeleteRecordResponse</h2>
<!-- backwards compatibility -->
<a id="schemadeleterecordresponse"></a>
<a id="schema_DeleteRecordResponse"></a>
<a id="tocSdeleterecordresponse"></a>
<a id="tocsdeleterecordresponse"></a>

```json
{}

```

### Properties

*None*

<h2 id="tocS_DeleteResourceRequest">DeleteResourceRequest</h2>
<!-- backwards compatibility -->
<a id="schemadeleteresourcerequest"></a>
<a id="schema_DeleteResourceRequest"></a>
<a id="tocSdeleteresourcerequest"></a>
<a id="tocsdeleteresourcerequest"></a>

```json
{
  "token": "string",
  "ids": [
    "string"
  ],
  "doMigration": true,
  "forceMigration": true,
  "annotations": {
    "property1": "string",
    "property2": "string"
  }
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|token|string|false|none|none|
|ids|[string]|false|none|none|
|doMigration|boolean|false|none|none|
|forceMigration|boolean|false|none|none|
|annotations|object|false|none|none|
|» **additionalProperties**|string|false|none|none|

<h2 id="tocS_DeleteResourceResponse">DeleteResourceResponse</h2>
<!-- backwards compatibility -->
<a id="schemadeleteresourceresponse"></a>
<a id="schema_DeleteResourceResponse"></a>
<a id="tocSdeleteresourceresponse"></a>
<a id="tocsdeleteresourceresponse"></a>

```json
{}

```

### Properties

*None*

<h2 id="tocS_DeleteUserResponse">DeleteUserResponse</h2>
<!-- backwards compatibility -->
<a id="schemadeleteuserresponse"></a>
<a id="schema_DeleteUserResponse"></a>
<a id="tocSdeleteuserresponse"></a>
<a id="tocsdeleteuserresponse"></a>

```json
{}

```

### Properties

*None*

<h2 id="tocS_Expression">Expression</h2>
<!-- backwards compatibility -->
<a id="schemaexpression"></a>
<a id="schema_Expression"></a>
<a id="tocSexpression"></a>
<a id="tocsexpression"></a>

```json
{
  "additionalProperties": [
    {
      "name": "string",
      "value": {
        "value": {
          "@type": "string"
        },
        "yaml": "string"
      }
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|additionalProperties|[[NamedAny](#schemanamedany)]|false|none|[Automatically-generated message used to represent maps of Any as ordered (name,value) pairs.]|

<h2 id="tocS_Extension">Extension</h2>
<!-- backwards compatibility -->
<a id="schemaextension"></a>
<a id="schema_Extension"></a>
<a id="tocSextension"></a>
<a id="tocsextension"></a>

```json
{
  "id": "string",
  "name": "string",
  "description": "string",
  "namespace": "string",
  "resource": "string",
  "before": {
    "all": {
      "functionCall": {
        "host": "string",
        "functionName": "string"
      },
      "httpCall": {
        "uri": "string",
        "method": "string"
      }
    },
    "create": {
      "functionCall": {
        "host": "string",
        "functionName": "string"
      },
      "httpCall": {
        "uri": "string",
        "method": "string"
      }
    },
    "update": {
      "functionCall": {
        "host": "string",
        "functionName": "string"
      },
      "httpCall": {
        "uri": "string",
        "method": "string"
      }
    },
    "delete": {
      "functionCall": {
        "host": "string",
        "functionName": "string"
      },
      "httpCall": {
        "uri": "string",
        "method": "string"
      }
    },
    "get": {
      "functionCall": {
        "host": "string",
        "functionName": "string"
      },
      "httpCall": {
        "uri": "string",
        "method": "string"
      }
    },
    "list": {
      "functionCall": {
        "host": "string",
        "functionName": "string"
      },
      "httpCall": {
        "uri": "string",
        "method": "string"
      }
    },
    "sync": true
  },
  "instead": {
    "all": {
      "functionCall": {
        "host": "string",
        "functionName": "string"
      },
      "httpCall": {
        "uri": "string",
        "method": "string"
      }
    },
    "create": {
      "functionCall": {
        "host": "string",
        "functionName": "string"
      },
      "httpCall": {
        "uri": "string",
        "method": "string"
      }
    },
    "update": {
      "functionCall": {
        "host": "string",
        "functionName": "string"
      },
      "httpCall": {
        "uri": "string",
        "method": "string"
      }
    },
    "delete": {
      "functionCall": {
        "host": "string",
        "functionName": "string"
      },
      "httpCall": {
        "uri": "string",
        "method": "string"
      }
    },
    "get": {
      "functionCall": {
        "host": "string",
        "functionName": "string"
      },
      "httpCall": {
        "uri": "string",
        "method": "string"
      }
    },
    "list": {
      "functionCall": {
        "host": "string",
        "functionName": "string"
      },
      "httpCall": {
        "uri": "string",
        "method": "string"
      }
    },
    "finalize": true
  },
  "after": {
    "all": {
      "functionCall": {
        "host": "string",
        "functionName": "string"
      },
      "httpCall": {
        "uri": "string",
        "method": "string"
      }
    },
    "create": {
      "functionCall": {
        "host": "string",
        "functionName": "string"
      },
      "httpCall": {
        "uri": "string",
        "method": "string"
      }
    },
    "update": {
      "functionCall": {
        "host": "string",
        "functionName": "string"
      },
      "httpCall": {
        "uri": "string",
        "method": "string"
      }
    },
    "delete": {
      "functionCall": {
        "host": "string",
        "functionName": "string"
      },
      "httpCall": {
        "uri": "string",
        "method": "string"
      }
    },
    "get": {
      "functionCall": {
        "host": "string",
        "functionName": "string"
      },
      "httpCall": {
        "uri": "string",
        "method": "string"
      }
    },
    "list": {
      "functionCall": {
        "host": "string",
        "functionName": "string"
      },
      "httpCall": {
        "uri": "string",
        "method": "string"
      }
    },
    "sync": true
  },
  "auditData": {
    "createdOn": "2019-08-24T14:15:22Z",
    "updatedOn": "2019-08-24T14:15:22Z",
    "createdBy": "string",
    "updatedBy": "string"
  },
  "version": 0
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|id|string|false|none|none|
|name|string|false|none|none|
|description|string|false|none|none|
|namespace|string|false|none|none|
|resource|string|false|none|none|
|before|[Extension_Before](#schemaextension_before)|false|none|none|
|instead|[Extension_Instead](#schemaextension_instead)|false|none|none|
|after|[Extension_After](#schemaextension_after)|false|none|none|
|auditData|[AuditData](#schemaauditdata)|false|none|none|
|version|integer(uint32)|false|none|none|

<h2 id="tocS_Extension_After">Extension_After</h2>
<!-- backwards compatibility -->
<a id="schemaextension_after"></a>
<a id="schema_Extension_After"></a>
<a id="tocSextension_after"></a>
<a id="tocsextension_after"></a>

```json
{
  "all": {
    "functionCall": {
      "host": "string",
      "functionName": "string"
    },
    "httpCall": {
      "uri": "string",
      "method": "string"
    }
  },
  "create": {
    "functionCall": {
      "host": "string",
      "functionName": "string"
    },
    "httpCall": {
      "uri": "string",
      "method": "string"
    }
  },
  "update": {
    "functionCall": {
      "host": "string",
      "functionName": "string"
    },
    "httpCall": {
      "uri": "string",
      "method": "string"
    }
  },
  "delete": {
    "functionCall": {
      "host": "string",
      "functionName": "string"
    },
    "httpCall": {
      "uri": "string",
      "method": "string"
    }
  },
  "get": {
    "functionCall": {
      "host": "string",
      "functionName": "string"
    },
    "httpCall": {
      "uri": "string",
      "method": "string"
    }
  },
  "list": {
    "functionCall": {
      "host": "string",
      "functionName": "string"
    },
    "httpCall": {
      "uri": "string",
      "method": "string"
    }
  },
  "sync": true
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|all|[ExternalCall](#schemaexternalcall)|false|none|none|
|create|[ExternalCall](#schemaexternalcall)|false|none|none|
|update|[ExternalCall](#schemaexternalcall)|false|none|none|
|delete|[ExternalCall](#schemaexternalcall)|false|none|none|
|get|[ExternalCall](#schemaexternalcall)|false|none|none|
|list|[ExternalCall](#schemaexternalcall)|false|none|none|
|sync|boolean|false|none|none|

<h2 id="tocS_Extension_Before">Extension_Before</h2>
<!-- backwards compatibility -->
<a id="schemaextension_before"></a>
<a id="schema_Extension_Before"></a>
<a id="tocSextension_before"></a>
<a id="tocsextension_before"></a>

```json
{
  "all": {
    "functionCall": {
      "host": "string",
      "functionName": "string"
    },
    "httpCall": {
      "uri": "string",
      "method": "string"
    }
  },
  "create": {
    "functionCall": {
      "host": "string",
      "functionName": "string"
    },
    "httpCall": {
      "uri": "string",
      "method": "string"
    }
  },
  "update": {
    "functionCall": {
      "host": "string",
      "functionName": "string"
    },
    "httpCall": {
      "uri": "string",
      "method": "string"
    }
  },
  "delete": {
    "functionCall": {
      "host": "string",
      "functionName": "string"
    },
    "httpCall": {
      "uri": "string",
      "method": "string"
    }
  },
  "get": {
    "functionCall": {
      "host": "string",
      "functionName": "string"
    },
    "httpCall": {
      "uri": "string",
      "method": "string"
    }
  },
  "list": {
    "functionCall": {
      "host": "string",
      "functionName": "string"
    },
    "httpCall": {
      "uri": "string",
      "method": "string"
    }
  },
  "sync": true
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|all|[ExternalCall](#schemaexternalcall)|false|none|none|
|create|[ExternalCall](#schemaexternalcall)|false|none|none|
|update|[ExternalCall](#schemaexternalcall)|false|none|none|
|delete|[ExternalCall](#schemaexternalcall)|false|none|none|
|get|[ExternalCall](#schemaexternalcall)|false|none|none|
|list|[ExternalCall](#schemaexternalcall)|false|none|none|
|sync|boolean|false|none|none|

<h2 id="tocS_Extension_Instead">Extension_Instead</h2>
<!-- backwards compatibility -->
<a id="schemaextension_instead"></a>
<a id="schema_Extension_Instead"></a>
<a id="tocSextension_instead"></a>
<a id="tocsextension_instead"></a>

```json
{
  "all": {
    "functionCall": {
      "host": "string",
      "functionName": "string"
    },
    "httpCall": {
      "uri": "string",
      "method": "string"
    }
  },
  "create": {
    "functionCall": {
      "host": "string",
      "functionName": "string"
    },
    "httpCall": {
      "uri": "string",
      "method": "string"
    }
  },
  "update": {
    "functionCall": {
      "host": "string",
      "functionName": "string"
    },
    "httpCall": {
      "uri": "string",
      "method": "string"
    }
  },
  "delete": {
    "functionCall": {
      "host": "string",
      "functionName": "string"
    },
    "httpCall": {
      "uri": "string",
      "method": "string"
    }
  },
  "get": {
    "functionCall": {
      "host": "string",
      "functionName": "string"
    },
    "httpCall": {
      "uri": "string",
      "method": "string"
    }
  },
  "list": {
    "functionCall": {
      "host": "string",
      "functionName": "string"
    },
    "httpCall": {
      "uri": "string",
      "method": "string"
    }
  },
  "finalize": true
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|all|[ExternalCall](#schemaexternalcall)|false|none|none|
|create|[ExternalCall](#schemaexternalcall)|false|none|none|
|update|[ExternalCall](#schemaexternalcall)|false|none|none|
|delete|[ExternalCall](#schemaexternalcall)|false|none|none|
|get|[ExternalCall](#schemaexternalcall)|false|none|none|
|list|[ExternalCall](#schemaexternalcall)|false|none|none|
|finalize|boolean|false|none|none|

<h2 id="tocS_ExternalCall">ExternalCall</h2>
<!-- backwards compatibility -->
<a id="schemaexternalcall"></a>
<a id="schema_ExternalCall"></a>
<a id="tocSexternalcall"></a>
<a id="tocsexternalcall"></a>

```json
{
  "functionCall": {
    "host": "string",
    "functionName": "string"
  },
  "httpCall": {
    "uri": "string",
    "method": "string"
  }
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|functionCall|[FunctionCall](#schemafunctioncall)|false|none|none|
|httpCall|[HttpCall](#schemahttpcall)|false|none|none|

<h2 id="tocS_FunctionCall">FunctionCall</h2>
<!-- backwards compatibility -->
<a id="schemafunctioncall"></a>
<a id="schema_FunctionCall"></a>
<a id="tocSfunctioncall"></a>
<a id="tocsfunctioncall"></a>

```json
{
  "host": "string",
  "functionName": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|host|string|false|none|none|
|functionName|string|false|none|none|

<h2 id="tocS_GetDataSourceResponse">GetDataSourceResponse</h2>
<!-- backwards compatibility -->
<a id="schemagetdatasourceresponse"></a>
<a id="schema_GetDataSourceResponse"></a>
<a id="tocSgetdatasourceresponse"></a>
<a id="tocsgetdatasourceresponse"></a>

```json
{
  "dataSource": {
    "id": "string",
    "backend": "POSTGRESQL",
    "name": "string",
    "description": "string",
    "postgresqlParams": {
      "username": "string",
      "password": "string",
      "host": "string",
      "port": 0,
      "dbName": "string",
      "defaultSchema": "string"
    },
    "mysqlParams": {
      "username": "string",
      "password": "string",
      "host": "string",
      "port": 0,
      "dbName": "string",
      "defaultSchema": "string"
    },
    "virtualParams": {
      "mode": "DISCARD"
    },
    "redisParams": {
      "addr": "string",
      "password": "string",
      "db": 0
    },
    "mongoParams": {
      "uri": "string",
      "dbName": "string"
    },
    "auditData": {
      "createdOn": "2019-08-24T14:15:22Z",
      "updatedOn": "2019-08-24T14:15:22Z",
      "createdBy": "string",
      "updatedBy": "string"
    },
    "version": 0
  }
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|dataSource|[DataSource](#schemadatasource)|false|none|none|

<h2 id="tocS_GetExtensionResponse">GetExtensionResponse</h2>
<!-- backwards compatibility -->
<a id="schemagetextensionresponse"></a>
<a id="schema_GetExtensionResponse"></a>
<a id="tocSgetextensionresponse"></a>
<a id="tocsgetextensionresponse"></a>

```json
{
  "extension": {
    "id": "string",
    "name": "string",
    "description": "string",
    "namespace": "string",
    "resource": "string",
    "before": {
      "all": {
        "functionCall": {
          "host": "string",
          "functionName": "string"
        },
        "httpCall": {
          "uri": "string",
          "method": "string"
        }
      },
      "create": {
        "functionCall": {
          "host": "string",
          "functionName": "string"
        },
        "httpCall": {
          "uri": "string",
          "method": "string"
        }
      },
      "update": {
        "functionCall": {
          "host": "string",
          "functionName": "string"
        },
        "httpCall": {
          "uri": "string",
          "method": "string"
        }
      },
      "delete": {
        "functionCall": {
          "host": "string",
          "functionName": "string"
        },
        "httpCall": {
          "uri": "string",
          "method": "string"
        }
      },
      "get": {
        "functionCall": {
          "host": "string",
          "functionName": "string"
        },
        "httpCall": {
          "uri": "string",
          "method": "string"
        }
      },
      "list": {
        "functionCall": {
          "host": "string",
          "functionName": "string"
        },
        "httpCall": {
          "uri": "string",
          "method": "string"
        }
      },
      "sync": true
    },
    "instead": {
      "all": {
        "functionCall": {
          "host": "string",
          "functionName": "string"
        },
        "httpCall": {
          "uri": "string",
          "method": "string"
        }
      },
      "create": {
        "functionCall": {
          "host": "string",
          "functionName": "string"
        },
        "httpCall": {
          "uri": "string",
          "method": "string"
        }
      },
      "update": {
        "functionCall": {
          "host": "string",
          "functionName": "string"
        },
        "httpCall": {
          "uri": "string",
          "method": "string"
        }
      },
      "delete": {
        "functionCall": {
          "host": "string",
          "functionName": "string"
        },
        "httpCall": {
          "uri": "string",
          "method": "string"
        }
      },
      "get": {
        "functionCall": {
          "host": "string",
          "functionName": "string"
        },
        "httpCall": {
          "uri": "string",
          "method": "string"
        }
      },
      "list": {
        "functionCall": {
          "host": "string",
          "functionName": "string"
        },
        "httpCall": {
          "uri": "string",
          "method": "string"
        }
      },
      "finalize": true
    },
    "after": {
      "all": {
        "functionCall": {
          "host": "string",
          "functionName": "string"
        },
        "httpCall": {
          "uri": "string",
          "method": "string"
        }
      },
      "create": {
        "functionCall": {
          "host": "string",
          "functionName": "string"
        },
        "httpCall": {
          "uri": "string",
          "method": "string"
        }
      },
      "update": {
        "functionCall": {
          "host": "string",
          "functionName": "string"
        },
        "httpCall": {
          "uri": "string",
          "method": "string"
        }
      },
      "delete": {
        "functionCall": {
          "host": "string",
          "functionName": "string"
        },
        "httpCall": {
          "uri": "string",
          "method": "string"
        }
      },
      "get": {
        "functionCall": {
          "host": "string",
          "functionName": "string"
        },
        "httpCall": {
          "uri": "string",
          "method": "string"
        }
      },
      "list": {
        "functionCall": {
          "host": "string",
          "functionName": "string"
        },
        "httpCall": {
          "uri": "string",
          "method": "string"
        }
      },
      "sync": true
    },
    "auditData": {
      "createdOn": "2019-08-24T14:15:22Z",
      "updatedOn": "2019-08-24T14:15:22Z",
      "createdBy": "string",
      "updatedBy": "string"
    },
    "version": 0
  }
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|extension|[Extension](#schemaextension)|false|none|none|

<h2 id="tocS_GetNamespaceResponse">GetNamespaceResponse</h2>
<!-- backwards compatibility -->
<a id="schemagetnamespaceresponse"></a>
<a id="schema_GetNamespaceResponse"></a>
<a id="tocSgetnamespaceresponse"></a>
<a id="tocsgetnamespaceresponse"></a>

```json
{
  "Namespace": {
    "id": "string",
    "name": "string",
    "description": "string",
    "details": {},
    "securityContext": {
      "constraints": [
        {
          "namespace": "string",
          "resource": "string",
          "property": "string",
          "before": "2019-08-24T14:15:22Z",
          "after": "2019-08-24T14:15:22Z",
          "principal": "string",
          "recordIds": [
            "string"
          ],
          "operation": "OPERATION_TYPE_READ",
          "permit": "PERMIT_TYPE_ALLOW"
        }
      ]
    },
    "auditData": {
      "createdOn": "2019-08-24T14:15:22Z",
      "updatedOn": "2019-08-24T14:15:22Z",
      "createdBy": "string",
      "updatedBy": "string"
    },
    "version": 0
  }
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|Namespace|[Namespace](#schemanamespace)|false|none|none|

<h2 id="tocS_GetRecordResponse">GetRecordResponse</h2>
<!-- backwards compatibility -->
<a id="schemagetrecordresponse"></a>
<a id="schema_GetRecordResponse"></a>
<a id="tocSgetrecordresponse"></a>
<a id="tocsgetrecordresponse"></a>

```json
{
  "record": {
    "id": "string",
    "properties": {
      "property1": null,
      "property2": null
    },
    "propertiesPacked": [
      null
    ]
  }
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|record|[Record](#schemarecord)|false|none|none|

<h2 id="tocS_GetResourceByNameResponse">GetResourceByNameResponse</h2>
<!-- backwards compatibility -->
<a id="schemagetresourcebynameresponse"></a>
<a id="schema_GetResourceByNameResponse"></a>
<a id="tocSgetresourcebynameresponse"></a>
<a id="tocsgetresourcebynameresponse"></a>

```json
{
  "resource": {
    "id": "string",
    "name": "string",
    "namespace": "string",
    "sourceConfig": {
      "dataSource": "string",
      "catalog": "string",
      "entity": "string"
    },
    "properties": [
      {
        "id": "string",
        "name": "string",
        "type": "BOOL",
        "mapping": "string",
        "required": true,
        "primary": true,
        "length": 0,
        "unique": true,
        "immutable": true,
        "securityContext": {
          "constraints": [
            {
              "namespace": "string",
              "resource": "string",
              "property": "string",
              "before": "2019-08-24T14:15:22Z",
              "after": "2019-08-24T14:15:22Z",
              "principal": "string",
              "recordIds": [
                "string"
              ],
              "operation": "OPERATION_TYPE_READ",
              "permit": "PERMIT_TYPE_ALLOW"
            }
          ]
        },
        "defaultValue": null,
        "exampleValue": null,
        "enumValues": [
          null
        ],
        "reference": {
          "referencedResource": "string",
          "cascade": true
        },
        "properties": [
          {}
        ],
        "Item": {},
        "title": "string",
        "description": "string",
        "annotations": {
          "property1": "string",
          "property2": "string"
        }
      }
    ],
    "indexes": [
      {
        "properties": [
          {
            "name": "string",
            "order": "ORDER_UNKNOWN"
          }
        ],
        "indexType": "BTREE",
        "unique": true,
        "annotations": {
          "property1": "string",
          "property2": "string"
        }
      }
    ],
    "securityContext": {
      "constraints": [
        {
          "namespace": "string",
          "resource": "string",
          "property": "string",
          "before": "2019-08-24T14:15:22Z",
          "after": "2019-08-24T14:15:22Z",
          "principal": "string",
          "recordIds": [
            "string"
          ],
          "operation": "OPERATION_TYPE_READ",
          "permit": "PERMIT_TYPE_ALLOW"
        }
      ]
    },
    "virtual": true,
    "immutable": true,
    "abstract": true,
    "title": "string",
    "description": "string",
    "auditData": {
      "createdOn": "2019-08-24T14:15:22Z",
      "updatedOn": "2019-08-24T14:15:22Z",
      "createdBy": "string",
      "updatedBy": "string"
    },
    "version": 0,
    "annotations": {
      "property1": "string",
      "property2": "string"
    }
  }
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|resource|[Resource](#schemaresource)|false|none|none|

<h2 id="tocS_GetResourceResponse">GetResourceResponse</h2>
<!-- backwards compatibility -->
<a id="schemagetresourceresponse"></a>
<a id="schema_GetResourceResponse"></a>
<a id="tocSgetresourceresponse"></a>
<a id="tocsgetresourceresponse"></a>

```json
{
  "resource": {
    "id": "string",
    "name": "string",
    "namespace": "string",
    "sourceConfig": {
      "dataSource": "string",
      "catalog": "string",
      "entity": "string"
    },
    "properties": [
      {
        "id": "string",
        "name": "string",
        "type": "BOOL",
        "mapping": "string",
        "required": true,
        "primary": true,
        "length": 0,
        "unique": true,
        "immutable": true,
        "securityContext": {
          "constraints": [
            {
              "namespace": "string",
              "resource": "string",
              "property": "string",
              "before": "2019-08-24T14:15:22Z",
              "after": "2019-08-24T14:15:22Z",
              "principal": "string",
              "recordIds": [
                "string"
              ],
              "operation": "OPERATION_TYPE_READ",
              "permit": "PERMIT_TYPE_ALLOW"
            }
          ]
        },
        "defaultValue": null,
        "exampleValue": null,
        "enumValues": [
          null
        ],
        "reference": {
          "referencedResource": "string",
          "cascade": true
        },
        "properties": [
          {}
        ],
        "Item": {},
        "title": "string",
        "description": "string",
        "annotations": {
          "property1": "string",
          "property2": "string"
        }
      }
    ],
    "indexes": [
      {
        "properties": [
          {
            "name": "string",
            "order": "ORDER_UNKNOWN"
          }
        ],
        "indexType": "BTREE",
        "unique": true,
        "annotations": {
          "property1": "string",
          "property2": "string"
        }
      }
    ],
    "securityContext": {
      "constraints": [
        {
          "namespace": "string",
          "resource": "string",
          "property": "string",
          "before": "2019-08-24T14:15:22Z",
          "after": "2019-08-24T14:15:22Z",
          "principal": "string",
          "recordIds": [
            "string"
          ],
          "operation": "OPERATION_TYPE_READ",
          "permit": "PERMIT_TYPE_ALLOW"
        }
      ]
    },
    "virtual": true,
    "immutable": true,
    "abstract": true,
    "title": "string",
    "description": "string",
    "auditData": {
      "createdOn": "2019-08-24T14:15:22Z",
      "updatedOn": "2019-08-24T14:15:22Z",
      "createdBy": "string",
      "updatedBy": "string"
    },
    "version": 0,
    "annotations": {
      "property1": "string",
      "property2": "string"
    }
  }
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|resource|[Resource](#schemaresource)|false|none|none|

<h2 id="tocS_GetUserResponse">GetUserResponse</h2>
<!-- backwards compatibility -->
<a id="schemagetuserresponse"></a>
<a id="schema_GetUserResponse"></a>
<a id="tocSgetuserresponse"></a>
<a id="tocsgetuserresponse"></a>

```json
{
  "user": {
    "id": "string",
    "username": "string",
    "password": "string",
    "securityContext": {
      "constraints": [
        {
          "namespace": "string",
          "resource": "string",
          "property": "string",
          "before": "2019-08-24T14:15:22Z",
          "after": "2019-08-24T14:15:22Z",
          "principal": "string",
          "recordIds": [
            "string"
          ],
          "operation": "OPERATION_TYPE_READ",
          "permit": "PERMIT_TYPE_ALLOW"
        }
      ]
    },
    "details": {},
    "signKey": "string",
    "auditData": {
      "createdOn": "2019-08-24T14:15:22Z",
      "updatedOn": "2019-08-24T14:15:22Z",
      "createdBy": "string",
      "updatedBy": "string"
    },
    "version": 0
  }
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|user|[User](#schemauser)|false|none|none|

<h2 id="tocS_GoogleProtobufAny">GoogleProtobufAny</h2>
<!-- backwards compatibility -->
<a id="schemagoogleprotobufany"></a>
<a id="schema_GoogleProtobufAny"></a>
<a id="tocSgoogleprotobufany"></a>
<a id="tocsgoogleprotobufany"></a>

```json
{
  "@type": "string"
}

```

Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|@type|string|false|none|The type of the serialized message.|

<h2 id="tocS_GoogleProtobufValue">GoogleProtobufValue</h2>
<!-- backwards compatibility -->
<a id="schemagoogleprotobufvalue"></a>
<a id="schema_GoogleProtobufValue"></a>
<a id="tocSgoogleprotobufvalue"></a>
<a id="tocsgoogleprotobufvalue"></a>

```json
null

```

Represents a dynamically typed value which can be either null, a number, a string, a boolean, a recursive struct value, or a list of values.

### Properties

*None*

<h2 id="tocS_HttpCall">HttpCall</h2>
<!-- backwards compatibility -->
<a id="schemahttpcall"></a>
<a id="schema_HttpCall"></a>
<a id="tocShttpcall"></a>
<a id="tocshttpcall"></a>

```json
{
  "uri": "string",
  "method": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|uri|string|false|none|none|
|method|string|false|none|none|

<h2 id="tocS_ListDataSourceResponse">ListDataSourceResponse</h2>
<!-- backwards compatibility -->
<a id="schemalistdatasourceresponse"></a>
<a id="schema_ListDataSourceResponse"></a>
<a id="tocSlistdatasourceresponse"></a>
<a id="tocslistdatasourceresponse"></a>

```json
{
  "content": [
    {
      "id": "string",
      "backend": "POSTGRESQL",
      "name": "string",
      "description": "string",
      "postgresqlParams": {
        "username": "string",
        "password": "string",
        "host": "string",
        "port": 0,
        "dbName": "string",
        "defaultSchema": "string"
      },
      "mysqlParams": {
        "username": "string",
        "password": "string",
        "host": "string",
        "port": 0,
        "dbName": "string",
        "defaultSchema": "string"
      },
      "virtualParams": {
        "mode": "DISCARD"
      },
      "redisParams": {
        "addr": "string",
        "password": "string",
        "db": 0
      },
      "mongoParams": {
        "uri": "string",
        "dbName": "string"
      },
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|content|[[DataSource](#schemadatasource)]|false|none|none|

<h2 id="tocS_ListEntitiesResponse">ListEntitiesResponse</h2>
<!-- backwards compatibility -->
<a id="schemalistentitiesresponse"></a>
<a id="schema_ListEntitiesResponse"></a>
<a id="tocSlistentitiesresponse"></a>
<a id="tocslistentitiesresponse"></a>

```json
{
  "catalogs": [
    {
      "name": "string",
      "entities": [
        {
          "name": "string",
          "readOnly": true
        }
      ]
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|catalogs|[[DataSourceCatalog](#schemadatasourcecatalog)]|false|none|none|

<h2 id="tocS_ListExtensionResponse">ListExtensionResponse</h2>
<!-- backwards compatibility -->
<a id="schemalistextensionresponse"></a>
<a id="schema_ListExtensionResponse"></a>
<a id="tocSlistextensionresponse"></a>
<a id="tocslistextensionresponse"></a>

```json
{
  "content": [
    {
      "id": "string",
      "name": "string",
      "description": "string",
      "namespace": "string",
      "resource": "string",
      "before": {
        "all": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "create": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "update": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "delete": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "get": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "list": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "sync": true
      },
      "instead": {
        "all": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "create": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "update": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "delete": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "get": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "list": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "finalize": true
      },
      "after": {
        "all": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "create": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "update": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "delete": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "get": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "list": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "sync": true
      },
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|content|[[Extension](#schemaextension)]|false|none|none|

<h2 id="tocS_ListNamespaceResponse">ListNamespaceResponse</h2>
<!-- backwards compatibility -->
<a id="schemalistnamespaceresponse"></a>
<a id="schema_ListNamespaceResponse"></a>
<a id="tocSlistnamespaceresponse"></a>
<a id="tocslistnamespaceresponse"></a>

```json
{
  "content": [
    {
      "id": "string",
      "name": "string",
      "description": "string",
      "details": {},
      "securityContext": {
        "constraints": [
          {
            "namespace": "string",
            "resource": "string",
            "property": "string",
            "before": "2019-08-24T14:15:22Z",
            "after": "2019-08-24T14:15:22Z",
            "principal": "string",
            "recordIds": [
              "string"
            ],
            "operation": "OPERATION_TYPE_READ",
            "permit": "PERMIT_TYPE_ALLOW"
          }
        ]
      },
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|content|[[Namespace](#schemanamespace)]|false|none|none|

<h2 id="tocS_ListRecordResponse">ListRecordResponse</h2>
<!-- backwards compatibility -->
<a id="schemalistrecordresponse"></a>
<a id="schema_ListRecordResponse"></a>
<a id="tocSlistrecordresponse"></a>
<a id="tocslistrecordresponse"></a>

```json
{
  "total": 0,
  "content": [
    {
      "id": "string",
      "properties": {
        "property1": null,
        "property2": null
      },
      "propertiesPacked": [
        null
      ]
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|total|integer(uint32)|false|none|none|
|content|[[Record](#schemarecord)]|false|none|none|

<h2 id="tocS_ListResourceResponse">ListResourceResponse</h2>
<!-- backwards compatibility -->
<a id="schemalistresourceresponse"></a>
<a id="schema_ListResourceResponse"></a>
<a id="tocSlistresourceresponse"></a>
<a id="tocslistresourceresponse"></a>

```json
{
  "resources": [
    {
      "id": "string",
      "name": "string",
      "namespace": "string",
      "sourceConfig": {
        "dataSource": "string",
        "catalog": "string",
        "entity": "string"
      },
      "properties": [
        {
          "id": "string",
          "name": "string",
          "type": "BOOL",
          "mapping": "string",
          "required": true,
          "primary": true,
          "length": 0,
          "unique": true,
          "immutable": true,
          "securityContext": {
            "constraints": [
              {
                "namespace": "string",
                "resource": "string",
                "property": "string",
                "before": "2019-08-24T14:15:22Z",
                "after": "2019-08-24T14:15:22Z",
                "principal": "string",
                "recordIds": [
                  "string"
                ],
                "operation": "OPERATION_TYPE_READ",
                "permit": "PERMIT_TYPE_ALLOW"
              }
            ]
          },
          "defaultValue": null,
          "exampleValue": null,
          "enumValues": [
            null
          ],
          "reference": {
            "referencedResource": "string",
            "cascade": true
          },
          "properties": [
            {}
          ],
          "Item": {},
          "title": "string",
          "description": "string",
          "annotations": {
            "property1": "string",
            "property2": "string"
          }
        }
      ],
      "indexes": [
        {
          "properties": [
            {
              "name": "string",
              "order": "ORDER_UNKNOWN"
            }
          ],
          "indexType": "BTREE",
          "unique": true,
          "annotations": {
            "property1": "string",
            "property2": "string"
          }
        }
      ],
      "securityContext": {
        "constraints": [
          {
            "namespace": "string",
            "resource": "string",
            "property": "string",
            "before": "2019-08-24T14:15:22Z",
            "after": "2019-08-24T14:15:22Z",
            "principal": "string",
            "recordIds": [
              "string"
            ],
            "operation": "OPERATION_TYPE_READ",
            "permit": "PERMIT_TYPE_ALLOW"
          }
        ]
      },
      "virtual": true,
      "immutable": true,
      "abstract": true,
      "title": "string",
      "description": "string",
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0,
      "annotations": {
        "property1": "string",
        "property2": "string"
      }
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|resources|[[Resource](#schemaresource)]|false|none|none|

<h2 id="tocS_ListUserResponse">ListUserResponse</h2>
<!-- backwards compatibility -->
<a id="schemalistuserresponse"></a>
<a id="schema_ListUserResponse"></a>
<a id="tocSlistuserresponse"></a>
<a id="tocslistuserresponse"></a>

```json
{
  "content": [
    {
      "id": "string",
      "username": "string",
      "password": "string",
      "securityContext": {
        "constraints": [
          {
            "namespace": "string",
            "resource": "string",
            "property": "string",
            "before": "2019-08-24T14:15:22Z",
            "after": "2019-08-24T14:15:22Z",
            "principal": "string",
            "recordIds": [
              "string"
            ],
            "operation": "OPERATION_TYPE_READ",
            "permit": "PERMIT_TYPE_ALLOW"
          }
        ]
      },
      "details": {},
      "signKey": "string",
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|content|[[User](#schemauser)]|false|none|none|

<h2 id="tocS_MongoParams">MongoParams</h2>
<!-- backwards compatibility -->
<a id="schemamongoparams"></a>
<a id="schema_MongoParams"></a>
<a id="tocSmongoparams"></a>
<a id="tocsmongoparams"></a>

```json
{
  "uri": "string",
  "dbName": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|uri|string|false|none|none|
|dbName|string|false|none|none|

<h2 id="tocS_MysqlParams">MysqlParams</h2>
<!-- backwards compatibility -->
<a id="schemamysqlparams"></a>
<a id="schema_MysqlParams"></a>
<a id="tocSmysqlparams"></a>
<a id="tocsmysqlparams"></a>

```json
{
  "username": "string",
  "password": "string",
  "host": "string",
  "port": 0,
  "dbName": "string",
  "defaultSchema": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|username|string|false|none|none|
|password|string|false|none|none|
|host|string|false|none|none|
|port|integer(uint32)|false|none|none|
|dbName|string|false|none|none|
|defaultSchema|string|false|none|none|

<h2 id="tocS_NamedAny">NamedAny</h2>
<!-- backwards compatibility -->
<a id="schemanamedany"></a>
<a id="schema_NamedAny"></a>
<a id="tocSnamedany"></a>
<a id="tocsnamedany"></a>

```json
{
  "name": "string",
  "value": {
    "value": {
      "@type": "string"
    },
    "yaml": "string"
  }
}

```

Automatically-generated message used to represent maps of Any as ordered (name,value) pairs.

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|name|string|false|none|Map key|
|value|[Any](#schemaany)|false|none|none|

<h2 id="tocS_Namespace">Namespace</h2>
<!-- backwards compatibility -->
<a id="schemanamespace"></a>
<a id="schema_Namespace"></a>
<a id="tocSnamespace"></a>
<a id="tocsnamespace"></a>

```json
{
  "id": "string",
  "name": "string",
  "description": "string",
  "details": {},
  "securityContext": {
    "constraints": [
      {
        "namespace": "string",
        "resource": "string",
        "property": "string",
        "before": "2019-08-24T14:15:22Z",
        "after": "2019-08-24T14:15:22Z",
        "principal": "string",
        "recordIds": [
          "string"
        ],
        "operation": "OPERATION_TYPE_READ",
        "permit": "PERMIT_TYPE_ALLOW"
      }
    ]
  },
  "auditData": {
    "createdOn": "2019-08-24T14:15:22Z",
    "updatedOn": "2019-08-24T14:15:22Z",
    "createdBy": "string",
    "updatedBy": "string"
  },
  "version": 0
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|id|string|false|none|none|
|name|string|false|none|none|
|description|string|false|none|none|
|details|object|false|none|none|
|securityContext|[SecurityContext](#schemasecuritycontext)|false|none|none|
|auditData|[AuditData](#schemaauditdata)|false|none|none|
|version|integer(uint32)|false|none|none|

<h2 id="tocS_PairExpression">PairExpression</h2>
<!-- backwards compatibility -->
<a id="schemapairexpression"></a>
<a id="schema_PairExpression"></a>
<a id="tocSpairexpression"></a>
<a id="tocspairexpression"></a>

```json
{
  "left": {
    "additionalProperties": [
      {
        "name": "string",
        "value": {
          "value": {
            "@type": "string"
          },
          "yaml": "string"
        }
      }
    ]
  },
  "right": {
    "additionalProperties": [
      {
        "name": "string",
        "value": {
          "value": {
            "@type": "string"
          },
          "yaml": "string"
        }
      }
    ]
  }
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|left|[Expression](#schemaexpression)|false|none|none|
|right|[Expression](#schemaexpression)|false|none|none|

<h2 id="tocS_PostgresqlParams">PostgresqlParams</h2>
<!-- backwards compatibility -->
<a id="schemapostgresqlparams"></a>
<a id="schema_PostgresqlParams"></a>
<a id="tocSpostgresqlparams"></a>
<a id="tocspostgresqlparams"></a>

```json
{
  "username": "string",
  "password": "string",
  "host": "string",
  "port": 0,
  "dbName": "string",
  "defaultSchema": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|username|string|false|none|none|
|password|string|false|none|none|
|host|string|false|none|none|
|port|integer(uint32)|false|none|none|
|dbName|string|false|none|none|
|defaultSchema|string|false|none|none|

<h2 id="tocS_PrepareResourceFromEntityRequest">PrepareResourceFromEntityRequest</h2>
<!-- backwards compatibility -->
<a id="schemaprepareresourcefromentityrequest"></a>
<a id="schema_PrepareResourceFromEntityRequest"></a>
<a id="tocSprepareresourcefromentityrequest"></a>
<a id="tocsprepareresourcefromentityrequest"></a>

```json
{
  "token": "string",
  "id": "string",
  "catalog": "string",
  "entity": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|token|string|false|none|none|
|id|string|false|none|none|
|catalog|string|false|none|none|
|entity|string|false|none|none|

<h2 id="tocS_PrepareResourceFromEntityResponse">PrepareResourceFromEntityResponse</h2>
<!-- backwards compatibility -->
<a id="schemaprepareresourcefromentityresponse"></a>
<a id="schema_PrepareResourceFromEntityResponse"></a>
<a id="tocSprepareresourcefromentityresponse"></a>
<a id="tocsprepareresourcefromentityresponse"></a>

```json
{
  "resource": {
    "id": "string",
    "name": "string",
    "namespace": "string",
    "sourceConfig": {
      "dataSource": "string",
      "catalog": "string",
      "entity": "string"
    },
    "properties": [
      {
        "id": "string",
        "name": "string",
        "type": "BOOL",
        "mapping": "string",
        "required": true,
        "primary": true,
        "length": 0,
        "unique": true,
        "immutable": true,
        "securityContext": {
          "constraints": [
            {
              "namespace": "string",
              "resource": "string",
              "property": "string",
              "before": "2019-08-24T14:15:22Z",
              "after": "2019-08-24T14:15:22Z",
              "principal": "string",
              "recordIds": [
                "string"
              ],
              "operation": "OPERATION_TYPE_READ",
              "permit": "PERMIT_TYPE_ALLOW"
            }
          ]
        },
        "defaultValue": null,
        "exampleValue": null,
        "enumValues": [
          null
        ],
        "reference": {
          "referencedResource": "string",
          "cascade": true
        },
        "properties": [
          {}
        ],
        "Item": {},
        "title": "string",
        "description": "string",
        "annotations": {
          "property1": "string",
          "property2": "string"
        }
      }
    ],
    "indexes": [
      {
        "properties": [
          {
            "name": "string",
            "order": "ORDER_UNKNOWN"
          }
        ],
        "indexType": "BTREE",
        "unique": true,
        "annotations": {
          "property1": "string",
          "property2": "string"
        }
      }
    ],
    "securityContext": {
      "constraints": [
        {
          "namespace": "string",
          "resource": "string",
          "property": "string",
          "before": "2019-08-24T14:15:22Z",
          "after": "2019-08-24T14:15:22Z",
          "principal": "string",
          "recordIds": [
            "string"
          ],
          "operation": "OPERATION_TYPE_READ",
          "permit": "PERMIT_TYPE_ALLOW"
        }
      ]
    },
    "virtual": true,
    "immutable": true,
    "abstract": true,
    "title": "string",
    "description": "string",
    "auditData": {
      "createdOn": "2019-08-24T14:15:22Z",
      "updatedOn": "2019-08-24T14:15:22Z",
      "createdBy": "string",
      "updatedBy": "string"
    },
    "version": 0,
    "annotations": {
      "property1": "string",
      "property2": "string"
    }
  }
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|resource|[Resource](#schemaresource)|false|none|none|

<h2 id="tocS_PrepareResourceMigrationPlanRequest">PrepareResourceMigrationPlanRequest</h2>
<!-- backwards compatibility -->
<a id="schemaprepareresourcemigrationplanrequest"></a>
<a id="schema_PrepareResourceMigrationPlanRequest"></a>
<a id="tocSprepareresourcemigrationplanrequest"></a>
<a id="tocsprepareresourcemigrationplanrequest"></a>

```json
{
  "token": "string",
  "prepareFromDataSource": true,
  "resources": [
    {
      "id": "string",
      "name": "string",
      "namespace": "string",
      "sourceConfig": {
        "dataSource": "string",
        "catalog": "string",
        "entity": "string"
      },
      "properties": [
        {
          "id": "string",
          "name": "string",
          "type": "BOOL",
          "mapping": "string",
          "required": true,
          "primary": true,
          "length": 0,
          "unique": true,
          "immutable": true,
          "securityContext": {
            "constraints": [
              {
                "namespace": "string",
                "resource": "string",
                "property": "string",
                "before": "2019-08-24T14:15:22Z",
                "after": "2019-08-24T14:15:22Z",
                "principal": "string",
                "recordIds": [
                  "string"
                ],
                "operation": "OPERATION_TYPE_READ",
                "permit": "PERMIT_TYPE_ALLOW"
              }
            ]
          },
          "defaultValue": null,
          "exampleValue": null,
          "enumValues": [
            null
          ],
          "reference": {
            "referencedResource": "string",
            "cascade": true
          },
          "properties": [
            {}
          ],
          "Item": {},
          "title": "string",
          "description": "string",
          "annotations": {
            "property1": "string",
            "property2": "string"
          }
        }
      ],
      "indexes": [
        {
          "properties": [
            {
              "name": "string",
              "order": "ORDER_UNKNOWN"
            }
          ],
          "indexType": "BTREE",
          "unique": true,
          "annotations": {
            "property1": "string",
            "property2": "string"
          }
        }
      ],
      "securityContext": {
        "constraints": [
          {
            "namespace": "string",
            "resource": "string",
            "property": "string",
            "before": "2019-08-24T14:15:22Z",
            "after": "2019-08-24T14:15:22Z",
            "principal": "string",
            "recordIds": [
              "string"
            ],
            "operation": "OPERATION_TYPE_READ",
            "permit": "PERMIT_TYPE_ALLOW"
          }
        ]
      },
      "virtual": true,
      "immutable": true,
      "abstract": true,
      "title": "string",
      "description": "string",
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0,
      "annotations": {
        "property1": "string",
        "property2": "string"
      }
    }
  ],
  "annotations": {
    "property1": "string",
    "property2": "string"
  }
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|token|string|false|none|none|
|prepareFromDataSource|boolean|false|none|none|
|resources|[[Resource](#schemaresource)]|false|none|none|
|annotations|object|false|none|none|
|» **additionalProperties**|string|false|none|none|

<h2 id="tocS_PrepareResourceMigrationPlanResponse">PrepareResourceMigrationPlanResponse</h2>
<!-- backwards compatibility -->
<a id="schemaprepareresourcemigrationplanresponse"></a>
<a id="schema_PrepareResourceMigrationPlanResponse"></a>
<a id="tocSprepareresourcemigrationplanresponse"></a>
<a id="tocsprepareresourcemigrationplanresponse"></a>

```json
{
  "plans": [
    {
      "existingResource": {
        "id": "string",
        "name": "string",
        "namespace": "string",
        "sourceConfig": {
          "dataSource": "string",
          "catalog": "string",
          "entity": "string"
        },
        "properties": [
          {
            "id": "string",
            "name": "string",
            "type": "BOOL",
            "mapping": "string",
            "required": true,
            "primary": true,
            "length": 0,
            "unique": true,
            "immutable": true,
            "securityContext": {
              "constraints": [
                {
                  "namespace": "string",
                  "resource": "string",
                  "property": "string",
                  "before": "2019-08-24T14:15:22Z",
                  "after": "2019-08-24T14:15:22Z",
                  "principal": "string",
                  "recordIds": [
                    "string"
                  ],
                  "operation": "OPERATION_TYPE_READ",
                  "permit": "PERMIT_TYPE_ALLOW"
                }
              ]
            },
            "defaultValue": null,
            "exampleValue": null,
            "enumValues": [
              null
            ],
            "reference": {
              "referencedResource": "string",
              "cascade": true
            },
            "properties": [
              {}
            ],
            "Item": {},
            "title": "string",
            "description": "string",
            "annotations": {
              "property1": "string",
              "property2": "string"
            }
          }
        ],
        "indexes": [
          {
            "properties": [
              {
                "name": "string",
                "order": "ORDER_UNKNOWN"
              }
            ],
            "indexType": "BTREE",
            "unique": true,
            "annotations": {
              "property1": "string",
              "property2": "string"
            }
          }
        ],
        "securityContext": {
          "constraints": [
            {
              "namespace": "string",
              "resource": "string",
              "property": "string",
              "before": "2019-08-24T14:15:22Z",
              "after": "2019-08-24T14:15:22Z",
              "principal": "string",
              "recordIds": [
                "string"
              ],
              "operation": "OPERATION_TYPE_READ",
              "permit": "PERMIT_TYPE_ALLOW"
            }
          ]
        },
        "virtual": true,
        "immutable": true,
        "abstract": true,
        "title": "string",
        "description": "string",
        "auditData": {
          "createdOn": "2019-08-24T14:15:22Z",
          "updatedOn": "2019-08-24T14:15:22Z",
          "createdBy": "string",
          "updatedBy": "string"
        },
        "version": 0,
        "annotations": {
          "property1": "string",
          "property2": "string"
        }
      },
      "currentResource": {
        "id": "string",
        "name": "string",
        "namespace": "string",
        "sourceConfig": {
          "dataSource": "string",
          "catalog": "string",
          "entity": "string"
        },
        "properties": [
          {
            "id": "string",
            "name": "string",
            "type": "BOOL",
            "mapping": "string",
            "required": true,
            "primary": true,
            "length": 0,
            "unique": true,
            "immutable": true,
            "securityContext": {
              "constraints": [
                {
                  "namespace": "string",
                  "resource": "string",
                  "property": "string",
                  "before": "2019-08-24T14:15:22Z",
                  "after": "2019-08-24T14:15:22Z",
                  "principal": "string",
                  "recordIds": [
                    "string"
                  ],
                  "operation": "OPERATION_TYPE_READ",
                  "permit": "PERMIT_TYPE_ALLOW"
                }
              ]
            },
            "defaultValue": null,
            "exampleValue": null,
            "enumValues": [
              null
            ],
            "reference": {
              "referencedResource": "string",
              "cascade": true
            },
            "properties": [
              {}
            ],
            "Item": {},
            "title": "string",
            "description": "string",
            "annotations": {
              "property1": "string",
              "property2": "string"
            }
          }
        ],
        "indexes": [
          {
            "properties": [
              {
                "name": "string",
                "order": "ORDER_UNKNOWN"
              }
            ],
            "indexType": "BTREE",
            "unique": true,
            "annotations": {
              "property1": "string",
              "property2": "string"
            }
          }
        ],
        "securityContext": {
          "constraints": [
            {
              "namespace": "string",
              "resource": "string",
              "property": "string",
              "before": "2019-08-24T14:15:22Z",
              "after": "2019-08-24T14:15:22Z",
              "principal": "string",
              "recordIds": [
                "string"
              ],
              "operation": "OPERATION_TYPE_READ",
              "permit": "PERMIT_TYPE_ALLOW"
            }
          ]
        },
        "virtual": true,
        "immutable": true,
        "abstract": true,
        "title": "string",
        "description": "string",
        "auditData": {
          "createdOn": "2019-08-24T14:15:22Z",
          "updatedOn": "2019-08-24T14:15:22Z",
          "createdBy": "string",
          "updatedBy": "string"
        },
        "version": 0,
        "annotations": {
          "property1": "string",
          "property2": "string"
        }
      },
      "steps": [
        {
          "createResource": {},
          "deleteResource": {},
          "updateResource": {
            "changedFields": [
              "string"
            ]
          },
          "createProperty": {
            "property": "string"
          },
          "deleteProperty": {
            "existingProperty": "string"
          },
          "updateProperty": {
            "existingProperty": "string",
            "property": "string",
            "changedFields": [
              "string"
            ]
          },
          "createIndex": {
            "index": 0
          },
          "deleteIndex": {
            "existingIndex": 0
          }
        }
      ]
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|plans|[[ResourceMigrationPlan](#schemaresourcemigrationplan)]|false|none|none|

<h2 id="tocS_ReadStreamRequest">ReadStreamRequest</h2>
<!-- backwards compatibility -->
<a id="schemareadstreamrequest"></a>
<a id="schema_ReadStreamRequest"></a>
<a id="tocSreadstreamrequest"></a>
<a id="tocsreadstreamrequest"></a>

```json
{
  "token": "string",
  "namespace": "string",
  "resource": "string",
  "query": {
    "and": {
      "expressions": [
        {}
      ]
    },
    "or": {
      "expressions": [
        {}
      ]
    },
    "not": {},
    "equal": {
      "left": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "right": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    },
    "lessThan": {
      "left": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "right": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    },
    "greaterThan": {
      "left": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "right": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    },
    "lessThanOrEqual": {
      "left": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "right": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    },
    "greaterThanOrEqual": {
      "left": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "right": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    },
    "in": {
      "left": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "right": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    },
    "isNull": {
      "additionalProperties": [
        {
          "name": "string",
          "value": {
            "value": {
              "@type": "string"
            },
            "yaml": "string"
          }
        }
      ]
    },
    "regexMatch": {
      "pattern": "string",
      "expression": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    }
  },
  "limit": 0,
  "offset": 0,
  "useHistory": true,
  "resolveReferences": [
    "string"
  ],
  "useTransaction": true,
  "packRecords": true,
  "annotations": {
    "property1": "string",
    "property2": "string"
  }
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|token|string|false|none|none|
|namespace|string|false|none|none|
|resource|string|false|none|none|
|query|[BooleanExpression](#schemabooleanexpression)|false|none|none|
|limit|integer(uint32)|false|none|none|
|offset|integer(uint64)|false|none|none|
|useHistory|boolean|false|none|none|
|resolveReferences|[string]|false|none|none|
|useTransaction|boolean|false|none|none|
|packRecords|boolean|false|none|none|
|annotations|object|false|none|none|
|» **additionalProperties**|string|false|none|none|

<h2 id="tocS_Record">Record</h2>
<!-- backwards compatibility -->
<a id="schemarecord"></a>
<a id="schema_Record"></a>
<a id="tocSrecord"></a>
<a id="tocsrecord"></a>

```json
{
  "id": "string",
  "properties": {
    "property1": null,
    "property2": null
  },
  "propertiesPacked": [
    null
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|id|string|false|none|none|
|properties|object|false|none|none|
|» **additionalProperties**|[GoogleProtobufValue](#schemagoogleprotobufvalue)|false|none|Represents a dynamically typed value which can be either null, a number, a string, a boolean, a recursive struct value, or a list of values.|
|propertiesPacked|[[GoogleProtobufValue](#schemagoogleprotobufvalue)]|false|none|[Represents a dynamically typed value which can be either null, a number, a string, a boolean, a recursive struct value, or a list of values.]|

<h2 id="tocS_RedisParams">RedisParams</h2>
<!-- backwards compatibility -->
<a id="schemaredisparams"></a>
<a id="schema_RedisParams"></a>
<a id="tocSredisparams"></a>
<a id="tocsredisparams"></a>

```json
{
  "addr": "string",
  "password": "string",
  "db": 0
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|addr|string|false|none|none|
|password|string|false|none|none|
|db|integer(int32)|false|none|none|

<h2 id="tocS_Reference">Reference</h2>
<!-- backwards compatibility -->
<a id="schemareference"></a>
<a id="schema_Reference"></a>
<a id="tocSreference"></a>
<a id="tocsreference"></a>

```json
{
  "referencedResource": "string",
  "cascade": true
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|referencedResource|string|false|none|none|
|cascade|boolean|false|none|none|

<h2 id="tocS_RegexMatchExpression">RegexMatchExpression</h2>
<!-- backwards compatibility -->
<a id="schemaregexmatchexpression"></a>
<a id="schema_RegexMatchExpression"></a>
<a id="tocSregexmatchexpression"></a>
<a id="tocsregexmatchexpression"></a>

```json
{
  "pattern": "string",
  "expression": {
    "additionalProperties": [
      {
        "name": "string",
        "value": {
          "value": {
            "@type": "string"
          },
          "yaml": "string"
        }
      }
    ]
  }
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|pattern|string|false|none|none|
|expression|[Expression](#schemaexpression)|false|none|none|

<h2 id="tocS_RenewTokenRequest">RenewTokenRequest</h2>
<!-- backwards compatibility -->
<a id="schemarenewtokenrequest"></a>
<a id="schema_RenewTokenRequest"></a>
<a id="tocSrenewtokenrequest"></a>
<a id="tocsrenewtokenrequest"></a>

```json
{
  "token": "string",
  "term": "SHORT"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|token|string|false|none|none|
|term|string(enum)|false|none|none|

#### Enumerated Values

|Property|Value|
|---|---|
|term|SHORT|
|term|MIDDLE|
|term|LONG|
|term|VERY_LONG|

<h2 id="tocS_RenewTokenResponse">RenewTokenResponse</h2>
<!-- backwards compatibility -->
<a id="schemarenewtokenresponse"></a>
<a id="schema_RenewTokenResponse"></a>
<a id="tocSrenewtokenresponse"></a>
<a id="tocsrenewtokenresponse"></a>

```json
{
  "token": {
    "term": "SHORT",
    "content": "string",
    "expiration": "2019-08-24T14:15:22Z"
  }
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|token|[Token](#schematoken)|false|none|none|

<h2 id="tocS_Resource">Resource</h2>
<!-- backwards compatibility -->
<a id="schemaresource"></a>
<a id="schema_Resource"></a>
<a id="tocSresource"></a>
<a id="tocsresource"></a>

```json
{
  "id": "string",
  "name": "string",
  "namespace": "string",
  "sourceConfig": {
    "dataSource": "string",
    "catalog": "string",
    "entity": "string"
  },
  "properties": [
    {
      "id": "string",
      "name": "string",
      "type": "BOOL",
      "mapping": "string",
      "required": true,
      "primary": true,
      "length": 0,
      "unique": true,
      "immutable": true,
      "securityContext": {
        "constraints": [
          {
            "namespace": "string",
            "resource": "string",
            "property": "string",
            "before": "2019-08-24T14:15:22Z",
            "after": "2019-08-24T14:15:22Z",
            "principal": "string",
            "recordIds": [
              "string"
            ],
            "operation": "OPERATION_TYPE_READ",
            "permit": "PERMIT_TYPE_ALLOW"
          }
        ]
      },
      "defaultValue": null,
      "exampleValue": null,
      "enumValues": [
        null
      ],
      "reference": {
        "referencedResource": "string",
        "cascade": true
      },
      "properties": [
        {}
      ],
      "Item": {},
      "title": "string",
      "description": "string",
      "annotations": {
        "property1": "string",
        "property2": "string"
      }
    }
  ],
  "indexes": [
    {
      "properties": [
        {
          "name": "string",
          "order": "ORDER_UNKNOWN"
        }
      ],
      "indexType": "BTREE",
      "unique": true,
      "annotations": {
        "property1": "string",
        "property2": "string"
      }
    }
  ],
  "securityContext": {
    "constraints": [
      {
        "namespace": "string",
        "resource": "string",
        "property": "string",
        "before": "2019-08-24T14:15:22Z",
        "after": "2019-08-24T14:15:22Z",
        "principal": "string",
        "recordIds": [
          "string"
        ],
        "operation": "OPERATION_TYPE_READ",
        "permit": "PERMIT_TYPE_ALLOW"
      }
    ]
  },
  "virtual": true,
  "immutable": true,
  "abstract": true,
  "title": "string",
  "description": "string",
  "auditData": {
    "createdOn": "2019-08-24T14:15:22Z",
    "updatedOn": "2019-08-24T14:15:22Z",
    "createdBy": "string",
    "updatedBy": "string"
  },
  "version": 0,
  "annotations": {
    "property1": "string",
    "property2": "string"
  }
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|id|string|false|none|none|
|name|string|false|none|none|
|namespace|string|false|none|none|
|sourceConfig|[ResourceSourceConfig](#schemaresourcesourceconfig)|false|none|source config is to configure resource and bind it to data-source and an entity inside data source. An entity is like a table on sql databases or collection on mongodb etc.|
|properties|[[ResourceProperty](#schemaresourceproperty)]|false|none|[Resource properties is used to describe its schema. Each resource property is corresponding to a field in a record API Brew is responsible to validate data according to property types. For example, when you call create record andif you send 123.45 for int64]|
|indexes|[[ResourceIndex](#schemaresourceindex)]|false|none|none|
|securityContext|[SecurityContext](#schemasecuritycontext)|false|none|none|
|virtual|boolean|false|none|none|
|immutable|boolean|false|none|none|
|abstract|boolean|false|none|none|
|title|string|false|none|none|
|description|string|false|none|none|
|auditData|[AuditData](#schemaauditdata)|false|none|none|
|version|integer(uint32)|false|none|none|
|annotations|object|false|none|none|
|» **additionalProperties**|string|false|none|none|

<h2 id="tocS_ResourceIndex">ResourceIndex</h2>
<!-- backwards compatibility -->
<a id="schemaresourceindex"></a>
<a id="schema_ResourceIndex"></a>
<a id="tocSresourceindex"></a>
<a id="tocsresourceindex"></a>

```json
{
  "properties": [
    {
      "name": "string",
      "order": "ORDER_UNKNOWN"
    }
  ],
  "indexType": "BTREE",
  "unique": true,
  "annotations": {
    "property1": "string",
    "property2": "string"
  }
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|properties|[[ResourceIndexProperty](#schemaresourceindexproperty)]|false|none|none|
|indexType|string(enum)|false|none|none|
|unique|boolean|false|none|none|
|annotations|object|false|none|none|
|» **additionalProperties**|string|false|none|none|

#### Enumerated Values

|Property|Value|
|---|---|
|indexType|BTREE|
|indexType|HASH|

<h2 id="tocS_ResourceIndexProperty">ResourceIndexProperty</h2>
<!-- backwards compatibility -->
<a id="schemaresourceindexproperty"></a>
<a id="schema_ResourceIndexProperty"></a>
<a id="tocSresourceindexproperty"></a>
<a id="tocsresourceindexproperty"></a>

```json
{
  "name": "string",
  "order": "ORDER_UNKNOWN"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|name|string|false|none|none|
|order|string(enum)|false|none|none|

#### Enumerated Values

|Property|Value|
|---|---|
|order|ORDER_UNKNOWN|
|order|ORDER_ASC|
|order|ORDER_DESC|

<h2 id="tocS_ResourceMigrationCreateIndex">ResourceMigrationCreateIndex</h2>
<!-- backwards compatibility -->
<a id="schemaresourcemigrationcreateindex"></a>
<a id="schema_ResourceMigrationCreateIndex"></a>
<a id="tocSresourcemigrationcreateindex"></a>
<a id="tocsresourcemigrationcreateindex"></a>

```json
{
  "index": 0
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|index|integer(uint32)|false|none|none|

<h2 id="tocS_ResourceMigrationCreateProperty">ResourceMigrationCreateProperty</h2>
<!-- backwards compatibility -->
<a id="schemaresourcemigrationcreateproperty"></a>
<a id="schema_ResourceMigrationCreateProperty"></a>
<a id="tocSresourcemigrationcreateproperty"></a>
<a id="tocsresourcemigrationcreateproperty"></a>

```json
{
  "property": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|property|string|false|none|none|

<h2 id="tocS_ResourceMigrationCreateResource">ResourceMigrationCreateResource</h2>
<!-- backwards compatibility -->
<a id="schemaresourcemigrationcreateresource"></a>
<a id="schema_ResourceMigrationCreateResource"></a>
<a id="tocSresourcemigrationcreateresource"></a>
<a id="tocsresourcemigrationcreateresource"></a>

```json
{}

```

### Properties

*None*

<h2 id="tocS_ResourceMigrationDeleteIndex">ResourceMigrationDeleteIndex</h2>
<!-- backwards compatibility -->
<a id="schemaresourcemigrationdeleteindex"></a>
<a id="schema_ResourceMigrationDeleteIndex"></a>
<a id="tocSresourcemigrationdeleteindex"></a>
<a id="tocsresourcemigrationdeleteindex"></a>

```json
{
  "existingIndex": 0
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|existingIndex|integer(uint32)|false|none|none|

<h2 id="tocS_ResourceMigrationDeleteProperty">ResourceMigrationDeleteProperty</h2>
<!-- backwards compatibility -->
<a id="schemaresourcemigrationdeleteproperty"></a>
<a id="schema_ResourceMigrationDeleteProperty"></a>
<a id="tocSresourcemigrationdeleteproperty"></a>
<a id="tocsresourcemigrationdeleteproperty"></a>

```json
{
  "existingProperty": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|existingProperty|string|false|none|none|

<h2 id="tocS_ResourceMigrationDeleteResource">ResourceMigrationDeleteResource</h2>
<!-- backwards compatibility -->
<a id="schemaresourcemigrationdeleteresource"></a>
<a id="schema_ResourceMigrationDeleteResource"></a>
<a id="tocSresourcemigrationdeleteresource"></a>
<a id="tocsresourcemigrationdeleteresource"></a>

```json
{}

```

### Properties

*None*

<h2 id="tocS_ResourceMigrationPlan">ResourceMigrationPlan</h2>
<!-- backwards compatibility -->
<a id="schemaresourcemigrationplan"></a>
<a id="schema_ResourceMigrationPlan"></a>
<a id="tocSresourcemigrationplan"></a>
<a id="tocsresourcemigrationplan"></a>

```json
{
  "existingResource": {
    "id": "string",
    "name": "string",
    "namespace": "string",
    "sourceConfig": {
      "dataSource": "string",
      "catalog": "string",
      "entity": "string"
    },
    "properties": [
      {
        "id": "string",
        "name": "string",
        "type": "BOOL",
        "mapping": "string",
        "required": true,
        "primary": true,
        "length": 0,
        "unique": true,
        "immutable": true,
        "securityContext": {
          "constraints": [
            {
              "namespace": "string",
              "resource": "string",
              "property": "string",
              "before": "2019-08-24T14:15:22Z",
              "after": "2019-08-24T14:15:22Z",
              "principal": "string",
              "recordIds": [
                "string"
              ],
              "operation": "OPERATION_TYPE_READ",
              "permit": "PERMIT_TYPE_ALLOW"
            }
          ]
        },
        "defaultValue": null,
        "exampleValue": null,
        "enumValues": [
          null
        ],
        "reference": {
          "referencedResource": "string",
          "cascade": true
        },
        "properties": [
          {}
        ],
        "Item": {},
        "title": "string",
        "description": "string",
        "annotations": {
          "property1": "string",
          "property2": "string"
        }
      }
    ],
    "indexes": [
      {
        "properties": [
          {
            "name": "string",
            "order": "ORDER_UNKNOWN"
          }
        ],
        "indexType": "BTREE",
        "unique": true,
        "annotations": {
          "property1": "string",
          "property2": "string"
        }
      }
    ],
    "securityContext": {
      "constraints": [
        {
          "namespace": "string",
          "resource": "string",
          "property": "string",
          "before": "2019-08-24T14:15:22Z",
          "after": "2019-08-24T14:15:22Z",
          "principal": "string",
          "recordIds": [
            "string"
          ],
          "operation": "OPERATION_TYPE_READ",
          "permit": "PERMIT_TYPE_ALLOW"
        }
      ]
    },
    "virtual": true,
    "immutable": true,
    "abstract": true,
    "title": "string",
    "description": "string",
    "auditData": {
      "createdOn": "2019-08-24T14:15:22Z",
      "updatedOn": "2019-08-24T14:15:22Z",
      "createdBy": "string",
      "updatedBy": "string"
    },
    "version": 0,
    "annotations": {
      "property1": "string",
      "property2": "string"
    }
  },
  "currentResource": {
    "id": "string",
    "name": "string",
    "namespace": "string",
    "sourceConfig": {
      "dataSource": "string",
      "catalog": "string",
      "entity": "string"
    },
    "properties": [
      {
        "id": "string",
        "name": "string",
        "type": "BOOL",
        "mapping": "string",
        "required": true,
        "primary": true,
        "length": 0,
        "unique": true,
        "immutable": true,
        "securityContext": {
          "constraints": [
            {
              "namespace": "string",
              "resource": "string",
              "property": "string",
              "before": "2019-08-24T14:15:22Z",
              "after": "2019-08-24T14:15:22Z",
              "principal": "string",
              "recordIds": [
                "string"
              ],
              "operation": "OPERATION_TYPE_READ",
              "permit": "PERMIT_TYPE_ALLOW"
            }
          ]
        },
        "defaultValue": null,
        "exampleValue": null,
        "enumValues": [
          null
        ],
        "reference": {
          "referencedResource": "string",
          "cascade": true
        },
        "properties": [
          {}
        ],
        "Item": {},
        "title": "string",
        "description": "string",
        "annotations": {
          "property1": "string",
          "property2": "string"
        }
      }
    ],
    "indexes": [
      {
        "properties": [
          {
            "name": "string",
            "order": "ORDER_UNKNOWN"
          }
        ],
        "indexType": "BTREE",
        "unique": true,
        "annotations": {
          "property1": "string",
          "property2": "string"
        }
      }
    ],
    "securityContext": {
      "constraints": [
        {
          "namespace": "string",
          "resource": "string",
          "property": "string",
          "before": "2019-08-24T14:15:22Z",
          "after": "2019-08-24T14:15:22Z",
          "principal": "string",
          "recordIds": [
            "string"
          ],
          "operation": "OPERATION_TYPE_READ",
          "permit": "PERMIT_TYPE_ALLOW"
        }
      ]
    },
    "virtual": true,
    "immutable": true,
    "abstract": true,
    "title": "string",
    "description": "string",
    "auditData": {
      "createdOn": "2019-08-24T14:15:22Z",
      "updatedOn": "2019-08-24T14:15:22Z",
      "createdBy": "string",
      "updatedBy": "string"
    },
    "version": 0,
    "annotations": {
      "property1": "string",
      "property2": "string"
    }
  },
  "steps": [
    {
      "createResource": {},
      "deleteResource": {},
      "updateResource": {
        "changedFields": [
          "string"
        ]
      },
      "createProperty": {
        "property": "string"
      },
      "deleteProperty": {
        "existingProperty": "string"
      },
      "updateProperty": {
        "existingProperty": "string",
        "property": "string",
        "changedFields": [
          "string"
        ]
      },
      "createIndex": {
        "index": 0
      },
      "deleteIndex": {
        "existingIndex": 0
      }
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|existingResource|[Resource](#schemaresource)|false|none|none|
|currentResource|[Resource](#schemaresource)|false|none|none|
|steps|[[ResourceMigrationStep](#schemaresourcemigrationstep)]|false|none|none|

<h2 id="tocS_ResourceMigrationStep">ResourceMigrationStep</h2>
<!-- backwards compatibility -->
<a id="schemaresourcemigrationstep"></a>
<a id="schema_ResourceMigrationStep"></a>
<a id="tocSresourcemigrationstep"></a>
<a id="tocsresourcemigrationstep"></a>

```json
{
  "createResource": {},
  "deleteResource": {},
  "updateResource": {
    "changedFields": [
      "string"
    ]
  },
  "createProperty": {
    "property": "string"
  },
  "deleteProperty": {
    "existingProperty": "string"
  },
  "updateProperty": {
    "existingProperty": "string",
    "property": "string",
    "changedFields": [
      "string"
    ]
  },
  "createIndex": {
    "index": 0
  },
  "deleteIndex": {
    "existingIndex": 0
  }
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|createResource|[ResourceMigrationCreateResource](#schemaresourcemigrationcreateresource)|false|none|none|
|deleteResource|[ResourceMigrationDeleteResource](#schemaresourcemigrationdeleteresource)|false|none|none|
|updateResource|[ResourceMigrationUpdateResource](#schemaresourcemigrationupdateresource)|false|none|none|
|createProperty|[ResourceMigrationCreateProperty](#schemaresourcemigrationcreateproperty)|false|none|none|
|deleteProperty|[ResourceMigrationDeleteProperty](#schemaresourcemigrationdeleteproperty)|false|none|none|
|updateProperty|[ResourceMigrationUpdateProperty](#schemaresourcemigrationupdateproperty)|false|none|none|
|createIndex|[ResourceMigrationCreateIndex](#schemaresourcemigrationcreateindex)|false|none|none|
|deleteIndex|[ResourceMigrationDeleteIndex](#schemaresourcemigrationdeleteindex)|false|none|none|

<h2 id="tocS_ResourceMigrationUpdateProperty">ResourceMigrationUpdateProperty</h2>
<!-- backwards compatibility -->
<a id="schemaresourcemigrationupdateproperty"></a>
<a id="schema_ResourceMigrationUpdateProperty"></a>
<a id="tocSresourcemigrationupdateproperty"></a>
<a id="tocsresourcemigrationupdateproperty"></a>

```json
{
  "existingProperty": "string",
  "property": "string",
  "changedFields": [
    "string"
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|existingProperty|string|false|none|none|
|property|string|false|none|none|
|changedFields|[string]|false|none|none|

<h2 id="tocS_ResourceMigrationUpdateResource">ResourceMigrationUpdateResource</h2>
<!-- backwards compatibility -->
<a id="schemaresourcemigrationupdateresource"></a>
<a id="schema_ResourceMigrationUpdateResource"></a>
<a id="tocSresourcemigrationupdateresource"></a>
<a id="tocsresourcemigrationupdateresource"></a>

```json
{
  "changedFields": [
    "string"
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|changedFields|[string]|false|none|none|

<h2 id="tocS_ResourceProperty">ResourceProperty</h2>
<!-- backwards compatibility -->
<a id="schemaresourceproperty"></a>
<a id="schema_ResourceProperty"></a>
<a id="tocSresourceproperty"></a>
<a id="tocsresourceproperty"></a>

```json
{
  "id": "string",
  "name": "string",
  "type": "BOOL",
  "mapping": "string",
  "required": true,
  "primary": true,
  "length": 0,
  "unique": true,
  "immutable": true,
  "securityContext": {
    "constraints": [
      {
        "namespace": "string",
        "resource": "string",
        "property": "string",
        "before": "2019-08-24T14:15:22Z",
        "after": "2019-08-24T14:15:22Z",
        "principal": "string",
        "recordIds": [
          "string"
        ],
        "operation": "OPERATION_TYPE_READ",
        "permit": "PERMIT_TYPE_ALLOW"
      }
    ]
  },
  "defaultValue": null,
  "exampleValue": null,
  "enumValues": [
    null
  ],
  "reference": {
    "referencedResource": "string",
    "cascade": true
  },
  "properties": [
    {
      "id": "string",
      "name": "string",
      "type": "BOOL",
      "mapping": "string",
      "required": true,
      "primary": true,
      "length": 0,
      "unique": true,
      "immutable": true,
      "securityContext": {
        "constraints": [
          {
            "namespace": "string",
            "resource": "string",
            "property": "string",
            "before": "2019-08-24T14:15:22Z",
            "after": "2019-08-24T14:15:22Z",
            "principal": "string",
            "recordIds": [
              "string"
            ],
            "operation": "OPERATION_TYPE_READ",
            "permit": "PERMIT_TYPE_ALLOW"
          }
        ]
      },
      "defaultValue": null,
      "exampleValue": null,
      "enumValues": [
        null
      ],
      "reference": {
        "referencedResource": "string",
        "cascade": true
      },
      "properties": [],
      "Item": {},
      "title": "string",
      "description": "string",
      "annotations": {
        "property1": "string",
        "property2": "string"
      }
    }
  ],
  "Item": {
    "id": "string",
    "name": "string",
    "type": "BOOL",
    "mapping": "string",
    "required": true,
    "primary": true,
    "length": 0,
    "unique": true,
    "immutable": true,
    "securityContext": {
      "constraints": [
        {
          "namespace": "string",
          "resource": "string",
          "property": "string",
          "before": "2019-08-24T14:15:22Z",
          "after": "2019-08-24T14:15:22Z",
          "principal": "string",
          "recordIds": [
            "string"
          ],
          "operation": "OPERATION_TYPE_READ",
          "permit": "PERMIT_TYPE_ALLOW"
        }
      ]
    },
    "defaultValue": null,
    "exampleValue": null,
    "enumValues": [
      null
    ],
    "reference": {
      "referencedResource": "string",
      "cascade": true
    },
    "properties": [
      {}
    ],
    "Item": {},
    "title": "string",
    "description": "string",
    "annotations": {
      "property1": "string",
      "property2": "string"
    }
  },
  "title": "string",
  "description": "string",
  "annotations": {
    "property1": "string",
    "property2": "string"
  }
}

```

Resource properties is used to describe its schema. Each resource property is corresponding to a field in a record API Brew is responsible to validate data according to property types. For example, when you call create record andif you send 123.45 for int64

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|id|string|false|none|none|
|name|string|false|none|none|
|type|string(enum)|false|none|none|
|mapping|string|false|none|none|
|required|boolean|false|none|none|
|primary|boolean|false|none|none|
|length|integer(uint32)|false|none|none|
|unique|boolean|false|none|none|
|immutable|boolean|false|none|none|
|securityContext|[SecurityContext](#schemasecuritycontext)|false|none|none|
|defaultValue|[GoogleProtobufValue](#schemagoogleprotobufvalue)|false|none|Represents a dynamically typed value which can be either null, a number, a string, a boolean, a recursive struct value, or a list of values.|
|exampleValue|[GoogleProtobufValue](#schemagoogleprotobufvalue)|false|none|Represents a dynamically typed value which can be either null, a number, a string, a boolean, a recursive struct value, or a list of values.|
|enumValues|[[GoogleProtobufValue](#schemagoogleprotobufvalue)]|false|none|[Represents a dynamically typed value which can be either null, a number, a string, a boolean, a recursive struct value, or a list of values.]|
|reference|[Reference](#schemareference)|false|none|none|
|properties|[[ResourceProperty](#schemaresourceproperty)]|false|none|[Resource properties is used to describe its schema. Each resource property is corresponding to a field in a record API Brew is responsible to validate data according to property types. For example, when you call create record andif you send 123.45 for int64]|
|Item|[ResourceProperty](#schemaresourceproperty)|false|none|Resource properties is used to describe its schema. Each resource property is corresponding to a field in a record API Brew is responsible to validate data according to property types. For example, when you call create record andif you send 123.45 for int64|
|title|string|false|none|none|
|description|string|false|none|none|
|annotations|object|false|none|none|
|» **additionalProperties**|string|false|none|none|

#### Enumerated Values

|Property|Value|
|---|---|
|type|BOOL|
|type|STRING|
|type|FLOAT32|
|type|FLOAT64|
|type|INT32|
|type|INT64|
|type|BYTES|
|type|UUID|
|type|DATE|
|type|TIME|
|type|TIMESTAMP|
|type|OBJECT|
|type|MAP|
|type|LIST|
|type|REFERENCE|
|type|ENUM|
|type|STRUCT|

<h2 id="tocS_ResourceSourceConfig">ResourceSourceConfig</h2>
<!-- backwards compatibility -->
<a id="schemaresourcesourceconfig"></a>
<a id="schema_ResourceSourceConfig"></a>
<a id="tocSresourcesourceconfig"></a>
<a id="tocsresourcesourceconfig"></a>

```json
{
  "dataSource": "string",
  "catalog": "string",
  "entity": "string"
}

```

source config is to configure resource and bind it to data-source and an entity inside data source. An entity is like a table on sql databases or collection on mongodb etc.

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|dataSource|string|false|none|none|
|catalog|string|false|none|none|
|entity|string|false|none|none|

<h2 id="tocS_SearchRecordRequest">SearchRecordRequest</h2>
<!-- backwards compatibility -->
<a id="schemasearchrecordrequest"></a>
<a id="schema_SearchRecordRequest"></a>
<a id="tocSsearchrecordrequest"></a>
<a id="tocssearchrecordrequest"></a>

```json
{
  "token": "string",
  "namespace": "string",
  "resource": "string",
  "query": {
    "and": {
      "expressions": [
        {}
      ]
    },
    "or": {
      "expressions": [
        {}
      ]
    },
    "not": {},
    "equal": {
      "left": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "right": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    },
    "lessThan": {
      "left": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "right": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    },
    "greaterThan": {
      "left": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "right": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    },
    "lessThanOrEqual": {
      "left": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "right": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    },
    "greaterThanOrEqual": {
      "left": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "right": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    },
    "in": {
      "left": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      },
      "right": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    },
    "isNull": {
      "additionalProperties": [
        {
          "name": "string",
          "value": {
            "value": {
              "@type": "string"
            },
            "yaml": "string"
          }
        }
      ]
    },
    "regexMatch": {
      "pattern": "string",
      "expression": {
        "additionalProperties": [
          {
            "name": "string",
            "value": {
              "value": {
                "@type": "string"
              },
              "yaml": "string"
            }
          }
        ]
      }
    }
  },
  "limit": 0,
  "offset": 0,
  "useHistory": true,
  "resolveReferences": [
    "string"
  ],
  "annotations": {
    "property1": "string",
    "property2": "string"
  }
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|token|string|false|none|none|
|namespace|string|false|none|none|
|resource|string|false|none|none|
|query|[BooleanExpression](#schemabooleanexpression)|false|none|none|
|limit|integer(uint32)|false|none|none|
|offset|integer(uint64)|false|none|none|
|useHistory|boolean|false|none|none|
|resolveReferences|[string]|false|none|none|
|annotations|object|false|none|none|
|» **additionalProperties**|string|false|none|none|

<h2 id="tocS_SearchRecordResponse">SearchRecordResponse</h2>
<!-- backwards compatibility -->
<a id="schemasearchrecordresponse"></a>
<a id="schema_SearchRecordResponse"></a>
<a id="tocSsearchrecordresponse"></a>
<a id="tocssearchrecordresponse"></a>

```json
{
  "total": 0,
  "content": [
    {
      "id": "string",
      "properties": {
        "property1": null,
        "property2": null
      },
      "propertiesPacked": [
        null
      ]
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|total|integer(uint32)|false|none|none|
|content|[[Record](#schemarecord)]|false|none|none|

<h2 id="tocS_SecurityConstraint">SecurityConstraint</h2>
<!-- backwards compatibility -->
<a id="schemasecurityconstraint"></a>
<a id="schema_SecurityConstraint"></a>
<a id="tocSsecurityconstraint"></a>
<a id="tocssecurityconstraint"></a>

```json
{
  "namespace": "string",
  "resource": "string",
  "property": "string",
  "before": "2019-08-24T14:15:22Z",
  "after": "2019-08-24T14:15:22Z",
  "principal": "string",
  "recordIds": [
    "string"
  ],
  "operation": "OPERATION_TYPE_READ",
  "permit": "PERMIT_TYPE_ALLOW"
}

```

SecurityConstraint is a rule

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|namespace|string|false|none|none|
|resource|string|false|none|none|
|property|string|false|none|none|
|before|string(date-time)|false|none|none|
|after|string(date-time)|false|none|none|
|principal|string|false|none|none|
|recordIds|[string]|false|none|none|
|operation|string(enum)|false|none|none|
|permit|string(enum)|false|none|none|

#### Enumerated Values

|Property|Value|
|---|---|
|operation|OPERATION_TYPE_READ|
|operation|OPERATION_TYPE_CREATE|
|operation|OPERATION_TYPE_UPDATE|
|operation|OPERATION_TYPE_DELETE|
|operation|FULL|
|permit|PERMIT_TYPE_ALLOW|
|permit|PERMIT_TYPE_REJECT|
|permit|PERMIT_TYPE_UNKNOWN|

<h2 id="tocS_SecurityContext">SecurityContext</h2>
<!-- backwards compatibility -->
<a id="schemasecuritycontext"></a>
<a id="schema_SecurityContext"></a>
<a id="tocSsecuritycontext"></a>
<a id="tocssecuritycontext"></a>

```json
{
  "constraints": [
    {
      "namespace": "string",
      "resource": "string",
      "property": "string",
      "before": "2019-08-24T14:15:22Z",
      "after": "2019-08-24T14:15:22Z",
      "principal": "string",
      "recordIds": [
        "string"
      ],
      "operation": "OPERATION_TYPE_READ",
      "permit": "PERMIT_TYPE_ALLOW"
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|constraints|[[SecurityConstraint](#schemasecurityconstraint)]|false|none|[SecurityConstraint is a rule]|

<h2 id="tocS_Status">Status</h2>
<!-- backwards compatibility -->
<a id="schemastatus"></a>
<a id="schema_Status"></a>
<a id="tocSstatus"></a>
<a id="tocsstatus"></a>

```json
{
  "code": 0,
  "message": "string",
  "details": [
    {
      "@type": "string"
    }
  ]
}

```

The `Status` type defines a logical error model that is suitable for different programming environments, including REST APIs and RPC APIs. It is used by [gRPC](https://github.com/grpc). Each `Status` message contains three pieces of data: error code, error message, and error details. You can find out more about this error model and how to work with it in the [API Design Guide](https://cloud.google.com/apis/design/errors).

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|code|integer(int32)|false|none|The status code, which should be an enum value of [google.rpc.Code][google.rpc.Code].|
|message|string|false|none|A developer-facing error message, which should be in English. Any user-facing error message should be localized and sent in the [google.rpc.Status.details][google.rpc.Status.details] field, or localized by the client.|
|details|[[GoogleProtobufAny](#schemagoogleprotobufany)]|false|none|A list of messages that carry the error details.  There is a common set of message types for APIs to use.|

<h2 id="tocS_StatusResponse">StatusResponse</h2>
<!-- backwards compatibility -->
<a id="schemastatusresponse"></a>
<a id="schema_StatusResponse"></a>
<a id="tocSstatusresponse"></a>
<a id="tocsstatusresponse"></a>

```json
{
  "connectionAlreadyInitiated": true,
  "testConnection": true
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|connectionAlreadyInitiated|boolean|false|none|none|
|testConnection|boolean|false|none|none|

<h2 id="tocS_Token">Token</h2>
<!-- backwards compatibility -->
<a id="schematoken"></a>
<a id="schema_Token"></a>
<a id="tocStoken"></a>
<a id="tocstoken"></a>

```json
{
  "term": "SHORT",
  "content": "string",
  "expiration": "2019-08-24T14:15:22Z"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|term|string(enum)|false|none|none|
|content|string|false|none|none|
|expiration|string(date-time)|false|none|none|

#### Enumerated Values

|Property|Value|
|---|---|
|term|SHORT|
|term|MIDDLE|
|term|LONG|
|term|VERY_LONG|

<h2 id="tocS_UpdateDataSourceRequest">UpdateDataSourceRequest</h2>
<!-- backwards compatibility -->
<a id="schemaupdatedatasourcerequest"></a>
<a id="schema_UpdateDataSourceRequest"></a>
<a id="tocSupdatedatasourcerequest"></a>
<a id="tocsupdatedatasourcerequest"></a>

```json
{
  "token": "string",
  "dataSources": [
    {
      "id": "string",
      "backend": "POSTGRESQL",
      "name": "string",
      "description": "string",
      "postgresqlParams": {
        "username": "string",
        "password": "string",
        "host": "string",
        "port": 0,
        "dbName": "string",
        "defaultSchema": "string"
      },
      "mysqlParams": {
        "username": "string",
        "password": "string",
        "host": "string",
        "port": 0,
        "dbName": "string",
        "defaultSchema": "string"
      },
      "virtualParams": {
        "mode": "DISCARD"
      },
      "redisParams": {
        "addr": "string",
        "password": "string",
        "db": 0
      },
      "mongoParams": {
        "uri": "string",
        "dbName": "string"
      },
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|token|string|false|none|none|
|dataSources|[[DataSource](#schemadatasource)]|false|none|none|

<h2 id="tocS_UpdateDataSourceResponse">UpdateDataSourceResponse</h2>
<!-- backwards compatibility -->
<a id="schemaupdatedatasourceresponse"></a>
<a id="schema_UpdateDataSourceResponse"></a>
<a id="tocSupdatedatasourceresponse"></a>
<a id="tocsupdatedatasourceresponse"></a>

```json
{
  "dataSources": [
    {
      "id": "string",
      "backend": "POSTGRESQL",
      "name": "string",
      "description": "string",
      "postgresqlParams": {
        "username": "string",
        "password": "string",
        "host": "string",
        "port": 0,
        "dbName": "string",
        "defaultSchema": "string"
      },
      "mysqlParams": {
        "username": "string",
        "password": "string",
        "host": "string",
        "port": 0,
        "dbName": "string",
        "defaultSchema": "string"
      },
      "virtualParams": {
        "mode": "DISCARD"
      },
      "redisParams": {
        "addr": "string",
        "password": "string",
        "db": 0
      },
      "mongoParams": {
        "uri": "string",
        "dbName": "string"
      },
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|dataSources|[[DataSource](#schemadatasource)]|false|none|none|

<h2 id="tocS_UpdateExtensionRequest">UpdateExtensionRequest</h2>
<!-- backwards compatibility -->
<a id="schemaupdateextensionrequest"></a>
<a id="schema_UpdateExtensionRequest"></a>
<a id="tocSupdateextensionrequest"></a>
<a id="tocsupdateextensionrequest"></a>

```json
{
  "token": "string",
  "extensions": [
    {
      "id": "string",
      "name": "string",
      "description": "string",
      "namespace": "string",
      "resource": "string",
      "before": {
        "all": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "create": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "update": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "delete": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "get": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "list": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "sync": true
      },
      "instead": {
        "all": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "create": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "update": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "delete": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "get": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "list": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "finalize": true
      },
      "after": {
        "all": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "create": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "update": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "delete": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "get": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "list": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "sync": true
      },
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|token|string|false|none|none|
|extensions|[[Extension](#schemaextension)]|false|none|none|

<h2 id="tocS_UpdateExtensionResponse">UpdateExtensionResponse</h2>
<!-- backwards compatibility -->
<a id="schemaupdateextensionresponse"></a>
<a id="schema_UpdateExtensionResponse"></a>
<a id="tocSupdateextensionresponse"></a>
<a id="tocsupdateextensionresponse"></a>

```json
{
  "extensions": [
    {
      "id": "string",
      "name": "string",
      "description": "string",
      "namespace": "string",
      "resource": "string",
      "before": {
        "all": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "create": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "update": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "delete": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "get": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "list": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "sync": true
      },
      "instead": {
        "all": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "create": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "update": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "delete": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "get": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "list": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "finalize": true
      },
      "after": {
        "all": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "create": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "update": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "delete": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "get": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "list": {
          "functionCall": {
            "host": "string",
            "functionName": "string"
          },
          "httpCall": {
            "uri": "string",
            "method": "string"
          }
        },
        "sync": true
      },
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|extensions|[[Extension](#schemaextension)]|false|none|none|

<h2 id="tocS_UpdateMultiRecordResponse">UpdateMultiRecordResponse</h2>
<!-- backwards compatibility -->
<a id="schemaupdatemultirecordresponse"></a>
<a id="schema_UpdateMultiRecordResponse"></a>
<a id="tocSupdatemultirecordresponse"></a>
<a id="tocsupdatemultirecordresponse"></a>

```json
{
  "record": {
    "id": "string",
    "properties": {
      "property1": null,
      "property2": null
    },
    "propertiesPacked": [
      null
    ]
  },
  "records": [
    {
      "id": "string",
      "properties": {
        "property1": null,
        "property2": null
      },
      "propertiesPacked": [
        null
      ]
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|record|[Record](#schemarecord)|false|none|none|
|records|[[Record](#schemarecord)]|false|none|none|

<h2 id="tocS_UpdateNamespaceRequest">UpdateNamespaceRequest</h2>
<!-- backwards compatibility -->
<a id="schemaupdatenamespacerequest"></a>
<a id="schema_UpdateNamespaceRequest"></a>
<a id="tocSupdatenamespacerequest"></a>
<a id="tocsupdatenamespacerequest"></a>

```json
{
  "token": "string",
  "Namespaces": [
    {
      "id": "string",
      "name": "string",
      "description": "string",
      "details": {},
      "securityContext": {
        "constraints": [
          {
            "namespace": "string",
            "resource": "string",
            "property": "string",
            "before": "2019-08-24T14:15:22Z",
            "after": "2019-08-24T14:15:22Z",
            "principal": "string",
            "recordIds": [
              "string"
            ],
            "operation": "OPERATION_TYPE_READ",
            "permit": "PERMIT_TYPE_ALLOW"
          }
        ]
      },
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|token|string|false|none|none|
|Namespaces|[[Namespace](#schemanamespace)]|false|none|none|

<h2 id="tocS_UpdateNamespaceResponse">UpdateNamespaceResponse</h2>
<!-- backwards compatibility -->
<a id="schemaupdatenamespaceresponse"></a>
<a id="schema_UpdateNamespaceResponse"></a>
<a id="tocSupdatenamespaceresponse"></a>
<a id="tocsupdatenamespaceresponse"></a>

```json
{
  "Namespaces": [
    {
      "id": "string",
      "name": "string",
      "description": "string",
      "details": {},
      "securityContext": {
        "constraints": [
          {
            "namespace": "string",
            "resource": "string",
            "property": "string",
            "before": "2019-08-24T14:15:22Z",
            "after": "2019-08-24T14:15:22Z",
            "principal": "string",
            "recordIds": [
              "string"
            ],
            "operation": "OPERATION_TYPE_READ",
            "permit": "PERMIT_TYPE_ALLOW"
          }
        ]
      },
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|Namespaces|[[Namespace](#schemanamespace)]|false|none|none|

<h2 id="tocS_UpdateRecordRequest_PropertiesEntry">UpdateRecordRequest_PropertiesEntry</h2>
<!-- backwards compatibility -->
<a id="schemaupdaterecordrequest_propertiesentry"></a>
<a id="schema_UpdateRecordRequest_PropertiesEntry"></a>
<a id="tocSupdaterecordrequest_propertiesentry"></a>
<a id="tocsupdaterecordrequest_propertiesentry"></a>

```json
{
  "key": "string",
  "value": null
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|key|string|false|none|none|
|value|[GoogleProtobufValue](#schemagoogleprotobufvalue)|false|none|Represents a dynamically typed value which can be either null, a number, a string, a boolean, a recursive struct value, or a list of values.|

<h2 id="tocS_UpdateRecordResponse">UpdateRecordResponse</h2>
<!-- backwards compatibility -->
<a id="schemaupdaterecordresponse"></a>
<a id="schema_UpdateRecordResponse"></a>
<a id="tocSupdaterecordresponse"></a>
<a id="tocsupdaterecordresponse"></a>

```json
{
  "record": {
    "id": "string",
    "properties": {
      "property1": null,
      "property2": null
    },
    "propertiesPacked": [
      null
    ]
  },
  "records": [
    {
      "id": "string",
      "properties": {
        "property1": null,
        "property2": null
      },
      "propertiesPacked": [
        null
      ]
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|record|[Record](#schemarecord)|false|none|none|
|records|[[Record](#schemarecord)]|false|none|none|

<h2 id="tocS_UpdateResourceRequest">UpdateResourceRequest</h2>
<!-- backwards compatibility -->
<a id="schemaupdateresourcerequest"></a>
<a id="schema_UpdateResourceRequest"></a>
<a id="tocSupdateresourcerequest"></a>
<a id="tocsupdateresourcerequest"></a>

```json
{
  "token": "string",
  "resources": [
    {
      "id": "string",
      "name": "string",
      "namespace": "string",
      "sourceConfig": {
        "dataSource": "string",
        "catalog": "string",
        "entity": "string"
      },
      "properties": [
        {
          "id": "string",
          "name": "string",
          "type": "BOOL",
          "mapping": "string",
          "required": true,
          "primary": true,
          "length": 0,
          "unique": true,
          "immutable": true,
          "securityContext": {
            "constraints": [
              {
                "namespace": "string",
                "resource": "string",
                "property": "string",
                "before": "2019-08-24T14:15:22Z",
                "after": "2019-08-24T14:15:22Z",
                "principal": "string",
                "recordIds": [
                  "string"
                ],
                "operation": "OPERATION_TYPE_READ",
                "permit": "PERMIT_TYPE_ALLOW"
              }
            ]
          },
          "defaultValue": null,
          "exampleValue": null,
          "enumValues": [
            null
          ],
          "reference": {
            "referencedResource": "string",
            "cascade": true
          },
          "properties": [
            {}
          ],
          "Item": {},
          "title": "string",
          "description": "string",
          "annotations": {
            "property1": "string",
            "property2": "string"
          }
        }
      ],
      "indexes": [
        {
          "properties": [
            {
              "name": "string",
              "order": "ORDER_UNKNOWN"
            }
          ],
          "indexType": "BTREE",
          "unique": true,
          "annotations": {
            "property1": "string",
            "property2": "string"
          }
        }
      ],
      "securityContext": {
        "constraints": [
          {
            "namespace": "string",
            "resource": "string",
            "property": "string",
            "before": "2019-08-24T14:15:22Z",
            "after": "2019-08-24T14:15:22Z",
            "principal": "string",
            "recordIds": [
              "string"
            ],
            "operation": "OPERATION_TYPE_READ",
            "permit": "PERMIT_TYPE_ALLOW"
          }
        ]
      },
      "virtual": true,
      "immutable": true,
      "abstract": true,
      "title": "string",
      "description": "string",
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0,
      "annotations": {
        "property1": "string",
        "property2": "string"
      }
    }
  ],
  "doMigration": true,
  "forceMigration": true,
  "annotations": {
    "property1": "string",
    "property2": "string"
  }
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|token|string|false|none|none|
|resources|[[Resource](#schemaresource)]|false|none|none|
|doMigration|boolean|false|none|none|
|forceMigration|boolean|false|none|none|
|annotations|object|false|none|none|
|» **additionalProperties**|string|false|none|none|

<h2 id="tocS_UpdateResourceResponse">UpdateResourceResponse</h2>
<!-- backwards compatibility -->
<a id="schemaupdateresourceresponse"></a>
<a id="schema_UpdateResourceResponse"></a>
<a id="tocSupdateresourceresponse"></a>
<a id="tocsupdateresourceresponse"></a>

```json
{
  "resources": [
    {
      "id": "string",
      "name": "string",
      "namespace": "string",
      "sourceConfig": {
        "dataSource": "string",
        "catalog": "string",
        "entity": "string"
      },
      "properties": [
        {
          "id": "string",
          "name": "string",
          "type": "BOOL",
          "mapping": "string",
          "required": true,
          "primary": true,
          "length": 0,
          "unique": true,
          "immutable": true,
          "securityContext": {
            "constraints": [
              {
                "namespace": "string",
                "resource": "string",
                "property": "string",
                "before": "2019-08-24T14:15:22Z",
                "after": "2019-08-24T14:15:22Z",
                "principal": "string",
                "recordIds": [
                  "string"
                ],
                "operation": "OPERATION_TYPE_READ",
                "permit": "PERMIT_TYPE_ALLOW"
              }
            ]
          },
          "defaultValue": null,
          "exampleValue": null,
          "enumValues": [
            null
          ],
          "reference": {
            "referencedResource": "string",
            "cascade": true
          },
          "properties": [
            {}
          ],
          "Item": {},
          "title": "string",
          "description": "string",
          "annotations": {
            "property1": "string",
            "property2": "string"
          }
        }
      ],
      "indexes": [
        {
          "properties": [
            {
              "name": "string",
              "order": "ORDER_UNKNOWN"
            }
          ],
          "indexType": "BTREE",
          "unique": true,
          "annotations": {
            "property1": "string",
            "property2": "string"
          }
        }
      ],
      "securityContext": {
        "constraints": [
          {
            "namespace": "string",
            "resource": "string",
            "property": "string",
            "before": "2019-08-24T14:15:22Z",
            "after": "2019-08-24T14:15:22Z",
            "principal": "string",
            "recordIds": [
              "string"
            ],
            "operation": "OPERATION_TYPE_READ",
            "permit": "PERMIT_TYPE_ALLOW"
          }
        ]
      },
      "virtual": true,
      "immutable": true,
      "abstract": true,
      "title": "string",
      "description": "string",
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0,
      "annotations": {
        "property1": "string",
        "property2": "string"
      }
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|resources|[[Resource](#schemaresource)]|false|none|none|

<h2 id="tocS_UpdateUserResponse">UpdateUserResponse</h2>
<!-- backwards compatibility -->
<a id="schemaupdateuserresponse"></a>
<a id="schema_UpdateUserResponse"></a>
<a id="tocSupdateuserresponse"></a>
<a id="tocsupdateuserresponse"></a>

```json
{
  "user": {
    "id": "string",
    "username": "string",
    "password": "string",
    "securityContext": {
      "constraints": [
        {
          "namespace": "string",
          "resource": "string",
          "property": "string",
          "before": "2019-08-24T14:15:22Z",
          "after": "2019-08-24T14:15:22Z",
          "principal": "string",
          "recordIds": [
            "string"
          ],
          "operation": "OPERATION_TYPE_READ",
          "permit": "PERMIT_TYPE_ALLOW"
        }
      ]
    },
    "details": {},
    "signKey": "string",
    "auditData": {
      "createdOn": "2019-08-24T14:15:22Z",
      "updatedOn": "2019-08-24T14:15:22Z",
      "createdBy": "string",
      "updatedBy": "string"
    },
    "version": 0
  },
  "users": [
    {
      "id": "string",
      "username": "string",
      "password": "string",
      "securityContext": {
        "constraints": [
          {
            "namespace": "string",
            "resource": "string",
            "property": "string",
            "before": "2019-08-24T14:15:22Z",
            "after": "2019-08-24T14:15:22Z",
            "principal": "string",
            "recordIds": [
              "string"
            ],
            "operation": "OPERATION_TYPE_READ",
            "permit": "PERMIT_TYPE_ALLOW"
          }
        ]
      },
      "details": {},
      "signKey": "string",
      "auditData": {
        "createdOn": "2019-08-24T14:15:22Z",
        "updatedOn": "2019-08-24T14:15:22Z",
        "createdBy": "string",
        "updatedBy": "string"
      },
      "version": 0
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|user|[User](#schemauser)|false|none|none|
|users|[[User](#schemauser)]|false|none|none|

<h2 id="tocS_User">User</h2>
<!-- backwards compatibility -->
<a id="schemauser"></a>
<a id="schema_User"></a>
<a id="tocSuser"></a>
<a id="tocsuser"></a>

```json
{
  "id": "string",
  "username": "string",
  "password": "string",
  "securityContext": {
    "constraints": [
      {
        "namespace": "string",
        "resource": "string",
        "property": "string",
        "before": "2019-08-24T14:15:22Z",
        "after": "2019-08-24T14:15:22Z",
        "principal": "string",
        "recordIds": [
          "string"
        ],
        "operation": "OPERATION_TYPE_READ",
        "permit": "PERMIT_TYPE_ALLOW"
      }
    ]
  },
  "details": {},
  "signKey": "string",
  "auditData": {
    "createdOn": "2019-08-24T14:15:22Z",
    "updatedOn": "2019-08-24T14:15:22Z",
    "createdBy": "string",
    "updatedBy": "string"
  },
  "version": 0
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|id|string|false|none|none|
|username|string|false|none|none|
|password|string|false|none|none|
|securityContext|[SecurityContext](#schemasecuritycontext)|false|none|none|
|details|object|false|none|none|
|signKey|string|false|none|none|
|auditData|[AuditData](#schemaauditdata)|false|none|none|
|version|integer(uint32)|false|none|none|

<h2 id="tocS_VirtualParams">VirtualParams</h2>
<!-- backwards compatibility -->
<a id="schemavirtualparams"></a>
<a id="schema_VirtualParams"></a>
<a id="tocSvirtualparams"></a>
<a id="tocsvirtualparams"></a>

```json
{
  "mode": "DISCARD"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|mode|string(enum)|false|none|none|

#### Enumerated Values

|Property|Value|
|---|---|
|mode|DISCARD|
|mode|ERROR|

<h2 id="tocS_WatchMessage">WatchMessage</h2>
<!-- backwards compatibility -->
<a id="schemawatchmessage"></a>
<a id="schema_WatchMessage"></a>
<a id="tocSwatchmessage"></a>
<a id="tocswatchmessage"></a>

```json
{
  "changes": {},
  "recordIds": [
    "string"
  ],
  "event": "CREATE",
  "eventOn": "2019-08-24T14:15:22Z"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|changes|object|false|none|none|
|recordIds|[string]|false|none|none|
|event|string(enum)|false|none|none|
|eventOn|string(date-time)|false|none|none|

#### Enumerated Values

|Property|Value|
|---|---|
|event|CREATE|
|event|UPDATE|
|event|DELETE|
|event|GET|
|event|LIST|

