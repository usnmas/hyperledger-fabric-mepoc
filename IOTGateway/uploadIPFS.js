const IPFS = require("ipfs-mini");
const ipfs = new IPFS({
  host: "ipfs.infura.io",
  port: 5001,
  protocol: "https"
});
const data = "This is my first message upload to IPFS";

ipfs.add(data, (err, hash) => {
  if (err) {
    return console.log(err);
  }
  console.log("https://ipfs.infura.io/ipfs/" + hash);
});
