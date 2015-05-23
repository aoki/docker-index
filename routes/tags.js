var express = require('express');
var fs = require('fs');
var router = express.Router();

router.get('/:user/:repo', function(req, res) {
  'use strict';
  const REPOSITORY_DIR = './sample_repo';

  let result = fs.readdirSync(
      REPOSITORY_DIR + '/' + req.params.user + '/' + req.params.repo
  ).filter((f) => {
    if (/tag_(.+)/.test(f)) { return f; }
  }).map((f) => {
    return f.match(/tag_(.+)/)[1];
  });

  res.status(200).json(
    {'tags': result}
  );
});

module.exports = router;
