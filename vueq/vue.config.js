module.exports = {

  // Whether to perform lint-on-save during development using eslint-loader.
  // This value is respected only when @vue/cli-plugin-eslint is installed.
  lintOnSave: false,

  // Whether to use the build of Vue core that includes the runtime compiler.
  // Setting it to true will allow you to use the template option in Vue components, but will incur around an extra 10kb payload for your app.
  runtimeCompiler: undefined,

  // Setting this to false can speed up production builds if you don't need source maps for production.
  productionSourceMap: false,

  // Whether to use thread-loader for Babel or TypeScript transpilation.
  // This is enabled for production builds when the system has more than 1 CPU cores.
  // Passing a number will define the amount of workers used.
  parallel: undefined,

  // If your frontend app and the backend API server are not running on the same host,
  // you will need to proxy API requests to the API server during development.
  // proxy API requests to Laravel during development
  devServer: {
    proxy: 'https://dashboard.productlead.evx',
    progress: false // https://github.com/vuejs/vue-cli/issues/4557
  },

}
