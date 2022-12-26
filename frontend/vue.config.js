// 详细配置请参考https://cli.vuejs.org/zh/config/#vue-config-js
const outputDir = process.env.VUE_APP_CONTEXT_PATH.substring(1, process.env.VUE_APP_CONTEXT_PATH.length - 1)
module.exports = {
  publicPath: process.env.VUE_APP_CONTEXT_PATH,
  outputDir: outputDir === '/' ? 'dist' : outputDir,
  assetsDir: 'static',
  lintOnSave: false,
  devServer: {
    host: '0.0.0.0',
    port: 10084,
    proxy: {
      [process.env.VUE_APP_API_PREFIX]: {
        target: 'http://192.168.3.104:8887',
        changeOrigin: true,
        pathRewrite: {
          [`^${[process.env.VUE_APP_API_PREFIX]}`]: ''
        }
      }
    }
  }
}
