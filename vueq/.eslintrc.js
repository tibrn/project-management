module.exports = {
  root: true,

  parserOptions: {
    parser: "@typescript-eslint/parser",
    sourceType: "module",
    ecmaFeatures: {
      legacyDecorators: true
    }
  },

  env: {
    browser: true
  },

  extends: [
    // "airbnb-base",
    // // Uncomment any of the lines below to choose desired strictness,
    // // but leave only one uncommented!
    // // See https://eslint.vuejs.org/rules/#available-rules
    // "plugin:vue/essential" // Priority A: Essential (Error Prevention)
    // // 'plugin:vue/strongly-recommended' // Priority B: Strongly Recommended (Improving Readability)
    // // 'plugin:vue/recommended' // Priority C: Recommended (Minimizing Arbitrary Choices and Cognitive Overhead)

    'plugin:vue/essential',
    // '@vue/airbnb',
    '@vue/typescript/recommended',

    '@vue/prettier',
    '@vue/prettier/@typescript-eslint'
  ],

  // required to lint *.vue files
  plugins: ["vue"],

  globals: {
    ga: true, // Google Analytics
    cordova: true,
    __statics: true,
    process: true,
    Capacitor: true,
    chrome: true
  },

  // add your custom rules here
  rules: {
    "no-param-reassign": "off",

    "import/first": "off",
    "import/named": "error",
    "import/namespace": "error",
    "import/default": "error",
    "import/export": "error",
    "import/extensions": "off",
    "import/no-unresolved": "off",
    "import/no-extraneous-dependencies": "off",
    "import/prefer-default-export": "off",
    "prefer-promise-reject-errors": "off",

    "quotes": "off",
    camelcase: "off",
    "no-empty": "off",
    "max-len": ["warn", { code: 150 }],
    // allow debugger during development only
    "no-debugger": process.env.NODE_ENV === "production" ? "error" : "off",
    "no-unused-vars": [
      "warn",
      {
        vars: "all",
        argsIgnorePattern: "^_",
        varsIgnorePattern: "[interface|extends]\\w+"
      }
    ],
    //TODO: RESEARCH FOR THIS TO SOLVE CONFLICT OF RULES
    "space-before-function-paren": "off",
    "comma-dangle": "off",
    "arrow-parens": "off",
    "no-underscore-dangle": "off",
    "implicit-arrow-linebreak": "off",
    "object-curly-newline": "off",
    // "prettier/prettier": "off",
    // "vue/html-closing-bracket-spacing": [
    //   "error",
    //   {
    //     selfClosingTag: "never"
    //   }
    // ]
  }
};
