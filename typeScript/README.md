## 撒点儿调料——TypeScript

### PS:

- 公司项目需要, 暂定需要用ts写脚本, 写一些机器人, 本项目作为学习ts的一个记录

## 使用webpack打包ts

- 初始化package.json (打包配置文件) `npm init -y`
- 安装webpack `npm i -D webpack webpack-cli typescript ts-loader`
- 编写webpack配置文件 [webpack.config.js](webpack.config.js)
- 编写ts配置文件 [tsconfig.json](tsconfig.json)
- 修改[package.json](package.json)添加打包命令
    - `"dev": "webpack --mode development",`
    - `"build": "webpack --mode production"`