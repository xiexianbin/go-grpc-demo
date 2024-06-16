
- get swagger-ui from https://github.com/swagger-api/swagger-ui/tree/master/dist
- update swagger-ui/swagger-initializer.js

```
sed 's#https://petstore.swagger.io/v2/swagger.json#/demo/v1/demo.swagger.json#g' -i swagger-ui/swagger-initializer.js
```
