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
- 配置完成后, 执行 `npm run dev` 即可打包ts文件

## 配置webpack

- 下载插件 `npm i -D html-webpack-plugin`
    - 修改[webpack.config.js](webpack.config.js)配置文件
    - 添加插件 `const HtmlWebpackPlugin = require('html-webpack-plugin');`
    - 添加插件配置

    ```js
    plugins: [
        new HtmlWebpackPlugin({
            template: './src/index.html'
        })
    ]
    ```
- 配置webpack服务器 `npm i -D webpack-dev-server`
    - 修改[package.json](package.json)添加启动命令
        - `"start": "webpack server --open"`
- 下载插件 `npm i -D clean-webpack-plugin`
    - 修改[webpack.config.js](webpack.config.js)配置文件
    - 添加插件 `const { CleanWebpackPlugin } = require('clean-webpack-plugin');`
    - 添加插件配置

    ```js
    plugins: [
        new CleanWebpackPlugin(),
    ]
    ```
- 设置引用模块
    - 修改[webpack.config.js](webpack.config.js)配置文件
    - 添加模块配置 以ts，js结尾的文件可以作为模块使用

    ```js
    resolve: {
        extensions: ['.ts', '.js']
    }
    ```
- 配置babel, [哔哩哔哩](https://www.bilibili.com/video/BV1Xy4y1v7S2?p=12)