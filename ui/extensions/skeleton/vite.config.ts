import react from '@vitejs/plugin-react';
import path from 'node:path';
import {defineConfig} from 'vite';

export default defineConfig({
    plugins: [
        react(),
    ],
    logLevel: 'info',
    build: {
        minify: false,
        sourcemap: true,
        lib: {
            entry: path.resolve(__dirname, 'src/proto.ts'),
            name: 'Skeleton',
            formats: ['cjs','es', 'umd'],
            fileName: (format) => `skeleton.${format}.js`,
        },
        rollupOptions: {
            external: (source: string, importer: string | undefined, isResolved: boolean) => {
                if (source.indexOf('/skeleton/') !== -1) {
                    return false
                }

                if (source.indexOf('.') === 0) {
                    return false
                }

                console.log(source, importer, isResolved);

                if (isResolved) {
                    return true
                }

                return true
            },
            output: {
                globals: (name) => {
                    console.log(name);
                    return name
                },
            },
        },
    },
})
;
