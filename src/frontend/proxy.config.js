var PROXY_CONFIG = [
  {
    context: [
      "/login",
      "/logout",
      "/currentuser",
      "/api",
      "/openapi",
      "/ws",
      "/healthz",
      "/session"
    ],
    target: "http://192.168.128.94:8080",
    secure: false,
    changeOrigin: true
  }
];

module.exports = PROXY_CONFIG;
