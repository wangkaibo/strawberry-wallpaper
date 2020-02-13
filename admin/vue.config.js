const path = require('path')

module.exports = {
    outputDir: path.resolve(__dirname, '../static'),
    // module: {
    //     rules: [
    //         {
    //             test: /\.less$/,
    //             use: [
    //                 'vue-style-loader',
    //                 'css-loader',
    //                 'less-loader'
    //             ]
    //         },
    //     ]
    // }
}
