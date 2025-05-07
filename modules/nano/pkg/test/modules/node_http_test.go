package modules

import (
	"github.com/apibrew/apibrew/modules/nano/pkg/nodejs"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
	"testing"
)

func TestNodeHttp(t *testing.T) {
	jsCode := `
const https = require('https');

https.get('https://jsonplaceholder.typicode.com/users', res => {
  let data = [];
  const headerDate = res.headers && res.headers.date ? res.headers.date : 'no response date';
  console.log('Status Code:', res.statusCode);
  console.log('Date in Response header:', headerDate);

  res.on('data', chunk => {
    data.push(chunk);
  });

  res.on('end', () => {
    console.log('Response ended: ');
    const users = JSON.parse(Buffer.concat(data).toString());

    for(user of users) {
      console.log('Got user with id: ${user.id}, name: ${user.name}');
    }
  });
}).on('error', err => {
  console.log('Error: ', err.message);
});
`

	vm := goja.New()

	var req = require.NewRegistry()
	req.Enable(vm)

	nodejs.RegisterModules(req)

	res, err := vm.RunString(jsCode)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(res)

}
