const TerserPlugin = require("terser-webpack-plugin");

console.log(process.env.VUE_APP_URL);

module.exports = {
  transpileDependencies: [/node_modules[/\\\\]vuetify[/\\\\]/],
  filenameHashing: false,
  lintOnSave: false,
  productionSourceMap: false,

  // proxy API requests to Laravel during development
  devServer: {
    proxy: "http://baron.test:3000"
  },
  // paths
  outputDir: "../golang/public/",
  //   publicPath: (process.env.NODE_ENV === 'production') ? process.env.VUE_APP_URL : undefined,
  assetsDir: "backend/assets",

  // configure webpack / terser
  configureWebpack: {
    optimization: {
      splitChunks: {
        cacheGroups: {
          vendors: {
            name: "vendors",
            test: /[\\/]node_modules[\\/]/,
            priority: -10,
            chunks: "initial"
          }
        }
      },
      minimizer: [
        new TerserPlugin({
          cache: true,
          parallel: true,
          sourceMap: false, // Must be set to true if using source-maps in production
          terserOptions: {
            // https://github.com/webpack-contrib/terser-webpack-plugin#terseroptions
            output: { comments: false }
          }
        })
      ]
    },
    output: {
      filename: "backend/assets/js/baron.[name].js",
      chunkFilename: "backend/assets/js/baron.[name].js",
      jsonpFunction: "jsonpFunction"
    }
  },

  // delete HTML related webpack plugins
  chainWebpack: config => {
    if (process.env.NODE_ENV === "production") {
      config.plugins.delete("html");
      config.plugins.delete("preload");
      config.plugins.delete("prefetch");
      config.plugins.delete("copy");
    }
  },

  runtimeCompiler: false,
  parallel: undefined
};
