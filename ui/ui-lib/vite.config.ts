import react from '@vitejs/plugin-react';
import path from 'node:path';
import {defineConfig} from 'vite';
import dts from 'vite-plugin-dts';

export default defineConfig({
    plugins: [
        react(),
        dts({
            insertTypesEntry: true,
        }),
    ],
    logLevel: 'info',
    build: {
        minify: false,
        sourcemap: true,
        lib: {
            entry: path.resolve(__dirname, 'src/proto.ts'),
            name: 'CoreLib',
            formats: ['es', 'umd'],
            fileName: (format) => `ui-lib.${format}.js`,
        },
        rollupOptions: {
            external: [/react/, 'react-dom', 'styled-components', /^@mui\//, 'axios', ],
            input: [
                './src/'
            ],
            output: {
                preserveModules: false,
                globals: (name) => {
                    return name
                },
            },
        },
    },
});
