'use strict'

const path = require('path');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const CssMinimizerPlugin = require("css-minimizer-webpack-plugin");
const WebpackObfuscator = require('webpack-obfuscator');

module.exports = {
    watch: false,
    mode: 'production',
    entry: {
        ebook: './middleware/ebook/main.js',
    },
    output: {
        path: path.resolve(__dirname, 'static'),
        filename: 'js/[name].bundle.js',
    },
    module: {
        rules: [
            {
                test: /\.css$/,
                use: [MiniCssExtractPlugin.loader, 'css-loader']
            },
            {
                test: /\.js$/,
                exclude: /node_modules/,
                use: {
                    loader: 'babel-loader',
                    options: {
                        presets: ['@babel/preset-env'],
                    }
                },
            }
        ]
    },
    optimization: {
        usedExports: true,
        minimizer: [
            new CssMinimizerPlugin(),
        ],
        minimize: true,
    },
    plugins: [
        new MiniCssExtractPlugin({ filename: 'css/[name].css' }),
        new WebpackObfuscator({
            rotateStringArray: true,
          })
    ],
    devtool: false,
};