var express = require('express');
var fs = require('fs');
var router = express.Router();

/* GET users listing. */
router.get('/', function(req, res, next) {
  const REPOSITORY_DIR = './sample_repo';

  console.dir(fs.readdirSync(REPOSITORY_DIR));
  var x = fs.readdirSync(REPOSITORY_DIR).reduce(function(acc, user){
    console.log(user);
    return acc[user] = fs.readdirSync(REPOSITORY_DIR + '/' + user);
  });

  res.status(200).json(
    {'users': fs.readdirSync(REPOSITORY_DIR)}
  );
});

module.exports = router;
