var gulp = require('gulp'),
  config = require('../config'),
  sass = require('gulp-sass'),
  merge = require('merge-stream'),
  concat = require('gulp-concat');

gulp.task('styles', function() {

  var scssStream = gulp.src(config.styles.src)
    .pipe(sass({style: 'expanded'}));

  var cssStream = gulp.src(config.styles.cssFiles);

  return merge(scssStream, cssStream)
    .pipe(concat(config.styles.dst))
    .pipe(gulp.dest(config.dst));
});
