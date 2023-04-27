module.exports = {
    env: {
        browser: true,
        es2021: true
    },
    extends: [
        'plugin:react/recommended',
        'standard-with-typescript'
    ],
    overrides: [],
    parserOptions: {
        ecmaVersion: 'latest',
        sourceType: 'module',
        project: ["tsconfig.json"]
    },
    plugins: [
        'react'
    ],
    rules: {
        "@typescript-eslint/space-before-function-paren": "off",
        "@typescript-eslint/indent": ["error", 4],
    },
    settings: {
        "react": {
            "createClass": "createReactClass", // Regex for Component Factory to use,
                                               // default to "createReactClass"
            "pragma": "React",  // Pragma to use, default to "React"
            "fragment": "Fragment",  // Fragment to use (may be a property of <pragma>), default to "Fragment"
            "version": "detect", // React version. "detect" automatically picks the version you have installed.
                                 // You can also use `16.0`, `16.3`, etc, if you want to override the detected value.
                                 // It will default to "latest" and warn if missing, and to "detect" in the future
            "flowVersion": "0.53" // Flow version
        },
    }
}
