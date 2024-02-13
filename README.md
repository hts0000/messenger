## Frontend
```bash
npx create-next-app@latest
npm install
npm run dev

npm install axios

npm install @types/long
npm install protobufjs protobufjs-cli
# npm install ts-protoc-gen # 另一种生成js/ts代码的方案

npm install react-hook-form react-hot-toast react-icons

npm install @tailwindcss/forms
# tailwind.config.js
plugins: [
  require("@tailwindcss/forms")({
    strategy: "class",
  }),
],

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

go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

## MySQL
```bash
# 初始化
# -e MYSQL_ROOT_PASSWORD=123456: 注入环境变量，配置root用户密码
# -e MYSQL_DATABASE=messenger: 注入环境变量，自动创建messenger库
# -v /my/custom:/etc/mysql/conf.d: 将本地配置文件挂载到docker中
# -v /my/own/datadir:/var/lib/mysql: 使用本地目录作为数据库存储目录
# -v /my/initsql:/docker-entrypoint-initdb.d: 挂载到/docker-entrypoint-initdb.d目录的脚本，容器初始化时会自动执行
# -p 13306:3306: 将容器3306端口映射到本机13306端口
# --character-set-server=utf8mb4: 配置数据库字符集
# --collation-server=utf8mb4_unicode_ci: 配置数据表字符集
# --restart on-failures: 容器异常退出时，总是重启容器
docker run --name messenger-mysql \
            --restart on-failures \
            -v /my/custom:/etc/mysql/conf.d \
            -v /my/own/datadir:/var/lib/mysql \
            -v /somepath/messenger/mysql/sql:/docker-entrypoint-initdb.d \
            -v /somepath/messenger/mysql/sql:/docker-entrypoint-initdb.d \
            -e MYSQL_ROOT_PASSWORD=123456 \
            -e MYSQL_DATABASE=messenger \
            -p 13306:3306 \
            -d mysql:8.3.0 \
            --character-set-server=utf8mb4 \
            --collation-server=utf8mb4_unicode_ci

# 测试使用
docker run --name messenger-mysql-test \
            --restart on-failures \
            -v /somepath/messenger/mysql/conf.d:/etc/mysql/conf.d \
            -v /somepath/messenger/mysql/sql:/docker-entrypoint-initdb.d \
            -e MYSQL_ROOT_PASSWORD=123456 \
            -e MYSQL_DATABASE=messenger \
            -p 13306:3306 \
            -d mysql:8.3.0 \
            --character-set-server=utf8mb4 \
            --collation-server=utf8mb4_unicode_ci

# 连接
docker exec -it messenger-mysql mysql -uroot -p123456
```

## Git Commit规范
| 类型     | 描述                                               |
| -------- | -------------------------------------------------- |
| feat     | 新增feature                                        |
| fix      | 修复bug                                            |
| docs     | 修改文档，如readme.md                              |
| style    | 修改代码格式，不改变代码逻辑，如逗号、缩进、空格等 |
| refactor | 代码重构，没有新增功能或修复bug                    |
| perf     | 优化相关，如提升性能、用户体验等                   |
| test     | 测试用例，包括单元测试、集成测试                   |
| ci       | 修改ci配置文件或脚本，如jenkins fastlame           |
| chore    | 修改构建脚本、或者增加依赖库、工具等               |
| revert   | 回滚之前的commit                                   |

## TODO
1. zap logger自定义打印格式，打印微服务名称
2. 双token刷新验证