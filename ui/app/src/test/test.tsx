import {BACKEND_URL} from '../config'
import {PageLayout} from '../layout/PageLayout'
import Button from '@mui/material/Button'
import {PlusOneOutlined} from '@mui/icons-material'
import React, {JSX} from 'react'

export function Test(): JSX.Element {
    const code = `
    function _interopDefault (ex) { return (ex && (typeof ex === 'object') && 'default' in ex) ? ex['default'] : ex; }

var React = require('react');
var React__default = _interopDefault(React);

var styles = {"test":"_3ybTi"};

function Component1() {
  return React__default.createElement("div", null, "Hello World ", React__default.createElement("span", null, "!"));
}

var ExampleComponent = function ExampleComponent(_ref) {
  var text = _ref.text;
  return React.createElement("div", {
    className: styles.test
  }, "Example Component: ", text);
};

exports.Component1 = Component1;
exports.ExampleComponent = ExampleComponent;
//# sourceMappingURL=index.js.map
    `
    const params: { [key: string]: any } = {
        exports: {},
        React: React,
        require: (name: string) => {
            switch (name) {
                case 'react':
                    return React;
                case '@mui/material/Button':
                    return Button;
            }

            throw new Error(`Cannot find module '${name}'`);
        }
    }

    const paramKeys = Object.keys(params)
    const paramValues = paramKeys.map(key => params[key])

    const func = new Function(...paramKeys, `return (function(){${code}})()`);

    func(...paramValues);

    const Component1 = params.exports.Component1;

    return (
        <PageLayout pageTitle={'Test Page'} actions={<>
            <Button variant={'contained'} color='success' startIcon={<PlusOneOutlined/>}>New Item</Button>
        </>}>
            <>
                Hello World {BACKEND_URL}
                <br/>
                <Component1/>
            </>
        </PageLayout>
    )
}
