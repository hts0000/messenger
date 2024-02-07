## Frontend
```bash
npx create-next-app@latest
npm install
npm run dev

npm install axios

npm install @types/long
npm install protobufjs protobufjs-cli
# npm install ts-protoc-gen # 另一种生成js/ts代码的方案

# tailwind.config.js
/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}

# global.css
@tailwind base;
@tailwind components;
@tailwind utilities;

# shadcn ui
https://ui.shadcn.com/docs/installation/vite
```

## Backend
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
```