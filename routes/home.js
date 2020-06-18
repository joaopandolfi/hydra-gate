const router = require('express').Router();
const navigationController = require('../controller/navigationController')
const authController = require("../controller/authController")


router.get('/',authController.CheckValidToken, navigationController.Home)

module.exports = router