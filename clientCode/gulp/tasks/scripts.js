var gulp = require('gulp'),
  config = require('../config'),
  browserify = require('browserify'),
  source = require('vinyl-source-stream');

gulp.task('scripts', function() {
  var bundler = browserify({
    debug: config.debug,
    noparse: ['angular', 'angular-bootstrap', 'angular-cookies', 'angular-messages', 'angular-resources', 'angular-route', 'angular-sanitize', 'angular-ui-router', 'angular-ui-select']
  });

  if (!config.debug) {
    bundler.transform({
      global: true,
      mangle: true,
      compress: {
        sequences: true,
        dead_code: true,
        drop_debugger: true,
        properties: true,
        unsafe: true,
        conditionals: true,             // optimize if-s and conditional expressions
        comparisons: true,              // optimize comparisons
        evaluate: true,                 // evaluate constant expressions
        booleans: true,                 // optimize boolean expressions
        loops: true,                    // optimize loops
        unused: true,                   // drop unused variables/functions
        hoist_funs: true,               // hoist function declarations
        hoist_vars: false,              // hoist variable declarations
        if_return: true,                // optimize if-s followed by return/continue
        join_vars: true,                // join var declarations
        cascade: true,                  // try to cascade 'right' into 'left' in sequences
        side_effects: true,             // drop side-effect-free statements
        warnings: false                // warn about potentially dangerous optimizations/code
      }
    }, 'uglifyify');
  }

  bundler.add(config.scripts.src);
  return bundler
    .bundle()
    .pipe(source(config.scripts.dst))
    .pipe(gulp.dest(config.dst))
});
