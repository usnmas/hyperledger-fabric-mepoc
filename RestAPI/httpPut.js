var https = require("follow-redirects").https;
var fs = require("fs");

var options = {
  method: "PUT",
  hostname: "130.195.10.173",
  port: 8080,
  path: "/api/changeowner/CAR4",
  headers: {
    "Content-Type": "application/json"
  },
  maxRedirects: 20
};

var req = https.request(options, function(res) {
  var chunks = [];

  res.on("data", function(chunk) {
    chunks.push(chunk);
  });

  res.on("end", function(chunk) {
    var body = Buffer.concat(chunks);
    console.log(body.toString());
  });

  res.on("error", function(error) {
    console.error(error);
  });
});

var postData = JSON.stringify({ owner: "XYZ" });

req.write(postData);

req.end();
