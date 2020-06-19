const authConstants = require('../configurations/pass')
const constants = require('../configurations/constants')
const authController = {}



authController.CheckValidToken = (req,res,next) =>{
    return next()
}

authController.CheckWorkerToken = token =>{
    return true
}

authController.Forbidden = (req,res) =>{
    res.render('forbidden.hbs')
}

authController.NewToken = (req,res) =>{
    res.send({success:true})
}

module.exports = authController 