Library app
=======

This is a simple library app that demonstrates how to use the API Brew.

Prerequisites
 - install apbr
 - install API Brew server
 - clone/download schema file [schema.yaml](schema.yaml)

Let's apply our schema

```bash
apbr apply -f schema.yaml
```

Now let's see our Rest API documentation through swagger. Open http://localhost:9009/docs

