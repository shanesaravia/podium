const sveltePreprocess = require("svelte-preprocess");

module.exports = {
  preprocess: sveltePreprocess({
    // sourceMap: !production,
    scss: {
      prependData: `@import 'client/scss/theme.scss';`,
    },
  }),
};
