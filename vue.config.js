const path = require("path");
module.exports = {
    outputDir: path.resolve(__dirname, "./web/public/dist"),
    pluginOptions: {
        sourceDir: path.resolve(__dirname, "./web/public/src")
    }
};