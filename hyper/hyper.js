// See https://hyper.is#cfg for all currently supported options.

module.exports = {
  config: {
    // "stable" or "canary"
    updateChannel: 'stable',

    fontSize: 13,
    fontFamily: 'Fira Code, Consolas, monospace',

    // "BEAM", "UNDERLINE", "BLOCK"
    cursorShape: 'UNDERLINE',
    cursorColor: 'rgba(70, 252, 62, .8)',
    cursorBlink: true,

    foregroundColor: '#fff',
    backgroundColor: '#000',
    borderColor: 'rgba(70, 252, 62, .4)',

    css: '',
    termCSS: '',
    padding: '12px 14px',

    colors: {
      black: '#000000',
      red: '#ff0000',
      green: '#33ff00',
      yellow: '#ffff00',
      blue: '#0066ff',
      magenta: '#cc00ff',
      cyan: '#00ffff',
      white: '#d0d0d0',
      lightBlack: '#808080',
      lightRed: '#ff0000',
      lightGreen: '#33ff00',
      lightYellow: '#ffff00',
      lightBlue: '#0066ff',
      lightMagenta: '#cc00ff',
      lightCyan: '#00ffff',
      lightWhite: '#ffffff'
    },

    // hyperBorder: {
    //   borderColors: ['#2ca164', '#c6ff00'],
    //   blurredColors: ['#e6e6e6'],
    //   borderWidth: '1px',
    // },

    shell: '/usr/local/bin/fish',
    shellArgs: ['--login'],

    bell: 'SOUND',
    env: {},
  },

  // Plugins from npm
  plugins: ['hyper-ligatures'],
  localPlugins: [],

  keymaps: {
    // 'window:devtools': 'cmd+alt+o',
  },
};
