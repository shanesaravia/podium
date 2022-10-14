module.exports = {
  parser: '@typescript-eslint/parser',
  parserOptions: {
    ecmaVersion: 2020,
    // project: 'tsconfig.json',
    // tsconfigRootDir: __dirname,
    sourceType: 'module',
    extraFileExtensions: ['.svelte'],
  },
  overrides: [
    {
      files: ['*.svelte'],
      parser: 'svelte-eslint-parser',
      parserOptions: {
        parser: '@typescript-eslint/parser',
      },
    },
  ],
  settings: {
    'import/resolver': {
      // typescript: {},
      // node: {
      //   paths: ['src'],
      //   extensions: ['.js', '.jsx', '.ts', '.tsx', '.svelte'],
      // },
      alias: {
        map: [['src', './src']],
        extensions: ['.js', '.jsx', '.ts', '.tsx', '.d.ts', '.svelte'],
        // '@': path.resolve(__dirname, 'src'),
      },
    },
    // 'svelte3/typescript': require('typescript'),
    // ignore style tags in Svelte because of Tailwind CSS
    // See https://github.com/sveltejs/eslint-plugin-svelte3/issues/70
    // 'svelte3/ignore-styles': () => true,
    'svelte3/typescript': true,
  },
  plugins: ['svelte3', '@typescript-eslint', 'html', 'prettier'],
  extends: [
    'plugin:svelte/recommended',
    'plugin:@typescript-eslint/recommended',
    'plugin:prettier/recommended',
    'plugin:import/errors',
    'plugin:import/typescript',
    'prettier'
  ],
  // root: true,
  env: {
    es6: true,
    browser: true,
    node: true,
    jest: true,
  },
  ignorePatterns: ['public/build/', '.eslintrc.js', 'node_modules', '*.js'],
  rules: {
    "prettier/prettier": [
      "error",
      {
        "svelteBracketNewLine": true,
        "jsxBracketSameLine": true,
      },
      {}
    ],
    'import/order': [
      'error',
      {
        groups: ['builtin', 'external', 'internal', ['parent', 'sibling']],
        pathGroups: [
          {
            pattern: 'react',
            group: 'external',
            position: 'before',
          },
        ],
        'newlines-between': 'always',
        alphabetize: {
          order: 'asc',
          caseInsensitive: true,
        },
      },
    ],
    '@typescript-eslint/member-delimiter-style': [
      'error',
      {
        multiline: {
          delimiter: 'none', // 'none' or 'semi' or 'comma'
          requireLast: true,
        },
        singleline: {
          delimiter: 'semi', // 'semi' or 'comma'
          requireLast: false,
        },
      },
    ],
    '@typescript-eslint/camelcase': 'off',
    '@typescript-eslint/explicit-function-return-type': 'off',
    '@typescript-eslint/explicit-module-boundary-types': 'off',
    '@typescript-eslint/no-explicit-any': 'off',
    '@typescript-eslint/no-use-before-define': 'off',
    '@typescript-eslint/triple-slash-reference': 'off',
    'no-unused-vars': 'off',
    '@typescript-eslint/no-unused-vars': [
      'error',
      {
        vars: 'all',
        args: 'after-used',
        ignoreRestSiblings: false,
      },
    ],
    'import/no-duplicates': 'warn',
  },
}
