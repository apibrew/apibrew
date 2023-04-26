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
    }
}
