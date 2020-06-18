var Pass = {
    Bcrypt:{
        hash: "&&&&%#@#@$#@!@#@$#@$@#*()*(&$(*DAME(*DJAM*(DJSA<D",
        salt:10,
    },

    mysql: {
        connectionLimit: 10,
        multipleStatements: true,
        host: 'localhost',
        user: 'root',
        password: 'rootpassword',
        database: 'hydra_db',
        port: 3306,//3311,
      },
      
    Socket:{
        use: true,
        tokenPatient: "",
        tokenDoctor: "",
        tokenAdmin: ""
    },

    session: {
        secret: 'aSCfa43427*# 326#@.$%213(--)832$$!',
        resave: false,
        saveUninitialized: false
    }
}

module.exports = Pass