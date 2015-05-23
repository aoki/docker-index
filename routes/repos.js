var express = require('express');
var fs = require('fs');
var router = express.Router();

router.get('/:id', function(req, res) {
  const REPOSITORY_DIR = './sample_repo';

  res.status(200).json(
    {'repos': fs.readdirSync(REPOSITORY_DIR + '/' + req.params.id)}
  );
});

module.exports = router;
