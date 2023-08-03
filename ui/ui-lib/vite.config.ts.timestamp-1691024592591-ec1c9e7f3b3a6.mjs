// vite.config.ts
import react from "file:///Users/taleh/Projects/apibrew/ui/node_modules/@vitejs/plugin-react/dist/index.mjs";
import path from "node:path";
import { defineConfig } from "file:///Users/taleh/Projects/apibrew/ui/node_modules/vite/dist/node/index.js";
import dts from "file:///Users/taleh/Projects/apibrew/ui/node_modules/vite-plugin-dts/dist/index.mjs";
var __vite_injected_original_dirname = "/Users/taleh/Projects/apibrew/ui/ui-lib";
var vite_config_default = defineConfig({
  plugins: [
    react(),
    dts({
      insertTypesEntry: true
    })
  ],
  logLevel: "info",
  build: {
    minify: false,
    sourcemap: true,
    lib: {
      entry: path.resolve(__vite_injected_original_dirname, "src/proto.ts"),
      name: "CoreLib",
      formats: ["es", "umd"],
      fileName: (format) => `ui-lib.${format}.js`
    },
    rollupOptions: {
      external: [/react/, "react-dom", "styled-components", /^@mui\//, "axios"],
      input: [
        "./src/"
      ],
      output: {
        preserveModules: false,
        globals: (name) => {
          return name;
        }
      }
    }
  }
});
export {
  vite_config_default as default
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsidml0ZS5jb25maWcudHMiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImNvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9kaXJuYW1lID0gXCIvVXNlcnMvdGFsZWgvUHJvamVjdHMvYXBpYnJldy91aS91aS1saWJcIjtjb25zdCBfX3ZpdGVfaW5qZWN0ZWRfb3JpZ2luYWxfZmlsZW5hbWUgPSBcIi9Vc2Vycy90YWxlaC9Qcm9qZWN0cy9hcGlicmV3L3VpL3VpLWxpYi92aXRlLmNvbmZpZy50c1wiO2NvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9pbXBvcnRfbWV0YV91cmwgPSBcImZpbGU6Ly8vVXNlcnMvdGFsZWgvUHJvamVjdHMvYXBpYnJldy91aS91aS1saWIvdml0ZS5jb25maWcudHNcIjtpbXBvcnQgcmVhY3QgZnJvbSAnQHZpdGVqcy9wbHVnaW4tcmVhY3QnO1xuaW1wb3J0IHBhdGggZnJvbSAnbm9kZTpwYXRoJztcbmltcG9ydCB7ZGVmaW5lQ29uZmlnfSBmcm9tICd2aXRlJztcbmltcG9ydCBkdHMgZnJvbSAndml0ZS1wbHVnaW4tZHRzJztcblxuZXhwb3J0IGRlZmF1bHQgZGVmaW5lQ29uZmlnKHtcbiAgICBwbHVnaW5zOiBbXG4gICAgICAgIHJlYWN0KCksXG4gICAgICAgIGR0cyh7XG4gICAgICAgICAgICBpbnNlcnRUeXBlc0VudHJ5OiB0cnVlLFxuICAgICAgICB9KSxcbiAgICBdLFxuICAgIGxvZ0xldmVsOiAnaW5mbycsXG4gICAgYnVpbGQ6IHtcbiAgICAgICAgbWluaWZ5OiBmYWxzZSxcbiAgICAgICAgc291cmNlbWFwOiB0cnVlLFxuICAgICAgICBsaWI6IHtcbiAgICAgICAgICAgIGVudHJ5OiBwYXRoLnJlc29sdmUoX19kaXJuYW1lLCAnc3JjL3Byb3RvLnRzJyksXG4gICAgICAgICAgICBuYW1lOiAnQ29yZUxpYicsXG4gICAgICAgICAgICBmb3JtYXRzOiBbJ2VzJywgJ3VtZCddLFxuICAgICAgICAgICAgZmlsZU5hbWU6IChmb3JtYXQpID0+IGB1aS1saWIuJHtmb3JtYXR9LmpzYCxcbiAgICAgICAgfSxcbiAgICAgICAgcm9sbHVwT3B0aW9uczoge1xuICAgICAgICAgICAgZXh0ZXJuYWw6IFsvcmVhY3QvLCAncmVhY3QtZG9tJywgJ3N0eWxlZC1jb21wb25lbnRzJywgL15AbXVpXFwvLywgJ2F4aW9zJywgXSxcbiAgICAgICAgICAgIGlucHV0OiBbXG4gICAgICAgICAgICAgICAgJy4vc3JjLydcbiAgICAgICAgICAgIF0sXG4gICAgICAgICAgICBvdXRwdXQ6IHtcbiAgICAgICAgICAgICAgICBwcmVzZXJ2ZU1vZHVsZXM6IGZhbHNlLFxuICAgICAgICAgICAgICAgIGdsb2JhbHM6IChuYW1lKSA9PiB7XG4gICAgICAgICAgICAgICAgICAgIHJldHVybiBuYW1lXG4gICAgICAgICAgICAgICAgfSxcbiAgICAgICAgICAgIH0sXG4gICAgICAgIH0sXG4gICAgfSxcbn0pO1xuIl0sCiAgIm1hcHBpbmdzIjogIjtBQUF1UyxPQUFPLFdBQVc7QUFDelQsT0FBTyxVQUFVO0FBQ2pCLFNBQVEsb0JBQW1CO0FBQzNCLE9BQU8sU0FBUztBQUhoQixJQUFNLG1DQUFtQztBQUt6QyxJQUFPLHNCQUFRLGFBQWE7QUFBQSxFQUN4QixTQUFTO0FBQUEsSUFDTCxNQUFNO0FBQUEsSUFDTixJQUFJO0FBQUEsTUFDQSxrQkFBa0I7QUFBQSxJQUN0QixDQUFDO0FBQUEsRUFDTDtBQUFBLEVBQ0EsVUFBVTtBQUFBLEVBQ1YsT0FBTztBQUFBLElBQ0gsUUFBUTtBQUFBLElBQ1IsV0FBVztBQUFBLElBQ1gsS0FBSztBQUFBLE1BQ0QsT0FBTyxLQUFLLFFBQVEsa0NBQVcsY0FBYztBQUFBLE1BQzdDLE1BQU07QUFBQSxNQUNOLFNBQVMsQ0FBQyxNQUFNLEtBQUs7QUFBQSxNQUNyQixVQUFVLENBQUMsV0FBVyxVQUFVO0FBQUEsSUFDcEM7QUFBQSxJQUNBLGVBQWU7QUFBQSxNQUNYLFVBQVUsQ0FBQyxTQUFTLGFBQWEscUJBQXFCLFdBQVcsT0FBUztBQUFBLE1BQzFFLE9BQU87QUFBQSxRQUNIO0FBQUEsTUFDSjtBQUFBLE1BQ0EsUUFBUTtBQUFBLFFBQ0osaUJBQWlCO0FBQUEsUUFDakIsU0FBUyxDQUFDLFNBQVM7QUFDZixpQkFBTztBQUFBLFFBQ1g7QUFBQSxNQUNKO0FBQUEsSUFDSjtBQUFBLEVBQ0o7QUFDSixDQUFDOyIsCiAgIm5hbWVzIjogW10KfQo=
