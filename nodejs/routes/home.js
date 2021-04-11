const router = require('express').Router();
const authController = require("../controller/authController")
const predictController = require("../controller/predictController")

router.post('/x',(req,res)=>{res.send(req.body)})
router.get('/x',predictController.ListWorkers)

router.all('/*',authController.CheckValidToken, predictController.Process) 

module.exports = router