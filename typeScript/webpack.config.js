// 引入一个包
const path = require('path');
// 引入html-webpack-plugin
const HtmlWebpackPlugin = require('html-webpack-plugin')
const {CleanWebpackPlugin} = require('clean-webpack-plugin');

// webpack中的所有配置信息都应该写在module.exports中
module.exports = {
    // 指定入口文件
    entry: './src/index.ts',
    // 指定输出文件
    output: {
        // 指定输出文件的目录
        path: path.resolve(__dirname, 'dist'),
        // 打包后输出文件的文件名
        filename: 'bundle.js'
    },
    // 设置mode development为开发模式 production为生产模式
    mode: 'development',
    // 指定webpack打包时要使用的模块
    module: {
        // 指定加载的规则
        rules: [
            {
                // test指定的是规则生效的文件
                test: /\.ts$/,
                // 要使用的loader
                use: 'ts-loader',
                // 要排除的文件
                exclude: /node_modules/
            }
        ]
    },

    // 配置webpack插件
    plugins: [
        new HtmlWebpackPlugin(
            {
                title: "Webpack Demo"
            }
        ),
        new CleanWebpackPlugin(),
    ],

    // 用来设置引用模块
    resolve: {
        extensions: ['.ts', '.js']
    },

};