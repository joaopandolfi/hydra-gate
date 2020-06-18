const router = require('express').Router();
const navigationController = require('../controller/navigationController')
const authController = require("../controller/authController")


router.get('/',authController.CheckValidToken, navigationController.Home)
router.post('/x',(req,res)=>{res.send(req.body)})

module.exports = router