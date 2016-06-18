var gulp = require('gulp'),
config = require('../config');

gulp.task('htmlPages', function() {
  return gulp.src(config.html, {base: config.src})
    .pipe(gulp.dest(config.dst));
});
