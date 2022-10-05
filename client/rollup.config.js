import alias from '@rollup/plugin-alias'
import commonjs from '@rollup/plugin-commonjs'
import copy from 'rollup-plugin-copy'
import css from 'rollup-plugin-css-only'
import livereload from 'rollup-plugin-livereload'
import replace from 'rollup-plugin-replace'
import resolve from '@rollup/plugin-node-resolve'
import svelte from 'rollup-plugin-svelte'
import sveltePreprocess from 'svelte-preprocess'
import { terser } from 'rollup-plugin-terser'
import typescript from '@rollup/plugin-typescript'

const production = !process.env.ROLLUP_WATCH

function serve() {
  let server

  function toExit() {
    if (server) server.kill(0)
  }

  return {
    writeBundle() {
      if (server) return
      server = require('child_process').spawn('npm', ['run', 'start', '--', '--dev'], {
        stdio: ['ignore', 'inherit', 'inherit'],
        shell: true,
      })

      process.on('SIGTERM', toExit)
      process.on('exit', toExit)
    },
  }
}

const aliases = alias({
  resolve: ['.svelte', '.js', '.ts'], //optional, by default this will just look for .js files or folders
  entries: [
    { find: 'src', replacement: 'src' },
    { find: 'pages', replacement: 'src/pages' },
    { find: 'components', replacement: 'src/components' },
    { find: 'utils', replacement: 'src/utils' },
    { find: 'configs', replacement: 'src/configs' },
  ],
})

export default {
  input: 'src/main.ts',
  output: {
    sourcemap: !production,
    format: 'iife',
    name: 'app',
    file: 'public/build/bundle.js',
  },
  plugins: [
    replace({ 'process.env.NODE_ENV': JSON.stringify('development') }),
    copy({
      targets: [
        {
          src: 'node_modules/bootstrap/dist/css/*',
          dest: 'public/vendor/bootstrap/css',
        },
        {
          src: 'node_modules/bootstrap/dist/js/*',
          dest: 'public/vendor/bootstrap/js',
        },
      ],
    }),
    svelte({
      preprocess: sveltePreprocess({
        sourceMap: !production,
        scss: {
          prependData: `@import '../scss/theme.scss';`,
        },
      }),
      compilerOptions: {
        // enable run-time checks when not in production
        dev: !production,
      },
      // onwarn: (warning, handler) => {
      //   const { code, frame } = warning
      //   if (code === 'css-unused-selector') return

      //   handler(warning)
      // },
    }),
    // scss(),
    aliases,
    // we'll extract any component CSS out into
    // a separate file - better for performance
    css({ output: 'bundle.css' }),

    // If you have external dependencies installed from
    // npm, you'll most likely need these plugins. In
    // some cases you'll need additional configuration -
    // consult the documentation for details:
    // https://github.com/rollup/plugins/tree/master/packages/commonjs
    resolve({
      browser: true,
      dedupe: ['svelte'],
    }),
    commonjs(),
    typescript({
      sourceMap: !production,
      inlineSources: !production,
      noUnusedLocals: false,
      noUnusedParameters: false,
    }),

    // In dev mode, call `npm run start` once
    // the bundle has been generated
    !production && serve(),

    // Watch the `public` directory and refresh the
    // browser on changes when not in production
    !production && livereload('public'),

    // If we're building for production (npm run build
    // instead of npm run dev), minify
    production && terser(),
  ],
  watch: {
    clearScreen: false,
  },
}
