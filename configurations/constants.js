
const Constants = {
    Ports:{
        http:8888,
        https:443
    },
    Debug: true,
    SSL:{
        Key: "/etc/letsencrypt/live/hydra/privkey.pem",
        Cert: "/etc/letsencrypt/live/hydra/fullchain.pem"
    }
}

module.exports = Constants