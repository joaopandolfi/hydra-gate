const router = require('express').Router();
const authController = require("../controller/authController")
const predictController = require("../controller/predictController")

router.post("/predict",authController.CheckValidToken,predictController.Process)

module.exports = router