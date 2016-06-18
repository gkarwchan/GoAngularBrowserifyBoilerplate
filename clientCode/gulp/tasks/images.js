var gulp = require('gulp'),
config = require('../config');

gulp.task('images', function() {
  return gulp.src(config.images, {base: config.src})
    .pipe(gulp.dest(config.dst));
});
