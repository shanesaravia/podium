const sveltePreprocess = require("svelte-preprocess");
const production = !process.env.ROLLUP_WATCH;

module.exports = {
  preprocess: sveltePreprocess({
    sourceMap: !production,
    scss: {
      prependData: `@import 'client/scss/theme.scss';`,
    },
  }),
  onwarn: (warning, handler) => {
    const { code, frame } = warning;
    if (code === "css-unused-selector") return;

    handler(warning);
  },
};
