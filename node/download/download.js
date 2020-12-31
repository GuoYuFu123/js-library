const util = require("util");
const stream = require("stream");
const pipeline = util.promisify(stream.pipeline);
const fs = require("fs");
const axios = require("axios");

const downloadFile = async () => {
  try {
    await pipeline(
      (await axios.get("url", { responseType: "stream" })).data,
      fs.createWriteStream("a.html")
    );
    console.log("donwload pdf pipeline successfull");
  } catch (error) {
    console.error("donwload pdf pipeline failed", error);
  }
};

downloadFile();
