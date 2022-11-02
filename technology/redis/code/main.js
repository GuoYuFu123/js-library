console.log("enter");

const redis = require("redis");

const init = async () => {
  const client = redis.createClient({
    url: "redis://:password@ip:port",
  });

  await client.connect();
  console.log(await client.ping());

  client.flushAll();

  // 字符串
  console.log("######------str-----######");
  await client.set("str", 1);
  const str = await client.get("str");
  console.log(str);
  client.incr("str");
  console.log(await client.get("str"));

  // list
  console.log("#####-----list-----######");
  await client.lPush("list", "1");
  await client.lPush("list", "2");
  await client.rPush("list", "3");
  await client.rPush("list", "4");
  console.log(await client.lRange("list", 0, -1));
  console.log(await client.lPop("list"));

  // set
  console.log("#####-----set----#####");
  await client.sAdd("set", ["m1", "m2", "m3", "m4"]);
  console.log(await client.sMembers("set"));

  // hash
  console.log("#####-----hash----#####");
  await client.hSet("hash", { a: 1, b: 2, c: "gg" });
  console.log(await client.hGetAll("hash"), await client.hGet("hash", "a"));

  // zset
  console.log("#####-----zset----#####");
  await client.zAdd("zset", { score: 100, value: "xiaoming" });
  await client.zAdd("zset", [
    { score: 101, value: "xiaoming2" },
    { score: 99, value: "xiaoming3" },
  ]);

  // done
  console.log('########----done----#######')

  console.log(await client.zCard("zset"), await client.zRange('zset',0, -1))

  client.on("error", (err) => console.log("Redis Client Error", err));
  client.quit();
};

init();
