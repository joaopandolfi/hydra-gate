const router = require('express').Router();
const authRouter = require('./auth')
const navigationRouter = require('./home')
const aliasRouter = require('./alias')

router.use('/', aliasRouter)
router.use('/', authRouter)
router.use('/', navigationRouter)

router.get('*', function(req, res){
  res.render('404.hbs')
});

module.exports = router