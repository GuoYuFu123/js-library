/**
 * csv处理utils
 */

var fs = require("fs");
var parse = require("csv-parse");
var stringify = require("csv-stringify");

var parser = parse({ columns: true }, function (err, records) {
    console.log(records);

  const arr = records.map((it) => {
    it.wRate = it._col7 / it._col6;
    return it;
  });
  
  // i,t,_col2,_col3,_col4,_col5,_col6,_col7,_col8,msg
  const arrFilter = arr.filter((it) => {
    return it.wRate >=3
  })
  stringify(
    arrFilter,
    {
      header: true,
      columns: [
        { key: "i", header: "i" },
        { key: "t", header: "时间" },
        { key: "_col2", header: "sid" },
        { key: "_col3", header: "tid" },
        { key: "_col4", header: "pid" },
        { key: "_col5", header: "contentType" },
        { key: "_col6", header: "clientWidth" },
        { key: "_col7", header: "overWidth" },
        { key: "wRate", header: "比率" },
      ],
    },
    function (err, output) {
      fs.writeFile(__dirname + "/index.target.csv", output, () => {});
    }
  );
});
console.log(__dirname);
fs.createReadStream(__dirname + "/index.csv").pipe(parser);
