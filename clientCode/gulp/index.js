var fs = require('fs');
var gracefulFs = require('graceful-fs');
var tasks = fs.readdirSync('./gulp/tasks');

gracefulFs.gracefulify(fs);

tasks.forEach(function(task) {
  require('./tasks/' + task);
});
