var webRoot = 'web',
  staticPath = '../output';

module.exports = {
  debug: true,
  src : webRoot,
  dst : staticPath,
  html: webRoot + '/**/*.html',
  styles : {
    src : webRoot + '/styles/main.scss',
    cssFiles : ['bower_components/ngToast/dist/*.min.css'],
    dst : staticPath + '/styles/main.css'
  },
  scripts: {
    src : webRoot + '/app/app.js',
    dst : 'scripts/main.js'
  },
  images: webRoot + '/images/**',
  fonts: {
      src: ['bower_components/bootstrap-sass/assets/fonts/**','bower_components/font-awesome/fonts/**'],
      dst: staticPath + '/fonts'
    }
};
