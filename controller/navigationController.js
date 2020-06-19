const baseController = require('./baseController');
const constants = require("../configurations/constants")
var NavigationController = {}
NavigationController = Object.create(baseController);


// ==> PAGES

NavigationController.Home = (req,res) =>{
	res.render("home.hbs")
}

module.exports = NavigationController
