const router = require('express').Router();
const authController = require("../controller/authController")

router.post('/token', authController.NewToken)

module.exports = router