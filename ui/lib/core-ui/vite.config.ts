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
            entry: path.resolve(__dirname, 'src/index.ts'),
            name: 'CoreUI',
            formats: ['es', 'umd'],
            fileName: (format) => `core-ui.${format}.js`,
        },
        rollupOptions: {
            external: [/react/, 'react-dom', 'styled-components', /^@mui\//, 'axios', ],
            output: {
                globals: (name) => {
                    return name
                },
            },
        },
    },
});
