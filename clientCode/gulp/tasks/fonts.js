var gulp = require('gulp'),
    debug = require('gulp-debug'),
    config = require('../config');

gulp.task('fonts', function() {
    return gulp.src(config.fonts.src)
               .pipe(gulp.dest(config.fonts.dst));
});
