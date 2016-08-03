var gulp = require('gulp'), 
    config = require('../config');

gulp.task('watch', ['clean', 'build'], function() {
    gulp.watch(config.fonts.src,    ['fonts']);
    gulp.watch(config.images,       ['images']);
    gulp.watch(config.styles.src,   ['styles']);

});
